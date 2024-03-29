// Code generated by goctl. DO NOT EDIT.
// Source: vm.proto

package vmservice

import (
	"context"

	"HorizonX/rpc/vm/vm"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	DeployVMInstanceReq  = vm.DeployVMInstanceReq
	DeployVMInstanceResp = vm.DeployVMInstanceResp
	GetFreeIPv4Req       = vm.GetFreeIPv4Req
	GetFreeIPv4Resp      = vm.GetFreeIPv4Resp
	GetFreeIPv6Req       = vm.GetFreeIPv6Req
	GetFreeIPv6Resp      = vm.GetFreeIPv6Resp
	IPv4Address          = vm.IPv4Address
	IPv6Address          = vm.IPv6Address

	VMService interface {
		DeployVMInstance(ctx context.Context, in *DeployVMInstanceReq, opts ...grpc.CallOption) (*DeployVMInstanceResp, error)
	}

	defaultVMService struct {
		cli zrpc.Client
	}
)

func NewVMService(cli zrpc.Client) VMService {
	return &defaultVMService{
		cli: cli,
	}
}

func (m *defaultVMService) DeployVMInstance(ctx context.Context, in *DeployVMInstanceReq, opts ...grpc.CallOption) (*DeployVMInstanceResp, error) {
	client := vm.NewVMServiceClient(m.cli.Conn())
	return client.DeployVMInstance(ctx, in, opts...)
}
