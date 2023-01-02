// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/assets/internal/transactions/revoke/service.v1.proto

package revoke

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

func init() {
	proto.RegisterFile("modules/assets/internal/transactions/revoke/service.v1.proto", fileDescriptor_e7c39db829a9fdd5)
}

var fileDescriptor_e7c39db829a9fdd5 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x31, 0x4b, 0x03, 0x41,
	0x10, 0x85, 0x73, 0x27, 0x44, 0xb8, 0xc2, 0x22, 0x20, 0xc2, 0x12, 0xb7, 0x88, 0x62, 0xb9, 0x4b,
	0xb4, 0x3b, 0xb4, 0x48, 0x1a, 0x45, 0x08, 0x84, 0x18, 0x2c, 0x24, 0xcd, 0xe4, 0x32, 0x9c, 0x87,
	0x77, 0xbb, 0x61, 0x77, 0x93, 0xce, 0xc6, 0x5f, 0x20, 0xd8, 0x5b, 0x58, 0xfa, 0x4b, 0xc4, 0x2a,
	0x60, 0x63, 0x29, 0x17, 0x2b, 0x7f, 0x85, 0x24, 0x93, 0x90, 0x6b, 0x12, 0x48, 0xfd, 0xbe, 0xc7,
	0x7b, 0xf3, 0x26, 0x38, 0xcf, 0xf4, 0x60, 0x94, 0xa2, 0x95, 0x60, 0x2d, 0x3a, 0x2b, 0x13, 0xe5,
	0xd0, 0x28, 0x48, 0xa5, 0x33, 0xa0, 0x2c, 0x44, 0x2e, 0xd1, 0xca, 0x4a, 0x83, 0x63, 0xfd, 0x80,
	0xd2, 0xa2, 0x19, 0x27, 0x11, 0x8a, 0x71, 0x5d, 0x0c, 0x8d, 0x76, 0xba, 0xc2, 0xc8, 0x25, 0x8a,
	0xb0, 0x20, 0x98, 0x55, 0x63, 0xad, 0xe3, 0x14, 0x25, 0x0c, 0x13, 0x09, 0x4a, 0x69, 0x07, 0x24,
	0xcf, 0x9d, 0x6c, 0xab, 0xdc, 0x0c, 0xad, 0x85, 0x78, 0x95, 0xcb, 0x2e, 0xb6, 0x71, 0x1b, 0xb4,
	0x43, 0xad, 0xec, 0xca, 0x7e, 0xfa, 0x18, 0xec, 0xde, 0xd0, 0x29, 0x15, 0x13, 0x94, 0xaf, 0x40,
	0x0d, 0x52, 0xac, 0x1c, 0x89, 0xf5, 0xc7, 0x88, 0x16, 0x35, 0x60, 0xc7, 0x9b, 0xa0, 0xce, 0x22,
	0xa8, 0x76, 0xf8, 0xf4, 0xf5, 0xfb, 0xe2, 0x1f, 0xd4, 0xf6, 0x65, 0x06, 0xca, 0xcd, 0x36, 0xa0,
	0x9a, 0xc4, 0x35, 0x5f, 0xfd, 0x8f, 0x9c, 0x7b, 0x93, 0x9c, 0x7b, 0x3f, 0x39, 0xf7, 0x9e, 0xa7,
	0xbc, 0x34, 0x99, 0xf2, 0xd2, 0xf7, 0x94, 0x97, 0x02, 0x1e, 0xe9, 0x6c, 0x43, 0x44, 0x73, 0x6f,
	0xd1, 0xfb, 0xb6, 0xde, 0x9e, 0x5d, 0xd2, 0xf6, 0xee, 0xae, 0xe3, 0xc4, 0xdd, 0x8f, 0xfa, 0x22,
	0xd2, 0x99, 0x6c, 0xcc, 0x8c, 0x2d, 0x8a, 0x5c, 0x2e, 0xb4, 0xc5, 0x52, 0x6f, 0xfe, 0x4e, 0xa3,
	0xdb, 0x79, 0xf7, 0x59, 0x83, 0x0a, 0x74, 0x8b, 0x05, 0x3a, 0x73, 0xe4, 0x73, 0x29, 0xf6, 0x8a,
	0x62, 0x8f, 0xc4, 0xdc, 0x3f, 0x59, 0x2f, 0xf6, 0x2e, 0xdb, 0xcd, 0x16, 0x3a, 0x18, 0x80, 0x83,
	0x3f, 0xbf, 0x4a, 0x60, 0x18, 0x16, 0xc9, 0x30, 0x24, 0xb4, 0x5f, 0x9e, 0xbf, 0xe9, 0xec, 0x3f,
	0x00, 0x00, 0xff, 0xff, 0xd7, 0x37, 0x34, 0xde, 0x9d, 0x02, 0x00, 0x00,
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
	Handle(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Response, error)
}

type serviceClient struct {
	cc grpc1.ClientConn
}

func NewServiceClient(cc grpc1.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Handle(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/assets.transactions.revoke.Service/Handle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	Handle(context.Context, *Message) (*Response, error)
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) Handle(ctx context.Context, req *Message) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handle not implemented")
}

func RegisterServiceServer(s grpc1.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_Handle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assets.transactions.revoke.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assets.transactions.revoke.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/assets/internal/transactions/revoke/service.v1.proto",
}