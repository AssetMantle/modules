// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: AssetMantle/modules/x/identities/transactions/update/service.proto

package update

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
	proto.RegisterFile("AssetMantle/modules/x/identities/transactions/update/service.proto", fileDescriptor_188ef3d964dd38f9)
}

var fileDescriptor_188ef3d964dd38f9 = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x3f, 0x4b, 0xc3, 0x40,
	0x14, 0xc0, 0x93, 0x14, 0x3b, 0x04, 0xa7, 0x2e, 0x42, 0x90, 0x0c, 0x8e, 0x0e, 0x77, 0xf8, 0x67,
	0xd0, 0x03, 0xd1, 0x66, 0xd1, 0x0e, 0xc1, 0xa2, 0x2d, 0x04, 0x09, 0xc8, 0x35, 0x39, 0x62, 0x20,
	0xb9, 0x0b, 0x79, 0xd7, 0xd2, 0xd1, 0x8f, 0xe0, 0xe0, 0xe2, 0xaa, 0x9b, 0x93, 0x1f, 0x43, 0x9c,
	0x3a, 0x3a, 0x4a, 0x3a, 0x08, 0x7e, 0x0a, 0x69, 0x2f, 0x90, 0x74, 0xbd, 0x29, 0xf0, 0xc2, 0xef,
	0x77, 0x3f, 0x1e, 0xcf, 0xf6, 0xfa, 0x00, 0x4c, 0xfa, 0x94, 0xcb, 0x8c, 0xe1, 0x5c, 0xc4, 0xd3,
	0x8c, 0x01, 0x9e, 0xe3, 0x34, 0x66, 0x5c, 0xa6, 0x32, 0x65, 0x80, 0x65, 0x49, 0x39, 0xd0, 0x48,
	0xa6, 0x82, 0x03, 0x9e, 0x16, 0x31, 0x95, 0x0c, 0x03, 0x2b, 0x67, 0x69, 0xc4, 0x50, 0x51, 0x0a,
	0x29, 0x7a, 0xc7, 0x2d, 0x07, 0xaa, 0x1d, 0x68, 0x8e, 0x1a, 0x07, 0x6a, 0x3b, 0x90, 0x72, 0x38,
	0x3b, 0x91, 0x80, 0x5c, 0x00, 0xce, 0x21, 0xc1, 0xb3, 0x83, 0xd5, 0x47, 0xe9, 0x9c, 0xdd, 0x44,
	0x88, 0x24, 0x63, 0x98, 0x16, 0x29, 0xa6, 0x9c, 0x0b, 0x49, 0x15, 0xa7, 0xfe, 0xea, 0x05, 0xe7,
	0x0c, 0x80, 0x26, 0x75, 0xb0, 0x73, 0xad, 0xe5, 0x68, 0xcd, 0xee, 0x4b, 0x06, 0x85, 0xe0, 0x50,
	0x0b, 0x0f, 0xdf, 0x4c, 0xbb, 0xe3, 0x43, 0xd2, 0x7b, 0x36, 0xed, 0xee, 0x15, 0xe5, 0x71, 0xc6,
	0x7a, 0x67, 0x48, 0x67, 0x2b, 0xc8, 0x57, 0xa1, 0xce, 0x40, 0x0f, 0x1f, 0x35, 0xb3, 0x9b, 0x3a,
	0x71, 0xcf, 0x70, 0xb6, 0x1e, 0x7f, 0x3f, 0xf6, 0x4d, 0xef, 0xa5, 0xf3, 0x59, 0xb9, 0xe6, 0xa2,
	0x72, 0xcd, 0x9f, 0xca, 0x35, 0x9f, 0x96, 0xae, 0xb1, 0x58, 0xba, 0xc6, 0xf7, 0xd2, 0x35, 0xec,
	0x93, 0x48, 0xe4, 0x5a, 0x2f, 0x7a, 0xdb, 0xb7, 0xea, 0x16, 0x86, 0xab, 0x45, 0x0c, 0xcd, 0xbb,
	0xf3, 0x24, 0x95, 0x0f, 0xd3, 0x09, 0x8a, 0x44, 0x8e, 0x75, 0xd6, 0xfc, 0x6a, 0x75, 0xfb, 0x7e,
	0x30, 0x18, 0x8d, 0xdf, 0xad, 0x8d, 0xbb, 0xf2, 0xeb, 0xa0, 0x00, 0x0d, 0x9a, 0xa0, 0x51, 0x3b,
	0x68, 0xbc, 0xe6, 0xbf, 0x36, 0xb0, 0xb0, 0xc6, 0xc2, 0x20, 0x6c, 0xb0, 0xb0, 0x8d, 0x85, 0x0a,
	0xab, 0xac, 0x0b, 0x1d, 0x2c, 0xbc, 0x1c, 0x7a, 0x3e, 0x93, 0x34, 0xa6, 0x92, 0xfe, 0x59, 0xa7,
	0x2d, 0x05, 0x21, 0xb5, 0x83, 0x90, 0x80, 0x90, 0xc6, 0x42, 0x48, 0x5b, 0x43, 0x88, 0xf2, 0x4c,
	0xba, 0xeb, 0x43, 0x3a, 0xfa, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x31, 0x3c, 0x77, 0x90, 0x03,
	0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/AssetMantle.modules.x.identities.transactions.update.Msg/Handle", in, out, opts...)
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
		FullMethod: "/AssetMantle.modules.x.identities.transactions.update.Msg/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AssetMantle.modules.x.identities.transactions.update.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Msg_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "AssetMantle/modules/x/identities/transactions/update/service.proto",
}
