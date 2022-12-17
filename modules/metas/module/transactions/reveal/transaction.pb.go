// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/metas/module/transactions/reveal/transaction.proto

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
	proto.RegisterFile("modules/metas/module/transactions/reveal/transaction.proto", fileDescriptor_f8ff0aa578548ac1)
}

var fileDescriptor_f8ff0aa578548ac1 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0xca, 0xcd, 0x4f, 0x29,
	0xcd, 0x49, 0x2d, 0xd6, 0xcf, 0x4d, 0x2d, 0x49, 0x2c, 0xd6, 0x87, 0xf0, 0xf4, 0x4b, 0x8a, 0x12,
	0xf3, 0x8a, 0x13, 0x93, 0x4b, 0x32, 0xf3, 0xf3, 0x8a, 0xf5, 0x8b, 0x52, 0xcb, 0x52, 0x13, 0x73,
	0x90, 0xc5, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xd8, 0x20, 0x32, 0x52, 0x66, 0x44, 0x9b,
	0x91, 0x9b, 0x5a, 0x5c, 0x9c, 0x98, 0x9e, 0x0a, 0xd1, 0x2f, 0xe5, 0x44, 0x8e, 0xdd, 0x41, 0xa9,
	0xc5, 0x05, 0xf9, 0x79, 0xc5, 0x30, 0x33, 0x64, 0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x13,
	0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0xc1, 0x1a, 0x21, 0xb2, 0x46, 0xc9, 0x5c,
	0xdc, 0x21, 0x08, 0xad, 0x42, 0x21, 0x5c, 0x6c, 0x41, 0x60, 0x03, 0x85, 0xf8, 0xf5, 0x20, 0x26,
	0xeb, 0xf9, 0x42, 0x5c, 0x24, 0x25, 0x0d, 0x13, 0x08, 0xc1, 0xb4, 0x4a, 0x49, 0xa6, 0xe9, 0xf2,
	0x93, 0xc9, 0x4c, 0x62, 0x4a, 0x22, 0xfa, 0xb9, 0x89, 0x79, 0x25, 0x39, 0xa9, 0x50, 0x07, 0x43,
	0x74, 0x38, 0x6d, 0x67, 0x3c, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4,
	0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x06, 0x2e, 0xae,
	0xe4, 0xfc, 0x5c, 0xa8, 0xc1, 0x4e, 0x02, 0x48, 0x26, 0x07, 0x80, 0x5c, 0x17, 0xc0, 0x18, 0xe5,
	0x9e, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0xef, 0x58, 0x5c, 0x9c, 0x5a,
	0xe2, 0x0b, 0x35, 0x1f, 0x16, 0x30, 0x44, 0x06, 0xd0, 0x22, 0x26, 0xe6, 0xa0, 0x88, 0x88, 0x55,
	0x4c, 0x50, 0xef, 0x9d, 0x82, 0x31, 0x1e, 0x31, 0x09, 0x41, 0x18, 0x31, 0xee, 0x01, 0x4e, 0xbe,
	0xa9, 0x25, 0x89, 0x29, 0x89, 0x25, 0x89, 0xaf, 0x60, 0xb2, 0x49, 0x6c, 0xe0, 0x50, 0x32, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x95, 0xed, 0x70, 0x05, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TransactionClient is the client API for Transaction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransactionClient interface {
	Reveal(ctx context.Context, in *Message, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type transactionClient struct {
	cc grpc1.ClientConn
}

func NewTransactionClient(cc grpc1.ClientConn) TransactionClient {
	return &transactionClient{cc}
}

func (c *transactionClient) Reveal(ctx context.Context, in *Message, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/reveal.Transaction/Reveal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServer is the server API for Transaction service.
type TransactionServer interface {
	Reveal(context.Context, *Message) (*TransactionResponse, error)
}

// UnimplementedTransactionServer can be embedded to have forward compatible implementations.
type UnimplementedTransactionServer struct {
}

func (*UnimplementedTransactionServer) Reveal(ctx context.Context, req *Message) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reveal not implemented")
}

func RegisterTransactionServer(s grpc1.Server, srv TransactionServer) {
	s.RegisterService(&_Transaction_serviceDesc, srv)
}

func _Transaction_Reveal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).Reveal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reveal.Transaction/Reveal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).Reveal(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Transaction_serviceDesc = grpc.ServiceDesc{
	ServiceName: "reveal.Transaction",
	HandlerType: (*TransactionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Reveal",
			Handler:    _Transaction_Reveal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/metas/module/transactions/reveal/transaction.proto",
}