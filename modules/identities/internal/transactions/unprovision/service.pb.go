// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/identities/internal/transactions/unprovision/service.proto

package unprovision

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
	proto.RegisterFile("modules/identities/internal/transactions/unprovision/service.proto", fileDescriptor_50c445e015ae4a61)
}

var fileDescriptor_50c445e015ae4a61 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x31, 0x4b, 0xf3, 0x40,
	0x18, 0xc7, 0x9b, 0xbc, 0xd0, 0x17, 0x82, 0x53, 0xc7, 0x20, 0x87, 0xb4, 0xa8, 0x8b, 0xe4, 0x40,
	0x17, 0xc9, 0x66, 0x16, 0xed, 0x50, 0x29, 0xb6, 0x5d, 0xa4, 0xcb, 0x35, 0x39, 0xe2, 0x41, 0xf2,
	0x3c, 0x21, 0xf7, 0xb4, 0x1f, 0xc0, 0x4f, 0x20, 0xb8, 0x38, 0x8b, 0x93, 0x9f, 0x44, 0x9c, 0x8a,
	0x2e, 0x8e, 0x92, 0x3a, 0xf9, 0x29, 0xa4, 0x4d, 0x25, 0x07, 0x2e, 0x21, 0xdb, 0xc1, 0xfd, 0x7f,
	0xcf, 0xfd, 0xfe, 0x0f, 0xe7, 0x04, 0x29, 0x46, 0xf3, 0x44, 0x6a, 0xae, 0x22, 0x09, 0xa4, 0x48,
	0xad, 0x8f, 0x40, 0x32, 0x07, 0x91, 0x70, 0xca, 0x05, 0x68, 0x11, 0x92, 0x42, 0xd0, 0x7c, 0x0e,
	0x59, 0x8e, 0x0b, 0xa5, 0x15, 0x02, 0xd7, 0x32, 0x5f, 0xa8, 0x50, 0x7a, 0x59, 0x8e, 0x84, 0x9d,
	0x5e, 0xc5, 0x7a, 0x26, 0xe2, 0x19, 0x88, 0xbb, 0x1b, 0x23, 0xc6, 0x89, 0xe4, 0x22, 0x53, 0x5c,
	0x00, 0x20, 0x89, 0x32, 0xb3, 0x19, 0xe1, 0x36, 0xd3, 0x48, 0xa5, 0xd6, 0x22, 0xde, 0x6a, 0xb8,
	0x97, 0x8d, 0x66, 0x18, 0x17, 0x57, 0x52, 0x67, 0x08, 0x7a, 0x3b, 0xef, 0xf8, 0xc9, 0x72, 0xfe,
	0x8f, 0xca, 0xa2, 0x9d, 0x07, 0xcb, 0x69, 0x5f, 0x08, 0x88, 0x12, 0xd9, 0x39, 0xf2, 0x6a, 0xd4,
	0xf5, 0x06, 0xa5, 0x9a, 0x7b, 0x5a, 0x2b, 0x3d, 0xfe, 0x2b, 0xd1, 0x3d, 0xb8, 0x7d, 0xff, 0xba,
	0xb7, 0xf7, 0xba, 0x8c, 0xa7, 0x02, 0x28, 0x91, 0x66, 0x2d, 0x83, 0x0d, 0xde, 0xec, 0x97, 0x82,
	0x59, 0xcb, 0x82, 0x59, 0x9f, 0x05, 0xb3, 0xee, 0x56, 0xac, 0xb5, 0x5c, 0xb1, 0xd6, 0xc7, 0x8a,
	0xb5, 0x9c, 0xc3, 0x10, 0xd3, 0x3a, 0xef, 0x07, 0x3b, 0xdb, 0x9e, 0xc3, 0x75, 0xf1, 0xa1, 0x75,
	0x3d, 0x8a, 0x15, 0xdd, 0xcc, 0x67, 0x5e, 0x88, 0x29, 0x3f, 0xd3, 0x5a, 0xd2, 0xa0, 0x54, 0xf8,
	0xdd, 0x70, 0x93, 0x4d, 0x3f, 0xda, 0xff, 0xfa, 0xe3, 0xc9, 0xb3, 0xdd, 0xeb, 0x57, 0x42, 0x63,
	0x53, 0x68, 0x52, 0x65, 0x5f, 0xcd, 0xd4, 0xd4, 0x4c, 0x4d, 0x8d, 0x54, 0x61, 0xf3, 0x1a, 0xa9,
	0xe9, 0xf9, 0x30, 0x18, 0x48, 0x12, 0x91, 0x20, 0xf1, 0x6d, 0xef, 0x57, 0x84, 0xef, 0x9b, 0x88,
	0xef, 0x1b, 0xcc, 0xac, 0xbd, 0xf9, 0x02, 0x27, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x90, 0xf5,
	0xe7, 0x4e, 0x1f, 0x03, 0x00, 0x00,
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
	Handle(ctx context.Context, in *Message, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type serviceClient struct {
	cc grpc1.ClientConn
}

func NewServiceClient(cc grpc1.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Handle(ctx context.Context, in *Message, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/identities.transactions.unprovision.Service/Handle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	Handle(context.Context, *Message) (*TransactionResponse, error)
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) Handle(ctx context.Context, req *Message) (*TransactionResponse, error) {
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
		FullMethod: "/identities.transactions.unprovision.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "identities.transactions.unprovision.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/identities/internal/transactions/unprovision/service.proto",
}