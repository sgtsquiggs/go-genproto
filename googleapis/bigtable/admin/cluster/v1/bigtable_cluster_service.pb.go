// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/bigtable/admin/cluster/v1/bigtable_cluster_service.proto

package cluster // import "github.com/sgtsquiggs/go-genproto/googleapis/bigtable/admin/cluster/v1"

import proto "github.com/sgtsquiggs/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/sgtsquiggs/protobuf/ptypes/empty"
import _ "github.com/sgtsquiggs/go-genproto/googleapis/api/annotations"
import longrunning "github.com/sgtsquiggs/go-genproto/googleapis/longrunning"

import (
	context "golang.org/x/net/context"
	grpc "github.com/sgtsquiggs/grpc-go"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BigtableClusterServiceClient is the client API for BigtableClusterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/github.com/sgtsquiggs/grpc-go#ClientConn.NewStream.
type BigtableClusterServiceClient interface {
	// Lists the supported zones for the given project.
	ListZones(ctx context.Context, in *ListZonesRequest, opts ...grpc.CallOption) (*ListZonesResponse, error)
	// Gets information about a particular cluster.
	GetCluster(ctx context.Context, in *GetClusterRequest, opts ...grpc.CallOption) (*Cluster, error)
	// Lists all clusters in the given project, along with any zones for which
	// cluster information could not be retrieved.
	ListClusters(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (*ListClustersResponse, error)
	// Creates a cluster and begins preparing it to begin serving. The returned
	// cluster embeds as its "current_operation" a long-running operation which
	// can be used to track the progress of turning up the new cluster.
	// Immediately upon completion of this request:
	//  * The cluster will be readable via the API, with all requested attributes
	//    but no allocated resources.
	// Until completion of the embedded operation:
	//  * Cancelling the operation will render the cluster immediately unreadable
	//    via the API.
	//  * All other attempts to modify or delete the cluster will be rejected.
	// Upon completion of the embedded operation:
	//  * Billing for all successfully-allocated resources will begin (some types
	//    may have lower than the requested levels).
	//  * New tables can be created in the cluster.
	//  * The cluster's allocated resource levels will be readable via the API.
	// The embedded operation's "metadata" field type is
	// [CreateClusterMetadata][google.bigtable.admin.cluster.v1.CreateClusterMetadata] The embedded operation's "response" field type is
	// [Cluster][google.bigtable.admin.cluster.v1.Cluster], if successful.
	CreateCluster(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*Cluster, error)
	// Updates a cluster, and begins allocating or releasing resources as
	// requested. The returned cluster embeds as its "current_operation" a
	// long-running operation which can be used to track the progress of updating
	// the cluster.
	// Immediately upon completion of this request:
	//  * For resource types where a decrease in the cluster's allocation has been
	//    requested, billing will be based on the newly-requested level.
	// Until completion of the embedded operation:
	//  * Cancelling the operation will set its metadata's "cancelled_at_time",
	//    and begin restoring resources to their pre-request values. The operation
	//    is guaranteed to succeed at undoing all resource changes, after which
	//    point it will terminate with a CANCELLED status.
	//  * All other attempts to modify or delete the cluster will be rejected.
	//  * Reading the cluster via the API will continue to give the pre-request
	//    resource levels.
	// Upon completion of the embedded operation:
	//  * Billing will begin for all successfully-allocated resources (some types
	//    may have lower than the requested levels).
	//  * All newly-reserved resources will be available for serving the cluster's
	//    tables.
	//  * The cluster's new resource levels will be readable via the API.
	// [UpdateClusterMetadata][google.bigtable.admin.cluster.v1.UpdateClusterMetadata] The embedded operation's "response" field type is
	// [Cluster][google.bigtable.admin.cluster.v1.Cluster], if successful.
	UpdateCluster(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*Cluster, error)
	// Marks a cluster and all of its tables for permanent deletion in 7 days.
	// Immediately upon completion of the request:
	//  * Billing will cease for all of the cluster's reserved resources.
	//  * The cluster's "delete_time" field will be set 7 days in the future.
	// Soon afterward:
	//  * All tables within the cluster will become unavailable.
	// Prior to the cluster's "delete_time":
	//  * The cluster can be recovered with a call to UndeleteCluster.
	//  * All other attempts to modify or delete the cluster will be rejected.
	// At the cluster's "delete_time":
	//  * The cluster and *all of its tables* will immediately and irrevocably
	//    disappear from the API, and their data will be permanently deleted.
	DeleteCluster(ctx context.Context, in *DeleteClusterRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Cancels the scheduled deletion of an cluster and begins preparing it to
	// resume serving. The returned operation will also be embedded as the
	// cluster's "current_operation".
	// Immediately upon completion of this request:
	//  * The cluster's "delete_time" field will be unset, protecting it from
	//    automatic deletion.
	// Until completion of the returned operation:
	//  * The operation cannot be cancelled.
	// Upon completion of the returned operation:
	//  * Billing for the cluster's resources will resume.
	//  * All tables within the cluster will be available.
	// [UndeleteClusterMetadata][google.bigtable.admin.cluster.v1.UndeleteClusterMetadata] The embedded operation's "response" field type is
	// [Cluster][google.bigtable.admin.cluster.v1.Cluster], if successful.
	UndeleteCluster(ctx context.Context, in *UndeleteClusterRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
}

type bigtableClusterServiceClient struct {
	cc *grpc.ClientConn
}

func NewBigtableClusterServiceClient(cc *grpc.ClientConn) BigtableClusterServiceClient {
	return &bigtableClusterServiceClient{cc}
}

func (c *bigtableClusterServiceClient) ListZones(ctx context.Context, in *ListZonesRequest, opts ...grpc.CallOption) (*ListZonesResponse, error) {
	out := new(ListZonesResponse)
	err := c.cc.Invoke(ctx, "/google.bigtable.admin.cluster.v1.BigtableClusterService/ListZones", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableClusterServiceClient) GetCluster(ctx context.Context, in *GetClusterRequest, opts ...grpc.CallOption) (*Cluster, error) {
	out := new(Cluster)
	err := c.cc.Invoke(ctx, "/google.bigtable.admin.cluster.v1.BigtableClusterService/GetCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableClusterServiceClient) ListClusters(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (*ListClustersResponse, error) {
	out := new(ListClustersResponse)
	err := c.cc.Invoke(ctx, "/google.bigtable.admin.cluster.v1.BigtableClusterService/ListClusters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableClusterServiceClient) CreateCluster(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*Cluster, error) {
	out := new(Cluster)
	err := c.cc.Invoke(ctx, "/google.bigtable.admin.cluster.v1.BigtableClusterService/CreateCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableClusterServiceClient) UpdateCluster(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*Cluster, error) {
	out := new(Cluster)
	err := c.cc.Invoke(ctx, "/google.bigtable.admin.cluster.v1.BigtableClusterService/UpdateCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableClusterServiceClient) DeleteCluster(ctx context.Context, in *DeleteClusterRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.bigtable.admin.cluster.v1.BigtableClusterService/DeleteCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableClusterServiceClient) UndeleteCluster(ctx context.Context, in *UndeleteClusterRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/google.bigtable.admin.cluster.v1.BigtableClusterService/UndeleteCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BigtableClusterServiceServer is the server API for BigtableClusterService service.
type BigtableClusterServiceServer interface {
	// Lists the supported zones for the given project.
	ListZones(context.Context, *ListZonesRequest) (*ListZonesResponse, error)
	// Gets information about a particular cluster.
	GetCluster(context.Context, *GetClusterRequest) (*Cluster, error)
	// Lists all clusters in the given project, along with any zones for which
	// cluster information could not be retrieved.
	ListClusters(context.Context, *ListClustersRequest) (*ListClustersResponse, error)
	// Creates a cluster and begins preparing it to begin serving. The returned
	// cluster embeds as its "current_operation" a long-running operation which
	// can be used to track the progress of turning up the new cluster.
	// Immediately upon completion of this request:
	//  * The cluster will be readable via the API, with all requested attributes
	//    but no allocated resources.
	// Until completion of the embedded operation:
	//  * Cancelling the operation will render the cluster immediately unreadable
	//    via the API.
	//  * All other attempts to modify or delete the cluster will be rejected.
	// Upon completion of the embedded operation:
	//  * Billing for all successfully-allocated resources will begin (some types
	//    may have lower than the requested levels).
	//  * New tables can be created in the cluster.
	//  * The cluster's allocated resource levels will be readable via the API.
	// The embedded operation's "metadata" field type is
	// [CreateClusterMetadata][google.bigtable.admin.cluster.v1.CreateClusterMetadata] The embedded operation's "response" field type is
	// [Cluster][google.bigtable.admin.cluster.v1.Cluster], if successful.
	CreateCluster(context.Context, *CreateClusterRequest) (*Cluster, error)
	// Updates a cluster, and begins allocating or releasing resources as
	// requested. The returned cluster embeds as its "current_operation" a
	// long-running operation which can be used to track the progress of updating
	// the cluster.
	// Immediately upon completion of this request:
	//  * For resource types where a decrease in the cluster's allocation has been
	//    requested, billing will be based on the newly-requested level.
	// Until completion of the embedded operation:
	//  * Cancelling the operation will set its metadata's "cancelled_at_time",
	//    and begin restoring resources to their pre-request values. The operation
	//    is guaranteed to succeed at undoing all resource changes, after which
	//    point it will terminate with a CANCELLED status.
	//  * All other attempts to modify or delete the cluster will be rejected.
	//  * Reading the cluster via the API will continue to give the pre-request
	//    resource levels.
	// Upon completion of the embedded operation:
	//  * Billing will begin for all successfully-allocated resources (some types
	//    may have lower than the requested levels).
	//  * All newly-reserved resources will be available for serving the cluster's
	//    tables.
	//  * The cluster's new resource levels will be readable via the API.
	// [UpdateClusterMetadata][google.bigtable.admin.cluster.v1.UpdateClusterMetadata] The embedded operation's "response" field type is
	// [Cluster][google.bigtable.admin.cluster.v1.Cluster], if successful.
	UpdateCluster(context.Context, *Cluster) (*Cluster, error)
	// Marks a cluster and all of its tables for permanent deletion in 7 days.
	// Immediately upon completion of the request:
	//  * Billing will cease for all of the cluster's reserved resources.
	//  * The cluster's "delete_time" field will be set 7 days in the future.
	// Soon afterward:
	//  * All tables within the cluster will become unavailable.
	// Prior to the cluster's "delete_time":
	//  * The cluster can be recovered with a call to UndeleteCluster.
	//  * All other attempts to modify or delete the cluster will be rejected.
	// At the cluster's "delete_time":
	//  * The cluster and *all of its tables* will immediately and irrevocably
	//    disappear from the API, and their data will be permanently deleted.
	DeleteCluster(context.Context, *DeleteClusterRequest) (*empty.Empty, error)
	// Cancels the scheduled deletion of an cluster and begins preparing it to
	// resume serving. The returned operation will also be embedded as the
	// cluster's "current_operation".
	// Immediately upon completion of this request:
	//  * The cluster's "delete_time" field will be unset, protecting it from
	//    automatic deletion.
	// Until completion of the returned operation:
	//  * The operation cannot be cancelled.
	// Upon completion of the returned operation:
	//  * Billing for the cluster's resources will resume.
	//  * All tables within the cluster will be available.
	// [UndeleteClusterMetadata][google.bigtable.admin.cluster.v1.UndeleteClusterMetadata] The embedded operation's "response" field type is
	// [Cluster][google.bigtable.admin.cluster.v1.Cluster], if successful.
	UndeleteCluster(context.Context, *UndeleteClusterRequest) (*longrunning.Operation, error)
}

func RegisterBigtableClusterServiceServer(s *grpc.Server, srv BigtableClusterServiceServer) {
	s.RegisterService(&_BigtableClusterService_serviceDesc, srv)
}

func _BigtableClusterService_ListZones_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListZonesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableClusterServiceServer).ListZones(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.cluster.v1.BigtableClusterService/ListZones",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableClusterServiceServer).ListZones(ctx, req.(*ListZonesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableClusterService_GetCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableClusterServiceServer).GetCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.cluster.v1.BigtableClusterService/GetCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableClusterServiceServer).GetCluster(ctx, req.(*GetClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableClusterService_ListClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableClusterServiceServer).ListClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.cluster.v1.BigtableClusterService/ListClusters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableClusterServiceServer).ListClusters(ctx, req.(*ListClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableClusterService_CreateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableClusterServiceServer).CreateCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.cluster.v1.BigtableClusterService/CreateCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableClusterServiceServer).CreateCluster(ctx, req.(*CreateClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableClusterService_UpdateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableClusterServiceServer).UpdateCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.cluster.v1.BigtableClusterService/UpdateCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableClusterServiceServer).UpdateCluster(ctx, req.(*Cluster))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableClusterService_DeleteCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableClusterServiceServer).DeleteCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.cluster.v1.BigtableClusterService/DeleteCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableClusterServiceServer).DeleteCluster(ctx, req.(*DeleteClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableClusterService_UndeleteCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UndeleteClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableClusterServiceServer).UndeleteCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.cluster.v1.BigtableClusterService/UndeleteCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableClusterServiceServer).UndeleteCluster(ctx, req.(*UndeleteClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BigtableClusterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.bigtable.admin.cluster.v1.BigtableClusterService",
	HandlerType: (*BigtableClusterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListZones",
			Handler:    _BigtableClusterService_ListZones_Handler,
		},
		{
			MethodName: "GetCluster",
			Handler:    _BigtableClusterService_GetCluster_Handler,
		},
		{
			MethodName: "ListClusters",
			Handler:    _BigtableClusterService_ListClusters_Handler,
		},
		{
			MethodName: "CreateCluster",
			Handler:    _BigtableClusterService_CreateCluster_Handler,
		},
		{
			MethodName: "UpdateCluster",
			Handler:    _BigtableClusterService_UpdateCluster_Handler,
		},
		{
			MethodName: "DeleteCluster",
			Handler:    _BigtableClusterService_DeleteCluster_Handler,
		},
		{
			MethodName: "UndeleteCluster",
			Handler:    _BigtableClusterService_UndeleteCluster_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/bigtable/admin/cluster/v1/bigtable_cluster_service.proto",
}

func init() {
	proto.RegisterFile("google/bigtable/admin/cluster/v1/bigtable_cluster_service.proto", fileDescriptor_bigtable_cluster_service_eb6b3340e63f14bc)
}

var fileDescriptor_bigtable_cluster_service_eb6b3340e63f14bc = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x4f, 0x6b, 0x14, 0x31,
	0x18, 0xc6, 0x89, 0x07, 0xa1, 0xc1, 0x45, 0xc8, 0xa1, 0x87, 0x6d, 0x0b, 0x32, 0x15, 0xb1, 0x23,
	0x26, 0x6e, 0x17, 0xc5, 0xbf, 0x08, 0x5b, 0xa5, 0x1e, 0x04, 0x8b, 0xd2, 0x4b, 0x2f, 0x4b, 0x76,
	0xe7, 0x35, 0x8c, 0xcc, 0x24, 0x31, 0xc9, 0x2c, 0xa8, 0xf4, 0xe2, 0xcd, 0x93, 0x88, 0x27, 0x3d,
	0x78, 0xeb, 0xdd, 0xef, 0xe2, 0x57, 0xf0, 0x83, 0xc8, 0x64, 0x92, 0xb5, 0x2b, 0x6b, 0x77, 0xa6,
	0xb7, 0x99, 0xc9, 0xfb, 0xbc, 0xcf, 0x6f, 0x9e, 0x24, 0x2f, 0x7e, 0x2c, 0x94, 0x12, 0x05, 0xb0,
	0x49, 0x2e, 0x1c, 0x9f, 0x14, 0xc0, 0x78, 0x56, 0xe6, 0x92, 0x4d, 0x8b, 0xca, 0x3a, 0x30, 0x6c,
	0x36, 0x98, 0xaf, 0x8c, 0xc3, 0xb7, 0xb1, 0x05, 0x33, 0xcb, 0xa7, 0x40, 0xb5, 0x51, 0x4e, 0x91,
	0x2b, 0x4d, 0x03, 0x1a, 0xcb, 0xa8, 0x6f, 0x40, 0x43, 0x31, 0x9d, 0x0d, 0xfa, 0x9b, 0xc1, 0x82,
	0xeb, 0x9c, 0x71, 0x29, 0x95, 0xe3, 0x2e, 0x57, 0xd2, 0x36, 0xfa, 0xfe, 0xc3, 0xee, 0x00, 0x19,
	0x77, 0x3c, 0xa8, 0x9f, 0x9d, 0x1b, 0x7f, 0x5c, 0x82, 0xb5, 0x5c, 0x40, 0xe4, 0xd8, 0x0e, 0x9d,
	0x0a, 0x25, 0x85, 0xa9, 0xa4, 0xcc, 0xa5, 0x60, 0x4a, 0x83, 0x59, 0x80, 0xdd, 0x08, 0x45, 0xfe,
	0x6d, 0x52, 0xbd, 0x66, 0x50, 0x6a, 0xf7, 0xae, 0x59, 0xdc, 0xfd, 0xb4, 0x86, 0xd7, 0x47, 0xc1,
	0x6d, 0xaf, 0x31, 0x7b, 0xd5, 0x78, 0x91, 0x6f, 0x08, 0xaf, 0x3d, 0xcf, 0xad, 0x3b, 0x52, 0x12,
	0x2c, 0xd9, 0xa5, 0xab, 0x32, 0xa3, 0xf3, 0xe2, 0x97, 0xf0, 0xb6, 0x02, 0xeb, 0xfa, 0xc3, 0x4e,
	0x1a, 0xab, 0x95, 0xb4, 0x90, 0x6c, 0x7f, 0xfc, 0xf5, 0xfb, 0xeb, 0x85, 0x2d, 0xb2, 0x51, 0x07,
	0xf1, 0x41, 0xf2, 0x12, 0x1e, 0x69, 0xa3, 0xde, 0xc0, 0xd4, 0x59, 0x96, 0x1e, 0xb3, 0xf7, 0x9e,
	0xe6, 0x07, 0xc2, 0x78, 0x1f, 0x5c, 0x20, 0x26, 0x2d, 0x8c, 0xfe, 0x56, 0x47, 0xba, 0x9d, 0xd5,
	0xa2, 0xa0, 0x48, 0x6e, 0x79, 0xa6, 0x94, 0x5c, 0x5f, 0xc6, 0xd4, 0x20, 0xb1, 0x34, 0x6e, 0x60,
	0x8d, 0x49, 0x7e, 0x22, 0x7c, 0xa9, 0xfe, 0xb7, 0xd0, 0xc1, 0x92, 0xdb, 0xed, 0xb2, 0x88, 0xf5,
	0x11, 0xf2, 0x4e, 0x57, 0x59, 0x48, 0x71, 0xe0, 0x89, 0x6f, 0x90, 0x9d, 0xe5, 0x29, 0x72, 0x21,
	0x0c, 0x08, 0xee, 0x20, 0x9b, 0x53, 0x93, 0x13, 0x84, 0x7b, 0x7b, 0x06, 0xb8, 0x8b, 0x07, 0x81,
	0xb4, 0x30, 0x5f, 0x10, 0x9c, 0x23, 0xd9, 0xc0, 0x99, 0x5c, 0x3b, 0x2b, 0xd9, 0xe3, 0x39, 0xe4,
	0x7d, 0x94, 0x92, 0xef, 0x08, 0xf7, 0x0e, 0x75, 0x76, 0x8a, 0xb3, 0xbd, 0x5f, 0x17, 0xb4, 0xa1,
	0x47, 0xbb, 0xd9, 0x6f, 0xbd, 0xe9, 0x35, 0xdc, 0x17, 0x84, 0x7b, 0x4f, 0xa0, 0x80, 0x4e, 0x21,
	0x2e, 0x08, 0x62, 0x88, 0xeb, 0x51, 0x17, 0xef, 0x2d, 0x7d, 0x5a, 0xdf, 0xdb, 0x78, 0x16, 0xd3,
	0xf6, 0x67, 0xf1, 0x04, 0xe1, 0xcb, 0x87, 0x32, 0x5b, 0xa0, 0xba, 0xbb, 0x9a, 0xea, 0x1f, 0x49,
	0xe4, 0xda, 0x8a, 0xca, 0x53, 0x43, 0x87, 0xbe, 0x88, 0x43, 0x27, 0xb9, 0xe7, 0xf1, 0x86, 0xc9,
	0xa0, 0x75, 0x6a, 0x55, 0xf0, 0x19, 0x7d, 0x46, 0xf8, 0xea, 0x54, 0x95, 0x2b, 0xc9, 0x46, 0x9b,
	0xcb, 0x27, 0x96, 0x3d, 0xa8, 0x93, 0x3a, 0x40, 0x47, 0xfb, 0xa1, 0x83, 0x50, 0x05, 0x97, 0x82,
	0x2a, 0x23, 0x98, 0x00, 0xe9, 0x73, 0x64, 0xcd, 0x12, 0xd7, 0xb9, 0xfd, 0xff, 0xfc, 0x7d, 0x10,
	0x1e, 0x27, 0x17, 0xbd, 0x66, 0xf8, 0x27, 0x00, 0x00, 0xff, 0xff, 0x50, 0x92, 0x91, 0x86, 0x71,
	0x06, 0x00, 0x00,
}
