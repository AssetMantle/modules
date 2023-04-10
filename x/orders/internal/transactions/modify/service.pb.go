// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/orders/internal/transactions/modify/service.proto

package modify

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
	proto.RegisterFile("x/orders/internal/transactions/modify/service.proto", fileDescriptor_ffa8dcff470c04a5)
}

var fileDescriptor_ffa8dcff470c04a5 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x4b, 0xc3, 0x40,
	0x18, 0x86, 0x9b, 0x08, 0x15, 0x82, 0x53, 0x41, 0x84, 0x50, 0x6f, 0xa8, 0xe0, 0x98, 0x03, 0xbb,
	0x65, 0x11, 0x8b, 0xa0, 0x4b, 0x68, 0xd1, 0x4e, 0xd2, 0xe5, 0x6b, 0xf2, 0x19, 0x0f, 0x72, 0xf7,
	0x95, 0xbb, 0xab, 0xe8, 0xe0, 0xe2, 0xe4, 0x28, 0x08, 0xfe, 0x00, 0x47, 0x7f, 0x89, 0x38, 0x15,
	0x5c, 0x1c, 0x25, 0x75, 0xf2, 0x57, 0x88, 0xbd, 0x16, 0x02, 0xd2, 0xd2, 0x35, 0xef, 0xf3, 0xe6,
	0xf9, 0xde, 0x24, 0x68, 0xdf, 0x70, 0xd2, 0x19, 0x6a, 0xc3, 0x85, 0xb2, 0xa8, 0x15, 0x14, 0xdc,
	0x6a, 0x50, 0x06, 0x52, 0x2b, 0x48, 0x19, 0x2e, 0x29, 0x13, 0x97, 0xb7, 0xdc, 0xa0, 0xbe, 0x16,
	0x29, 0x46, 0x23, 0x4d, 0x96, 0x1a, 0xa1, 0xab, 0x44, 0x55, 0x32, 0x72, 0x64, 0xd8, 0xcc, 0x89,
	0xf2, 0x02, 0x39, 0x8c, 0x04, 0x07, 0xa5, 0xc8, 0x82, 0x8b, 0x67, 0xcd, 0x70, 0x4d, 0x9d, 0x44,
	0x63, 0x20, 0x9f, 0xeb, 0xc2, 0xc3, 0xf5, 0x4a, 0x95, 0x67, 0x67, 0x68, 0x46, 0xa4, 0xcc, 0xfc,
	0x05, 0x07, 0x0f, 0x5e, 0xb0, 0x79, 0xee, 0x16, 0x34, 0xee, 0x82, 0xfa, 0x29, 0xa8, 0xac, 0xc0,
	0xc6, 0x5e, 0xb4, 0x7c, 0x46, 0x94, 0xb8, 0x0b, 0x42, 0xbe, 0x0a, 0xea, 0xff, 0x37, 0xb6, 0x76,
	0xef, 0x3f, 0xbe, 0x9f, 0xfc, 0x9d, 0xd6, 0x36, 0x97, 0xa0, 0x6c, 0x81, 0x8b, 0xd3, 0x5d, 0xa5,
	0xf3, 0xec, 0xbf, 0x95, 0xcc, 0x9b, 0x94, 0xcc, 0xfb, 0x2a, 0x99, 0xf7, 0x38, 0x65, 0xb5, 0xc9,
	0x94, 0xd5, 0x3e, 0xa7, 0xac, 0x16, 0xb0, 0x94, 0xe4, 0x0a, 0x5b, 0x67, 0x6b, 0x3e, 0xa1, 0xf7,
	0xb7, 0xa9, 0xe7, 0x5d, 0x1c, 0xe7, 0xc2, 0x5e, 0x8d, 0x87, 0x51, 0x4a, 0x92, 0x1f, 0x19, 0x83,
	0x36, 0x71, 0x42, 0x49, 0xd9, 0xb8, 0x40, 0xc3, 0xd7, 0xfa, 0x6a, 0x2f, 0xfe, 0x46, 0xb7, 0x9f,
	0xbc, 0xfa, 0x61, 0xd7, 0x89, 0xfb, 0x55, 0x71, 0x32, 0x43, 0xde, 0x17, 0xe1, 0xa0, 0x1a, 0x0e,
	0x5c, 0x58, 0xfa, 0xfb, 0xcb, 0xc3, 0xc1, 0x49, 0xaf, 0x93, 0xa0, 0x85, 0x0c, 0x2c, 0xfc, 0xf8,
	0x4d, 0x07, 0xc6, 0x71, 0x95, 0x8c, 0x63, 0x87, 0x0e, 0xeb, 0xb3, 0x5f, 0xd5, 0xfe, 0x0d, 0x00,
	0x00, 0xff, 0xff, 0xae, 0xfe, 0xb1, 0x67, 0x91, 0x02, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/orders.transactions.modify.Service/Handle", in, out, opts...)
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
		FullMethod: "/orders.transactions.modify.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "orders.transactions.modify.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "x/orders/internal/transactions/modify/service.proto",
}
