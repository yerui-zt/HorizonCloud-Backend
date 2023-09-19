package pvectl

import (
	"HorizonX/model"
	"context"
	"crypto/tls"
	"github.com/luthermonson/go-proxmox"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type PVECtl struct {
	ctx    context.Context
	Client *proxmox.Client
	Node   *proxmox.Node
}

func NewPVECtl(ctx context.Context, node *model.HypervisorNode) *PVECtl {
	return &PVECtl{
		ctx:    ctx,
		Client: newPVECtlClient(node),
		Node:   newPVECtlNode(ctx, node),
	}
}
func newPVECtlClient(node *model.HypervisorNode) *proxmox.Client {
	endPoint := node.PveApi
	tokenID := node.PveTokenId
	secret := node.PveSecret

	insecureHTTPClient := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	client := proxmox.NewClient(endPoint,
		// todo 安全 - 是否需要跳过证书验证
		proxmox.WithClient(&insecureHTTPClient),

		proxmox.WithAPIToken(tokenID, secret),
	)
	return client
}
func newPVECtlNode(ctx context.Context, node *model.HypervisorNode) *proxmox.Node {
	nodeName := node.PveNodeName
	client := newPVECtlClient(node)

	pvenode, err := client.Node(nodeName)
	if err != nil {
		logx.WithContext(ctx).Errorf("get proxmox node failed: %v", err)
		return nil
	}
	return pvenode
}

// todo: 重构，go-proxmox库的Delete没有附带
//		purge=1&destroy-unreferenced-disks=1
// 		https://178.173.230.163:8006/api2/extjs/nodes/pvetest/qemu/1001?purge=1&destroy-unreferenced-disks=1
//func (l *PVECtl) DeleteVM()

func (l *PVECtl) CleanupVirtualMachine(vm *proxmox.VirtualMachine) {
	task, err := vm.Stop()
	if err != nil {
		logx.WithContext(l.ctx).Errorf("create stop vm task failed: %v", err)
	}
	_, c, err := task.WaitForCompleteStatus(60)
	if err != nil || !c {
		logx.WithContext(l.ctx).Errorf("exec stop vm task failed: %v", err)
		return
	}

	//todo: 可能有bug
	//  https://github.com/luthermonson/go-proxmox/blob/37d0caadb112464943857328b2971abc38014f8f/tests/integration/virtual_machines_test.go#L86
	//if vm.VirtualMachineConfig != nil && vm.VirtualMachineConfig.IDE2 != "" {
	//	s := strings.Split(vm.VirtualMachineConfig.IDE2, ",")
	//	if len(s) > 2 {
	//		iso, err := l.Node.StorageISO()
	//
	//		task, err := iso.Delete()
	//		require.NoError(t, err)
	//		require.NoError(t, task.Wait(1*time.Second, 10*time.Second))
	//	}
	//}
	// 删除
	_, err = vm.Delete()
	if err != nil {
		logx.WithContext(l.ctx).Errorf("create delete vm task failed: %v", err)
	}

}
