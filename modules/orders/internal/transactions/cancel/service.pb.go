// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/orders/internal/transactions/cancel/service.proto

package cancel

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
	proto.RegisterFile("modules/orders/internal/transactions/cancel/service.proto", fileDescriptor_311a812641b1e50e)
}

var fileDescriptor_311a812641b1e50e = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x31, 0x4b, 0x33, 0x31,
	0x1c, 0xc6, 0x7b, 0x79, 0xa1, 0x2f, 0x1c, 0x4e, 0x05, 0x11, 0x42, 0xcd, 0x50, 0xc1, 0x31, 0x01,
	0x9d, 0xbc, 0xcd, 0x16, 0x51, 0x84, 0xd2, 0xa2, 0x9d, 0xa4, 0xcb, 0xbf, 0xb9, 0x3f, 0xe7, 0x41,
	0x2e, 0x29, 0x49, 0xea, 0xe6, 0xe2, 0xe4, 0x28, 0xb8, 0x3a, 0x39, 0xfa, 0x49, 0xc4, 0xa9, 0xe0,
	0xe2, 0x28, 0x57, 0x27, 0x3f, 0x85, 0xd8, 0xb4, 0x70, 0x20, 0x2d, 0x74, 0xbd, 0xe7, 0xf7, 0xdc,
	0x2f, 0x4f, 0x12, 0x1f, 0x15, 0x26, 0x9d, 0x28, 0x74, 0xc2, 0xd8, 0x14, 0xad, 0x13, 0xb9, 0xf6,
	0x68, 0x35, 0x28, 0xe1, 0x2d, 0x68, 0x07, 0xd2, 0xe7, 0x46, 0x3b, 0x21, 0x41, 0x4b, 0x54, 0xc2,
	0xa1, 0xbd, 0xc9, 0x25, 0xf2, 0xb1, 0x35, 0xde, 0x34, 0x68, 0xa8, 0xf0, 0x2a, 0xc9, 0x03, 0x49,
	0x9b, 0x99, 0x31, 0x99, 0x42, 0x01, 0xe3, 0x5c, 0x80, 0xd6, 0xc6, 0x43, 0x88, 0xe7, 0x4d, 0xba,
	0x91, 0xb4, 0x40, 0xe7, 0x20, 0x5b, 0x48, 0xe9, 0xc9, 0x26, 0xd5, 0xca, 0xb7, 0x0b, 0x74, 0x63,
	0xa3, 0xdd, 0xe2, 0x37, 0x07, 0xf7, 0x51, 0xfc, 0xff, 0x32, 0xac, 0x69, 0xdc, 0xc6, 0xf5, 0x33,
	0xd0, 0xa9, 0xc2, 0xc6, 0x1e, 0x5f, 0x3d, 0x89, 0x77, 0xc3, 0x39, 0xa8, 0x58, 0x07, 0x0d, 0xfe,
	0x1a, 0x5b, 0xbb, 0x77, 0xef, 0x5f, 0x8f, 0x64, 0xa7, 0xb5, 0x2d, 0x0a, 0xd0, 0x5e, 0xe1, 0xf2,
	0xe8, 0xa1, 0xd2, 0x7e, 0x22, 0xaf, 0x25, 0x8b, 0xa6, 0x25, 0x8b, 0x3e, 0x4b, 0x16, 0x3d, 0xcc,
	0x58, 0x6d, 0x3a, 0x63, 0xb5, 0x8f, 0x19, 0xab, 0xc5, 0x4c, 0x9a, 0x62, 0x8d, 0xad, 0xbd, 0xb5,
	0x98, 0xd0, 0xff, 0xdd, 0xd4, 0x8f, 0xae, 0xce, 0xb3, 0xdc, 0x5f, 0x4f, 0x46, 0x5c, 0x9a, 0x42,
	0x1c, 0x3b, 0x87, 0xbe, 0x1b, 0x84, 0xcb, 0x3b, 0xdb, 0xe0, 0xee, 0x9e, 0xc9, 0xbf, 0xde, 0xa0,
	0xf3, 0x42, 0x68, 0x2f, 0xe8, 0x07, 0x55, 0x7d, 0x67, 0x8e, 0xbc, 0x2d, 0xc3, 0x61, 0x35, 0x1c,
	0x86, 0xb0, 0x24, 0xfb, 0xab, 0xc3, 0xe1, 0x69, 0xbf, 0xdd, 0x45, 0x0f, 0x29, 0x78, 0xf8, 0x26,
	0xcd, 0x00, 0x26, 0x49, 0x95, 0x4c, 0x92, 0x80, 0x8e, 0xea, 0xf3, 0x07, 0x3b, 0xfc, 0x09, 0x00,
	0x00, 0xff, 0xff, 0xe0, 0x4b, 0x82, 0x6c, 0xa9, 0x02, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/orders.transactions.cancel.Service/Handle", in, out, opts...)
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
		FullMethod: "/orders.transactions.cancel.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "orders.transactions.cancel.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/orders/internal/transactions/cancel/service.proto",
}