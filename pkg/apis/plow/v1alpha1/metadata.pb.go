// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metadata.proto

package v1alpha1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RegionListRequest struct {
	CloudCredential string `protobuf:"bytes,1,opt,name=cloud_credential,json=cloudCredential" json:"cloud_credential,omitempty"`
}

func (m *RegionListRequest) Reset()                    { *m = RegionListRequest{} }
func (m *RegionListRequest) String() string            { return proto.CompactTextString(m) }
func (*RegionListRequest) ProtoMessage()               {}
func (*RegionListRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *RegionListRequest) GetCloudCredential() string {
	if m != nil {
		return m.CloudCredential
	}
	return ""
}

type RegionListResponse struct {
	Regions []string `protobuf:"bytes,1,rep,name=regions" json:"regions,omitempty"`
}

func (m *RegionListResponse) Reset()                    { *m = RegionListResponse{} }
func (m *RegionListResponse) String() string            { return proto.CompactTextString(m) }
func (*RegionListResponse) ProtoMessage()               {}
func (*RegionListResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *RegionListResponse) GetRegions() []string {
	if m != nil {
		return m.Regions
	}
	return nil
}

type ZoneListRequest struct {
	CloudCredential string `protobuf:"bytes,1,opt,name=cloud_credential,json=cloudCredential" json:"cloud_credential,omitempty"`
	Region          string `protobuf:"bytes,2,opt,name=region" json:"region,omitempty"`
}

func (m *ZoneListRequest) Reset()                    { *m = ZoneListRequest{} }
func (m *ZoneListRequest) String() string            { return proto.CompactTextString(m) }
func (*ZoneListRequest) ProtoMessage()               {}
func (*ZoneListRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *ZoneListRequest) GetCloudCredential() string {
	if m != nil {
		return m.CloudCredential
	}
	return ""
}

func (m *ZoneListRequest) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

type ZoneListResponse struct {
	Zones []string `protobuf:"bytes,1,rep,name=zones" json:"zones,omitempty"`
}

func (m *ZoneListResponse) Reset()                    { *m = ZoneListResponse{} }
func (m *ZoneListResponse) String() string            { return proto.CompactTextString(m) }
func (*ZoneListResponse) ProtoMessage()               {}
func (*ZoneListResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *ZoneListResponse) GetZones() []string {
	if m != nil {
		return m.Zones
	}
	return nil
}

type BucketListRequest struct {
	CloudCredential string `protobuf:"bytes,1,opt,name=cloud_credential,json=cloudCredential" json:"cloud_credential,omitempty"`
	GceProject      string `protobuf:"bytes,2,opt,name=gce_project,json=gceProject" json:"gce_project,omitempty"`
	ClusterUid      string `protobuf:"bytes,3,opt,name=cluster_uid,json=clusterUid" json:"cluster_uid,omitempty"`
	SecretNamespace string `protobuf:"bytes,4,opt,name=secret_namespace,json=secretNamespace" json:"secret_namespace,omitempty"`
	SecretName      string `protobuf:"bytes,5,opt,name=secret_name,json=secretName" json:"secret_name,omitempty"`
	Provider        string `protobuf:"bytes,6,opt,name=provider" json:"provider,omitempty"`
}

func (m *BucketListRequest) Reset()                    { *m = BucketListRequest{} }
func (m *BucketListRequest) String() string            { return proto.CompactTextString(m) }
func (*BucketListRequest) ProtoMessage()               {}
func (*BucketListRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *BucketListRequest) GetCloudCredential() string {
	if m != nil {
		return m.CloudCredential
	}
	return ""
}

func (m *BucketListRequest) GetGceProject() string {
	if m != nil {
		return m.GceProject
	}
	return ""
}

func (m *BucketListRequest) GetClusterUid() string {
	if m != nil {
		return m.ClusterUid
	}
	return ""
}

func (m *BucketListRequest) GetSecretNamespace() string {
	if m != nil {
		return m.SecretNamespace
	}
	return ""
}

func (m *BucketListRequest) GetSecretName() string {
	if m != nil {
		return m.SecretName
	}
	return ""
}

func (m *BucketListRequest) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

type BucketListResponse struct {
	Names []string `protobuf:"bytes,1,rep,name=names" json:"names,omitempty"`
}

func (m *BucketListResponse) Reset()                    { *m = BucketListResponse{} }
func (m *BucketListResponse) String() string            { return proto.CompactTextString(m) }
func (*BucketListResponse) ProtoMessage()               {}
func (*BucketListResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *BucketListResponse) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func init() {
	proto.RegisterType((*RegionListRequest)(nil), "appscode.plow.v1alpha1.RegionListRequest")
	proto.RegisterType((*RegionListResponse)(nil), "appscode.plow.v1alpha1.RegionListResponse")
	proto.RegisterType((*ZoneListRequest)(nil), "appscode.plow.v1alpha1.ZoneListRequest")
	proto.RegisterType((*ZoneListResponse)(nil), "appscode.plow.v1alpha1.ZoneListResponse")
	proto.RegisterType((*BucketListRequest)(nil), "appscode.plow.v1alpha1.BucketListRequest")
	proto.RegisterType((*BucketListResponse)(nil), "appscode.plow.v1alpha1.BucketListResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Metadata service

type MetadataClient interface {
	ListRegions(ctx context.Context, in *RegionListRequest, opts ...grpc.CallOption) (*RegionListResponse, error)
	ListZones(ctx context.Context, in *ZoneListRequest, opts ...grpc.CallOption) (*ZoneListResponse, error)
	ListBuckets(ctx context.Context, in *BucketListRequest, opts ...grpc.CallOption) (*BucketListResponse, error)
}

type metadataClient struct {
	cc *grpc.ClientConn
}

func NewMetadataClient(cc *grpc.ClientConn) MetadataClient {
	return &metadataClient{cc}
}

func (c *metadataClient) ListRegions(ctx context.Context, in *RegionListRequest, opts ...grpc.CallOption) (*RegionListResponse, error) {
	out := new(RegionListResponse)
	err := grpc.Invoke(ctx, "/appscode.plow.v1alpha1.Metadata/ListRegions", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataClient) ListZones(ctx context.Context, in *ZoneListRequest, opts ...grpc.CallOption) (*ZoneListResponse, error) {
	out := new(ZoneListResponse)
	err := grpc.Invoke(ctx, "/appscode.plow.v1alpha1.Metadata/ListZones", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataClient) ListBuckets(ctx context.Context, in *BucketListRequest, opts ...grpc.CallOption) (*BucketListResponse, error) {
	out := new(BucketListResponse)
	err := grpc.Invoke(ctx, "/appscode.plow.v1alpha1.Metadata/ListBuckets", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Metadata service

type MetadataServer interface {
	ListRegions(context.Context, *RegionListRequest) (*RegionListResponse, error)
	ListZones(context.Context, *ZoneListRequest) (*ZoneListResponse, error)
	ListBuckets(context.Context, *BucketListRequest) (*BucketListResponse, error)
}

func RegisterMetadataServer(s *grpc.Server, srv MetadataServer) {
	s.RegisterService(&_Metadata_serviceDesc, srv)
}

func _Metadata_ListRegions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegionListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServer).ListRegions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.plow.v1alpha1.Metadata/ListRegions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServer).ListRegions(ctx, req.(*RegionListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metadata_ListZones_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ZoneListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServer).ListZones(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.plow.v1alpha1.Metadata/ListZones",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServer).ListZones(ctx, req.(*ZoneListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metadata_ListBuckets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BucketListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServer).ListBuckets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.plow.v1alpha1.Metadata/ListBuckets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServer).ListBuckets(ctx, req.(*BucketListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Metadata_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.plow.v1alpha1.Metadata",
	HandlerType: (*MetadataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRegions",
			Handler:    _Metadata_ListRegions_Handler,
		},
		{
			MethodName: "ListZones",
			Handler:    _Metadata_ListZones_Handler,
		},
		{
			MethodName: "ListBuckets",
			Handler:    _Metadata_ListBuckets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metadata.proto",
}

func init() { proto.RegisterFile("metadata.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0x13, 0x1a, 0x92, 0xa9, 0xa0, 0xed, 0x0a, 0x55, 0x96, 0x85, 0x4a, 0xe5, 0x0b, 0x69,
	0x84, 0x6c, 0x15, 0x2e, 0xa0, 0x4a, 0x1c, 0xc2, 0x81, 0x0b, 0xa0, 0x28, 0x82, 0x4b, 0x2f, 0xd6,
	0x76, 0x3d, 0x32, 0x5b, 0x9c, 0xdd, 0xc5, 0xbb, 0x2e, 0x12, 0x88, 0x4b, 0x7f, 0x81, 0x3b, 0x57,
	0xbe, 0x80, 0x2f, 0xe1, 0x17, 0x38, 0xf0, 0x05, 0x9c, 0x91, 0x77, 0xbd, 0x4d, 0x4a, 0x5a, 0x68,
	0x73, 0xcb, 0xcc, 0xbe, 0x99, 0xf7, 0xf2, 0xe6, 0xc9, 0x70, 0x7b, 0x86, 0x86, 0xe6, 0xd4, 0xd0,
	0x44, 0x55, 0xd2, 0x48, 0xb2, 0x4d, 0x95, 0xd2, 0x4c, 0xe6, 0x98, 0xa8, 0x52, 0x7e, 0x48, 0x4e,
	0xf6, 0x69, 0xa9, 0xde, 0xd2, 0xfd, 0xe8, 0x6e, 0x21, 0x65, 0x51, 0x62, 0x4a, 0x15, 0x4f, 0xa9,
	0x10, 0xd2, 0x50, 0xc3, 0xa5, 0xd0, 0x6e, 0x2a, 0xda, 0xf1, 0x53, 0x17, 0xbf, 0xc7, 0x4f, 0x61,
	0x6b, 0x8a, 0x05, 0x97, 0xe2, 0x05, 0xd7, 0x66, 0x8a, 0xef, 0x6b, 0xd4, 0x86, 0xec, 0xc1, 0x26,
	0x2b, 0x65, 0x9d, 0x67, 0xac, 0xc2, 0x1c, 0x85, 0xe1, 0xb4, 0x0c, 0x83, 0xdd, 0x60, 0x38, 0x98,
	0x6e, 0xd8, 0xfe, 0xb3, 0xb3, 0x76, 0x9c, 0x00, 0x59, 0x9c, 0xd7, 0x4a, 0x0a, 0x8d, 0x24, 0x84,
	0x9b, 0x95, 0xed, 0xea, 0x30, 0xd8, 0xed, 0x0e, 0x07, 0x53, 0x5f, 0xc6, 0xaf, 0x61, 0xe3, 0x50,
	0x0a, 0x5c, 0x8d, 0x8d, 0x6c, 0x43, 0xcf, 0x2d, 0x0a, 0x3b, 0x16, 0xd0, 0x56, 0xf1, 0x10, 0x36,
	0xe7, 0x5b, 0x5b, 0x0d, 0x77, 0x60, 0xed, 0xa3, 0x14, 0xe8, 0x15, 0xb8, 0x22, 0xfe, 0x15, 0xc0,
	0xd6, 0xb8, 0x66, 0xef, 0xd0, 0xac, 0x28, 0xe1, 0x1e, 0xac, 0x17, 0x0c, 0x33, 0x55, 0xc9, 0x63,
	0x64, 0xa6, 0xd5, 0x01, 0x05, 0xc3, 0x89, 0xeb, 0x34, 0x00, 0x56, 0xd6, 0xda, 0x60, 0x95, 0xd5,
	0x3c, 0x0f, 0xbb, 0x0e, 0xd0, 0xb6, 0xde, 0xf0, 0xbc, 0x21, 0xd3, 0xc8, 0x2a, 0x34, 0x99, 0xa0,
	0x33, 0xd4, 0x8a, 0x32, 0x0c, 0x6f, 0x38, 0x32, 0xd7, 0x7f, 0xe5, 0xdb, 0xcd, 0xae, 0x05, 0x68,
	0xb8, 0xe6, 0x76, 0xcd, 0x51, 0x24, 0x82, 0xbe, 0xaa, 0xe4, 0x09, 0xcf, 0xb1, 0x0a, 0x7b, 0xf6,
	0xf5, 0xac, 0x8e, 0x47, 0x40, 0x16, 0xff, 0xe9, 0xdc, 0x16, 0x4b, 0xeb, 0x6d, 0xb1, 0xc5, 0xc3,
	0xdf, 0x5d, 0xe8, 0xbf, 0x6c, 0xf3, 0x46, 0xbe, 0x06, 0xb0, 0xee, 0x66, 0xec, 0xcd, 0xc8, 0x5e,
	0x72, 0x71, 0xf4, 0x92, 0xa5, 0xe4, 0x44, 0xa3, 0xab, 0x40, 0x9d, 0x92, 0xf8, 0xf1, 0xe9, 0xf7,
	0xb0, 0xd3, 0x0f, 0x4e, 0x7f, 0xfc, 0xfc, 0xd2, 0x79, 0x40, 0x46, 0x69, 0x76, 0x2e, 0xaa, 0xd6,
	0xf7, 0xd4, 0xcf, 0xa7, 0x6d, 0x7e, 0xd2, 0x63, 0x2d, 0x05, 0xf9, 0x16, 0xc0, 0xa0, 0x59, 0xd5,
	0xdc, 0x5c, 0x93, 0xfb, 0x97, 0x71, 0xfe, 0x15, 0xb4, 0x68, 0xf8, 0x7f, 0x60, 0x2b, 0xed, 0xf9,
	0x82, 0xb4, 0x03, 0xf2, 0xe4, 0x6a, 0xd2, 0x3e, 0xb9, 0x1f, 0x9f, 0x53, 0x9b, 0x34, 0xa7, 0xd4,
	0x5b, 0xe9, 0x0e, 0xf1, 0x0f, 0x2b, 0x97, 0x32, 0x79, 0xb9, 0x95, 0xcb, 0x47, 0xbd, 0x9e, 0x95,
	0x47, 0x4e, 0x8b, 0x15, 0x38, 0x3e, 0x80, 0x1d, 0x26, 0x67, 0x73, 0x2a, 0xaa, 0xf8, 0x79, 0xba,
	0xf1, 0x2d, 0x9f, 0x8b, 0x49, 0xf3, 0xc1, 0x98, 0x04, 0x87, 0x7d, 0xff, 0x74, 0xd4, 0xb3, 0xdf,
	0x90, 0x47, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x0d, 0xb1, 0x7b, 0xa5, 0xab, 0x04, 0x00, 0x00,
}
