// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/identities/internal/queries/identity/service.proto

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
	proto.RegisterFile("x/identities/internal/queries/identity/service.proto", fileDescriptor_dd6933fe3afc0ee1)
}

var fileDescriptor_dd6933fe3afc0ee1 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xb1, 0x4b, 0xc3, 0x40,
	0x18, 0xc5, 0x9b, 0x13, 0x2a, 0x04, 0xa7, 0x8c, 0xad, 0x9e, 0xa0, 0x88, 0xa8, 0x70, 0x47, 0xd5,
	0xc5, 0x6c, 0x16, 0x44, 0x3b, 0x14, 0x5a, 0xdd, 0xa4, 0xcb, 0xb5, 0xf9, 0xa8, 0x07, 0xc9, 0x5d,
	0x9b, 0xfb, 0x52, 0xec, 0xea, 0xec, 0x20, 0xb8, 0x38, 0x38, 0x39, 0xfa, 0x97, 0x88, 0x53, 0xc1,
	0xc5, 0x51, 0x52, 0x27, 0xff, 0x0a, 0x69, 0x73, 0xb5, 0x9d, 0x4a, 0x5c, 0xbf, 0xf7, 0x1e, 0xbf,
	0xf7, 0xee, 0xdc, 0xe3, 0x5b, 0x2e, 0x03, 0x50, 0x28, 0x51, 0x82, 0xe1, 0x52, 0x21, 0xc4, 0x4a,
	0x84, 0xbc, 0x9f, 0x40, 0x3c, 0x3d, 0x64, 0xda, 0x90, 0x1b, 0x88, 0x07, 0xb2, 0x03, 0xac, 0x17,
	0x6b, 0xd4, 0x5e, 0x79, 0x9e, 0x61, 0xd6, 0xca, 0x66, 0xd6, 0xd2, 0x7a, 0x57, 0xeb, 0x6e, 0x08,
	0x5c, 0xf4, 0x24, 0x17, 0x4a, 0x69, 0x14, 0x28, 0xb5, 0x32, 0x59, 0xb4, 0x74, 0x92, 0x13, 0x38,
	0x39, 0x0c, 0x2f, 0xa1, 0x9f, 0x80, 0x41, 0x1b, 0xf5, 0xff, 0x17, 0x35, 0x3d, 0xad, 0x8c, 0x6d,
	0x7c, 0xf8, 0xe4, 0xb8, 0xab, 0x57, 0xd9, 0x06, 0xef, 0xde, 0x71, 0x8b, 0x17, 0x42, 0x05, 0x21,
	0x78, 0x7b, 0x6c, 0xc9, 0x12, 0xd6, 0x5c, 0xe8, 0x50, 0xda, 0xcf, 0x63, 0xcd, 0x98, 0x5b, 0x07,
	0x77, 0x1f, 0xdf, 0x8f, 0x64, 0xc7, 0xdb, 0xe6, 0x91, 0x50, 0x18, 0xc2, 0x62, 0xef, 0x41, 0xa5,
	0x0d, 0x28, 0x2a, 0x7f, 0x75, 0xab, 0xcf, 0xe4, 0x2d, 0xa5, 0xce, 0x28, 0xa5, 0xce, 0x57, 0x4a,
	0x9d, 0x87, 0x31, 0x2d, 0x8c, 0xc6, 0xb4, 0xf0, 0x39, 0xa6, 0x05, 0x77, 0xb3, 0xa3, 0xa3, 0x65,
	0xd8, 0xea, 0x9a, 0xdd, 0xd4, 0x98, 0x8c, 0x6c, 0x38, 0xd7, 0x67, 0x5d, 0x89, 0x37, 0x49, 0x9b,
	0x75, 0x74, 0xc4, 0x4f, 0x8d, 0x01, 0xac, 0x67, 0xfc, 0x48, 0x07, 0x49, 0x08, 0x86, 0xe7, 0x7b,
	0xc1, 0x17, 0xb2, 0x52, 0x6b, 0xd6, 0x5e, 0x49, 0xb9, 0x36, 0x47, 0x37, 0x2d, 0xda, 0x9e, 0x86,
	0xef, 0x8b, 0x6a, 0xcb, 0xaa, 0xad, 0x99, 0x9a, 0x92, 0xdd, 0x25, 0x6a, 0xeb, 0xbc, 0x51, 0xad,
	0x03, 0x8a, 0x40, 0xa0, 0xf8, 0x21, 0x1b, 0x73, 0xa7, 0xef, 0x5b, 0xab, 0xef, 0xcf, 0xbc, 0xed,
	0xe2, 0xf4, 0x03, 0x8f, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x6c, 0xe7, 0x46, 0x3f, 0xaa, 0x02,
	0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/identities.queries.identity.Service/Handle", in, out, opts...)
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
		FullMethod: "/identities.queries.identity.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "identities.queries.identity.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "x/identities/internal/queries/identity/service.proto",
}
