// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/maintainers/internal/queries/maintainer/service.proto

package maintainer

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
	proto.RegisterFile("x/maintainers/internal/queries/maintainer/service.proto", fileDescriptor_f77c6eede371cc07)
}

var fileDescriptor_f77c6eede371cc07 = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x31, 0x4b, 0xfb, 0x40,
	0x18, 0xc6, 0x9b, 0xfc, 0xf9, 0x57, 0x08, 0x4e, 0x19, 0x8b, 0x64, 0x70, 0x10, 0xa7, 0x3b, 0x5a,
	0xa1, 0x42, 0x50, 0x31, 0x5d, 0xec, 0x72, 0xd0, 0xea, 0x26, 0x59, 0xde, 0xb6, 0x2f, 0xf5, 0x20,
	0xb9, 0x6b, 0x73, 0x97, 0xa2, 0xab, 0x9f, 0x40, 0x10, 0xfc, 0x00, 0x8e, 0x7e, 0x07, 0x77, 0x75,
	0x2a, 0xb8, 0x38, 0x4a, 0xea, 0xe4, 0xa7, 0x90, 0x9a, 0x03, 0xaf, 0x5b, 0xa3, 0xeb, 0x13, 0x9e,
	0x1f, 0xbf, 0xe7, 0xcd, 0x79, 0xfb, 0x97, 0x34, 0x05, 0x2e, 0x34, 0x70, 0x81, 0x99, 0xa2, 0x5c,
	0x68, 0xcc, 0x04, 0x24, 0x74, 0x9a, 0x63, 0xc6, 0x51, 0x59, 0x1f, 0xa9, 0xc2, 0x6c, 0xc6, 0x87,
	0x48, 0x26, 0x99, 0xd4, 0xd2, 0x6f, 0x81, 0x52, 0xa8, 0x53, 0x10, 0x3a, 0x41, 0x92, 0xca, 0x51,
	0x9e, 0xa0, 0x22, 0x16, 0x8a, 0x18, 0x82, 0x95, 0x35, 0xb6, 0xc6, 0x52, 0x8e, 0x13, 0xa4, 0x30,
	0xe1, 0x14, 0x84, 0x90, 0x1a, 0x34, 0x97, 0x42, 0x95, 0xc4, 0xc6, 0xc1, 0xfa, 0x2a, 0xcb, 0xe8,
	0xea, 0x14, 0xa7, 0x39, 0x2a, 0x6d, 0xda, 0x87, 0x95, 0xdb, 0x6a, 0x22, 0x85, 0x32, 0x73, 0x5a,
	0xcf, 0x8e, 0xb7, 0x71, 0x56, 0x0e, 0xf4, 0x1f, 0x1d, 0xaf, 0xde, 0x05, 0x31, 0x4a, 0xd0, 0x3f,
	0x26, 0xd5, 0x67, 0x92, 0xbe, 0x65, 0xd7, 0x88, 0xfe, 0x40, 0x28, 0x0d, 0xb7, 0xc9, 0xf5, 0xeb,
	0xc7, 0xad, 0xbb, 0xeb, 0xef, 0xd0, 0x92, 0xb2, 0x32, 0x74, 0xd6, 0x1c, 0xa0, 0x86, 0xa6, 0x95,
	0x75, 0xee, 0xfe, 0x3d, 0x15, 0x81, 0x33, 0x2f, 0x02, 0xe7, 0xbd, 0x08, 0x9c, 0x9b, 0x45, 0x50,
	0x9b, 0x2f, 0x82, 0xda, 0xdb, 0x22, 0xa8, 0x79, 0xed, 0xa1, 0x4c, 0x7f, 0x21, 0xd4, 0xd9, 0x34,
	0xb7, 0xe9, 0x2d, 0x8f, 0xd5, 0x73, 0xce, 0xbb, 0x63, 0xae, 0x2f, 0xf2, 0x01, 0x19, 0xca, 0x94,
	0x46, 0x4b, 0x1c, 0x33, 0x66, 0x25, 0x8e, 0xae, 0xfd, 0x33, 0xee, 0xdd, 0xff, 0x11, 0x63, 0x7d,
	0xf6, 0xe0, 0xb6, 0x22, 0x4b, 0x8b, 0x19, 0x2d, 0x66, 0x69, 0xf5, 0x8d, 0xd6, 0x4f, 0xf6, 0xb2,
	0x52, 0x8a, 0x4d, 0x29, 0xb6, 0x4a, 0xb1, 0x29, 0x59, 0x59, 0xe1, 0x1e, 0x55, 0x2f, 0xc5, 0x27,
	0xbd, 0x0e, 0x43, 0x0d, 0x23, 0xd0, 0xf0, 0xe9, 0xb6, 0x2d, 0x40, 0x18, 0x1a, 0x42, 0x18, 0x5a,
	0x88, 0x30, 0x34, 0x0c, 0x3b, 0x1d, 0xd4, 0xbf, 0xdf, 0xda, 0xde, 0x57, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x80, 0x71, 0x84, 0x61, 0x75, 0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/assetmantle.modules.maintainers.queries.maintainer.Service/Handle", in, out, opts...)
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
		FullMethod: "/assetmantle.modules.maintainers.queries.maintainer.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assetmantle.modules.maintainers.queries.maintainer.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "x/maintainers/internal/queries/maintainer/service.proto",
}
