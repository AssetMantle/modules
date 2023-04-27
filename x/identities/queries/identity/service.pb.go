// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: identities/queries/identity/service.proto

package identity

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
	proto.RegisterFile("identities/queries/identity/service.proto", fileDescriptor_02e93da1e07ea84b)
}

var fileDescriptor_02e93da1e07ea84b = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x3f, 0x4b, 0xfb, 0x40,
	0x18, 0xc7, 0x9b, 0xc0, 0xaf, 0x3f, 0x08, 0x4e, 0x19, 0x8b, 0x64, 0x50, 0x1c, 0x44, 0xc8, 0x59,
	0xbb, 0xc5, 0x3f, 0xd0, 0x2e, 0xda, 0x21, 0xd0, 0xea, 0x26, 0x01, 0xb9, 0x36, 0x0f, 0xf5, 0x20,
	0xb9, 0x6b, 0x73, 0x97, 0x62, 0x56, 0x5f, 0x81, 0xe0, 0x3b, 0x70, 0x74, 0xf6, 0x05, 0x38, 0x8a,
	0x53, 0xc1, 0xc5, 0x51, 0x52, 0x27, 0xdf, 0x82, 0x8b, 0xc4, 0x7b, 0xda, 0xd4, 0xa5, 0x58, 0xa7,
	0x83, 0x2f, 0x7c, 0xbe, 0xf7, 0x79, 0x9e, 0x3b, 0x6b, 0x9b, 0x85, 0xc0, 0x15, 0x53, 0x0c, 0x24,
	0x19, 0xa5, 0x90, 0x14, 0x27, 0x46, 0x19, 0x91, 0x90, 0x8c, 0x59, 0x1f, 0xdc, 0x61, 0x22, 0x94,
	0xb0, 0x09, 0x95, 0x12, 0x54, 0x4c, 0xb9, 0x8a, 0xc0, 0x8d, 0x45, 0x98, 0x46, 0x20, 0xdd, 0x12,
	0x77, 0x11, 0x9f, 0x45, 0x59, 0x6d, 0x7d, 0x20, 0xc4, 0x20, 0x02, 0x42, 0x87, 0x8c, 0x50, 0xce,
	0x85, 0xa2, 0x8a, 0x09, 0x2e, 0x75, 0x5d, 0x8d, 0x2c, 0xbb, 0xb9, 0x08, 0xb2, 0x8b, 0x04, 0x46,
	0x29, 0x48, 0x85, 0xc0, 0xee, 0x6f, 0x00, 0x39, 0x14, 0x5c, 0xa2, 0xf1, 0xde, 0xa3, 0x61, 0xfd,
	0x3f, 0xd3, 0x33, 0xd8, 0x0f, 0x86, 0x55, 0x3d, 0xa1, 0x3c, 0x8c, 0xc0, 0x3e, 0x74, 0x57, 0x9c,
	0xc4, 0xed, 0x16, 0xed, 0xa7, 0xda, 0xa6, 0x76, 0xf4, 0x57, 0x5c, 0xbb, 0x6d, 0xec, 0x5c, 0xbf,
	0xbc, 0xdf, 0x9a, 0x5b, 0xf6, 0x26, 0xd1, 0x15, 0x8b, 0x8b, 0x18, 0xd7, 0x7b, 0xa0, 0x68, 0x7d,
	0x3e, 0x57, 0xeb, 0xd3, 0x7c, 0xca, 0x1d, 0x63, 0x92, 0x3b, 0xc6, 0x5b, 0xee, 0x18, 0x37, 0x53,
	0xa7, 0x32, 0x99, 0x3a, 0x95, 0xd7, 0xa9, 0x53, 0xb1, 0x1a, 0x7d, 0x11, 0xaf, 0xaa, 0xd2, 0x5a,
	0xc3, 0x7d, 0x74, 0x8a, 0x05, 0x75, 0x8c, 0xf3, 0x83, 0x01, 0x53, 0x97, 0x69, 0xcf, 0xed, 0x8b,
	0x98, 0x34, 0x8b, 0x2e, 0x5f, 0x3b, 0x61, 0x17, 0xb9, 0x5a, 0xf6, 0x4c, 0x77, 0xe6, 0xbf, 0xa6,
	0xdf, 0xee, 0xb6, 0xef, 0x4d, 0xcd, 0xa2, 0x87, 0x8f, 0x1e, 0xed, 0xd2, 0xa3, 0x8b, 0x1e, 0x18,
	0x65, 0xcf, 0x3f, 0x88, 0x00, 0x89, 0xa0, 0x24, 0x02, 0x24, 0x66, 0x51, 0x96, 0x9b, 0xfb, 0x2b,
	0x12, 0xc1, 0x71, 0xa7, 0xe5, 0x83, 0xa2, 0x21, 0x55, 0xf4, 0xc3, 0x6c, 0x2c, 0xd0, 0x9e, 0x87,
	0xb8, 0xe7, 0x95, 0xbc, 0xe7, 0x61, 0xc1, 0x3c, 0xcc, 0x7a, 0xd5, 0xef, 0x7f, 0xd4, 0xf8, 0x0a,
	0x00, 0x00, 0xff, 0xff, 0x55, 0xa8, 0x11, 0xbd, 0x26, 0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/assetmantle.modules.identities.queries.identity.Service/Handle", in, out, opts...)
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
		FullMethod: "/assetmantle.modules.identities.queries.identity.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assetmantle.modules.identities.queries.identity.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "identities/queries/identity/service.proto",
}