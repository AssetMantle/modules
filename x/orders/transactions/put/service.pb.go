// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: orders/transactions/put/service.proto

package put

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
	proto.RegisterFile("orders/transactions/put/service.proto", fileDescriptor_58e16b380821a1be)
}

var fileDescriptor_58e16b380821a1be = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x31, 0x4b, 0xc3, 0x40,
	0x14, 0xc7, 0x9b, 0x88, 0x1d, 0x82, 0x53, 0x70, 0x0a, 0x92, 0xa1, 0xe0, 0x54, 0xb8, 0x83, 0x2a,
	0x28, 0x37, 0xd9, 0x2e, 0xba, 0x84, 0x06, 0xed, 0x24, 0x01, 0xb9, 0x26, 0x47, 0x0c, 0x24, 0x77,
	0x21, 0xf7, 0x4e, 0x9c, 0xfd, 0x04, 0x82, 0xab, 0x93, 0x38, 0xf9, 0x49, 0xc4, 0xa9, 0xe0, 0xe2,
	0x28, 0x89, 0x93, 0x9f, 0x42, 0xea, 0x1d, 0x78, 0x0e, 0x1d, 0xb2, 0xbe, 0xfc, 0x7f, 0xff, 0xdf,
	0xcb, 0x4b, 0xbc, 0x7d, 0xd1, 0x64, 0xac, 0x91, 0x18, 0x1a, 0xca, 0x25, 0x4d, 0xa1, 0x10, 0x5c,
	0xe2, 0x5a, 0x01, 0x96, 0xac, 0xb9, 0x29, 0x52, 0x86, 0xea, 0x46, 0x80, 0xf0, 0xc7, 0x54, 0x4a,
	0x06, 0x15, 0xe5, 0x50, 0x32, 0x54, 0x89, 0x4c, 0x95, 0x4c, 0x22, 0x8d, 0x22, 0x1b, 0x45, 0xb5,
	0x82, 0x60, 0x2f, 0x17, 0x22, 0x2f, 0x19, 0xa6, 0x75, 0x81, 0x29, 0xe7, 0x02, 0xa8, 0x79, 0xb6,
	0xae, 0x0a, 0x36, 0x1a, 0x2b, 0x26, 0x25, 0xcd, 0x8d, 0x31, 0x98, 0x6c, 0x8a, 0x59, 0x83, 0xab,
	0x86, 0xc9, 0x5a, 0x70, 0x69, 0x98, 0xc9, 0xb3, 0xe3, 0x6d, 0x45, 0x32, 0xf7, 0x1f, 0x1d, 0x6f,
	0x78, 0x46, 0x79, 0x56, 0x32, 0xff, 0x10, 0xf5, 0xd8, 0x1c, 0x45, 0x7a, 0x85, 0xe0, 0xa4, 0x17,
	0xb5, 0xf8, 0x1b, 0x9c, 0x9b, 0x85, 0x46, 0xc1, 0xdd, 0xfb, 0xd7, 0x83, 0xbb, 0x3b, 0xf2, 0xb1,
	0x2e, 0xc1, 0xe6, 0x75, 0x6a, 0x05, 0xb3, 0xce, 0x7d, 0x6d, 0x43, 0x67, 0xd5, 0x86, 0xce, 0x67,
	0x1b, 0x3a, 0xf7, 0x5d, 0x38, 0x58, 0x75, 0xe1, 0xe0, 0xa3, 0x0b, 0x07, 0x1e, 0x4e, 0x45, 0xd5,
	0xc7, 0x3d, 0xdb, 0xb9, 0xd0, 0xdf, 0x29, 0x5e, 0x1f, 0x20, 0x76, 0x2e, 0x8f, 0xf3, 0x02, 0xae,
	0xd5, 0x12, 0xa5, 0xa2, 0xc2, 0xd3, 0x75, 0x4f, 0xa4, 0xf5, 0xa6, 0x07, 0xdf, 0xe2, 0x0d, 0x77,
	0x7d, 0x72, 0xb7, 0xa7, 0xd1, 0x7c, 0x11, 0xbf, 0xb8, 0xe3, 0xa9, 0xe5, 0x8f, 0x8c, 0x7f, 0xae,
	0xfd, 0x0b, 0xdb, 0x1f, 0x2b, 0x78, 0xfb, 0x97, 0x4e, 0x4c, 0x3a, 0xd1, 0xe9, 0xc4, 0x4e, 0x27,
	0xb1, 0x82, 0xd6, 0x3d, 0xea, 0x91, 0x4e, 0x4e, 0xe3, 0x59, 0xc4, 0x80, 0x66, 0x14, 0xe8, 0xb7,
	0x8b, 0x2d, 0x92, 0x10, 0x83, 0x12, 0xa2, 0x59, 0x42, 0x6c, 0x98, 0x90, 0x58, 0xc1, 0x72, 0xf8,
	0xfb, 0x4f, 0x1c, 0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb8, 0xac, 0xc9, 0xb0, 0xe2, 0x02, 0x00,
	0x00,
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
	err := c.cc.Invoke(ctx, "/assetmantle.modules.orders.transactions.put.Msg/Handle", in, out, opts...)
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
		FullMethod: "/assetmantle.modules.orders.transactions.put.Msg/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assetmantle.modules.orders.transactions.put.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Msg_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orders/transactions/put/service.proto",
}
