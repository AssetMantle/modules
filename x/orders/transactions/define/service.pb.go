// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: orders/transactions/define/service.proto

package define

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
	proto.RegisterFile("orders/transactions/define/service.proto", fileDescriptor_7538497e79fca1c1)
}

var fileDescriptor_7538497e79fca1c1 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xd2, 0x31, 0x4b, 0xc3, 0x40,
	0x14, 0x07, 0xf0, 0x26, 0x60, 0x85, 0xe0, 0x54, 0x10, 0x21, 0x68, 0x86, 0x4e, 0x4e, 0x77, 0x50,
	0x11, 0xe1, 0x9c, 0x5a, 0x0b, 0xba, 0x84, 0x16, 0xed, 0x24, 0x01, 0xb9, 0x26, 0xcf, 0x18, 0x48,
	0xee, 0x4a, 0xee, 0x2a, 0xce, 0x7e, 0x02, 0xc1, 0x6f, 0x20, 0xb8, 0xb8, 0xf9, 0x2d, 0xc4, 0xa9,
	0xe0, 0xe2, 0x28, 0xa9, 0x93, 0xbb, 0xbb, 0xb4, 0xef, 0xc0, 0x73, 0xb0, 0x90, 0xf5, 0x92, 0xdf,
	0xff, 0xfe, 0xef, 0x71, 0xde, 0xae, 0x2c, 0x13, 0x28, 0x15, 0xd5, 0x25, 0x17, 0x8a, 0xc7, 0x3a,
	0x93, 0x42, 0xd1, 0x04, 0x2e, 0x33, 0x01, 0x54, 0x41, 0x79, 0x9d, 0xc5, 0x40, 0x26, 0xa5, 0xd4,
	0xb2, 0x45, 0xb8, 0x52, 0xa0, 0x0b, 0x2e, 0x74, 0x0e, 0xa4, 0x90, 0xc9, 0x34, 0x07, 0x45, 0x50,
	0x13, 0x5b, 0x13, 0xd4, 0xfe, 0x76, 0x2a, 0x65, 0x9a, 0x03, 0xe5, 0x93, 0x8c, 0x72, 0x21, 0xa4,
	0xe6, 0xf8, 0x79, 0x99, 0xe6, 0xaf, 0xba, 0xb7, 0x00, 0xa5, 0x78, 0x6a, 0xee, 0xf5, 0xf7, 0x57,
	0xfc, 0x69, 0x9d, 0x5d, 0x94, 0xa0, 0x26, 0x52, 0x28, 0xc3, 0x3a, 0xcf, 0x8e, 0xb7, 0x7e, 0x86,
	0x03, 0xb4, 0x1e, 0x1d, 0xaf, 0x79, 0xc2, 0x45, 0x92, 0x43, 0xeb, 0xa0, 0xe6, 0x18, 0x24, 0xc4,
	0x32, 0xfe, 0x51, 0x5d, 0x38, 0xfa, 0x3d, 0x3b, 0x35, 0xd5, 0xda, 0x3b, 0xb7, 0x6f, 0x9f, 0xf7,
	0xee, 0x56, 0x7b, 0x93, 0x62, 0x0e, 0x35, 0xb3, 0x21, 0xe9, 0x7d, 0xbb, 0x2f, 0x55, 0xe0, 0xcc,
	0xaa, 0xc0, 0xf9, 0xa8, 0x02, 0xe7, 0x6e, 0x1e, 0x34, 0x66, 0xf3, 0xa0, 0xf1, 0x3e, 0x0f, 0x1a,
	0x5e, 0x27, 0x96, 0x45, 0xcd, 0x06, 0xbd, 0x0d, 0x33, 0xff, 0x70, 0xb1, 0x90, 0xa1, 0x73, 0x7e,
	0x98, 0x66, 0xfa, 0x6a, 0x3a, 0x26, 0xb1, 0x2c, 0x68, 0x77, 0x11, 0x15, 0x62, 0x09, 0x13, 0x45,
	0x6f, 0xe8, 0xff, 0xab, 0x7e, 0x70, 0xd7, 0xba, 0xe1, 0x60, 0xd4, 0x7f, 0x72, 0x49, 0xd7, 0x6a,
	0x11, 0x9a, 0x16, 0x03, 0x6c, 0x31, 0xb2, 0x5b, 0xf4, 0x97, 0xf0, 0xf5, 0x0f, 0x88, 0x0c, 0x88,
	0x10, 0x44, 0x36, 0x88, 0x10, 0x54, 0x2e, 0xab, 0x07, 0xa2, 0xe3, 0x61, 0x2f, 0x04, 0xcd, 0x13,
	0xae, 0xf9, 0x97, 0xdb, 0xb1, 0x30, 0x63, 0x46, 0x33, 0x86, 0x9c, 0x31, 0xdb, 0x33, 0x86, 0x01,
	0xe3, 0xe6, 0xf2, 0xc9, 0xec, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff, 0x28, 0x72, 0xac, 0x65, 0x0d,
	0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/assetmantle.modules.orders.transactions.define.Service/Handle", in, out, opts...)
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
		FullMethod: "/assetmantle.modules.orders.transactions.define.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assetmantle.modules.orders.transactions.define.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orders/transactions/define/service.proto",
}
