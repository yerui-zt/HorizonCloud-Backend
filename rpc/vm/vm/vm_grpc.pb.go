// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: vm.proto

package vm

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// VMServiceClient is the client API for VMService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VMServiceClient interface {
	DeployVMInstance(ctx context.Context, in *DeployVMInstanceReq, opts ...grpc.CallOption) (*DeployVMInstanceResp, error)
}

type vMServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVMServiceClient(cc grpc.ClientConnInterface) VMServiceClient {
	return &vMServiceClient{cc}
}

func (c *vMServiceClient) DeployVMInstance(ctx context.Context, in *DeployVMInstanceReq, opts ...grpc.CallOption) (*DeployVMInstanceResp, error) {
	out := new(DeployVMInstanceResp)
	err := c.cc.Invoke(ctx, "/vm.VMService/DeployVMInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VMServiceServer is the server API for VMService service.
// All implementations must embed UnimplementedVMServiceServer
// for forward compatibility
type VMServiceServer interface {
	DeployVMInstance(context.Context, *DeployVMInstanceReq) (*DeployVMInstanceResp, error)
	mustEmbedUnimplementedVMServiceServer()
}

// UnimplementedVMServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVMServiceServer struct {
}

func (UnimplementedVMServiceServer) DeployVMInstance(context.Context, *DeployVMInstanceReq) (*DeployVMInstanceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeployVMInstance not implemented")
}
func (UnimplementedVMServiceServer) mustEmbedUnimplementedVMServiceServer() {}

// UnsafeVMServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VMServiceServer will
// result in compilation errors.
type UnsafeVMServiceServer interface {
	mustEmbedUnimplementedVMServiceServer()
}

func RegisterVMServiceServer(s grpc.ServiceRegistrar, srv VMServiceServer) {
	s.RegisterService(&VMService_ServiceDesc, srv)
}

func _VMService_DeployVMInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeployVMInstanceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMServiceServer).DeployVMInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vm.VMService/DeployVMInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMServiceServer).DeployVMInstance(ctx, req.(*DeployVMInstanceReq))
	}
	return interceptor(ctx, in, info, handler)
}

// VMService_ServiceDesc is the grpc.ServiceDesc for VMService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VMService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vm.VMService",
	HandlerType: (*VMServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeployVMInstance",
			Handler:    _VMService_DeployVMInstance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vm.proto",
}
