// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: orders/transactions/cancel/service.proto

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
	proto.RegisterFile("orders/transactions/cancel/service.proto", fileDescriptor_af989d8259f117c5)
}

var fileDescriptor_af989d8259f117c5 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xd2, 0x31, 0x4b, 0xc3, 0x40,
	0x14, 0x07, 0xf0, 0x26, 0x62, 0x87, 0xe0, 0x54, 0x10, 0x21, 0x68, 0x86, 0x4e, 0x4e, 0x77, 0x50,
	0x11, 0xe1, 0x9c, 0xda, 0x0e, 0xba, 0x84, 0x16, 0xed, 0x24, 0x01, 0x79, 0x4d, 0x8e, 0x18, 0x48,
	0xee, 0x4a, 0xee, 0x2a, 0xce, 0x7e, 0x02, 0xc1, 0x6f, 0x20, 0xb8, 0x08, 0x7e, 0x0f, 0x71, 0x2a,
	0xb8, 0x38, 0x4a, 0xea, 0xe4, 0xee, 0x2e, 0xcd, 0x3b, 0xf0, 0x1c, 0x2c, 0x64, 0xbd, 0xe4, 0xf7,
	0xbf, 0xff, 0x7b, 0x89, 0xb7, 0x2f, 0xcb, 0x84, 0x97, 0x8a, 0xea, 0x12, 0x84, 0x82, 0x58, 0x67,
	0x52, 0x28, 0x1a, 0x83, 0x88, 0x79, 0x4e, 0x15, 0x2f, 0xaf, 0xb3, 0x98, 0x93, 0x59, 0x29, 0xb5,
	0xec, 0x10, 0x50, 0x8a, 0xeb, 0x02, 0x84, 0xce, 0x39, 0x29, 0x64, 0x32, 0xcf, 0xb9, 0x22, 0xa8,
	0x89, 0xad, 0x09, 0x6a, 0x7f, 0x37, 0x95, 0x32, 0xcd, 0x39, 0x85, 0x59, 0x46, 0x41, 0x08, 0xa9,
	0x01, 0x1f, 0xd7, 0x69, 0xfe, 0xba, 0x7b, 0x0b, 0xae, 0x14, 0xa4, 0xe6, 0x5e, 0xff, 0x70, 0xcd,
	0x9b, 0xd6, 0xd9, 0x65, 0xc9, 0xd5, 0x4c, 0x0a, 0x65, 0x58, 0xef, 0xd9, 0xf1, 0x36, 0x42, 0x95,
	0x76, 0x1e, 0x1d, 0xaf, 0x7d, 0x0a, 0x22, 0xc9, 0x79, 0xe7, 0xa8, 0xe1, 0x08, 0x24, 0xc4, 0x22,
	0xfe, 0xb0, 0x29, 0x9c, 0xfc, 0x9e, 0x9d, 0x99, 0x5a, 0xdd, 0xbd, 0xdb, 0xb7, 0xcf, 0x7b, 0x77,
	0xa7, 0xbb, 0x4d, 0x31, 0x87, 0x9a, 0xb9, 0x90, 0x0c, 0xbe, 0xdd, 0x97, 0x2a, 0x70, 0x16, 0x55,
	0xe0, 0x7c, 0x54, 0x81, 0x73, 0xb7, 0x0c, 0x5a, 0x8b, 0x65, 0xd0, 0x7a, 0x5f, 0x06, 0x2d, 0xaf,
	0x17, 0xcb, 0xa2, 0x61, 0x83, 0xc1, 0xd6, 0x39, 0x7e, 0xbc, 0xf1, 0x6a, 0x19, 0x63, 0xe7, 0xe2,
	0x38, 0xcd, 0xf4, 0xd5, 0x7c, 0x4a, 0x62, 0x59, 0xd0, 0xfe, 0x2a, 0x2a, 0xc4, 0x12, 0x26, 0x8a,
	0xde, 0xd0, 0xff, 0xd7, 0xfc, 0xe0, 0x6e, 0xf6, 0xc3, 0xd1, 0x64, 0xf8, 0xe4, 0x92, 0xbe, 0xd5,
	0x22, 0x34, 0x2d, 0x46, 0xd8, 0x62, 0x62, 0xb7, 0x18, 0xd6, 0xf0, 0xf5, 0x0f, 0x88, 0x0c, 0x88,
	0x10, 0x44, 0x36, 0x88, 0x10, 0x54, 0x2e, 0x6b, 0x06, 0xa2, 0x93, 0xf1, 0x20, 0xe4, 0x1a, 0x12,
	0xd0, 0xf0, 0xe5, 0xf6, 0x2c, 0xcc, 0x98, 0xd1, 0x8c, 0x21, 0x67, 0xcc, 0xf6, 0x8c, 0x61, 0xc0,
	0xb4, 0x5d, 0xff, 0x2e, 0x07, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x63, 0xc2, 0x86, 0x09,
	0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/assetmantle.modules.orders.transactions.cancel.Msg/Handle", in, out, opts...)
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
		FullMethod: "/assetmantle.modules.orders.transactions.cancel.Msg/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assetmantle.modules.orders.transactions.cancel.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Msg_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orders/transactions/cancel/service.proto",
}
