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
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x31, 0x4b, 0xf3, 0x40,
	0x18, 0xc7, 0x9b, 0xbc, 0xbc, 0x1d, 0xc2, 0x3b, 0x85, 0x77, 0x0a, 0x92, 0xa1, 0xe0, 0x54, 0xb8,
	0x83, 0x2a, 0x28, 0x37, 0xd9, 0x2e, 0x71, 0x09, 0x2d, 0xda, 0x49, 0x02, 0x72, 0x4d, 0x1e, 0xce,
	0x40, 0x92, 0x2b, 0xb9, 0xa7, 0xe2, 0xec, 0x27, 0x10, 0x5c, 0x9d, 0xc4, 0xc9, 0x4f, 0x22, 0x4e,
	0x05, 0x17, 0x47, 0x49, 0x9c, 0xfc, 0x14, 0xd2, 0xde, 0x81, 0xe7, 0xd0, 0x21, 0xeb, 0x93, 0xff,
	0xef, 0xff, 0x7b, 0xf2, 0x24, 0xde, 0xbe, 0xac, 0x33, 0xa8, 0x15, 0xc5, 0x9a, 0x57, 0x8a, 0xa7,
	0x98, 0xcb, 0x4a, 0x51, 0x01, 0x48, 0x15, 0xd4, 0xd7, 0x79, 0x0a, 0x64, 0x59, 0x4b, 0x94, 0xfe,
	0x90, 0x2b, 0x05, 0x58, 0xf2, 0x0a, 0x0b, 0x20, 0xa5, 0xcc, 0x56, 0x05, 0x28, 0xa2, 0x51, 0x62,
	0xa3, 0x44, 0x00, 0x06, 0x7b, 0x42, 0x4a, 0x51, 0x00, 0xe5, 0xcb, 0x9c, 0xf2, 0xaa, 0x92, 0xc8,
	0xf5, 0xb3, 0x6d, 0x55, 0xb0, 0xd3, 0x58, 0x82, 0x52, 0x5c, 0x18, 0x63, 0x30, 0xda, 0x15, 0xb3,
	0x06, 0x97, 0x35, 0xa8, 0xa5, 0xac, 0x94, 0x61, 0x46, 0x4f, 0x8e, 0xf7, 0x27, 0x56, 0xc2, 0x7f,
	0x70, 0xbc, 0xfe, 0x29, 0xaf, 0xb2, 0x02, 0xfc, 0x43, 0xd2, 0x61, 0x73, 0x12, 0xeb, 0x15, 0x82,
	0x93, 0x4e, 0xd4, 0xfc, 0x67, 0x70, 0x66, 0x16, 0x1a, 0x04, 0xb7, 0x6f, 0x9f, 0xf7, 0xee, 0xff,
	0x81, 0x4f, 0x75, 0x09, 0x35, 0xaf, 0x23, 0x00, 0x27, 0xad, 0xfb, 0xd2, 0x84, 0xce, 0xba, 0x09,
	0x9d, 0x8f, 0x26, 0x74, 0xee, 0xda, 0xb0, 0xb7, 0x6e, 0xc3, 0xde, 0x7b, 0x1b, 0xf6, 0x3c, 0x9a,
	0xca, 0xb2, 0x8b, 0x7b, 0xf2, 0xef, 0x5c, 0x7f, 0xa7, 0xd9, 0xe6, 0x00, 0x33, 0xe7, 0xe2, 0x58,
	0xe4, 0x78, 0xb5, 0x5a, 0x90, 0x54, 0x96, 0x74, 0xbc, 0xe9, 0x89, 0xb5, 0xde, 0xf4, 0xd0, 0x1b,
	0xba, 0xe3, 0xae, 0x8f, 0xee, 0xdf, 0x71, 0x3c, 0x9d, 0x47, 0xcf, 0xee, 0x70, 0x6c, 0xf9, 0x63,
	0xe3, 0x9f, 0x6a, 0xff, 0xdc, 0xf6, 0x47, 0x80, 0xaf, 0xbf, 0xd2, 0x89, 0x49, 0x27, 0x3a, 0x9d,
	0xd8, 0xe9, 0x24, 0x02, 0x6c, 0xdc, 0xa3, 0x0e, 0xe9, 0x24, 0x9a, 0x4d, 0x62, 0x40, 0x9e, 0x71,
	0xe4, 0x5f, 0x2e, 0xb5, 0x48, 0xc6, 0x0c, 0xca, 0x98, 0x66, 0x19, 0xb3, 0x61, 0xc6, 0x22, 0xc0,
	0x45, 0x7f, 0xfb, 0x4f, 0x1c, 0x7c, 0x07, 0x00, 0x00, 0xff, 0xff, 0xa0, 0xea, 0xde, 0x7b, 0xe2,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	Handle(ctx context.Context, in *Message, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Handle(ctx context.Context, in *Message, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/assetmantle.modules.orders.transactions.get.Msg/Handle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Handle(context.Context, *Message) (*TransactionResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Handle(ctx context.Context, req *Message) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handle not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Handle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assetmantle.modules.orders.transactions.get.Msg/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assetmantle.modules.orders.transactions.get.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Msg_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orders/transactions/get/service.proto",
}
