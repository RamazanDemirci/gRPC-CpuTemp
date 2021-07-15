// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc

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

// CpuTempClient is the client API for CpuTemp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CpuTempClient interface {
	// A simple RPC.
	//
	// Obtains the current location
	//
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	GetCpuTemperature(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Units, error)
}

type cpuTempClient struct {
	cc grpc.ClientConnInterface
}

func NewCpuTempClient(cc grpc.ClientConnInterface) CpuTempClient {
	return &cpuTempClient{cc}
}

func (c *cpuTempClient) GetCpuTemperature(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Units, error) {
	out := new(Units)
	err := c.cc.Invoke(ctx, "/cputemp.CpuTemp/GetCpuTemperature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CpuTempServer is the server API for CpuTemp service.
// All implementations must embed UnimplementedCpuTempServer
// for forward compatibility
type CpuTempServer interface {
	// A simple RPC.
	//
	// Obtains the current location
	//
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	GetCpuTemperature(context.Context, *Empty) (*Units, error)
	mustEmbedUnimplementedCpuTempServer()
}

// UnimplementedCpuTempServer must be embedded to have forward compatible implementations.
type UnimplementedCpuTempServer struct {
}

func (UnimplementedCpuTempServer) GetCpuTemperature(context.Context, *Empty) (*Units, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCpuTemperature not implemented")
}
func (UnimplementedCpuTempServer) mustEmbedUnimplementedCpuTempServer() {}

// UnsafeCpuTempServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CpuTempServer will
// result in compilation errors.
type UnsafeCpuTempServer interface {
	mustEmbedUnimplementedCpuTempServer()
}

func RegisterCpuTempServer(s grpc.ServiceRegistrar, srv CpuTempServer) {
	s.RegisterService(&CpuTemp_ServiceDesc, srv)
}

func _CpuTemp_GetCpuTemperature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CpuTempServer).GetCpuTemperature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cputemp.CpuTemp/GetCpuTemperature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CpuTempServer).GetCpuTemperature(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// CpuTemp_ServiceDesc is the grpc.ServiceDesc for CpuTemp service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CpuTemp_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cputemp.CpuTemp",
	HandlerType: (*CpuTempServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCpuTemperature",
			Handler:    _CpuTemp_GetCpuTemperature_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc_Cpu/protobuf/cpu_temp.proto",
}