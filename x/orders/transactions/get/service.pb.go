// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: orders/transactions/get/service.proto

package get

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
	proto.RegisterFile("orders/transactions/get/service.proto", fileDescriptor_2a3461b1dbdc9864)
}

var fileDescriptor_2a3461b1dbdc9864 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x31, 0x4b, 0xf3, 0x40,
	0x1c, 0xc6, 0x9b, 0xc0, 0xdb, 0x17, 0xc2, 0x3b, 0x85, 0x77, 0x0a, 0x92, 0xa1, 0xe0, 0x54, 0xb8,
	0x83, 0x2a, 0x28, 0x37, 0xd9, 0x2e, 0x71, 0x09, 0x2d, 0xda, 0x49, 0x02, 0x72, 0x4d, 0xfe, 0xc4,
	0x40, 0x72, 0x57, 0x72, 0xff, 0x8a, 0xb3, 0x9f, 0x40, 0x70, 0x75, 0x72, 0x12, 0x3f, 0x89, 0x38,
	0x15, 0x5c, 0x1c, 0x25, 0x71, 0xf2, 0x53, 0x48, 0x7b, 0x07, 0x9e, 0x43, 0x87, 0xac, 0x97, 0xe7,
	0xf7, 0xfc, 0x9e, 0x3b, 0xe2, 0xed, 0xcb, 0x3a, 0x83, 0x5a, 0x51, 0xac, 0xb9, 0x50, 0x3c, 0xc5,
	0x42, 0x0a, 0x45, 0x73, 0x40, 0xaa, 0xa0, 0xbe, 0x2e, 0x52, 0x20, 0xcb, 0x5a, 0xa2, 0xf4, 0x87,
	0x5c, 0x29, 0xc0, 0x8a, 0x0b, 0x2c, 0x81, 0x54, 0x32, 0x5b, 0x95, 0xa0, 0x88, 0x46, 0x89, 0x8d,
	0x92, 0x1c, 0x30, 0xd8, 0xcb, 0xa5, 0xcc, 0x4b, 0xa0, 0x7c, 0x59, 0x50, 0x2e, 0x84, 0x44, 0xae,
	0xbf, 0x6d, 0xab, 0x82, 0x9d, 0xc6, 0x0a, 0x94, 0xe2, 0xb9, 0x31, 0x06, 0xa3, 0x5d, 0x31, 0xeb,
	0xe0, 0xb2, 0x06, 0xb5, 0x94, 0x42, 0x19, 0x66, 0xf4, 0xe4, 0x78, 0x7f, 0xcf, 0xf5, 0x6e, 0xff,
	0xc1, 0xf1, 0xfa, 0xa7, 0x5c, 0x64, 0x25, 0xf8, 0x87, 0xa4, 0xc3, 0x7a, 0x12, 0xeb, 0x19, 0xc1,
	0x49, 0x27, 0x6a, 0xfe, 0x73, 0x70, 0x66, 0x46, 0x0d, 0x82, 0xdb, 0xb7, 0xcf, 0x7b, 0xf7, 0xff,
	0xc0, 0xa7, 0xba, 0x84, 0x9a, 0x2b, 0xe5, 0x80, 0x93, 0xd6, 0x7d, 0x69, 0x42, 0x67, 0xdd, 0x84,
	0xce, 0x47, 0x13, 0x3a, 0x77, 0x6d, 0xd8, 0x5b, 0xb7, 0x61, 0xef, 0xbd, 0x0d, 0x7b, 0x1e, 0x4d,
	0x65, 0xd5, 0xc5, 0x3d, 0xf9, 0x67, 0xee, 0x3c, 0xdb, 0x3c, 0xc2, 0xcc, 0xb9, 0x38, 0xce, 0x0b,
	0xbc, 0x5a, 0x2d, 0x48, 0x2a, 0x2b, 0x3a, 0xde, 0xf4, 0xc4, 0x5a, 0x6f, 0x7a, 0xe8, 0x0d, 0xdd,
	0xf1, 0xb6, 0x8f, 0xee, 0x9f, 0x71, 0x3c, 0x9d, 0x47, 0xcf, 0xee, 0x70, 0x6c, 0xf9, 0x63, 0xe3,
	0x9f, 0x6a, 0xff, 0xdc, 0xf6, 0x47, 0x80, 0xaf, 0xbf, 0xd2, 0x89, 0x49, 0x27, 0x3a, 0x9d, 0xd8,
	0xe9, 0x24, 0x02, 0x6c, 0xdc, 0xa3, 0x0e, 0xe9, 0x24, 0x9a, 0x4d, 0x62, 0x40, 0x9e, 0x71, 0xe4,
	0x5f, 0x2e, 0xb5, 0x48, 0xc6, 0x0c, 0xca, 0x98, 0x66, 0x19, 0xb3, 0x61, 0xc6, 0x22, 0xc0, 0x45,
	0x7f, 0xfb, 0x5f, 0x1c, 0x7c, 0x07, 0x00, 0x00, 0xff, 0xff, 0x24, 0x9f, 0xe5, 0x20, 0xe6, 0x02,
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
	err := c.cc.Invoke(ctx, "/assetmantle.modules.orders.transactions.get.Service/Handle", in, out, opts...)
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
		FullMethod: "/assetmantle.modules.orders.transactions.get.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assetmantle.modules.orders.transactions.get.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orders/transactions/get/service.proto",
}
