// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	common "github.com/multycloud/multy/api/proto/common"
	resources "github.com/multycloud/multy/api/proto/resources"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MultyResourceServiceClient is the client API for MultyResourceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MultyResourceServiceClient interface {
	CreateSubnet(ctx context.Context, in *resources.CreateSubnetRequest, opts ...grpc.CallOption) (*resources.SubnetResource, error)
	ReadSubnet(ctx context.Context, in *resources.ReadSubnetRequest, opts ...grpc.CallOption) (*resources.SubnetResource, error)
	UpdateSubnet(ctx context.Context, in *resources.UpdateSubnetRequest, opts ...grpc.CallOption) (*resources.SubnetResource, error)
	DeleteSubnet(ctx context.Context, in *resources.DeleteSubnetRequest, opts ...grpc.CallOption) (*common.Empty, error)
	CreateVirtualNetwork(ctx context.Context, in *resources.CreateVirtualNetworkRequest, opts ...grpc.CallOption) (*resources.VirtualNetworkResource, error)
	ReadVirtualNetwork(ctx context.Context, in *resources.ReadVirtualNetworkRequest, opts ...grpc.CallOption) (*resources.VirtualNetworkResource, error)
	UpdateVirtualNetwork(ctx context.Context, in *resources.UpdateVirtualNetworkRequest, opts ...grpc.CallOption) (*resources.VirtualNetworkResource, error)
	DeleteVirtualNetwork(ctx context.Context, in *resources.DeleteVirtualNetworkRequest, opts ...grpc.CallOption) (*common.Empty, error)
	CreateNetworkInterface(ctx context.Context, in *resources.CreateNetworkInterfaceRequest, opts ...grpc.CallOption) (*resources.NetworkInterfaceResource, error)
	ReadNetworkInterface(ctx context.Context, in *resources.ReadNetworkInterfaceRequest, opts ...grpc.CallOption) (*resources.NetworkInterfaceResource, error)
	UpdateNetworkInterface(ctx context.Context, in *resources.UpdateNetworkInterfaceRequest, opts ...grpc.CallOption) (*resources.NetworkInterfaceResource, error)
	DeleteNetworkInterface(ctx context.Context, in *resources.DeleteNetworkInterfaceRequest, opts ...grpc.CallOption) (*common.Empty, error)
	CreateRouteTable(ctx context.Context, in *resources.CreateRouteTableRequest, opts ...grpc.CallOption) (*resources.RouteTableResource, error)
	ReadRouteTable(ctx context.Context, in *resources.ReadRouteTableRequest, opts ...grpc.CallOption) (*resources.RouteTableResource, error)
	UpdateRouteTable(ctx context.Context, in *resources.UpdateRouteTableRequest, opts ...grpc.CallOption) (*resources.RouteTableResource, error)
	DeleteRouteTable(ctx context.Context, in *resources.DeleteRouteTableRequest, opts ...grpc.CallOption) (*common.Empty, error)
}

type multyResourceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMultyResourceServiceClient(cc grpc.ClientConnInterface) MultyResourceServiceClient {
	return &multyResourceServiceClient{cc}
}

func (c *multyResourceServiceClient) CreateSubnet(ctx context.Context, in *resources.CreateSubnetRequest, opts ...grpc.CallOption) (*resources.SubnetResource, error) {
	out := new(resources.SubnetResource)
	err := c.cc.Invoke(ctx, "/dev.multy.MultyResourceService/CreateSubnet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *multyResourceServiceClient) ReadSubnet(ctx context.Context, in *resources.ReadSubnetRequest, opts ...grpc.CallOption) (*resources.SubnetResource, error) {
	out := new(resources.SubnetResource)
	err := c.cc.Invoke(ctx, "/dev.multy.MultyResourceService/ReadSubnet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *multyResourceServiceClient) UpdateSubnet(ctx context.Context, in *resources.UpdateSubnetRequest, opts ...grpc.CallOption) (*resources.SubnetResource, error) {
	out := new(resources.SubnetResource)
	err := c.cc.Invoke(ctx, "/dev.multy.MultyResourceService/UpdateSubnet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *multyResourceServiceClient) DeleteSubnet(ctx context.Context, in *resources.DeleteSubnetRequest, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/dev.multy.MultyResourceService/DeleteSubnet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *multyResourceServiceClient) CreateVirtualNetwork(ctx context.Context, in *resources.CreateVirtualNetworkRequest, opts ...grpc.CallOption) (*resources.VirtualNetworkResource, error) {
	out := new(resources.VirtualNetworkResource)
	err := c.cc.Invoke(ctx, "/dev.multy.MultyResourceService/CreateVirtualNetwork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *multyResourceServiceClient) ReadVirtualNetwork(ctx context.Context, in *resources.ReadVirtualNetworkRequest, opts ...grpc.CallOption) (*resources.VirtualNetworkResource, error) {
	out := new(resources.VirtualNetworkResource)
	err := c.cc.Invoke(ctx, "/dev.multy.MultyResourceService/ReadVirtualNetwork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *multyResourceServiceClient) UpdateVirtualNetwork(ctx context.Context, in *resources.UpdateVirtualNetworkRequest, opts ...grpc.CallOption) (*resources.VirtualNetworkResource, error) {
	out := new(resources.VirtualNetworkResource)
	err := c.cc.Invoke(ctx, "/dev.multy.MultyResourceService/UpdateVirtualNetwork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *multyResourceServiceClient) DeleteVirtualNetwork(ctx context.Context, in *resources.DeleteVirtualNetworkRequest, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/dev.multy.MultyResourceService/DeleteVirtualNetwork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *multyResourceServiceClient) CreateNetworkInterface(ctx context.Context, in *resources.CreateNetworkInterfaceRequest, opts ...grpc.CallOption) (*resources.NetworkInterfaceResource, error) {
	out := new(resources.NetworkInterfaceResource)
	err := c.cc.Invoke(ctx, "/dev.multy.MultyResourceService/CreateNetworkInterface", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *multyResourceServiceClient) ReadNetworkInterface(ctx context.Context, in *resources.ReadNetworkInterfaceRequest, opts ...grpc.CallOption) (*resources.NetworkInterfaceResource, error) {
	out := new(resources.NetworkInterfaceResource)
	err := c.cc.Invoke(ctx, "/dev.multy.MultyResourceService/ReadNetworkInterface", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *multyResourceServiceClient) UpdateNetworkInterface(ctx context.Context, in *resources.UpdateNetworkInterfaceRequest, opts ...grpc.CallOption) (*resources.NetworkInterfaceResource, error) {
	out := new(resources.NetworkInterfaceResource)
	err := c.cc.Invoke(ctx, "/dev.multy.MultyResourceService/UpdateNetworkInterface", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *multyResourceServiceClient) DeleteNetworkInterface(ctx context.Context, in *resources.DeleteNetworkInterfaceRequest, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/dev.multy.MultyResourceService/DeleteNetworkInterface", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *multyResourceServiceClient) CreateRouteTable(ctx context.Context, in *resources.CreateRouteTableRequest, opts ...grpc.CallOption) (*resources.RouteTableResource, error) {
	out := new(resources.RouteTableResource)
	err := c.cc.Invoke(ctx, "/dev.multy.MultyResourceService/CreateRouteTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *multyResourceServiceClient) ReadRouteTable(ctx context.Context, in *resources.ReadRouteTableRequest, opts ...grpc.CallOption) (*resources.RouteTableResource, error) {
	out := new(resources.RouteTableResource)
	err := c.cc.Invoke(ctx, "/dev.multy.MultyResourceService/ReadRouteTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *multyResourceServiceClient) UpdateRouteTable(ctx context.Context, in *resources.UpdateRouteTableRequest, opts ...grpc.CallOption) (*resources.RouteTableResource, error) {
	out := new(resources.RouteTableResource)
	err := c.cc.Invoke(ctx, "/dev.multy.MultyResourceService/UpdateRouteTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *multyResourceServiceClient) DeleteRouteTable(ctx context.Context, in *resources.DeleteRouteTableRequest, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/dev.multy.MultyResourceService/DeleteRouteTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MultyResourceServiceServer is the server API for MultyResourceService service.
// All implementations must embed UnimplementedMultyResourceServiceServer
// for forward compatibility
type MultyResourceServiceServer interface {
	CreateSubnet(context.Context, *resources.CreateSubnetRequest) (*resources.SubnetResource, error)
	ReadSubnet(context.Context, *resources.ReadSubnetRequest) (*resources.SubnetResource, error)
	UpdateSubnet(context.Context, *resources.UpdateSubnetRequest) (*resources.SubnetResource, error)
	DeleteSubnet(context.Context, *resources.DeleteSubnetRequest) (*common.Empty, error)
	CreateVirtualNetwork(context.Context, *resources.CreateVirtualNetworkRequest) (*resources.VirtualNetworkResource, error)
	ReadVirtualNetwork(context.Context, *resources.ReadVirtualNetworkRequest) (*resources.VirtualNetworkResource, error)
	UpdateVirtualNetwork(context.Context, *resources.UpdateVirtualNetworkRequest) (*resources.VirtualNetworkResource, error)
	DeleteVirtualNetwork(context.Context, *resources.DeleteVirtualNetworkRequest) (*common.Empty, error)
	CreateNetworkInterface(context.Context, *resources.CreateNetworkInterfaceRequest) (*resources.NetworkInterfaceResource, error)
	ReadNetworkInterface(context.Context, *resources.ReadNetworkInterfaceRequest) (*resources.NetworkInterfaceResource, error)
	UpdateNetworkInterface(context.Context, *resources.UpdateNetworkInterfaceRequest) (*resources.NetworkInterfaceResource, error)
	DeleteNetworkInterface(context.Context, *resources.DeleteNetworkInterfaceRequest) (*common.Empty, error)
	CreateRouteTable(context.Context, *resources.CreateRouteTableRequest) (*resources.RouteTableResource, error)
	ReadRouteTable(context.Context, *resources.ReadRouteTableRequest) (*resources.RouteTableResource, error)
	UpdateRouteTable(context.Context, *resources.UpdateRouteTableRequest) (*resources.RouteTableResource, error)
	DeleteRouteTable(context.Context, *resources.DeleteRouteTableRequest) (*common.Empty, error)
	mustEmbedUnimplementedMultyResourceServiceServer()
}

// UnimplementedMultyResourceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMultyResourceServiceServer struct {
}

func (UnimplementedMultyResourceServiceServer) CreateSubnet(context.Context, *resources.CreateSubnetRequest) (*resources.SubnetResource, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSubnet not implemented")
}
func (UnimplementedMultyResourceServiceServer) ReadSubnet(context.Context, *resources.ReadSubnetRequest) (*resources.SubnetResource, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadSubnet not implemented")
}
func (UnimplementedMultyResourceServiceServer) UpdateSubnet(context.Context, *resources.UpdateSubnetRequest) (*resources.SubnetResource, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSubnet not implemented")
}
func (UnimplementedMultyResourceServiceServer) DeleteSubnet(context.Context, *resources.DeleteSubnetRequest) (*common.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSubnet not implemented")
}
func (UnimplementedMultyResourceServiceServer) CreateVirtualNetwork(context.Context, *resources.CreateVirtualNetworkRequest) (*resources.VirtualNetworkResource, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVirtualNetwork not implemented")
}
func (UnimplementedMultyResourceServiceServer) ReadVirtualNetwork(context.Context, *resources.ReadVirtualNetworkRequest) (*resources.VirtualNetworkResource, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadVirtualNetwork not implemented")
}
func (UnimplementedMultyResourceServiceServer) UpdateVirtualNetwork(context.Context, *resources.UpdateVirtualNetworkRequest) (*resources.VirtualNetworkResource, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVirtualNetwork not implemented")
}
func (UnimplementedMultyResourceServiceServer) DeleteVirtualNetwork(context.Context, *resources.DeleteVirtualNetworkRequest) (*common.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVirtualNetwork not implemented")
}
func (UnimplementedMultyResourceServiceServer) CreateNetworkInterface(context.Context, *resources.CreateNetworkInterfaceRequest) (*resources.NetworkInterfaceResource, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNetworkInterface not implemented")
}
func (UnimplementedMultyResourceServiceServer) ReadNetworkInterface(context.Context, *resources.ReadNetworkInterfaceRequest) (*resources.NetworkInterfaceResource, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadNetworkInterface not implemented")
}
func (UnimplementedMultyResourceServiceServer) UpdateNetworkInterface(context.Context, *resources.UpdateNetworkInterfaceRequest) (*resources.NetworkInterfaceResource, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNetworkInterface not implemented")
}
func (UnimplementedMultyResourceServiceServer) DeleteNetworkInterface(context.Context, *resources.DeleteNetworkInterfaceRequest) (*common.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNetworkInterface not implemented")
}
func (UnimplementedMultyResourceServiceServer) CreateRouteTable(context.Context, *resources.CreateRouteTableRequest) (*resources.RouteTableResource, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRouteTable not implemented")
}
func (UnimplementedMultyResourceServiceServer) ReadRouteTable(context.Context, *resources.ReadRouteTableRequest) (*resources.RouteTableResource, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadRouteTable not implemented")
}
func (UnimplementedMultyResourceServiceServer) UpdateRouteTable(context.Context, *resources.UpdateRouteTableRequest) (*resources.RouteTableResource, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRouteTable not implemented")
}
func (UnimplementedMultyResourceServiceServer) DeleteRouteTable(context.Context, *resources.DeleteRouteTableRequest) (*common.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRouteTable not implemented")
}
func (UnimplementedMultyResourceServiceServer) mustEmbedUnimplementedMultyResourceServiceServer() {}

// UnsafeMultyResourceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MultyResourceServiceServer will
// result in compilation errors.
type UnsafeMultyResourceServiceServer interface {
	mustEmbedUnimplementedMultyResourceServiceServer()
}

func RegisterMultyResourceServiceServer(s grpc.ServiceRegistrar, srv MultyResourceServiceServer) {
	s.RegisterService(&MultyResourceService_ServiceDesc, srv)
}

func _MultyResourceService_CreateSubnet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.CreateSubnetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultyResourceServiceServer).CreateSubnet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dev.multy.MultyResourceService/CreateSubnet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultyResourceServiceServer).CreateSubnet(ctx, req.(*resources.CreateSubnetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MultyResourceService_ReadSubnet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.ReadSubnetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultyResourceServiceServer).ReadSubnet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dev.multy.MultyResourceService/ReadSubnet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultyResourceServiceServer).ReadSubnet(ctx, req.(*resources.ReadSubnetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MultyResourceService_UpdateSubnet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.UpdateSubnetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultyResourceServiceServer).UpdateSubnet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dev.multy.MultyResourceService/UpdateSubnet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultyResourceServiceServer).UpdateSubnet(ctx, req.(*resources.UpdateSubnetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MultyResourceService_DeleteSubnet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.DeleteSubnetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultyResourceServiceServer).DeleteSubnet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dev.multy.MultyResourceService/DeleteSubnet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultyResourceServiceServer).DeleteSubnet(ctx, req.(*resources.DeleteSubnetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MultyResourceService_CreateVirtualNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.CreateVirtualNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultyResourceServiceServer).CreateVirtualNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dev.multy.MultyResourceService/CreateVirtualNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultyResourceServiceServer).CreateVirtualNetwork(ctx, req.(*resources.CreateVirtualNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MultyResourceService_ReadVirtualNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.ReadVirtualNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultyResourceServiceServer).ReadVirtualNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dev.multy.MultyResourceService/ReadVirtualNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultyResourceServiceServer).ReadVirtualNetwork(ctx, req.(*resources.ReadVirtualNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MultyResourceService_UpdateVirtualNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.UpdateVirtualNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultyResourceServiceServer).UpdateVirtualNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dev.multy.MultyResourceService/UpdateVirtualNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultyResourceServiceServer).UpdateVirtualNetwork(ctx, req.(*resources.UpdateVirtualNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MultyResourceService_DeleteVirtualNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.DeleteVirtualNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultyResourceServiceServer).DeleteVirtualNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dev.multy.MultyResourceService/DeleteVirtualNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultyResourceServiceServer).DeleteVirtualNetwork(ctx, req.(*resources.DeleteVirtualNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MultyResourceService_CreateNetworkInterface_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.CreateNetworkInterfaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultyResourceServiceServer).CreateNetworkInterface(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dev.multy.MultyResourceService/CreateNetworkInterface",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultyResourceServiceServer).CreateNetworkInterface(ctx, req.(*resources.CreateNetworkInterfaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MultyResourceService_ReadNetworkInterface_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.ReadNetworkInterfaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultyResourceServiceServer).ReadNetworkInterface(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dev.multy.MultyResourceService/ReadNetworkInterface",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultyResourceServiceServer).ReadNetworkInterface(ctx, req.(*resources.ReadNetworkInterfaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MultyResourceService_UpdateNetworkInterface_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.UpdateNetworkInterfaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultyResourceServiceServer).UpdateNetworkInterface(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dev.multy.MultyResourceService/UpdateNetworkInterface",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultyResourceServiceServer).UpdateNetworkInterface(ctx, req.(*resources.UpdateNetworkInterfaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MultyResourceService_DeleteNetworkInterface_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.DeleteNetworkInterfaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultyResourceServiceServer).DeleteNetworkInterface(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dev.multy.MultyResourceService/DeleteNetworkInterface",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultyResourceServiceServer).DeleteNetworkInterface(ctx, req.(*resources.DeleteNetworkInterfaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MultyResourceService_CreateRouteTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.CreateRouteTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultyResourceServiceServer).CreateRouteTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dev.multy.MultyResourceService/CreateRouteTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultyResourceServiceServer).CreateRouteTable(ctx, req.(*resources.CreateRouteTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MultyResourceService_ReadRouteTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.ReadRouteTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultyResourceServiceServer).ReadRouteTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dev.multy.MultyResourceService/ReadRouteTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultyResourceServiceServer).ReadRouteTable(ctx, req.(*resources.ReadRouteTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MultyResourceService_UpdateRouteTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.UpdateRouteTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultyResourceServiceServer).UpdateRouteTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dev.multy.MultyResourceService/UpdateRouteTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultyResourceServiceServer).UpdateRouteTable(ctx, req.(*resources.UpdateRouteTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MultyResourceService_DeleteRouteTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.DeleteRouteTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultyResourceServiceServer).DeleteRouteTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dev.multy.MultyResourceService/DeleteRouteTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultyResourceServiceServer).DeleteRouteTable(ctx, req.(*resources.DeleteRouteTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MultyResourceService_ServiceDesc is the grpc.ServiceDesc for MultyResourceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MultyResourceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dev.multy.MultyResourceService",
	HandlerType: (*MultyResourceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSubnet",
			Handler:    _MultyResourceService_CreateSubnet_Handler,
		},
		{
			MethodName: "ReadSubnet",
			Handler:    _MultyResourceService_ReadSubnet_Handler,
		},
		{
			MethodName: "UpdateSubnet",
			Handler:    _MultyResourceService_UpdateSubnet_Handler,
		},
		{
			MethodName: "DeleteSubnet",
			Handler:    _MultyResourceService_DeleteSubnet_Handler,
		},
		{
			MethodName: "CreateVirtualNetwork",
			Handler:    _MultyResourceService_CreateVirtualNetwork_Handler,
		},
		{
			MethodName: "ReadVirtualNetwork",
			Handler:    _MultyResourceService_ReadVirtualNetwork_Handler,
		},
		{
			MethodName: "UpdateVirtualNetwork",
			Handler:    _MultyResourceService_UpdateVirtualNetwork_Handler,
		},
		{
			MethodName: "DeleteVirtualNetwork",
			Handler:    _MultyResourceService_DeleteVirtualNetwork_Handler,
		},
		{
			MethodName: "CreateNetworkInterface",
			Handler:    _MultyResourceService_CreateNetworkInterface_Handler,
		},
		{
			MethodName: "ReadNetworkInterface",
			Handler:    _MultyResourceService_ReadNetworkInterface_Handler,
		},
		{
			MethodName: "UpdateNetworkInterface",
			Handler:    _MultyResourceService_UpdateNetworkInterface_Handler,
		},
		{
			MethodName: "DeleteNetworkInterface",
			Handler:    _MultyResourceService_DeleteNetworkInterface_Handler,
		},
		{
			MethodName: "CreateRouteTable",
			Handler:    _MultyResourceService_CreateRouteTable_Handler,
		},
		{
			MethodName: "ReadRouteTable",
			Handler:    _MultyResourceService_ReadRouteTable_Handler,
		},
		{
			MethodName: "UpdateRouteTable",
			Handler:    _MultyResourceService_UpdateRouteTable_Handler,
		},
		{
			MethodName: "DeleteRouteTable",
			Handler:    _MultyResourceService_DeleteRouteTable_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/multy_service.proto",
}
