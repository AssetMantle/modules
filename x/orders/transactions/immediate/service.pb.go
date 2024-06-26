// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: AssetMantle/modules/x/orders/transactions/immediate/service.proto

package immediate

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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
	proto.RegisterFile("AssetMantle/modules/x/orders/transactions/immediate/service.proto", fileDescriptor_0dcccd3b865dd390)
}

var fileDescriptor_0dcccd3b865dd390 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xb1, 0x4a, 0xfb, 0x40,
	0x18, 0xc0, 0x9b, 0x14, 0x3a, 0x84, 0xff, 0xd4, 0xe9, 0x4f, 0x90, 0x20, 0x7d, 0x80, 0x3b, 0xb0,
	0x83, 0x72, 0x88, 0xd2, 0x2e, 0xb6, 0x43, 0x6c, 0xd1, 0x0e, 0x41, 0x02, 0x72, 0x4d, 0x8e, 0x18,
	0xc8, 0xdd, 0x95, 0xdc, 0x55, 0x3a, 0xfb, 0x04, 0x42, 0xdf, 0xc0, 0xd1, 0xd1, 0xd1, 0x27, 0x10,
	0xa7, 0x82, 0x8b, 0xa3, 0x24, 0x4e, 0x3e, 0x85, 0xd8, 0x3b, 0xe3, 0x75, 0xbd, 0xf5, 0xb8, 0xdf,
	0x8f, 0xdf, 0xf7, 0xf1, 0x79, 0x83, 0x81, 0x10, 0x44, 0x86, 0x98, 0xc9, 0x82, 0x40, 0xca, 0xd3,
	0x65, 0x41, 0x04, 0x5c, 0x41, 0x5e, 0xa6, 0xa4, 0x14, 0x50, 0x96, 0x98, 0x09, 0x9c, 0xc8, 0x9c,
	0x33, 0x01, 0x73, 0x4a, 0x49, 0x9a, 0x63, 0x49, 0xa0, 0x20, 0xe5, 0x6d, 0x9e, 0x10, 0xb0, 0x28,
	0xb9, 0xe4, 0xdd, 0xbe, 0xa1, 0x00, 0x5a, 0x01, 0x56, 0x40, 0x29, 0x80, 0xa9, 0x00, 0x8d, 0xc2,
	0xdf, 0xcb, 0x38, 0xcf, 0x0a, 0x02, 0xf1, 0x22, 0x87, 0x98, 0x31, 0x2e, 0xb1, 0xfa, 0xb1, 0x55,
	0xfa, 0x56, 0x55, 0x94, 0x08, 0x81, 0x33, 0x5d, 0xe5, 0x9f, 0xdb, 0x28, 0x8c, 0xe7, 0xeb, 0x92,
	0x88, 0x05, 0x67, 0x42, 0xfb, 0x0e, 0x9e, 0x1d, 0xaf, 0x1d, 0x8a, 0xac, 0xfb, 0xe4, 0x78, 0x9d,
	0x11, 0x66, 0x69, 0x41, 0xba, 0xc7, 0xc0, 0x62, 0x72, 0x10, 0xaa, 0x4c, 0x7f, 0x64, 0x45, 0xcf,
	0xfe, 0x9e, 0x2f, 0x74, 0x60, 0x6f, 0xff, 0xee, 0xed, 0x73, 0xed, 0xfa, 0xbd, 0xff, 0x90, 0xaa,
	0x71, 0xf5, 0x90, 0x0d, 0x35, 0x5c, 0xb7, 0x5f, 0xaa, 0xc0, 0xd9, 0x54, 0x81, 0xf3, 0x51, 0x05,
	0xce, 0x7d, 0x1d, 0xb4, 0x36, 0x75, 0xd0, 0x7a, 0xaf, 0x83, 0x96, 0x77, 0x98, 0x70, 0x6a, 0x53,
	0x32, 0xfc, 0x77, 0xa9, 0xae, 0x60, 0xfa, 0xb3, 0x9e, 0xa9, 0x73, 0x75, 0x92, 0xe5, 0xf2, 0x66,
	0x39, 0x07, 0x09, 0xa7, 0xd0, 0x62, 0xf7, 0x0f, 0x6e, 0x67, 0x10, 0x46, 0x93, 0xd9, 0xf8, 0xd1,
	0xdd, 0x39, 0xa8, 0x50, 0xe7, 0x44, 0x60, 0xa2, 0x72, 0x66, 0x66, 0xce, 0xf8, 0x17, 0x7f, 0xdd,
	0xa1, 0x62, 0x4d, 0xc5, 0x51, 0xac, 0xa8, 0xd8, 0xa4, 0xe2, 0x86, 0xaa, 0xdc, 0x53, 0x0b, 0x2a,
	0x3e, 0x9b, 0x0e, 0x43, 0x22, 0x71, 0x8a, 0x25, 0xfe, 0x72, 0x8f, 0x0c, 0x03, 0x42, 0x5a, 0x81,
	0x50, 0x84, 0x90, 0x92, 0x20, 0x64, 0x5a, 0x10, 0x6a, 0x34, 0xf3, 0xce, 0xf6, 0xb2, 0xfa, 0xdf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x8d, 0xa2, 0xe0, 0x1f, 0x84, 0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/AssetMantle.modules.x.orders.transactions.immediate.Msg/Handle", in, out, opts...)
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
		FullMethod: "/AssetMantle.modules.x.orders.transactions.immediate.Msg/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AssetMantle.modules.x.orders.transactions.immediate.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Msg_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "AssetMantle/modules/x/orders/transactions/immediate/service.proto",
}
