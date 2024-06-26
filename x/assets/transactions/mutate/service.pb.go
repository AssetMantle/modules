// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: AssetMantle/modules/x/assets/transactions/mutate/service.proto

package mutate

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
	proto.RegisterFile("AssetMantle/modules/x/assets/transactions/mutate/service.proto", fileDescriptor_0cec9114e6ce7c03)
}

var fileDescriptor_0cec9114e6ce7c03 = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xd2, 0xcf, 0x4a, 0xeb, 0x40,
	0x14, 0x06, 0xf0, 0x26, 0x85, 0x2e, 0xc2, 0x5d, 0x15, 0x2e, 0x42, 0xd0, 0x2c, 0xfa, 0x00, 0x33,
	0xe2, 0x9f, 0x85, 0xa3, 0x08, 0x2d, 0x88, 0x82, 0x0c, 0x14, 0xed, 0x22, 0x48, 0x40, 0xa6, 0xe9,
	0x10, 0x03, 0xc9, 0x4c, 0xc9, 0x99, 0x4a, 0xb7, 0xfa, 0x04, 0x82, 0x6f, 0xe0, 0x4e, 0x77, 0xbe,
	0x85, 0xb8, 0x2a, 0xb8, 0x71, 0x29, 0xa9, 0x2b, 0x9f, 0x42, 0x9a, 0x19, 0x70, 0xba, 0x9c, 0xed,
	0x09, 0xbf, 0x8f, 0xef, 0x4c, 0x4e, 0x70, 0xdc, 0x07, 0xe0, 0x8a, 0x32, 0xa1, 0x0a, 0x8e, 0x4b,
	0x39, 0x99, 0x15, 0x1c, 0xf0, 0x1c, 0xb3, 0xd5, 0x14, 0xb0, 0xaa, 0x98, 0x00, 0x96, 0xaa, 0x5c,
	0x0a, 0xc0, 0xe5, 0x4c, 0x31, 0xc5, 0x31, 0xf0, 0xea, 0x36, 0x4f, 0x39, 0x9a, 0x56, 0x52, 0xc9,
	0xee, 0xb6, 0xe5, 0x91, 0xf1, 0x68, 0x8e, 0xb4, 0x47, 0xb6, 0x47, 0xda, 0x87, 0x9b, 0x99, 0x94,
	0x59, 0xc1, 0x31, 0x9b, 0xe6, 0x98, 0x09, 0x21, 0x15, 0xd3, 0x9f, 0x9b, 0xbc, 0xd0, 0xbd, 0x4f,
	0xc9, 0x01, 0x58, 0x66, 0xfa, 0x84, 0xe7, 0xce, 0xde, 0x9a, 0x5d, 0x57, 0x1c, 0xa6, 0x52, 0x80,
	0x09, 0xdb, 0x79, 0xf5, 0x82, 0x36, 0x85, 0xac, 0xfb, 0xec, 0x05, 0x9d, 0x33, 0x26, 0x26, 0x05,
	0xef, 0x1e, 0x20, 0xd7, 0x85, 0x11, 0xd5, 0x05, 0xc3, 0x13, 0x77, 0x3a, 0xfa, 0x9b, 0x5d, 0x98,
	0x6a, 0xbd, 0xad, 0xfb, 0x8f, 0xef, 0x47, 0x7f, 0xa3, 0xf7, 0x1f, 0x97, 0x7a, 0x4b, 0xb3, 0x9b,
	0x26, 0x83, 0xbb, 0xf6, 0x5b, 0x1d, 0x79, 0x8b, 0x3a, 0xf2, 0xbe, 0xea, 0xc8, 0x7b, 0x58, 0x46,
	0xad, 0xc5, 0x32, 0x6a, 0x7d, 0x2e, 0xa3, 0x56, 0xb0, 0x97, 0xca, 0xd2, 0xb9, 0xc3, 0xe0, 0xdf,
	0xa5, 0xfe, 0xe1, 0xc3, 0xd5, 0x93, 0x0c, 0xbd, 0xab, 0xc3, 0x2c, 0x57, 0x37, 0xb3, 0x31, 0x4a,
	0x65, 0x89, 0x5d, 0x1f, 0xfb, 0xc9, 0xef, 0xf4, 0x69, 0xdc, 0x1f, 0xd1, 0x17, 0x7f, 0xed, 0x70,
	0xa8, 0x29, 0x12, 0xeb, 0x7a, 0x60, 0x2f, 0x0e, 0x88, 0x36, 0xf6, 0x7d, 0x8d, 0x24, 0x86, 0x24,
	0x71, 0xa2, 0x49, 0x62, 0x93, 0x44, 0x93, 0xda, 0x3f, 0x72, 0x25, 0xc9, 0xe9, 0x70, 0x40, 0xb9,
	0x62, 0x13, 0xa6, 0xd8, 0x8f, 0xbf, 0x6f, 0x71, 0x42, 0x8c, 0x27, 0x24, 0x26, 0x44, 0x27, 0x10,
	0x62, 0x47, 0x10, 0xa2, 0x33, 0xc6, 0x9d, 0xe6, 0x7c, 0x76, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x80, 0x6e, 0xa3, 0xb9, 0x5d, 0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/AssetMantle.modules.x.assets.transactions.mutate.Msg/Handle", in, out, opts...)
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
		FullMethod: "/AssetMantle.modules.x.assets.transactions.mutate.Msg/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AssetMantle.modules.x.assets.transactions.mutate.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Msg_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "AssetMantle/modules/x/assets/transactions/mutate/service.proto",
}
