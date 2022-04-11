// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: api/proto/resourcespb/route_table_association.proto

package resourcespb

import (
	commonpb "github.com/multycloud/multy/api/proto/commonpb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateRouteTableAssociationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource *RouteTableAssociationArgs `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *CreateRouteTableAssociationRequest) Reset() {
	*x = CreateRouteTableAssociationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resourcespb_route_table_association_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRouteTableAssociationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRouteTableAssociationRequest) ProtoMessage() {}

func (x *CreateRouteTableAssociationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resourcespb_route_table_association_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRouteTableAssociationRequest.ProtoReflect.Descriptor instead.
func (*CreateRouteTableAssociationRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_route_table_association_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRouteTableAssociationRequest) GetResource() *RouteTableAssociationArgs {
	if x != nil {
		return x.Resource
	}
	return nil
}

type ReadRouteTableAssociationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *ReadRouteTableAssociationRequest) Reset() {
	*x = ReadRouteTableAssociationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resourcespb_route_table_association_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRouteTableAssociationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRouteTableAssociationRequest) ProtoMessage() {}

func (x *ReadRouteTableAssociationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resourcespb_route_table_association_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRouteTableAssociationRequest.ProtoReflect.Descriptor instead.
func (*ReadRouteTableAssociationRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_route_table_association_proto_rawDescGZIP(), []int{1}
}

func (x *ReadRouteTableAssociationRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type UpdateRouteTableAssociationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId string                     `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	Resource   *RouteTableAssociationArgs `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *UpdateRouteTableAssociationRequest) Reset() {
	*x = UpdateRouteTableAssociationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resourcespb_route_table_association_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRouteTableAssociationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRouteTableAssociationRequest) ProtoMessage() {}

func (x *UpdateRouteTableAssociationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resourcespb_route_table_association_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRouteTableAssociationRequest.ProtoReflect.Descriptor instead.
func (*UpdateRouteTableAssociationRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_route_table_association_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRouteTableAssociationRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *UpdateRouteTableAssociationRequest) GetResource() *RouteTableAssociationArgs {
	if x != nil {
		return x.Resource
	}
	return nil
}

type DeleteRouteTableAssociationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *DeleteRouteTableAssociationRequest) Reset() {
	*x = DeleteRouteTableAssociationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resourcespb_route_table_association_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRouteTableAssociationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRouteTableAssociationRequest) ProtoMessage() {}

func (x *DeleteRouteTableAssociationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resourcespb_route_table_association_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRouteTableAssociationRequest.ProtoReflect.Descriptor instead.
func (*DeleteRouteTableAssociationRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_route_table_association_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteRouteTableAssociationRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type RouteTableAssociationArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonParameters *commonpb.ChildResourceCommonArgs `protobuf:"bytes,1,opt,name=common_parameters,json=commonParameters,proto3" json:"common_parameters,omitempty"`
	SubnetId         string                            `protobuf:"bytes,2,opt,name=subnet_id,json=subnetId,proto3" json:"subnet_id,omitempty"`
	RouteTableId     string                            `protobuf:"bytes,3,opt,name=route_table_id,json=routeTableId,proto3" json:"route_table_id,omitempty"`
}

func (x *RouteTableAssociationArgs) Reset() {
	*x = RouteTableAssociationArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resourcespb_route_table_association_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteTableAssociationArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteTableAssociationArgs) ProtoMessage() {}

func (x *RouteTableAssociationArgs) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resourcespb_route_table_association_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteTableAssociationArgs.ProtoReflect.Descriptor instead.
func (*RouteTableAssociationArgs) Descriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_route_table_association_proto_rawDescGZIP(), []int{4}
}

func (x *RouteTableAssociationArgs) GetCommonParameters() *commonpb.ChildResourceCommonArgs {
	if x != nil {
		return x.CommonParameters
	}
	return nil
}

func (x *RouteTableAssociationArgs) GetSubnetId() string {
	if x != nil {
		return x.SubnetId
	}
	return ""
}

func (x *RouteTableAssociationArgs) GetRouteTableId() string {
	if x != nil {
		return x.RouteTableId
	}
	return ""
}

type RouteTableAssociationResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonParameters *commonpb.CommonChildResourceParameters `protobuf:"bytes,1,opt,name=common_parameters,json=commonParameters,proto3" json:"common_parameters,omitempty"`
	SubnetId         string                                  `protobuf:"bytes,2,opt,name=subnet_id,json=subnetId,proto3" json:"subnet_id,omitempty"`
	RouteTableId     string                                  `protobuf:"bytes,3,opt,name=route_table_id,json=routeTableId,proto3" json:"route_table_id,omitempty"`
}

func (x *RouteTableAssociationResource) Reset() {
	*x = RouteTableAssociationResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resourcespb_route_table_association_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteTableAssociationResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteTableAssociationResource) ProtoMessage() {}

func (x *RouteTableAssociationResource) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resourcespb_route_table_association_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteTableAssociationResource.ProtoReflect.Descriptor instead.
func (*RouteTableAssociationResource) Descriptor() ([]byte, []int) {
	return file_api_proto_resourcespb_route_table_association_proto_rawDescGZIP(), []int{5}
}

func (x *RouteTableAssociationResource) GetCommonParameters() *commonpb.CommonChildResourceParameters {
	if x != nil {
		return x.CommonParameters
	}
	return nil
}

func (x *RouteTableAssociationResource) GetSubnetId() string {
	if x != nil {
		return x.SubnetId
	}
	return ""
}

func (x *RouteTableAssociationResource) GetRouteTableId() string {
	if x != nil {
		return x.RouteTableId
	}
	return ""
}

var File_api_proto_resourcespb_route_table_association_proto protoreflect.FileDescriptor

var file_api_proto_resourcespb_route_table_association_proto_rawDesc = []byte{
	0x0a, 0x33, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x70, 0x62, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x5f, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x1f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x70, 0x0a, 0x22, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x41,
	0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x4a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x72, 0x67, 0x73, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x43, 0x0a,
	0x20, 0x52, 0x65, 0x61, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x41,
	0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x49, 0x64, 0x22, 0x91, 0x01, 0x0a, 0x22, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x4a, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x64,
	0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x73, 0x73,
	0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x72, 0x67, 0x73, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x45, 0x0a, 0x22, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0xb6, 0x01,
	0x0a, 0x19, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x73, 0x73, 0x6f,
	0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x72, 0x67, 0x73, 0x12, 0x56, 0x0a, 0x11, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c,
	0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x41, 0x72, 0x67,
	0x73, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x49, 0x64,
	0x12, 0x24, 0x0a, 0x0e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x22, 0xc0, 0x01, 0x0a, 0x1d, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x43, 0x68, 0x69,
	0x6c, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x73, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x62, 0x6e, 0x65,
	0x74, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x42, 0x5e, 0x0a, 0x17, 0x64, 0x65, 0x76,
	0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x42, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x75,
	0x6c, 0x74, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_proto_resourcespb_route_table_association_proto_rawDescOnce sync.Once
	file_api_proto_resourcespb_route_table_association_proto_rawDescData = file_api_proto_resourcespb_route_table_association_proto_rawDesc
)

func file_api_proto_resourcespb_route_table_association_proto_rawDescGZIP() []byte {
	file_api_proto_resourcespb_route_table_association_proto_rawDescOnce.Do(func() {
		file_api_proto_resourcespb_route_table_association_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_resourcespb_route_table_association_proto_rawDescData)
	})
	return file_api_proto_resourcespb_route_table_association_proto_rawDescData
}

var file_api_proto_resourcespb_route_table_association_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_proto_resourcespb_route_table_association_proto_goTypes = []interface{}{
	(*CreateRouteTableAssociationRequest)(nil),     // 0: dev.multy.resources.CreateRouteTableAssociationRequest
	(*ReadRouteTableAssociationRequest)(nil),       // 1: dev.multy.resources.ReadRouteTableAssociationRequest
	(*UpdateRouteTableAssociationRequest)(nil),     // 2: dev.multy.resources.UpdateRouteTableAssociationRequest
	(*DeleteRouteTableAssociationRequest)(nil),     // 3: dev.multy.resources.DeleteRouteTableAssociationRequest
	(*RouteTableAssociationArgs)(nil),              // 4: dev.multy.resources.RouteTableAssociationArgs
	(*RouteTableAssociationResource)(nil),          // 5: dev.multy.resources.RouteTableAssociationResource
	(*commonpb.ChildResourceCommonArgs)(nil),       // 6: dev.multy.common.ChildResourceCommonArgs
	(*commonpb.CommonChildResourceParameters)(nil), // 7: dev.multy.common.CommonChildResourceParameters
}
var file_api_proto_resourcespb_route_table_association_proto_depIdxs = []int32{
	4, // 0: dev.multy.resources.CreateRouteTableAssociationRequest.resource:type_name -> dev.multy.resources.RouteTableAssociationArgs
	4, // 1: dev.multy.resources.UpdateRouteTableAssociationRequest.resource:type_name -> dev.multy.resources.RouteTableAssociationArgs
	6, // 2: dev.multy.resources.RouteTableAssociationArgs.common_parameters:type_name -> dev.multy.common.ChildResourceCommonArgs
	7, // 3: dev.multy.resources.RouteTableAssociationResource.common_parameters:type_name -> dev.multy.common.CommonChildResourceParameters
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_proto_resourcespb_route_table_association_proto_init() }
func file_api_proto_resourcespb_route_table_association_proto_init() {
	if File_api_proto_resourcespb_route_table_association_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_resourcespb_route_table_association_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRouteTableAssociationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_resourcespb_route_table_association_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRouteTableAssociationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_resourcespb_route_table_association_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRouteTableAssociationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_resourcespb_route_table_association_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRouteTableAssociationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_resourcespb_route_table_association_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteTableAssociationArgs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_resourcespb_route_table_association_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteTableAssociationResource); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_resourcespb_route_table_association_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_resourcespb_route_table_association_proto_goTypes,
		DependencyIndexes: file_api_proto_resourcespb_route_table_association_proto_depIdxs,
		MessageInfos:      file_api_proto_resourcespb_route_table_association_proto_msgTypes,
	}.Build()
	File_api_proto_resourcespb_route_table_association_proto = out.File
	file_api_proto_resourcespb_route_table_association_proto_rawDesc = nil
	file_api_proto_resourcespb_route_table_association_proto_goTypes = nil
	file_api_proto_resourcespb_route_table_association_proto_depIdxs = nil
}
