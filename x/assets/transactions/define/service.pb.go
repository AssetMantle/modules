// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: AssetMantle/modules/x/assets/transactions/define/service.proto

package define

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
	proto.RegisterFile("AssetMantle/modules/x/assets/transactions/define/service.proto", fileDescriptor_1e8e48cb5e3393e4)
}

var fileDescriptor_1e8e48cb5e3393e4 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xd2, 0x3f, 0x4b, 0xfb, 0x40,
	0x18, 0x07, 0xf0, 0x5c, 0xca, 0xaf, 0x43, 0xf8, 0x4d, 0x5d, 0x84, 0x0c, 0x19, 0x1c, 0x1d, 0xee,
	0xfc, 0x3b, 0x78, 0x8a, 0xd0, 0x52, 0x51, 0x90, 0x40, 0xd1, 0x0e, 0x41, 0x02, 0x72, 0x4d, 0xce,
	0x18, 0x68, 0x72, 0x25, 0xcf, 0xb5, 0x74, 0xd4, 0xdd, 0xc1, 0xdd, 0xcd, 0xd1, 0xc9, 0x97, 0x21,
	0x4e, 0x1d, 0x1d, 0x25, 0x1d, 0x04, 0x5f, 0x85, 0xa4, 0x77, 0xe0, 0x75, 0xbc, 0x29, 0xf0, 0xc0,
	0xe7, 0x9b, 0xef, 0x93, 0x3c, 0xde, 0x49, 0x17, 0x80, 0xcb, 0x90, 0x95, 0x72, 0xcc, 0x49, 0x21,
	0xd2, 0xe9, 0x98, 0x03, 0x99, 0x13, 0xd6, 0x4c, 0x81, 0xc8, 0x8a, 0x95, 0xc0, 0x12, 0x99, 0x8b,
	0x12, 0x48, 0xca, 0x6f, 0xf3, 0x92, 0x13, 0xe0, 0xd5, 0x2c, 0x4f, 0x38, 0x9e, 0x54, 0x42, 0x8a,
	0xce, 0xb6, 0xe1, 0xb1, 0xf6, 0x78, 0x8e, 0x95, 0xc7, 0xa6, 0xc7, 0xca, 0xfb, 0x1b, 0x89, 0x80,
	0x42, 0x00, 0x29, 0x20, 0x23, 0xb3, 0x9d, 0xe6, 0xa1, 0xa2, 0x7c, 0xfb, 0x2a, 0x05, 0x07, 0x60,
	0x99, 0xae, 0xe2, 0x5f, 0x58, 0x7b, 0x63, 0x76, 0x53, 0x71, 0x98, 0x88, 0x12, 0x74, 0xd8, 0xee,
	0x33, 0xf2, 0x5a, 0x21, 0x64, 0x9d, 0x47, 0xe4, 0xb5, 0xcf, 0x59, 0x99, 0x8e, 0x79, 0xe7, 0x10,
	0xdb, 0xee, 0x8a, 0x43, 0x55, 0xd0, 0x3f, 0xb5, 0xa7, 0xc3, 0xbf, 0xd9, 0xa5, 0xae, 0xb6, 0xe9,
	0xf8, 0xff, 0xee, 0xbf, 0xdf, 0xb6, 0x50, 0xef, 0xa1, 0xf5, 0x5e, 0x07, 0x68, 0x51, 0x07, 0xe8,
	0xab, 0x0e, 0xd0, 0xd3, 0x32, 0x70, 0x16, 0xcb, 0xc0, 0xf9, 0x5c, 0x06, 0x8e, 0xb7, 0x9f, 0x88,
	0xc2, 0xfa, 0x6d, 0xbd, 0xff, 0x57, 0xea, 0xaf, 0x0e, 0x9a, 0xe5, 0x07, 0xe8, 0xfa, 0x28, 0xcb,
	0xe5, 0xdd, 0x74, 0x84, 0x13, 0x51, 0x10, 0xdb, 0xcf, 0xfa, 0xe2, 0xb6, 0xbb, 0x61, 0xd4, 0x1d,
	0xf6, 0x5f, 0xdd, 0xb5, 0xeb, 0x08, 0x75, 0x91, 0x48, 0xd5, 0x03, 0x73, 0x45, 0xc0, 0xfd, 0x95,
	0xfd, 0x58, 0x23, 0xb1, 0x26, 0x71, 0x14, 0x2b, 0x12, 0x9b, 0x24, 0x56, 0xa4, 0x76, 0x8f, 0x6d,
	0x49, 0x7c, 0x36, 0xe8, 0x85, 0x5c, 0xb2, 0x94, 0x49, 0xf6, 0xe3, 0x1e, 0x18, 0x9c, 0x52, 0xed,
	0x29, 0x8d, 0x28, 0x55, 0x09, 0x94, 0x9a, 0x11, 0x94, 0xaa, 0x8c, 0x51, 0x7b, 0x75, 0x28, 0x7b,
	0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x26, 0xd5, 0x35, 0x42, 0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/AssetMantle.modules.x.assets.transactions.define.Msg/Handle", in, out, opts...)
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
		FullMethod: "/AssetMantle.modules.x.assets.transactions.define.Msg/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var Msg_serviceDesc = _Msg_serviceDesc
var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AssetMantle.modules.x.assets.transactions.define.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Msg_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "AssetMantle/modules/x/assets/transactions/define/service.proto",
}
