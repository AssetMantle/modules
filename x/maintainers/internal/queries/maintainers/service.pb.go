// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/maintainers/internal/queries/maintainers/service.proto

package maintainers

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
	proto.RegisterFile("x/maintainers/internal/queries/maintainers/service.proto", fileDescriptor_b49540f3b3b07c28)
}

var fileDescriptor_b49540f3b3b07c28 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xb1, 0x4a, 0x03, 0x41,
	0x10, 0x86, 0x73, 0x27, 0x44, 0x38, 0xac, 0xae, 0x0c, 0xb2, 0x41, 0x2d, 0x44, 0xc4, 0x5d, 0xa2,
	0x8d, 0x1c, 0x28, 0x98, 0x46, 0x2d, 0x0e, 0x12, 0xed, 0x24, 0xcd, 0x24, 0x19, 0xe2, 0xc2, 0xdd,
	0x6e, 0x72, 0x3b, 0x17, 0xb4, 0xf5, 0x09, 0x04, 0x1b, 0x6b, 0xb1, 0xb2, 0xf6, 0x21, 0xc4, 0x2a,
	0x60, 0x63, 0x29, 0x17, 0x2b, 0x9f, 0x42, 0x62, 0xb6, 0xd8, 0xd8, 0x24, 0x69, 0xbf, 0xf9, 0x86,
	0xf9, 0xff, 0xdd, 0xe0, 0xf0, 0x46, 0xa4, 0x20, 0x15, 0x81, 0x54, 0x98, 0x19, 0x21, 0x15, 0x61,
	0xa6, 0x20, 0x11, 0x83, 0x1c, 0x33, 0x89, 0x66, 0x66, 0x68, 0x30, 0x1b, 0xca, 0x0e, 0xf2, 0x7e,
	0xa6, 0x49, 0x87, 0x55, 0x67, 0xc4, 0xad, 0xce, 0x1d, 0x56, 0x59, 0xef, 0x69, 0xdd, 0x4b, 0x50,
	0x40, 0x5f, 0x0a, 0x50, 0x4a, 0x13, 0x90, 0xd4, 0xca, 0x4c, 0xd7, 0x2b, 0x47, 0x4b, 0x1c, 0x9e,
	0xb0, 0xdb, 0x0b, 0x1c, 0xe4, 0x68, 0xc8, 0xae, 0x1f, 0x2f, 0xbf, 0x6e, 0xfa, 0x5a, 0x19, 0x9b,
	0x7e, 0xff, 0xd9, 0x0b, 0x56, 0x2f, 0xa7, 0x7d, 0xc2, 0x47, 0x2f, 0x28, 0x9f, 0x81, 0xea, 0x26,
	0x18, 0xee, 0xf1, 0x39, 0xad, 0x78, 0xd3, 0xc9, 0x52, 0xe1, 0x8b, 0xea, 0xd3, 0xdb, 0x9b, 0xe2,
	0xee, 0xe3, 0xfb, 0xc1, 0xdf, 0x09, 0xb7, 0x45, 0x0a, 0x8a, 0x12, 0x9c, 0x89, 0x39, 0xac, 0xb5,
	0x91, 0xa0, 0xe6, 0xb2, 0xfa, 0xab, 0xff, 0x56, 0x30, 0x6f, 0x54, 0x30, 0xef, 0xab, 0x60, 0xde,
	0xfd, 0x98, 0x95, 0x46, 0x63, 0x56, 0xfa, 0x1c, 0xb3, 0x52, 0xb0, 0xd5, 0xd1, 0xe9, 0xbc, 0xf3,
	0xf5, 0x35, 0xdb, 0xb1, 0x31, 0x29, 0xdd, 0xf0, 0xae, 0xce, 0x7b, 0x92, 0xae, 0xf3, 0x36, 0xef,
	0xe8, 0x54, 0x9c, 0x18, 0x83, 0x14, 0xdb, 0x1c, 0xba, 0x9b, 0x27, 0x68, 0xc4, 0xe2, 0xaf, 0xfa,
	0xe4, 0xaf, 0xc4, 0xcd, 0xf8, 0xc5, 0xaf, 0xc6, 0xff, 0xda, 0x4e, 0x22, 0x38, 0xec, 0x7d, 0xc6,
	0x68, 0x59, 0xa3, 0xe5, 0xb0, 0xc2, 0xdf, 0x9d, 0x63, 0xb4, 0x4e, 0x1b, 0xf5, 0x18, 0x09, 0xba,
	0x40, 0xf0, 0xe3, 0x6f, 0x38, 0x93, 0x28, 0xb2, 0x7a, 0x14, 0x39, 0xb4, 0x5d, 0xfe, 0xfb, 0xe4,
	0x83, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x56, 0x16, 0x03, 0x22, 0xde, 0x02, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/maintainers.queries.maintainers.Service/Handle", in, out, opts...)
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
		FullMethod: "/maintainers.queries.maintainers.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "maintainers.queries.maintainers.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "x/maintainers/internal/queries/maintainers/service.proto",
}