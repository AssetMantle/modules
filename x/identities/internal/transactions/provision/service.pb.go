// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/identities/internal/transactions/provision/service.proto

package provision

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
	proto.RegisterFile("x/identities/internal/transactions/provision/service.proto", fileDescriptor_c0111899f94208ff)
}

var fileDescriptor_c0111899f94208ff = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x31, 0x4b, 0xc3, 0x40,
	0x18, 0x86, 0x9b, 0x08, 0x15, 0x82, 0x53, 0xc7, 0x50, 0x0e, 0x5a, 0x75, 0x11, 0xbc, 0x03, 0x05,
	0x87, 0x6c, 0x76, 0x50, 0x8b, 0x14, 0x82, 0x76, 0x92, 0x2e, 0xd7, 0xe4, 0x23, 0x1e, 0x24, 0xf7,
	0x85, 0xdc, 0xd7, 0xe2, 0xec, 0x2f, 0x10, 0xdd, 0x1d, 0x1c, 0xfd, 0x15, 0x8e, 0xe2, 0x54, 0x70,
	0x71, 0x94, 0xd4, 0xc9, 0x5f, 0x21, 0xda, 0x9a, 0x1e, 0x38, 0xd4, 0xae, 0xc7, 0xf3, 0xdc, 0xfb,
	0xbe, 0x77, 0x5e, 0x70, 0x25, 0x54, 0x0c, 0x9a, 0x14, 0x29, 0x30, 0x42, 0x69, 0x82, 0x42, 0xcb,
	0x54, 0x50, 0x21, 0xb5, 0x91, 0x11, 0x29, 0xd4, 0x46, 0xe4, 0x05, 0x8e, 0x95, 0x51, 0xa8, 0x85,
	0x81, 0x62, 0xac, 0x22, 0xe0, 0x79, 0x81, 0x84, 0x8d, 0xd6, 0xc2, 0xe4, 0xb6, 0xc0, 0x2b, 0xc1,
	0x6f, 0x26, 0x88, 0x49, 0x0a, 0x42, 0xe6, 0x4a, 0x48, 0xad, 0x91, 0x64, 0x45, 0x10, 0xfa, 0xab,
	0x85, 0x67, 0x60, 0x8c, 0x4c, 0xe6, 0xe1, 0xfe, 0xd1, 0x4a, 0xae, 0x75, 0x7c, 0x06, 0x26, 0x47,
	0x6d, 0xe6, 0xf7, 0xec, 0xdd, 0x3b, 0xde, 0xfa, 0xf9, 0x6c, 0x56, 0xe3, 0xd6, 0xf1, 0xea, 0x27,
	0x52, 0xc7, 0x29, 0x34, 0x76, 0xf8, 0xd2, 0x71, 0xbc, 0x37, 0x2b, 0xe4, 0x1f, 0xfc, 0x83, 0xed,
	0xff, 0x2d, 0xd0, 0xde, 0xba, 0x7e, 0xfd, 0xb8, 0x73, 0x59, 0xbb, 0x29, 0x32, 0xa9, 0x29, 0x05,
	0x7b, 0x50, 0x65, 0x76, 0x9e, 0xdc, 0xe7, 0x92, 0x39, 0x93, 0x92, 0x39, 0xef, 0x25, 0x73, 0x6e,
	0xa6, 0xac, 0x36, 0x99, 0xb2, 0xda, 0xdb, 0x94, 0xd5, 0xbc, 0xed, 0x08, 0xb3, 0xe5, 0xd9, 0x9d,
	0x8d, 0xf9, 0xbe, 0xf0, 0x7b, 0x70, 0xe8, 0x5c, 0x9c, 0x26, 0x8a, 0x2e, 0x47, 0x43, 0x1e, 0x61,
	0x26, 0x0e, 0x8d, 0x01, 0xea, 0xcd, 0xe2, 0x33, 0x8c, 0x47, 0x29, 0x18, 0xb1, 0xca, 0xcb, 0x3e,
	0xb8, 0x6b, 0xdd, 0x7e, 0xf8, 0xe8, 0xb6, 0xba, 0x8b, 0x1a, 0x7d, 0xbb, 0x46, 0xf8, 0x4b, 0xbe,
	0xd8, 0xcc, 0xc0, 0x66, 0x06, 0x15, 0x53, 0xba, 0xbb, 0x4b, 0x99, 0xc1, 0x71, 0xd8, 0xe9, 0x01,
	0xc9, 0x58, 0x92, 0xfc, 0x74, 0x37, 0x17, 0x7c, 0x10, 0xd8, 0x42, 0x10, 0x54, 0xc6, 0xb0, 0xfe,
	0xf3, 0xd5, 0xfb, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x9c, 0x49, 0x26, 0xed, 0x02, 0x00,
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
	err := c.cc.Invoke(ctx, "/identities.transactions.provision.Service/Handle", in, out, opts...)
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
		FullMethod: "/identities.transactions.provision.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "identities.transactions.provision.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "x/identities/internal/transactions/provision/service.proto",
}
