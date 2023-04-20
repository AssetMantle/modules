// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/orders/internal/transactions/define/service.proto

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
	proto.RegisterFile("x/orders/internal/transactions/define/service.proto", fileDescriptor_08b33a9ed8d87fa6)
}

var fileDescriptor_08b33a9ed8d87fa6 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x4b, 0xc3, 0x40,
	0x18, 0x86, 0x9b, 0x13, 0x2a, 0x04, 0xa7, 0x82, 0x08, 0x47, 0xbd, 0xa1, 0x82, 0xe3, 0x1d, 0xd8,
	0x2d, 0x8b, 0x58, 0x0a, 0xba, 0x94, 0x16, 0xed, 0x24, 0x5d, 0xbe, 0x26, 0x9f, 0xf1, 0x20, 0xb9,
	0x2b, 0x77, 0x57, 0x71, 0x71, 0x71, 0x72, 0x14, 0x04, 0x7f, 0x80, 0xa3, 0xbf, 0x44, 0x9c, 0x0a,
	0x2e, 0x8e, 0x92, 0x3a, 0xf9, 0x2b, 0xc4, 0x5e, 0x0b, 0x01, 0x69, 0xc9, 0x9a, 0xf7, 0x79, 0xf3,
	0x7c, 0x6f, 0x12, 0xb6, 0x6f, 0x85, 0x36, 0x09, 0x1a, 0x2b, 0xa4, 0x72, 0x68, 0x14, 0x64, 0xc2,
	0x19, 0x50, 0x16, 0x62, 0x27, 0xb5, 0xb2, 0x22, 0xc1, 0x2b, 0xa9, 0x50, 0x58, 0x34, 0x37, 0x32,
	0x46, 0x3e, 0x31, 0xda, 0xe9, 0x06, 0xf5, 0x15, 0x5e, 0x26, 0xb9, 0x27, 0x69, 0x33, 0xd5, 0x3a,
	0xcd, 0x50, 0xc0, 0x44, 0x0a, 0x50, 0x4a, 0x3b, 0xf0, 0xf1, 0xa2, 0x49, 0x2b, 0xea, 0x72, 0xb4,
	0x16, 0xd2, 0xa5, 0x8e, 0x1e, 0x57, 0x2b, 0x95, 0x9e, 0x9d, 0xa3, 0x9d, 0x68, 0x65, 0x97, 0x2f,
	0x38, 0x7a, 0x08, 0xc2, 0xed, 0x0b, 0xbf, 0xa0, 0x71, 0x17, 0xd6, 0xcf, 0x40, 0x25, 0x19, 0x36,
	0x0e, 0xf8, 0xfa, 0x19, 0xbc, 0xe7, 0x2f, 0xa0, 0x62, 0x13, 0x34, 0xfc, 0x6f, 0x6c, 0xed, 0xdf,
	0x7f, 0x7c, 0x3f, 0x91, 0xbd, 0xd6, 0xae, 0xc8, 0x41, 0xb9, 0x0c, 0x57, 0xa7, 0xfb, 0x4a, 0xe7,
	0x99, 0xbc, 0x15, 0x2c, 0x98, 0x15, 0x2c, 0xf8, 0x2a, 0x58, 0xf0, 0x38, 0x67, 0xb5, 0xd9, 0x9c,
	0xd5, 0x3e, 0xe7, 0xac, 0x16, 0xb2, 0x58, 0xe7, 0x1b, 0x6c, 0x9d, 0x9d, 0xe5, 0x84, 0xc1, 0xdf,
	0xa6, 0x41, 0x70, 0xd9, 0x4d, 0xa5, 0xbb, 0x9e, 0x8e, 0x79, 0xac, 0x73, 0x71, 0x62, 0x2d, 0xba,
	0x9e, 0x17, 0xe6, 0x3a, 0x99, 0x66, 0x68, 0x45, 0xa5, 0xaf, 0xf6, 0x42, 0xb6, 0xfa, 0xc3, 0xee,
	0x2b, 0xa1, 0x7d, 0x2f, 0x1e, 0x96, 0xc5, 0xdd, 0x05, 0xf2, 0xbe, 0x0a, 0x47, 0xe5, 0x70, 0xe4,
	0xc3, 0x82, 0x1c, 0xae, 0x0f, 0x47, 0xa7, 0x83, 0x4e, 0x0f, 0x1d, 0x24, 0xe0, 0xe0, 0x87, 0x34,
	0x3d, 0x18, 0x45, 0x65, 0x32, 0x8a, 0x3c, 0x3a, 0xae, 0x2f, 0x7e, 0x55, 0xfb, 0x37, 0x00, 0x00,
	0xff, 0xff, 0x88, 0x3e, 0x55, 0x7d, 0x91, 0x02, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/orders.transactions.define.Service/Handle", in, out, opts...)
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
		FullMethod: "/orders.transactions.define.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "orders.transactions.define.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "x/orders/internal/transactions/define/service.proto",
}
