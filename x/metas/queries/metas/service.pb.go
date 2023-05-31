// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: metas/queries/metas/service.proto

package metas

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("metas/queries/metas/service.proto", fileDescriptor_60b84daf56824d07) }

var fileDescriptor_60b84daf56824d07 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x3f, 0x4b, 0xc3, 0x40,
	0x18, 0xc6, 0x9b, 0x80, 0x15, 0x82, 0x53, 0xc6, 0xaa, 0x01, 0x75, 0xa8, 0x50, 0xb8, 0xa3, 0x56,
	0x1d, 0xb2, 0xb5, 0x8b, 0x2e, 0x81, 0x56, 0x37, 0x09, 0xc8, 0xb5, 0x7d, 0xa9, 0x07, 0xc9, 0x5d,
	0x9b, 0xbb, 0x14, 0x5d, 0xfd, 0x04, 0x82, 0xab, 0x93, 0x93, 0xf8, 0x49, 0x44, 0x97, 0x82, 0x8b,
	0xa3, 0xa4, 0x4e, 0x7e, 0x0a, 0x49, 0xee, 0x55, 0x2a, 0x14, 0x89, 0x5b, 0x1e, 0xf2, 0x7b, 0xfe,
	0xdc, 0x9d, 0xb3, 0x15, 0x83, 0x66, 0x8a, 0x4e, 0x52, 0x48, 0x38, 0x28, 0x6a, 0x94, 0x82, 0x64,
	0xca, 0x07, 0x40, 0xc6, 0x89, 0xd4, 0xd2, 0xad, 0x33, 0xa5, 0x40, 0xc7, 0x4c, 0xe8, 0x08, 0x48,
	0x2c, 0x87, 0x69, 0x04, 0x8a, 0x14, 0x20, 0x41, 0x9b, 0x51, 0xb5, 0x8d, 0x91, 0x94, 0xa3, 0x08,
	0x28, 0x1b, 0x73, 0xca, 0x84, 0x90, 0x9a, 0x69, 0x2e, 0x85, 0x32, 0x31, 0xb5, 0xfa, 0xb2, 0xa6,
	0x5c, 0x5d, 0x9d, 0x27, 0x30, 0x49, 0x41, 0x69, 0x04, 0x77, 0xff, 0x02, 0xd5, 0x58, 0x0a, 0x85,
	0xcb, 0xf6, 0x1e, 0x2c, 0x67, 0xf5, 0xd4, 0x6c, 0x75, 0xef, 0x2c, 0xa7, 0x7a, 0xcc, 0xc4, 0x30,
	0x02, 0xf7, 0x80, 0x94, 0x5c, 0x4c, 0x7a, 0x79, 0xea, 0x89, 0x69, 0xaf, 0x1d, 0xfe, 0xd7, 0x66,
	0xb6, 0x6c, 0xef, 0x5c, 0xbf, 0x7e, 0xdc, 0xda, 0x9b, 0xee, 0x3a, 0x35, 0x56, 0x5c, 0x3c, 0x6d,
	0xf6, 0x41, 0xb3, 0xa6, 0x51, 0x9d, 0x17, 0xfb, 0x29, 0xf3, 0xac, 0x59, 0xe6, 0x59, 0xef, 0x99,
	0x67, 0xdd, 0xcc, 0xbd, 0xca, 0x6c, 0xee, 0x55, 0xde, 0xe6, 0x5e, 0xc5, 0x69, 0x0c, 0x64, 0x5c,
	0xb6, 0xba, 0xb3, 0x86, 0xe7, 0xed, 0xe6, 0x17, 0xd0, 0xb5, 0xce, 0xf6, 0x47, 0x5c, 0x5f, 0xa4,
	0x7d, 0x32, 0x90, 0x31, 0x6d, 0xe7, 0x19, 0x01, 0x6e, 0x30, 0x19, 0xf4, 0x92, 0x2e, 0xb9, 0xcd,
	0x7b, 0x7b, 0xa5, 0x1d, 0x04, 0xbd, 0xe0, 0xd1, 0xae, 0xb7, 0x17, 0x7a, 0x03, 0xec, 0x0d, 0x7e,
	0x0e, 0xc9, 0xbf, 0xd5, 0xf3, 0x2f, 0x32, 0x44, 0x32, 0x2c, 0xfe, 0x85, 0x48, 0x1a, 0x95, 0xd9,
	0xad, 0x92, 0x64, 0x78, 0xd4, 0xed, 0xe4, 0x1f, 0x43, 0xa6, 0xd9, 0xa7, 0xdd, 0x58, 0x70, 0xf9,
	0x3e, 0xda, 0x7c, 0xbf, 0x20, 0x7d, 0x1f, 0x8d, 0xa8, 0xfb, 0xd5, 0xe2, 0xfd, 0x5b, 0x5f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xec, 0x4d, 0x31, 0xdb, 0xbe, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	Handle(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
}

type serviceClient struct {
	cc grpc1.ClientConn
}

func NewServiceClient(cc grpc1.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Handle(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/assetmantle.modules.metas.queries.metas.Service/Handle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	Handle(context.Context, *QueryRequest) (*QueryResponse, error)
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) Handle(ctx context.Context, req *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handle not implemented")
}

func RegisterServiceServer(s grpc1.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_Handle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assetmantle.modules.metas.queries.metas.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assetmantle.modules.metas.queries.metas.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metas/queries/metas/service.proto",
}
