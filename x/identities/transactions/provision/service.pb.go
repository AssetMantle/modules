// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: AssetMantle/modules/x/identities/transactions/provision/service.proto

package provision

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
	proto.RegisterFile("AssetMantle/modules/x/identities/transactions/provision/service.proto", fileDescriptor_d21f8e25d50e279b)
}

var fileDescriptor_d21f8e25d50e279b = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0xd2, 0x31, 0x4b, 0xc3, 0x40,
	0x14, 0x07, 0xf0, 0x24, 0xc5, 0x0e, 0xc1, 0xa9, 0x8b, 0x90, 0x21, 0x83, 0xa3, 0xc3, 0x1d, 0xea,
	0x20, 0x9c, 0x0e, 0x36, 0x22, 0x5a, 0x30, 0x10, 0x6a, 0x87, 0x20, 0x01, 0xb9, 0x26, 0x47, 0x3c,
	0x68, 0x72, 0x25, 0xef, 0x5a, 0x3a, 0xfa, 0x11, 0x9c, 0x1d, 0xdd, 0xd4, 0xc5, 0x8f, 0x21, 0x4e,
	0x1d, 0x1d, 0x25, 0x1d, 0x04, 0x3f, 0x85, 0xd4, 0x9e, 0xcd, 0x75, 0xbd, 0x29, 0xf0, 0xe0, 0xff,
	0xcb, 0x9f, 0x77, 0xcf, 0x3d, 0xef, 0x02, 0x30, 0x19, 0xd2, 0x52, 0x8e, 0x18, 0x2e, 0x44, 0x36,
	0x19, 0x31, 0xc0, 0x33, 0xcc, 0x33, 0x56, 0x4a, 0x2e, 0x39, 0x03, 0x2c, 0x2b, 0x5a, 0x02, 0x4d,
	0x25, 0x17, 0x25, 0xe0, 0x71, 0x25, 0xa6, 0x1c, 0xb8, 0x28, 0x31, 0xb0, 0x6a, 0xca, 0x53, 0x86,
	0xc6, 0x95, 0x90, 0xa2, 0x73, 0xa4, 0x31, 0x48, 0x31, 0x68, 0x86, 0x1a, 0x06, 0xe9, 0x0c, 0x5a,
	0x33, 0xde, 0x4e, 0x2a, 0xa0, 0x10, 0x80, 0x0b, 0xc8, 0xf1, 0x74, 0x7f, 0xf9, 0x59, 0x89, 0x9e,
	0x71, 0xb1, 0x82, 0x01, 0xd0, 0x5c, 0x15, 0xf3, 0xfa, 0xa6, 0x8c, 0x36, 0xbe, 0xad, 0x18, 0x8c,
	0x45, 0x09, 0xca, 0x3c, 0x78, 0xb6, 0xdd, 0x56, 0x08, 0x79, 0xe7, 0xd1, 0x76, 0xdb, 0x97, 0xb4,
	0xcc, 0x46, 0xac, 0x73, 0x8a, 0x0c, 0x17, 0x80, 0xc2, 0x55, 0x5d, 0xef, 0xca, 0x58, 0x18, 0x34,
	0xe3, 0xbe, 0x2a, 0xba, 0x6b, 0x79, 0x5b, 0xf7, 0xdf, 0x6f, 0x7b, 0x76, 0xf0, 0xda, 0x7a, 0xaf,
	0x7d, 0x7b, 0x5e, 0xfb, 0xf6, 0x57, 0xed, 0xdb, 0x0f, 0x0b, 0xdf, 0x9a, 0x2f, 0x7c, 0xeb, 0x73,
	0xe1, 0x5b, 0xee, 0x71, 0x2a, 0x0a, 0xd3, 0x9f, 0x06, 0xdb, 0xd7, 0xab, 0xf7, 0x8f, 0x96, 0x1b,
	0x89, 0xec, 0x9b, 0x20, 0xe7, 0xf2, 0x6e, 0x32, 0x44, 0xa9, 0x28, 0xb0, 0xe1, 0xca, 0x9f, 0x9c,
	0x76, 0x37, 0x8c, 0x7b, 0x83, 0xe8, 0xc5, 0xd9, 0x38, 0xa7, 0x50, 0xd5, 0x8a, 0x51, 0xaf, 0xa9,
	0x35, 0xd0, 0x6b, 0x45, 0xff, 0xc4, 0xc7, 0x46, 0x32, 0x51, 0xc9, 0x24, 0x4e, 0x9a, 0x64, 0xa2,
	0x27, 0x93, 0x75, 0xb2, 0x76, 0xce, 0x0c, 0x93, 0xc9, 0x45, 0x14, 0x84, 0x4c, 0xd2, 0x8c, 0x4a,
	0xfa, 0xe3, 0x9c, 0x68, 0x0a, 0x21, 0x8a, 0x21, 0x24, 0x26, 0xa4, 0x81, 0x08, 0xd1, 0x25, 0x42,
	0xd6, 0xd4, 0xb0, 0xfd, 0x77, 0x60, 0x87, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x06, 0x2d, 0xbd,
	0xfe, 0x96, 0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/AssetMantle.modules.x.identities.transactions.provision.Msg/Handle", in, out, opts...)
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
		FullMethod: "/AssetMantle.modules.x.identities.transactions.provision.Msg/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AssetMantle.modules.x.identities.transactions.provision.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Msg_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "AssetMantle/modules/x/identities/transactions/provision/service.proto",
}
