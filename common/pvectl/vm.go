package pvectl

import (
	"HorizonX/common/xerr"
	"HorizonX/model"
	"HorizonX/rpc/vm/vm"
	"fmt"
	"github.com/luthermonson/go-proxmox"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
	"net"
	"regexp"
	"strconv"
	"strings"
)

func (l *PVECtl) NewVirtualMachine(imageVmID int, hostname string) (newVmID int, err error) {
	// 1. 创建虚拟机
	// 2. 配置虚拟机，如果发生错误

	// Step1: 创建虚拟机
	template, err := l.Node.VirtualMachine(imageVmID)
	if err != nil {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.OS_IMAGE_NOT_AVAILABLE), "get template failed: %v", err)
	}
	newid, task, err := template.Clone(&proxmox.VirtualMachineCloneOptions{
		Name: hostname,
		// Full clone
		Full: 1,
	})
	if err != nil {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.PROXMOX_VM_CREATE_FAILED), "create clone template task failed: %v", err)
	}
	logx.WithContext(l.ctx).Infof("create clone vm task success, [hypervisor node: %s] newid: %d", l.Node.Name, newid)

	_, completed, err := task.WaitForCompleteStatus(60)
	if err != nil || !completed {
		// 此时stop即可，pve会自动删除虚拟机
		err := task.Stop()
		if err != nil {
			logx.WithContext(l.ctx).Errorf("stop task failed: %v", err)
		}
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.PROXMOX_VM_CREATE_FAILED), "clone task failed: %v", err)
	}

	return newid, nil
}

func (l *PVECtl) UpdateVMResourceFromPlan(vm *proxmox.VirtualMachine, plan *model.VmPlan) error {
	task, err := vm.Config(
		// CPU
		proxmox.VirtualMachineOption{Name: "cores", Value: cast.ToInt(plan.Vcpu)},
		// 内存
		proxmox.VirtualMachineOption{Name: "memory", Value: cast.ToInt(plan.Memory)},
		// 带宽
		proxmox.VirtualMachineOption{Name: "net0", Value: l.setNetworkRate(vm.VirtualMachineConfig.Net0, cast.ToInt(plan.Bandwidth)/8)},
		// 磁盘
		proxmox.VirtualMachineOption{Name: "scsi0", Value: l.setSCSI0DiskSize(vm.VirtualMachineConfig.SCSI0, cast.ToInt(plan.Disk))},
	)
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.PROXMOX_VM_CONFIG_ERROR), "create config vm task failed: %v", err)
	}
	_, completed, err := task.WaitForCompleteStatus(30)
	if err != nil || !completed {
		return errors.Wrapf(xerr.NewErrCode(xerr.PROXMOX_VM_CONFIG_ERROR), "config task failed: %v", err)
	}
	return nil
}

func (l *PVECtl) SetSSHKey(v *proxmox.VirtualMachine, sshkey string) error {
	task, err := v.Config(
		proxmox.VirtualMachineOption{Name: "sshkeys", Value: sshkey},
	)
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.PROXMOX_VM_CONFIG_ERROR), "create config vm task failed: %v", err)
	}
	_, completed, err := task.WaitForCompleteStatus(30)
	if err != nil || !completed {
		return errors.Wrapf(xerr.NewErrCode(xerr.PROXMOX_VM_CONFIG_ERROR), "config task failed: %v", err)
	}
	return nil

}

// AddIPFilter 新增IP过滤，不会覆盖原有的设置

func (l *PVECtl) SetVmMainIPAddr(v *proxmox.VirtualMachine, ipv4addr []vm.IPv4Address, ipv6addr []vm.IPv6Address) error {
	var builder string
	if ipv4addr != nil {
		builder = fmt.Sprintf("ip=%s,gw=%s", ipv4addr[0].Ip, ipv4addr[0].Gateway)
	}
	if ipv6addr != nil {
		builder += fmt.Sprintf(",ip6=%s,gw6=%s", ipv6addr[0].Ip, ipv6addr[0].Gateway)
	}
	task, err := v.Config(
		proxmox.VirtualMachineOption{
			Name:  "ipconfig0",
			Value: builder,
		},
	)
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.PROXMOX_VM_CONFIG_ERROR), "create config vm task failed: %v", err)
	}
	_, completed, err := task.WaitForCompleteStatus(30)
	if err != nil || !completed {
		return errors.Wrapf(xerr.NewErrCode(xerr.PROXMOX_VM_CONFIG_ERROR), "config task failed: %v", err)
	}
	return nil
}

func (l *PVECtl) AddIPFilter(v *proxmox.VirtualMachine, ipv4addr []vm.IPv4Address, ipv6addr []vm.IPv6Address) error {
	type IpFilterJSON struct {
		CIDR    string `json:"cidr,omitempty"`
		NoMatch int    `json:"nomatch,omitempty"`
		Comment string `json:"comment,omitempty"`
	}
	if ipv4addr != nil {
		ipStr := ipv4addr[0].Ip
		ip, _, err := net.ParseCIDR(ipStr)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCodeMsg(400, "Invalid IP CIDR Format"), "parse ipv4 failed: %v", err)
		}
		v4 := IpFilterJSON{
			CIDR:    ip.String(),
			NoMatch: 0,
		}
		err = l.Client.Post(fmt.Sprintf("/nodes/%s/qemu/%d/firewall/ipset/ipfilter-net0", v.Node, v.VMID), v4, nil)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.PROXMOX_VM_CONFIG_ERROR), "set ipv4 filter failed: %v", err)
		}
	}
	if ipv6addr != nil {
		ipStr := ipv6addr[0].Ip
		ip, _, err := net.ParseCIDR(ipStr)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCodeMsg(400, "Invalid IP CIDR Format"), "parse ipv6 failed: %v", err)
		}
		v6 := IpFilterJSON{
			CIDR:    ip.String(),
			NoMatch: 0,
		}
		err = l.Client.Post(fmt.Sprintf("/nodes/%s/qemu/%d/firewall/ipset/ipfilter-net0", v.Node, v.VMID), v6, nil)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.PROXMOX_VM_CONFIG_ERROR), "set ipv6 filter failed: %v", err)
		}
	}
	return nil
}

func (l *PVECtl) setNetworkRate(netConfig string, newRate int) string {

	if newRate == 0 {
		if strings.Contains(netConfig, "rate=") {
			// 如果存在，删除rate
			r := regexp.MustCompile(`,rate=\d+`)
			netConfig = r.ReplaceAllString(netConfig, "")
			return netConfig
		} else {
			return netConfig
		}
	}

	if !strings.Contains(netConfig, "rate=") {
		// 如果不存在，在 oldStr 最后加上 ,rate=newRate
		netConfig += ",rate=" + strconv.Itoa(newRate)
	} else {
		// 如果存在，替换rate=newRate
		r := regexp.MustCompile(`rate=\d+`)
		netConfig = r.ReplaceAllString(netConfig, "rate="+strconv.Itoa(newRate))
	}
	return netConfig
}
func (l *PVECtl) setSCSI0DiskSize(scsi0Config string, newSize int) string {
	if newSize == 0 {
		return scsi0Config
	}
	// 使用正则表达式提取 Size 数字
	re := regexp.MustCompile(`size=(\d+)`)
	match := re.FindStringSubmatch(scsi0Config)
	if len(match) != 2 {
		return scsi0Config
	}
	oldSize, err := strconv.Atoi(match[1])
	if err != nil {
		return scsi0Config
	}
	if newSize <= oldSize {
		return scsi0Config
	}
	// 替换字符
	newConfig := strings.Replace(scsi0Config, fmt.Sprintf("size=%d", oldSize), fmt.Sprintf("size=%d", newSize), 1)
	return newConfig
}
