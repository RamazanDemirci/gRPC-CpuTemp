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

// OpenWeatherClient is the client API for OpenWeather service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OpenWeatherClient interface {
	// A simple RPC.
	//
	// Obtains the current location
	//
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	GetTemperatureByLocation(ctx context.Context, in *Location, opts ...grpc.CallOption) (*Temperature, error)
}

type openWeatherClient struct {
	cc grpc.ClientConnInterface
}

func NewOpenWeatherClient(cc grpc.ClientConnInterface) OpenWeatherClient {
	return &openWeatherClient{cc}
}

func (c *openWeatherClient) GetTemperatureByLocation(ctx context.Context, in *Location, opts ...grpc.CallOption) (*Temperature, error) {
	out := new(Temperature)
	err := c.cc.Invoke(ctx, "/openweather.OpenWeather/GetTemperatureByLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OpenWeatherServer is the server API for OpenWeather service.
// All implementations must embed UnimplementedOpenWeatherServer
// for forward compatibility
type OpenWeatherServer interface {
	// A simple RPC.
	//
	// Obtains the current location
	//
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	GetTemperatureByLocation(context.Context, *Location) (*Temperature, error)
	mustEmbedUnimplementedOpenWeatherServer()
}

// UnimplementedOpenWeatherServer must be embedded to have forward compatible implementations.
type UnimplementedOpenWeatherServer struct {
}

func (UnimplementedOpenWeatherServer) GetTemperatureByLocation(context.Context, *Location) (*Temperature, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTemperatureByLocation not implemented")
}
func (UnimplementedOpenWeatherServer) mustEmbedUnimplementedOpenWeatherServer() {}

// UnsafeOpenWeatherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OpenWeatherServer will
// result in compilation errors.
type UnsafeOpenWeatherServer interface {
	mustEmbedUnimplementedOpenWeatherServer()
}

func RegisterOpenWeatherServer(s grpc.ServiceRegistrar, srv OpenWeatherServer) {
	s.RegisterService(&OpenWeather_ServiceDesc, srv)
}

func _OpenWeather_GetTemperatureByLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Location)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenWeatherServer).GetTemperatureByLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openweather.OpenWeather/GetTemperatureByLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenWeatherServer).GetTemperatureByLocation(ctx, req.(*Location))
	}
	return interceptor(ctx, in, info, handler)
}

// OpenWeather_ServiceDesc is the grpc.ServiceDesc for OpenWeather service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OpenWeather_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "openweather.OpenWeather",
	HandlerType: (*OpenWeatherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTemperatureByLocation",
			Handler:    _OpenWeather_GetTemperatureByLocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc_OpenWeather/protobuf/open_weather.proto",
}
