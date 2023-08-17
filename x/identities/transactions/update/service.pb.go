// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: identities/transactions/update/service.proto

package update

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
	proto.RegisterFile("identities/transactions/update/service.proto", fileDescriptor_413e971d64eb69c7)
}

var fileDescriptor_413e971d64eb69c7 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xb1, 0x4a, 0xc3, 0x40,
	0x18, 0x80, 0x9b, 0x14, 0x3b, 0x04, 0xa7, 0x6e, 0x46, 0x09, 0xd8, 0x59, 0xee, 0xa0, 0x42, 0xc1,
	0x13, 0x94, 0x76, 0xa9, 0x1d, 0x0e, 0x8a, 0x76, 0x92, 0x80, 0x5c, 0x93, 0x23, 0x06, 0x92, 0xbb,
	0x90, 0xfb, 0x23, 0xce, 0xee, 0x82, 0xe0, 0x1b, 0x38, 0xba, 0x39, 0xfa, 0x06, 0xe2, 0x54, 0x70,
	0x71, 0x94, 0xd4, 0xc9, 0xa7, 0x90, 0xf6, 0x02, 0x39, 0x17, 0x4b, 0xbb, 0x5e, 0xf2, 0x7d, 0xf7,
	0xfd, 0x7f, 0xe2, 0x1c, 0xc4, 0x21, 0x17, 0x10, 0x43, 0xcc, 0x15, 0x86, 0x9c, 0x09, 0xc5, 0x02,
	0x88, 0xa5, 0x50, 0xb8, 0xc8, 0x42, 0x06, 0x1c, 0x2b, 0x9e, 0xdf, 0xc4, 0x01, 0x47, 0x59, 0x2e,
	0x41, 0xb6, 0xbb, 0x4c, 0x29, 0x0e, 0x29, 0x13, 0x90, 0x70, 0x94, 0xca, 0xb0, 0x48, 0xb8, 0x42,
	0xb5, 0x01, 0x99, 0x06, 0x94, 0x16, 0xc0, 0x80, 0xbb, 0x7b, 0x91, 0x94, 0x51, 0xc2, 0x31, 0xcb,
	0x62, 0xcc, 0x84, 0x90, 0xc0, 0xf4, 0xe3, 0xa5, 0xd1, 0x5d, 0x75, 0x7f, 0xca, 0x95, 0x62, 0x51,
	0x75, 0xbf, 0x7b, 0xb4, 0xe2, 0x6d, 0xe3, 0xec, 0x2a, 0xe7, 0x2a, 0x93, 0x42, 0x55, 0x68, 0xf7,
	0xd5, 0x72, 0x9a, 0x54, 0x45, 0xed, 0x17, 0xcb, 0x69, 0x9d, 0x31, 0x11, 0x26, 0xbc, 0x7d, 0x8c,
	0xd6, 0x1f, 0x07, 0x51, 0x1d, 0xe4, 0x0e, 0x37, 0x81, 0x27, 0xf5, 0xd9, 0x79, 0x95, 0xd7, 0xd9,
	0xbf, 0xfb, 0xf8, 0x7e, 0xb4, 0x77, 0x3b, 0x3b, 0x58, 0xbb, 0xb0, 0x31, 0xa3, 0x1e, 0x6b, 0x70,
	0xdf, 0x7c, 0x2b, 0x3d, 0x6b, 0x56, 0x7a, 0xd6, 0x57, 0xe9, 0x59, 0x0f, 0x73, 0xaf, 0x31, 0x9b,
	0x7b, 0x8d, 0xcf, 0xb9, 0xd7, 0x70, 0x7a, 0x81, 0x4c, 0x37, 0x28, 0x19, 0x6c, 0x5f, 0xe8, 0x0f,
	0x3b, 0x5e, 0x2c, 0x67, 0x6c, 0x5d, 0x9e, 0x46, 0x31, 0x5c, 0x17, 0x53, 0x14, 0xc8, 0x14, 0xf7,
	0x17, 0x3a, 0xaa, 0x63, 0x2a, 0x1d, 0xbe, 0xc5, 0xff, 0xaf, 0xfe, 0xc9, 0xde, 0xea, 0xd3, 0xd1,
	0x84, 0x3e, 0xdb, 0xdd, 0xbe, 0x51, 0x43, 0xab, 0x9a, 0x51, 0x5d, 0x33, 0x31, 0x6b, 0xe8, 0xb2,
	0xe6, 0xfd, 0x0f, 0xe4, 0x57, 0x90, 0x5f, 0x43, 0xbe, 0x09, 0xf9, 0x1a, 0x2a, 0xed, 0x93, 0xf5,
	0x21, 0x7f, 0x38, 0x1e, 0x50, 0x0e, 0x2c, 0x64, 0xc0, 0x7e, 0xec, 0x9e, 0x21, 0x20, 0xa4, 0x32,
	0x10, 0x52, 0x2b, 0x08, 0x31, 0x1d, 0x84, 0x68, 0xc9, 0xb4, 0xb5, 0xfc, 0xa5, 0x0e, 0x7f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0xd3, 0x99, 0xf5, 0x12, 0x3d, 0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/assetmantle.modules.identities.transactions.mutate.Msg/Handle", in, out, opts...)
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
		FullMethod: "/assetmantle.modules.identities.transactions.mutate.Msg/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assetmantle.modules.identities.transactions.mutate.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Msg_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "identities/transactions/update/service.proto",
}
