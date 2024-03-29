// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: identities/transactions/define/service.proto

package define

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
	proto.RegisterFile("identities/transactions/define/service.proto", fileDescriptor_efe0844c12310339)
}

var fileDescriptor_efe0844c12310339 = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x31, 0x4b, 0xeb, 0x40,
	0x1c, 0xc0, 0x9b, 0x94, 0xd7, 0x21, 0xbc, 0xa9, 0xdb, 0xcb, 0x7b, 0x04, 0x5e, 0x67, 0xb9, 0x83,
	0x0a, 0x05, 0x4f, 0x50, 0x5a, 0x0a, 0xb5, 0x43, 0xa0, 0x68, 0x27, 0x09, 0xc8, 0x35, 0xf9, 0x1b,
	0x0f, 0x92, 0xbb, 0x92, 0xbb, 0x8a, 0xb3, 0xbb, 0x20, 0xf8, 0x0d, 0x1c, 0xdd, 0x1c, 0xfd, 0x06,
	0xe2, 0x54, 0x70, 0x71, 0x94, 0xd4, 0xc9, 0x4f, 0x21, 0xed, 0x1d, 0xe4, 0x5c, 0x2c, 0xed, 0x7a,
	0xc9, 0xef, 0x77, 0xbf, 0xff, 0x3f, 0xf1, 0x76, 0x58, 0x02, 0x5c, 0x31, 0xc5, 0x40, 0x62, 0x55,
	0x50, 0x2e, 0x69, 0xac, 0x98, 0xe0, 0x12, 0x27, 0x70, 0xce, 0x38, 0x60, 0x09, 0xc5, 0x25, 0x8b,
	0x01, 0x4d, 0x0b, 0xa1, 0x44, 0xb3, 0x4d, 0xa5, 0x04, 0x95, 0x53, 0xae, 0x32, 0x40, 0xb9, 0x48,
	0x66, 0x19, 0x48, 0x54, 0x19, 0x90, 0x6d, 0x40, 0xda, 0xe0, 0xff, 0x4b, 0x85, 0x48, 0x33, 0xc0,
	0x74, 0xca, 0x30, 0xe5, 0x5c, 0x28, 0xaa, 0x1f, 0xaf, 0x8c, 0xfe, 0xba, 0xfb, 0x73, 0x90, 0x92,
	0xa6, 0xe6, 0x7e, 0x7f, 0x6f, 0xcd, 0xdb, 0xd6, 0xd9, 0x59, 0x01, 0x72, 0x2a, 0xb8, 0x34, 0x68,
	0xfb, 0xc9, 0xf1, 0xea, 0xa1, 0x4c, 0x9b, 0x8f, 0x8e, 0xd7, 0x38, 0xa2, 0x3c, 0xc9, 0xa0, 0xb9,
	0x8f, 0x36, 0x1f, 0x07, 0x85, 0x3a, 0xc8, 0x1f, 0x6c, 0x03, 0x8f, 0xab, 0xb3, 0x63, 0x93, 0xd7,
	0xfa, 0x7f, 0xfd, 0xfa, 0x71, 0xe7, 0xfe, 0x6d, 0xfd, 0xc1, 0xda, 0x85, 0xad, 0x19, 0x35, 0xd6,
	0xbb, 0xa9, 0x3f, 0x97, 0x81, 0x33, 0x2f, 0x03, 0xe7, 0xbd, 0x0c, 0x9c, 0xdb, 0x45, 0x50, 0x9b,
	0x2f, 0x82, 0xda, 0xdb, 0x22, 0xa8, 0x79, 0x9d, 0x58, 0xe4, 0x5b, 0x94, 0xf4, 0x7e, 0x9f, 0xe8,
	0x0f, 0x3b, 0x5a, 0x2e, 0x67, 0xe4, 0x9c, 0x1e, 0xa6, 0x4c, 0x5d, 0xcc, 0x26, 0x28, 0x16, 0x39,
	0xee, 0x2e, 0x75, 0xa1, 0x8e, 0x31, 0x3a, 0x7c, 0x85, 0x7f, 0x5e, 0xfd, 0xbd, 0xfb, 0xab, 0x1b,
	0x0e, 0xc7, 0xfd, 0x07, 0xb7, 0xdd, 0xb5, 0x6a, 0x42, 0x53, 0x33, 0xac, 0x6a, 0xc6, 0x76, 0x4d,
	0x7f, 0x05, 0xbf, 0x7c, 0x83, 0x22, 0x03, 0x45, 0x15, 0x14, 0xd9, 0x50, 0xa4, 0xa1, 0xd2, 0x3d,
	0xd8, 0x1c, 0x8a, 0x06, 0xa3, 0x5e, 0x08, 0x8a, 0x26, 0x54, 0xd1, 0x4f, 0xb7, 0x63, 0x09, 0x08,
	0x31, 0x06, 0x42, 0x2a, 0x05, 0x21, 0xb6, 0x83, 0x10, 0x2d, 0x99, 0x34, 0x56, 0xbf, 0xd4, 0xee,
	0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x51, 0x9c, 0x1b, 0x3d, 0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/assetmantle.modules.identities.transactions.define.Msg/Handle", in, out, opts...)
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
		FullMethod: "/assetmantle.modules.identities.transactions.define.Msg/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assetmantle.modules.identities.transactions.define.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Msg_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "identities/transactions/define/service.proto",
}
