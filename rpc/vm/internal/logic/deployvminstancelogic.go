package logic

import (
	"HorizonX/common/pvectl"
	"HorizonX/common/tools"
	"HorizonX/common/xerr"
	"HorizonX/model"
	"context"
	"github.com/luthermonson/go-proxmox"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"

	"HorizonX/rpc/vm/internal/svc"
	"HorizonX/rpc/vm/vm"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeployVMInstanceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeployVMInstanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeployVMInstanceLogic {
	return &DeployVMInstanceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeployVMInstanceLogic) DeployVMInstance(in *vm.DeployVMInstanceReq) (*vm.DeployVMInstanceResp, error) {
	// todo: 队列 maxRetry 设置
	// todo: 拆分 创建和配置，创建不可以重试，配置可以重试
	whereBuilder := l.svcCtx.HypervisorNodeModel.SelectBuilder().Where(
		"group_id = ? AND enable = 1 AND virt_type = 'pve'", in.GroupId)
	nodes, err := l.svcCtx.HypervisorNodeModel.FindAll(l.ctx, whereBuilder, "id ASC")
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get hypervisor node failed: %v", err)
	}
	node := nodes[0]
	// todo: 新增 node 权重，且目前只会使用第一个

	plan, err := l.svcCtx.VmPlanModel.FindOne(l.ctx, nil, in.PlanId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ORDER_PLAN_NOT_FOUND), "get vm plan failed: %v", err)
	}

	// 先创建数据库记录，分配IP （使用事务）
	// 然后再实际创建虚拟机
	var instance *model.VmInstance
	var ipv4Info []vm.IPv4Address
	var ipv6Info []vm.IPv6Address
	var newVM *proxmox.VirtualMachine
	var hypervisor *pvectl.PVECtl
	err = l.svcCtx.VmInstanceModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		now := time.Now()
		insertInstance := model.VmInstance{
			Status:           "pending",
			HypervisorNodeId: node.Id,
			Name:             in.Hostname,
			DueDate:          tools.CalculateDueDate(now, in.BillingCycle),
			BillingCycle:     in.BillingCycle,
			Price:            tools.CalculatePrice(plan, in.BillingCycle),
			PlanId:           plan.Id,
			Vcpu:             plan.Vcpu,
			Memory:           plan.Memory,
			Disk:             plan.Disk,
			Bandwidth:        plan.Bandwidth,
			DataTransfer:     plan.DataTransfer,
			Traffic:          0,
		}
		insertRes, err := l.svcCtx.VmInstanceModel.Insert(l.ctx, session, &insertInstance)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "create vm instance failed: %v", err)
		}
		insertId, _ := insertRes.LastInsertId()
		instance, err = l.svcCtx.VmInstanceModel.FindOne(l.ctx, session, insertId)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "find vm instance failed: %v", err)
		}

		// 查找ip地址
		ipGroup, err := l.svcCtx.IpGroupModel.FindOne(l.ctx, session, node.Id)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.IP_NO_AVAILABLE_ADDR_ERROR), "find ip group failed: %v", err)
		}
		var ipv4s []*model.IpAddress
		var ipv6s []*model.IpAddress
		if plan.Ipv4Num > 0 {
			v4WhereBuilder := l.svcCtx.IpAddressModel.SelectBuilder().
				Where("group_id = ? AND instance_id = 0 AND type = 'ipv4'", ipGroup.Id).
				Limit(cast.ToUint64(plan.Ipv4Num))
			ipv4s, err = l.svcCtx.IpAddressModel.FindAll(l.ctx, v4WhereBuilder, "id ASC")
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.IP_NO_AVAILABLE_ADDR_ERROR), "find ip address failed: %v", err)
			}
			if len(ipv4s) < cast.ToInt(plan.Ipv4Num) {
				return errors.Wrapf(xerr.NewErrCode(xerr.IP_NO_AVAILABLE_ADDR_ERROR), "no available ip address")
			}
			// 绑定 IP地址 - instance_id
			for _, ipv4 := range ipv4s {
				ipv4Info = append(ipv4Info, vm.IPv4Address{
					Ip:      ipv4.Ip,
					Gateway: ipv4.Gateway,
				})
				ipv4.InstanceId.Int64 = instance.Id
				ipv4.InstanceId.Valid = true
				_, err = l.svcCtx.IpAddressModel.Update(l.ctx, session, ipv4)
				if err != nil {
					return errors.Wrapf(xerr.NewErrCode(xerr.IP_NO_AVAILABLE_ADDR_ERROR), "update ip address failed: %v", err)
				}
			}
		}
		if plan.Ipv6Num > 0 {
			v6WhereBuilder := l.svcCtx.IpAddressModel.SelectBuilder().
				Where("group_id = ? AND instance_id = 0 AND type = 'ipv6'", ipGroup.Id).
				Limit(cast.ToUint64(plan.Ipv6Num))
			ipv6s, err = l.svcCtx.IpAddressModel.FindAll(l.ctx, v6WhereBuilder, "id ASC")
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.IP_NO_AVAILABLE_ADDR_ERROR), "find ip address failed: %v", err)
			}
			if len(ipv6s) < cast.ToInt(plan.Ipv6Num) {
				return errors.Wrapf(xerr.NewErrCode(xerr.IP_NO_AVAILABLE_ADDR_ERROR), "no available ip address")
			}
			for _, ipv6 := range ipv6s {
				ipv6Info = append(ipv6Info, vm.IPv6Address{
					Ip:      ipv6.Ip,
					Gateway: ipv6.Gateway,
				})
				ipv6.InstanceId.Int64 = instance.Id
				ipv6.InstanceId.Valid = true
				_, err = l.svcCtx.IpAddressModel.Update(l.ctx, session, ipv6)
				if err != nil {
					return errors.Wrapf(xerr.NewErrCode(xerr.IP_NO_AVAILABLE_ADDR_ERROR), "update binding ip address failed: %v", err)
				}
			}
		}

		// 调用pvectl实际创建虚拟机
		hypervisor = pvectl.NewPVECtl(l.ctx, node)
		// todo: 处理hostname，去掉特殊字符
		newVmID, err := hypervisor.NewVirtualMachine(cast.ToInt(in.ImageId), in.Hostname)
		if err != nil {
			return err
		}
		newVM, err = hypervisor.Node.VirtualMachine(newVmID)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.PROXMOX_VM_FETCH_ERROR), "get vm failed: %v", err)
		}

		// 更新vm_instance中的vm_id
		instance.Vmid = cast.ToInt64(newVmID)
		instance.Status = "active"
		_, err = l.svcCtx.VmInstanceModel.Update(l.ctx, session, instance)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "update vm instance failed: %v", err)
		}

		return nil

	})
	if err != nil {
		return nil, err
	}

	// 3.设置CPU、内存、磁盘、带宽
	err = hypervisor.UpdateVMResourceFromPlan(newVM, plan)
	if err != nil {
		return nil, err
	}

	// 4. 设置IP地址
	err = hypervisor.SetVmMainIPAddr(newVM, ipv4Info, ipv6Info)
	if err != nil {
		return nil, err
	}
	// 5. 设置IP过滤
	err = hypervisor.AddIPFilter(newVM, ipv4Info, ipv6Info)
	if err != nil {
		return nil, err
	}
	// 6. 设置用户 SSH 公钥
	err = hypervisor.SetSSHKey(newVM, tools.RawURLEncode(in.SshKey))
	if err != nil {
		return nil, err
	}

	newVM.Start()

	return &vm.DeployVMInstanceResp{
		InstanceId: cast.ToString(instance.Id),
	}, nil

}
