// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package enclavepb

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

// EnclaveClient is the client API for Enclave service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EnclaveClient interface {
	ReceiveOnboardTxn(ctx context.Context, in *OnboardTxn, opts ...grpc.CallOption) (*Status, error)
	ReceiveOffboardTxn(ctx context.Context, in *OffboardTxn, opts ...grpc.CallOption) (*Status, error)
}

type enclaveClient struct {
	cc grpc.ClientConnInterface
}

func NewEnclaveClient(cc grpc.ClientConnInterface) EnclaveClient {
	return &enclaveClient{cc}
}

func (c *enclaveClient) ReceiveOnboardTxn(ctx context.Context, in *OnboardTxn, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/enclavepb.Enclave/ReceiveOnboardTxn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enclaveClient) ReceiveOffboardTxn(ctx context.Context, in *OffboardTxn, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/enclavepb.Enclave/ReceiveOffboardTxn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnclaveServer is the server API for Enclave service.
// All implementations must embed UnimplementedEnclaveServer
// for forward compatibility
type EnclaveServer interface {
	ReceiveOnboardTxn(context.Context, *OnboardTxn) (*Status, error)
	ReceiveOffboardTxn(context.Context, *OffboardTxn) (*Status, error)
	mustEmbedUnimplementedEnclaveServer()
}

// UnimplementedEnclaveServer must be embedded to have forward compatible implementations.
type UnimplementedEnclaveServer struct {
}

func (UnimplementedEnclaveServer) ReceiveOnboardTxn(context.Context, *OnboardTxn) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveOnboardTxn not implemented")
}
func (UnimplementedEnclaveServer) ReceiveOffboardTxn(context.Context, *OffboardTxn) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveOffboardTxn not implemented")
}
func (UnimplementedEnclaveServer) mustEmbedUnimplementedEnclaveServer() {}

// UnsafeEnclaveServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EnclaveServer will
// result in compilation errors.
type UnsafeEnclaveServer interface {
	mustEmbedUnimplementedEnclaveServer()
}

func RegisterEnclaveServer(s grpc.ServiceRegistrar, srv EnclaveServer) {
	s.RegisterService(&Enclave_ServiceDesc, srv)
}

func _Enclave_ReceiveOnboardTxn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnboardTxn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnclaveServer).ReceiveOnboardTxn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/enclavepb.Enclave/ReceiveOnboardTxn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnclaveServer).ReceiveOnboardTxn(ctx, req.(*OnboardTxn))
	}
	return interceptor(ctx, in, info, handler)
}

func _Enclave_ReceiveOffboardTxn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OffboardTxn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnclaveServer).ReceiveOffboardTxn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/enclavepb.Enclave/ReceiveOffboardTxn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnclaveServer).ReceiveOffboardTxn(ctx, req.(*OffboardTxn))
	}
	return interceptor(ctx, in, info, handler)
}

// Enclave_ServiceDesc is the grpc.ServiceDesc for Enclave service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Enclave_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "enclavepb.Enclave",
	HandlerType: (*EnclaveServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReceiveOnboardTxn",
			Handler:    _Enclave_ReceiveOnboardTxn_Handler,
		},
		{
			MethodName: "ReceiveOffboardTxn",
			Handler:    _Enclave_ReceiveOffboardTxn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "enclave.proto",
}