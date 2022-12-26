// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: protos/applications_v1.proto

package protos

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

// ApplicationsClient is the client API for Applications service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApplicationsClient interface {
	GetApplications(ctx context.Context, in *ApplicationPageRequest, opts ...grpc.CallOption) (*ApplicationPageReply, error)
	GetApplicationById(ctx context.Context, in *ApplicationIdRequest, opts ...grpc.CallOption) (*ApplicationObjectReply, error)
	CreateApplication(ctx context.Context, in *ApplicationObjectRequest, opts ...grpc.CallOption) (*ApplicationObjectReply, error)
	UpdateApplication(ctx context.Context, in *ApplicationObjectRequest, opts ...grpc.CallOption) (*ApplicationObjectReply, error)
	DeleteApplicationById(ctx context.Context, in *ApplicationIdRequest, opts ...grpc.CallOption) (*ApplicationObjectReply, error)
}

type applicationsClient struct {
	cc grpc.ClientConnInterface
}

func NewApplicationsClient(cc grpc.ClientConnInterface) ApplicationsClient {
	return &applicationsClient{cc}
}

func (c *applicationsClient) GetApplications(ctx context.Context, in *ApplicationPageRequest, opts ...grpc.CallOption) (*ApplicationPageReply, error) {
	out := new(ApplicationPageReply)
	err := c.cc.Invoke(ctx, "/applications_v1.Applications/get_applications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationsClient) GetApplicationById(ctx context.Context, in *ApplicationIdRequest, opts ...grpc.CallOption) (*ApplicationObjectReply, error) {
	out := new(ApplicationObjectReply)
	err := c.cc.Invoke(ctx, "/applications_v1.Applications/get_application_by_id", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationsClient) CreateApplication(ctx context.Context, in *ApplicationObjectRequest, opts ...grpc.CallOption) (*ApplicationObjectReply, error) {
	out := new(ApplicationObjectReply)
	err := c.cc.Invoke(ctx, "/applications_v1.Applications/create_application", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationsClient) UpdateApplication(ctx context.Context, in *ApplicationObjectRequest, opts ...grpc.CallOption) (*ApplicationObjectReply, error) {
	out := new(ApplicationObjectReply)
	err := c.cc.Invoke(ctx, "/applications_v1.Applications/update_application", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationsClient) DeleteApplicationById(ctx context.Context, in *ApplicationIdRequest, opts ...grpc.CallOption) (*ApplicationObjectReply, error) {
	out := new(ApplicationObjectReply)
	err := c.cc.Invoke(ctx, "/applications_v1.Applications/delete_application_by_id", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApplicationsServer is the server API for Applications service.
// All implementations must embed UnimplementedApplicationsServer
// for forward compatibility
type ApplicationsServer interface {
	GetApplications(context.Context, *ApplicationPageRequest) (*ApplicationPageReply, error)
	GetApplicationById(context.Context, *ApplicationIdRequest) (*ApplicationObjectReply, error)
	CreateApplication(context.Context, *ApplicationObjectRequest) (*ApplicationObjectReply, error)
	UpdateApplication(context.Context, *ApplicationObjectRequest) (*ApplicationObjectReply, error)
	DeleteApplicationById(context.Context, *ApplicationIdRequest) (*ApplicationObjectReply, error)
	mustEmbedUnimplementedApplicationsServer()
}

// UnimplementedApplicationsServer must be embedded to have forward compatible implementations.
type UnimplementedApplicationsServer struct {
}

func (UnimplementedApplicationsServer) GetApplications(context.Context, *ApplicationPageRequest) (*ApplicationPageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApplications not implemented")
}
func (UnimplementedApplicationsServer) GetApplicationById(context.Context, *ApplicationIdRequest) (*ApplicationObjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApplicationById not implemented")
}
func (UnimplementedApplicationsServer) CreateApplication(context.Context, *ApplicationObjectRequest) (*ApplicationObjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApplication not implemented")
}
func (UnimplementedApplicationsServer) UpdateApplication(context.Context, *ApplicationObjectRequest) (*ApplicationObjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateApplication not implemented")
}
func (UnimplementedApplicationsServer) DeleteApplicationById(context.Context, *ApplicationIdRequest) (*ApplicationObjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteApplicationById not implemented")
}
func (UnimplementedApplicationsServer) mustEmbedUnimplementedApplicationsServer() {}

// UnsafeApplicationsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApplicationsServer will
// result in compilation errors.
type UnsafeApplicationsServer interface {
	mustEmbedUnimplementedApplicationsServer()
}

func RegisterApplicationsServer(s grpc.ServiceRegistrar, srv ApplicationsServer) {
	s.RegisterService(&Applications_ServiceDesc, srv)
}

func _Applications_GetApplications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationPageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationsServer).GetApplications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/applications_v1.Applications/get_applications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationsServer).GetApplications(ctx, req.(*ApplicationPageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Applications_GetApplicationById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationsServer).GetApplicationById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/applications_v1.Applications/get_application_by_id",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationsServer).GetApplicationById(ctx, req.(*ApplicationIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Applications_CreateApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationsServer).CreateApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/applications_v1.Applications/create_application",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationsServer).CreateApplication(ctx, req.(*ApplicationObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Applications_UpdateApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationsServer).UpdateApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/applications_v1.Applications/update_application",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationsServer).UpdateApplication(ctx, req.(*ApplicationObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Applications_DeleteApplicationById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationsServer).DeleteApplicationById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/applications_v1.Applications/delete_application_by_id",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationsServer).DeleteApplicationById(ctx, req.(*ApplicationIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Applications_ServiceDesc is the grpc.ServiceDesc for Applications service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Applications_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "applications_v1.Applications",
	HandlerType: (*ApplicationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get_applications",
			Handler:    _Applications_GetApplications_Handler,
		},
		{
			MethodName: "get_application_by_id",
			Handler:    _Applications_GetApplicationById_Handler,
		},
		{
			MethodName: "create_application",
			Handler:    _Applications_CreateApplication_Handler,
		},
		{
			MethodName: "update_application",
			Handler:    _Applications_UpdateApplication_Handler,
		},
		{
			MethodName: "delete_application_by_id",
			Handler:    _Applications_DeleteApplicationById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/applications_v1.proto",
}