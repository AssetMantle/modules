// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: AssetMantle/modules/x/identities/transactions/define/service.proto

package define

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
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
	proto.RegisterFile("AssetMantle/modules/x/identities/transactions/define/service.proto", fileDescriptor_0c544b9b622019f8)
}

var fileDescriptor_0c544b9b622019f8 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x3f, 0x4b, 0xf3, 0x40,
	0x18, 0xc0, 0x93, 0x94, 0xb7, 0x43, 0x78, 0xa7, 0x2e, 0x2f, 0x84, 0x97, 0x0c, 0xef, 0xf8, 0x0e,
	0x77, 0xf8, 0x67, 0xd0, 0x03, 0xd1, 0x86, 0x82, 0x76, 0x08, 0x16, 0xed, 0x10, 0x24, 0x20, 0xd7,
	0xe4, 0x8c, 0x07, 0xc9, 0x5d, 0xc9, 0x73, 0x2d, 0x1d, 0xfd, 0x08, 0x0e, 0x2e, 0xae, 0xba, 0x39,
	0xf9, 0x31, 0xc4, 0xa9, 0xa3, 0xa3, 0xa4, 0x83, 0xe0, 0xa7, 0x90, 0xf6, 0x02, 0x49, 0xd7, 0x9b,
	0x02, 0x4f, 0xf8, 0xfd, 0xee, 0xc7, 0xc3, 0xe3, 0x06, 0x7d, 0x00, 0xa6, 0x42, 0x2a, 0x54, 0xce,
	0x70, 0x21, 0xd3, 0x59, 0xce, 0x00, 0x2f, 0x30, 0x4f, 0x99, 0x50, 0x5c, 0x71, 0x06, 0x58, 0x95,
	0x54, 0x00, 0x4d, 0x14, 0x97, 0x02, 0x70, 0xca, 0x6e, 0xb8, 0x60, 0x18, 0x58, 0x39, 0xe7, 0x09,
	0x43, 0xd3, 0x52, 0x2a, 0xd9, 0xdb, 0x6f, 0x39, 0x50, 0xed, 0x40, 0x0b, 0xd4, 0x38, 0x50, 0xdb,
	0x81, 0xb4, 0xc3, 0xfb, 0x93, 0x48, 0x28, 0x24, 0xe0, 0x02, 0x32, 0x3c, 0xdf, 0x59, 0x7f, 0xb4,
	0xce, 0xfb, 0x9b, 0x49, 0x99, 0xe5, 0x0c, 0xd3, 0x29, 0xc7, 0x54, 0x08, 0xa9, 0xa8, 0xe6, 0xf4,
	0x5f, 0xb3, 0xe0, 0x82, 0x01, 0xd0, 0xac, 0x0e, 0xf6, 0xce, 0x8d, 0x1c, 0xad, 0xd9, 0x75, 0xc9,
	0x60, 0x2a, 0x05, 0xd4, 0xc2, 0xdd, 0x67, 0xdb, 0xed, 0x84, 0x90, 0xf5, 0x1e, 0x6c, 0xb7, 0x7b,
	0x46, 0x45, 0x9a, 0xb3, 0xde, 0x11, 0x32, 0xd9, 0x0a, 0x0a, 0x75, 0xa8, 0x37, 0x34, 0xc3, 0xc7,
	0xcd, 0xec, 0xa2, 0x4e, 0xfc, 0x67, 0x79, 0xbf, 0xee, 0xbe, 0x5e, 0xff, 0xdb, 0xc1, 0x63, 0xe7,
	0xad, 0xf2, 0xed, 0x65, 0xe5, 0xdb, 0x9f, 0x95, 0x6f, 0xdf, 0xaf, 0x7c, 0x6b, 0xb9, 0xf2, 0xad,
	0x8f, 0x95, 0x6f, 0xb9, 0x07, 0x89, 0x2c, 0x8c, 0x5e, 0x0c, 0x7e, 0x5f, 0xea, 0x5b, 0x18, 0xad,
	0x17, 0x31, 0xb2, 0xaf, 0x8e, 0x33, 0xae, 0x6e, 0x67, 0x13, 0x94, 0xc8, 0x02, 0x9b, 0xac, 0xf9,
	0xc9, 0xe9, 0xf6, 0xc3, 0x68, 0x38, 0x1e, 0xbc, 0x38, 0x5b, 0x77, 0x15, 0xd6, 0x41, 0x11, 0x1a,
	0x36, 0x41, 0xe3, 0x76, 0xd0, 0x60, 0xc3, 0xbf, 0x6f, 0x61, 0x71, 0x8d, 0xc5, 0x51, 0xdc, 0x60,
	0x71, 0x1b, 0x8b, 0x35, 0x56, 0x39, 0x27, 0x26, 0x58, 0x7c, 0x3a, 0x0a, 0x42, 0xa6, 0x68, 0x4a,
	0x15, 0xfd, 0x76, 0x0e, 0x5b, 0x0a, 0x42, 0x6a, 0x07, 0x21, 0x11, 0x21, 0x8d, 0x85, 0x90, 0xb6,
	0x86, 0x10, 0xed, 0x99, 0x74, 0x37, 0x87, 0xb4, 0xf7, 0x13, 0x00, 0x00, 0xff, 0xff, 0xef, 0xef,
	0xc1, 0xb5, 0x90, 0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/AssetMantle.modules.x.identities.transactions.define.Msg/Handle", in, out, opts...)
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
		FullMethod: "/AssetMantle.modules.x.identities.transactions.define.Msg/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AssetMantle.modules.x.identities.transactions.define.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Msg_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "AssetMantle/modules/x/identities/transactions/define/service.proto",
}
