// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/metas/internal/queries/meta/service.proto

package meta

import (
	context "context"
	fmt "fmt"
	math "math"

	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

func init() {
	proto.RegisterFile("x/metas/internal/queries/meta/service.proto", fileDescriptor_bf99570eed0a1b83)
}

var fileDescriptor_bf99570eed0a1b83 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x31, 0x4b, 0xfb, 0x40,
	0x18, 0xc6, 0x9b, 0xc0, 0xbf, 0x7f, 0x08, 0x4e, 0x19, 0x43, 0xc9, 0xd0, 0xa1, 0x43, 0x85, 0x3b,
	0x5b, 0x75, 0x09, 0x2e, 0xed, 0xa2, 0x4b, 0xa0, 0xd5, 0x4d, 0xb2, 0x5c, 0xdb, 0x97, 0x7a, 0x90,
	0xdc, 0xb5, 0xb9, 0x4b, 0xd1, 0xd5, 0x4f, 0x20, 0x38, 0xb9, 0x0a, 0x2e, 0x7e, 0x12, 0x71, 0x90,
	0x82, 0x8b, 0xa3, 0xa4, 0x4e, 0x7e, 0x0a, 0xc9, 0xdd, 0x3b, 0xd4, 0xa5, 0xc4, 0x29, 0x3c, 0xe1,
	0xf9, 0xbd, 0xcf, 0xf3, 0x24, 0xde, 0xfe, 0x35, 0xcd, 0x40, 0x33, 0x45, 0xb9, 0xd0, 0x90, 0x0b,
	0x96, 0xd2, 0x65, 0x01, 0x39, 0x07, 0x65, 0x5e, 0x53, 0x05, 0xf9, 0x8a, 0x4f, 0x81, 0x2c, 0x72,
	0xa9, 0xa5, 0xdf, 0x61, 0x4a, 0x81, 0xce, 0x98, 0xd0, 0x29, 0x90, 0x4c, 0xce, 0x8a, 0x14, 0x14,
	0x31, 0x38, 0x41, 0xca, 0xa8, 0xa0, 0x35, 0x97, 0x72, 0x9e, 0x02, 0x65, 0x0b, 0x4e, 0x99, 0x10,
	0x52, 0x33, 0xcd, 0xa5, 0x50, 0xf6, 0x4a, 0x70, 0xb0, 0x3b, 0xb2, 0x12, 0x37, 0xe7, 0xb0, 0x2c,
	0x40, 0x69, 0x24, 0x7a, 0xb5, 0x08, 0xb5, 0x90, 0x42, 0x61, 0xd5, 0xfe, 0x93, 0xe3, 0xfd, 0xbf,
	0xb0, 0xe5, 0xfd, 0x07, 0xc7, 0x6b, 0x9e, 0x31, 0x31, 0x4b, 0xc1, 0x3f, 0x22, 0xf5, 0x26, 0x90,
	0xf1, 0x56, 0x8b, 0xe0, 0xf8, 0x8f, 0x94, 0x6d, 0xd2, 0x6e, 0xdf, 0xbe, 0x7f, 0xdd, 0xbb, 0x2d,
	0x3f, 0xa0, 0x96, 0xc4, 0x11, 0xab, 0xde, 0x04, 0x34, 0xeb, 0x19, 0x35, 0x7c, 0x73, 0x5f, 0xca,
	0xd0, 0x59, 0x97, 0xa1, 0xf3, 0x59, 0x86, 0xce, 0xdd, 0x26, 0x6c, 0xac, 0x37, 0x61, 0xe3, 0x63,
	0x13, 0x36, 0xbc, 0xee, 0x54, 0x66, 0x35, 0x83, 0x87, 0x7b, 0xb8, 0x75, 0x54, 0x8d, 0x1f, 0x39,
	0x97, 0x27, 0x73, 0xae, 0xaf, 0x8a, 0x09, 0x99, 0xca, 0x8c, 0x0e, 0xaa, 0x13, 0x31, 0x36, 0xb0,
	0x27, 0xe8, 0xce, 0x0f, 0xfa, 0xe8, 0xfe, 0x1b, 0xc4, 0xf1, 0x38, 0x7e, 0x76, 0x3b, 0x83, 0xad,
	0xf8, 0x18, 0xe3, 0x63, 0x13, 0x3f, 0xc6, 0xf8, 0x4a, 0xbd, 0xfe, 0x32, 0x26, 0x68, 0x4c, 0x8c,
	0x31, 0x41, 0xa3, 0x51, 0xa5, 0xdb, 0xaf, 0x67, 0x4c, 0x4e, 0x47, 0xc3, 0xea, 0x39, 0x63, 0x9a,
	0x7d, 0xbb, 0xdd, 0x2d, 0x28, 0x8a, 0x90, 0x8a, 0x22, 0x83, 0x45, 0x11, 0x72, 0x56, 0x4f, 0x9a,
	0xe6, 0xff, 0x1f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x4d, 0x6b, 0x1a, 0xd9, 0x02, 0x00,
	0x00,
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
	err := c.cc.Invoke(ctx, "/assetmantle.modules.metas.queries.meta.Service/Handle", in, out, opts...)
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
		FullMethod: "/assetmantle.modules.metas.queries.meta.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assetmantle.modules.metas.queries.meta.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "x/metas/internal/queries/meta/service.proto",
}
