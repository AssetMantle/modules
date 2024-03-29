// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: identities/transactions/issue/service.proto

package issue

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
	proto.RegisterFile("identities/transactions/issue/service.proto", fileDescriptor_3c58d2643b032f4c)
}

var fileDescriptor_3c58d2643b032f4c = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xb1, 0x4a, 0xc3, 0x40,
	0x18, 0x80, 0x9b, 0x14, 0x3b, 0x04, 0xa7, 0x4e, 0x12, 0x24, 0x48, 0x57, 0xe1, 0x8e, 0x2a, 0x05,
	0x39, 0x50, 0x68, 0x07, 0x35, 0x43, 0xa0, 0x68, 0x27, 0x09, 0xc8, 0x35, 0x39, 0xe2, 0x41, 0x72,
	0x57, 0xf2, 0x5f, 0xc5, 0xd9, 0xc5, 0x55, 0xf0, 0x0d, 0x1c, 0x1d, 0x1c, 0x7d, 0x06, 0x71, 0x2a,
	0xb8, 0x38, 0x4a, 0xea, 0xe4, 0x53, 0x48, 0x93, 0x83, 0x9c, 0x4b, 0x25, 0xae, 0x97, 0x7c, 0xdf,
	0x7d, 0xff, 0x9f, 0x38, 0xbb, 0x3c, 0x66, 0x42, 0x71, 0xc5, 0x19, 0x60, 0x95, 0x53, 0x01, 0x34,
	0x52, 0x5c, 0x0a, 0xc0, 0x1c, 0x60, 0xce, 0x30, 0xb0, 0xfc, 0x9a, 0x47, 0x0c, 0xcd, 0x72, 0xa9,
	0x64, 0xb7, 0x4f, 0x01, 0x98, 0xca, 0xa8, 0x50, 0x29, 0x43, 0x99, 0x8c, 0xe7, 0x29, 0x03, 0x54,
	0x0b, 0x90, 0x29, 0x40, 0xa5, 0xc0, 0xdd, 0x4e, 0xa4, 0x4c, 0x52, 0x86, 0xe9, 0x8c, 0x63, 0x2a,
	0x84, 0x54, 0xb4, 0x7a, 0x5a, 0x0a, 0xdd, 0x3f, 0x6e, 0xcf, 0x18, 0x00, 0x4d, 0xf4, 0xed, 0xee,
	0xc1, 0xfa, 0x97, 0x8d, 0xa3, 0xcb, 0x9c, 0xc1, 0x4c, 0x0a, 0xd0, 0xe4, 0xde, 0x8b, 0xe5, 0xb4,
	0x03, 0x48, 0xba, 0xcf, 0x96, 0xd3, 0x39, 0xa5, 0x22, 0x4e, 0x59, 0x97, 0xa0, 0xc6, 0xb3, 0xa0,
	0xa0, 0xca, 0x71, 0x8f, 0xff, 0xc1, 0x4e, 0xea, 0xa3, 0x33, 0x1d, 0xd7, 0xdb, 0xb9, 0x7d, 0xff,
	0x7a, 0xb0, 0xdd, 0xde, 0x16, 0xae, 0x54, 0xd8, 0x18, 0xb0, 0xa4, 0x46, 0x77, 0xed, 0xd7, 0xc2,
	0xb3, 0x16, 0x85, 0x67, 0x7d, 0x16, 0x9e, 0x75, 0xbf, 0xf4, 0x5a, 0x8b, 0xa5, 0xd7, 0xfa, 0x58,
	0x7a, 0x2d, 0x67, 0x10, 0xc9, 0xac, 0x79, 0xc7, 0x68, 0xf3, 0xbc, 0xfa, 0xa2, 0xe3, 0xd5, 0x62,
	0xc6, 0xd6, 0xc5, 0x51, 0xc2, 0xd5, 0xd5, 0x7c, 0x8a, 0x22, 0x99, 0xe1, 0xe1, 0xca, 0x16, 0x54,
	0x29, 0xda, 0x86, 0x6f, 0xf0, 0xda, 0xad, 0x3f, 0xda, 0x1b, 0xc3, 0xc0, 0x9f, 0xf8, 0x4f, 0x76,
	0x7f, 0x68, 0xb4, 0x04, 0xba, 0xc5, 0xaf, 0x5b, 0x26, 0x66, 0x8b, 0xbf, 0x62, 0xdf, 0x7e, 0x31,
	0xa1, 0x66, 0xc2, 0x9a, 0x09, 0x4d, 0x26, 0x2c, 0x99, 0xc2, 0x3e, 0x6c, 0xcc, 0x84, 0x27, 0xe3,
	0x51, 0xc0, 0x14, 0x8d, 0xa9, 0xa2, 0xdf, 0xf6, 0xc0, 0xe0, 0x09, 0xd1, 0x02, 0x42, 0x6a, 0x03,
	0x21, 0xa6, 0x82, 0x90, 0xd2, 0x31, 0xed, 0x94, 0x7f, 0xd2, 0xfe, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xab, 0x05, 0xcf, 0x65, 0x30, 0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/assetmantle.modules.identities.transactions.issue.Msg/Handle", in, out, opts...)
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
		FullMethod: "/assetmantle.modules.identities.transactions.issue.Msg/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assetmantle.modules.identities.transactions.issue.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Msg_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "identities/transactions/issue/service.proto",
}
