// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: AssetMantle/modules/x/assets/transactions/revoke/service.proto

package revoke

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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
	proto.RegisterFile("AssetMantle/modules/x/assets/transactions/revoke/service.proto", fileDescriptor_5cc15353cc44c580)
}

var fileDescriptor_5cc15353cc44c580 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xd2, 0xbf, 0x4b, 0xc3, 0x40,
	0x14, 0x07, 0xf0, 0x5c, 0x8a, 0x1d, 0x82, 0x53, 0x17, 0x21, 0x43, 0x06, 0x47, 0x87, 0x3b, 0x7f,
	0x0e, 0x9e, 0x22, 0xb4, 0x20, 0x0a, 0x12, 0x28, 0xb5, 0x43, 0x90, 0x80, 0x5c, 0xd3, 0x23, 0x16,
	0x7b, 0xb9, 0x92, 0x77, 0x2d, 0x1d, 0x75, 0x77, 0x70, 0x77, 0x73, 0x74, 0xf2, 0xcf, 0x10, 0xa7,
	0x8e, 0x8e, 0x92, 0x0e, 0x82, 0x7f, 0x85, 0xa4, 0x77, 0xe0, 0x75, 0xbc, 0x29, 0xf0, 0xe0, 0xf3,
	0xcd, 0xf7, 0x25, 0x2f, 0x38, 0x6b, 0x03, 0x70, 0x15, 0xb3, 0x42, 0x8d, 0x39, 0x11, 0x72, 0x38,
	0x1d, 0x73, 0x20, 0x73, 0xc2, 0xea, 0x29, 0x10, 0x55, 0xb2, 0x02, 0x58, 0xa6, 0x46, 0xb2, 0x00,
	0x52, 0xf2, 0x99, 0xbc, 0xe7, 0x04, 0x78, 0x39, 0x1b, 0x65, 0x1c, 0x4f, 0x4a, 0xa9, 0x64, 0x6b,
	0xd7, 0xf2, 0xd8, 0x78, 0x3c, 0xc7, 0xda, 0x63, 0xdb, 0x63, 0xed, 0xc3, 0xad, 0x4c, 0x82, 0x90,
	0x40, 0x04, 0xe4, 0x64, 0xb6, 0x57, 0x3f, 0x74, 0x54, 0xe8, 0x5e, 0x45, 0x70, 0x00, 0x96, 0x9b,
	0x2a, 0xe1, 0x95, 0xb3, 0xb7, 0x66, 0xb7, 0x25, 0x87, 0x89, 0x2c, 0xc0, 0x84, 0xed, 0xbf, 0xa0,
	0xa0, 0x11, 0x43, 0xde, 0x7a, 0x42, 0x41, 0xf3, 0x92, 0x15, 0xc3, 0x31, 0x6f, 0x1d, 0x63, 0xd7,
	0x5d, 0x71, 0xac, 0x0b, 0x86, 0xe7, 0xee, 0xb4, 0xff, 0x3f, 0xeb, 0x99, 0x6a, 0xdb, 0x5e, 0xb8,
	0xf1, 0xf0, 0xf3, 0xbe, 0x83, 0x3a, 0x8f, 0x8d, 0x8f, 0x2a, 0x42, 0x8b, 0x2a, 0x42, 0xdf, 0x55,
	0x84, 0x9e, 0x97, 0x91, 0xb7, 0x58, 0x46, 0xde, 0xd7, 0x32, 0xf2, 0x82, 0xc3, 0x4c, 0x0a, 0xe7,
	0xb7, 0x75, 0x36, 0xaf, 0xf5, 0x5f, 0xed, 0xd6, 0xcb, 0x77, 0xd1, 0xcd, 0x49, 0x3e, 0x52, 0x77,
	0xd3, 0x01, 0xce, 0xa4, 0x20, 0xae, 0x9f, 0xf5, 0xd5, 0x6f, 0xb6, 0xe3, 0xa4, 0xdd, 0xef, 0xbd,
	0xf9, 0x6b, 0xd7, 0x11, 0x9b, 0x22, 0x89, 0xae, 0x07, 0xf6, 0x8a, 0x80, 0x7b, 0x2b, 0xfb, 0xb9,
	0x46, 0x52, 0x43, 0xd2, 0x24, 0xd5, 0x24, 0xb5, 0x49, 0xaa, 0x49, 0xe5, 0x9f, 0xba, 0x92, 0xf4,
	0xa2, 0xdb, 0x89, 0xb9, 0x62, 0x43, 0xa6, 0xd8, 0xaf, 0x7f, 0x64, 0x71, 0x4a, 0x8d, 0xa7, 0x34,
	0xa1, 0x54, 0x27, 0x50, 0x6a, 0x47, 0x50, 0xaa, 0x33, 0x06, 0xcd, 0xd5, 0xa1, 0x1c, 0xfc, 0x05,
	0x00, 0x00, 0xff, 0xff, 0x4b, 0x04, 0x61, 0x12, 0x42, 0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/AssetMantle.modules.x.assets.transactions.revoke.Msg/Handle", in, out, opts...)
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
		FullMethod: "/AssetMantle.modules.x.assets.transactions.revoke.Msg/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var Msg_serviceDesc = _Msg_serviceDesc
var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AssetMantle.modules.x.assets.transactions.revoke.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Msg_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "AssetMantle/modules/x/assets/transactions/revoke/service.proto",
}
