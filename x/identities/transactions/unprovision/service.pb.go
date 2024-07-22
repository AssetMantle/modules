// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: AssetMantle/modules/x/identities/transactions/unprovision/service.proto

package unprovision

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
	proto.RegisterFile("AssetMantle/modules/x/identities/transactions/unprovision/service.proto", fileDescriptor_cb37a005e91ee120)
}

var fileDescriptor_cb37a005e91ee120 = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0xd2, 0x31, 0x4b, 0xf3, 0x40,
	0x18, 0x07, 0xf0, 0x24, 0xe5, 0xed, 0x10, 0xde, 0xa9, 0xcb, 0x0b, 0xe1, 0x25, 0xc3, 0x3b, 0xbe,
	0xc3, 0x1d, 0xea, 0xe4, 0x81, 0x42, 0x83, 0x52, 0x3b, 0x44, 0x8a, 0xa6, 0x10, 0x24, 0x20, 0xd7,
	0xe4, 0x88, 0x07, 0xc9, 0x5d, 0xc8, 0x73, 0x2d, 0x1d, 0xfd, 0x08, 0x7e, 0x02, 0x07, 0x47, 0x41,
	0x70, 0xf3, 0x2b, 0x88, 0x53, 0x47, 0x47, 0x49, 0x07, 0xc1, 0x4f, 0x21, 0x35, 0x81, 0x5c, 0xd7,
	0x74, 0xba, 0xe1, 0xe1, 0xff, 0xbb, 0x3f, 0x0f, 0x8f, 0x3d, 0x1a, 0x02, 0x30, 0xe5, 0x53, 0xa1,
	0x32, 0x86, 0x73, 0x99, 0xcc, 0x33, 0x06, 0x78, 0x89, 0x79, 0xc2, 0x84, 0xe2, 0x8a, 0x33, 0xc0,
	0xaa, 0xa4, 0x02, 0x68, 0xac, 0xb8, 0x14, 0x80, 0xe7, 0xa2, 0x28, 0xe5, 0x82, 0x03, 0x97, 0x02,
	0x03, 0x2b, 0x17, 0x3c, 0x66, 0xa8, 0x28, 0xa5, 0x92, 0x83, 0x43, 0x0d, 0x42, 0x0d, 0x84, 0x96,
	0xa8, 0x85, 0x90, 0x0e, 0x21, 0x0d, 0x72, 0xfe, 0xc4, 0x12, 0x72, 0x09, 0x38, 0x87, 0x14, 0x2f,
	0xf6, 0x36, 0x4f, 0x6d, 0x3a, 0x7f, 0x53, 0x29, 0xd3, 0x8c, 0x61, 0x5a, 0x70, 0x4c, 0x85, 0x90,
	0x8a, 0xd6, 0xe1, 0x7a, 0xba, 0x43, 0xf5, 0x9c, 0x01, 0xd0, 0xb4, 0xa9, 0xee, 0x04, 0xdd, 0x21,
	0x6d, 0x70, 0x5d, 0x32, 0x28, 0xa4, 0x80, 0x46, 0xdd, 0x7f, 0x32, 0xed, 0x9e, 0x0f, 0xe9, 0xe0,
	0xde, 0xb4, 0xfb, 0x67, 0x54, 0x24, 0x19, 0x1b, 0x78, 0xa8, 0xf3, 0x92, 0x90, 0x5f, 0x57, 0x76,
	0xce, 0x77, 0x30, 0x82, 0x76, 0x70, 0xd1, 0x94, 0xfd, 0x67, 0x38, 0xbf, 0x6e, 0x3f, 0x9f, 0xff,
	0x9b, 0xde, 0x4b, 0xef, 0xb5, 0x72, 0xcd, 0x55, 0xe5, 0x9a, 0x1f, 0x95, 0x6b, 0xde, 0xad, 0x5d,
	0x63, 0xb5, 0x76, 0x8d, 0xf7, 0xb5, 0x6b, 0xd8, 0x47, 0xb1, 0xcc, 0xbb, 0x7f, 0xeb, 0xfd, 0xbe,
	0xac, 0x2f, 0x65, 0xb2, 0xd9, 0xcb, 0xc4, 0xbc, 0x3a, 0x49, 0xb9, 0xba, 0x99, 0xcf, 0x50, 0x2c,
	0x73, 0xdc, 0x79, 0xf5, 0x0f, 0x56, 0x7f, 0xe8, 0x87, 0xe3, 0x60, 0xfa, 0x68, 0x6d, 0x9d, 0x9e,
	0xdf, 0x54, 0x0b, 0xd1, 0xb8, 0xad, 0x16, 0xe8, 0xd5, 0xa6, 0x2d, 0xf2, 0xb6, 0x95, 0x8d, 0x9a,
	0x6c, 0x14, 0x46, 0x6d, 0x36, 0xd2, 0xb3, 0x91, 0x96, 0xad, 0xac, 0xd3, 0xce, 0xd9, 0x68, 0x34,
	0xf1, 0x7c, 0xa6, 0x68, 0x42, 0x15, 0xfd, 0xb2, 0x8e, 0x35, 0x87, 0x90, 0x06, 0x22, 0x24, 0x24,
	0xa4, 0xa5, 0x08, 0xd1, 0x2d, 0x42, 0x34, 0x6c, 0xd6, 0xff, 0x39, 0xb8, 0x83, 0xef, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xa7, 0x02, 0x03, 0xa6, 0xcc, 0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/AssetMantle.modules.x.identities.transactions.unprovision.Msg/Handle", in, out, opts...)
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
		FullMethod: "/AssetMantle.modules.x.identities.transactions.unprovision.Msg/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AssetMantle.modules.x.identities.transactions.unprovision.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Msg_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "AssetMantle/modules/x/identities/transactions/unprovision/service.proto",
}
