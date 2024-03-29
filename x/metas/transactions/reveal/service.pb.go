// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: metas/transactions/reveal/service.proto

package reveal

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
	proto.RegisterFile("metas/transactions/reveal/service.proto", fileDescriptor_d86eed0e8d50b71f)
}

var fileDescriptor_d86eed0e8d50b71f = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x31, 0x4b, 0xf3, 0x40,
	0x18, 0xc7, 0x9b, 0xbc, 0xbc, 0x1d, 0x82, 0x53, 0x11, 0x87, 0x50, 0x32, 0x74, 0x71, 0xea, 0x1d,
	0x55, 0x11, 0xbc, 0xad, 0x59, 0x74, 0x39, 0x28, 0xb5, 0x93, 0x04, 0xe4, 0x69, 0x7a, 0xc4, 0x40,
	0x72, 0x57, 0x72, 0xd7, 0xe2, 0xec, 0x27, 0x10, 0xfc, 0x06, 0x1d, 0xfb, 0x49, 0xc4, 0xa9, 0xe0,
	0xe2, 0x28, 0xa9, 0x93, 0x8b, 0x5f, 0x41, 0x7a, 0x77, 0xe0, 0x39, 0x54, 0xc8, 0x7a, 0xf7, 0xff,
	0xfd, 0x9f, 0x5f, 0x9e, 0x5c, 0x70, 0x5c, 0x32, 0x05, 0x12, 0xab, 0x0a, 0xb8, 0x84, 0x54, 0xe5,
	0x82, 0x4b, 0x5c, 0xb1, 0x25, 0x83, 0x02, 0x4b, 0x56, 0x2d, 0xf3, 0x94, 0xa1, 0x79, 0x25, 0x94,
	0xe8, 0xf4, 0x41, 0x4a, 0xa6, 0x4a, 0xe0, 0xaa, 0x60, 0xa8, 0x14, 0xb3, 0x45, 0xc1, 0x24, 0xd2,
	0x30, 0x72, 0x61, 0x64, 0xe0, 0xb0, 0x9b, 0x09, 0x91, 0x15, 0x0c, 0xc3, 0x3c, 0xc7, 0xc0, 0xb9,
	0x50, 0x60, 0xae, 0x75, 0x59, 0xf8, 0xc7, 0xd4, 0x92, 0x49, 0x09, 0x99, 0x9d, 0x1a, 0x9e, 0xed,
	0x0f, 0x3a, 0x67, 0xb7, 0x15, 0x93, 0x73, 0xc1, 0xa5, 0xa5, 0x4e, 0xd6, 0x5e, 0xf0, 0x8f, 0xca,
	0xac, 0xb3, 0xf2, 0x82, 0xf6, 0x15, 0xf0, 0x59, 0xc1, 0x3a, 0xe7, 0xa8, 0x91, 0x3f, 0xa2, 0x46,
	0x23, 0x8c, 0x1b, 0x72, 0x93, 0x9f, 0xb3, 0xb1, 0x95, 0xea, 0x75, 0x1f, 0x5e, 0x3f, 0x9e, 0xfc,
	0xa3, 0xde, 0x21, 0x36, 0x35, 0xd8, 0x7c, 0x94, 0x21, 0xe2, 0x2f, 0xff, 0xb9, 0x8e, 0xbc, 0x4d,
	0x1d, 0x79, 0xef, 0x75, 0xe4, 0x3d, 0x6e, 0xa3, 0xd6, 0x66, 0x1b, 0xb5, 0xde, 0xb6, 0x51, 0x2b,
	0x18, 0xa4, 0xa2, 0x6c, 0x36, 0x3f, 0x3e, 0xb8, 0x36, 0x7f, 0x6d, 0xb4, 0x5b, 0xc4, 0xc8, 0xbb,
	0x21, 0x59, 0xae, 0xee, 0x16, 0x53, 0x94, 0x8a, 0x12, 0x0f, 0x77, 0x4d, 0xd4, 0x2a, 0x98, 0x26,
	0x7c, 0x8f, 0xf7, 0x6e, 0x78, 0xe5, 0xff, 0x1f, 0x52, 0x3a, 0x19, 0xaf, 0xfd, 0xfe, 0xd0, 0x71,
	0xa0, 0xd6, 0x81, 0x6a, 0x87, 0x89, 0xeb, 0x30, 0xd6, 0xdc, 0xcb, 0xaf, 0x7c, 0x62, 0xf3, 0x89,
	0xce, 0x27, 0x6e, 0x3e, 0x31, 0xf9, 0xda, 0xbf, 0x68, 0x94, 0x4f, 0x2e, 0x47, 0xf1, 0xee, 0x72,
	0x06, 0x0a, 0x3e, 0xfd, 0x81, 0xc3, 0x12, 0x62, 0x61, 0x42, 0x34, 0x4d, 0x88, 0x8b, 0x13, 0x62,
	0xf8, 0x69, 0x5b, 0xbf, 0x92, 0xd3, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x6f, 0x2f, 0x1f,
	0xfc, 0x02, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/assetmantle.modules.metas.transactions.reveal.Msg/Handle", in, out, opts...)
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
		FullMethod: "/assetmantle.modules.metas.transactions.reveal.Msg/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assetmantle.modules.metas.transactions.reveal.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Msg_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metas/transactions/reveal/service.proto",
}
