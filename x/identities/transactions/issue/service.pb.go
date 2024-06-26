// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: AssetMantle/modules/x/identities/transactions/issue/service.proto

package issue

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
	proto.RegisterFile("AssetMantle/modules/x/identities/transactions/issue/service.proto", fileDescriptor_67f6e03c262c4d59)
}

var fileDescriptor_67f6e03c262c4d59 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xb1, 0x4a, 0xc3, 0x40,
	0x18, 0x80, 0x9b, 0x14, 0x3a, 0x04, 0xa7, 0x4e, 0x12, 0x24, 0x48, 0x1f, 0xe0, 0x0e, 0xec, 0xa0,
	0x1c, 0xa2, 0xb4, 0x8b, 0xcd, 0x10, 0x29, 0xda, 0x21, 0x48, 0x40, 0xae, 0xc9, 0x11, 0x0f, 0x92,
	0xbb, 0x92, 0xff, 0x2a, 0x9d, 0x7d, 0x02, 0xa1, 0x6f, 0xe0, 0xe8, 0xe8, 0xe8, 0x13, 0x88, 0x53,
	0xc1, 0xc5, 0x51, 0x52, 0x27, 0x9f, 0x42, 0x9a, 0x04, 0x72, 0x5d, 0x6f, 0xbd, 0xe3, 0xfb, 0xf8,
	0xfe, 0x9f, 0xdf, 0x19, 0x8d, 0x00, 0x98, 0x0a, 0xa8, 0x50, 0x19, 0xc3, 0xb9, 0x4c, 0x96, 0x19,
	0x03, 0xbc, 0xc2, 0x3c, 0x61, 0x42, 0x71, 0xc5, 0x19, 0x60, 0x55, 0x50, 0x01, 0x34, 0x56, 0x5c,
	0x0a, 0xc0, 0x1c, 0x60, 0xc9, 0x30, 0xb0, 0xe2, 0x91, 0xc7, 0x0c, 0x2d, 0x0a, 0xa9, 0x64, 0x7f,
	0xa8, 0x29, 0x50, 0xa3, 0x40, 0x2b, 0xd4, 0x2a, 0x90, 0xae, 0x40, 0x95, 0xc2, 0x3d, 0x4a, 0xa5,
	0x4c, 0x33, 0x86, 0xe9, 0x82, 0x63, 0x2a, 0x84, 0x54, 0xb4, 0xfe, 0xad, 0x94, 0xae, 0x51, 0x55,
	0xce, 0x00, 0x68, 0xda, 0x54, 0xb9, 0xd7, 0x26, 0x0a, 0xed, 0xe9, 0xbe, 0x60, 0xb0, 0x90, 0x02,
	0x1a, 0xdf, 0xc9, 0xbb, 0xe5, 0x74, 0x03, 0x48, 0xfb, 0x6f, 0x96, 0xd3, 0x9b, 0x50, 0x91, 0x64,
	0xac, 0x7f, 0x8e, 0x0c, 0x26, 0x47, 0x41, 0x9d, 0xe9, 0x4e, 0x8c, 0xe8, 0x59, 0xfb, 0x74, 0xd3,
	0x04, 0x0e, 0x8e, 0x9f, 0xbe, 0x7e, 0xd7, 0xb6, 0x3b, 0x38, 0xc4, 0x79, 0x3d, 0xae, 0x36, 0x64,
	0x45, 0x8d, 0xd7, 0xdd, 0x8f, 0xd2, 0xb3, 0x36, 0xa5, 0x67, 0xfd, 0x94, 0x9e, 0xf5, 0xbc, 0xf5,
	0x3a, 0x9b, 0xad, 0xd7, 0xf9, 0xde, 0x7a, 0x1d, 0xe7, 0x34, 0x96, 0xb9, 0x49, 0xc9, 0xf8, 0xe0,
	0xb6, 0xbe, 0x82, 0xe9, 0x6e, 0x3d, 0x53, 0xeb, 0xee, 0x22, 0xe5, 0xea, 0x61, 0x39, 0x47, 0xb1,
	0xcc, 0xb1, 0xc1, 0xee, 0x5f, 0xec, 0xde, 0x28, 0x08, 0xfd, 0x99, 0xff, 0x6a, 0xef, 0x1d, 0x54,
	0xd0, 0xe4, 0x84, 0xc8, 0x6f, 0x73, 0x66, 0x7a, 0x8e, 0xbf, 0xc3, 0x3f, 0xf7, 0xa8, 0xa8, 0xa1,
	0xa2, 0x30, 0x6a, 0xa9, 0x48, 0xa7, 0xa2, 0x8a, 0x2a, 0xed, 0x4b, 0x03, 0x2a, 0xba, 0x9a, 0x8e,
	0x03, 0xa6, 0x68, 0x42, 0x15, 0xfd, 0xb3, 0xcf, 0x34, 0x03, 0x21, 0x8d, 0x82, 0x90, 0x90, 0x90,
	0x56, 0x42, 0x88, 0x6e, 0x21, 0xa4, 0xd2, 0xcc, 0x7b, 0xd5, 0x65, 0x0d, 0xff, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x01, 0x5a, 0xf6, 0xd3, 0x84, 0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/AssetMantle.modules.x.identities.transactions.issue.Msg/Handle", in, out, opts...)
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
		FullMethod: "/AssetMantle.modules.x.identities.transactions.issue.Msg/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AssetMantle.modules.x.identities.transactions.issue.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Msg_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "AssetMantle/modules/x/identities/transactions/issue/service.proto",
}
