package vm

import (
	"HorizonX/api/internal/svc"
	"HorizonX/api/internal/types"
	"HorizonX/rpc/order/orderservice"
	"context"
	"github.com/spf13/cast"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeployVMInstanceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeployVMInstanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeployVMInstanceLogic {
	return &DeployVMInstanceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeployVMInstanceLogic) DeployVMInstance(req *types.DeployVMInstanceReq) (resp *types.DeployVMInstanceResp, err error) {
	uid := l.ctx.Value("uid").(string)
	// todo: 验证 req.SSHKey 合法性
	// 		可以考虑存储到数据库，然后再验证
	if err != nil {
		return nil, err
	}
	rpcResp, err := l.svcCtx.OrderRPC.CreateVMDeployOrder(l.ctx, &orderservice.CreateVMDeployOrderReq{
		Hostname:     req.HostName,
		Uid:          cast.ToInt64(uid),
		VmGroupId:    req.GroupId,
		PlanId:       req.PlanID,
		Image:        req.Image,
		BillingCycle: req.BillingCycle,
	})
	if err != nil {
		return nil, err
	}

	return &types.DeployVMInstanceResp{
		OrderNo: rpcResp.OrderNo,
	}, nil
}
