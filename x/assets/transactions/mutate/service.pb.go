// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: assets/transactions/mutate/service.proto

package mutate

import (
	context "context"
	fmt "fmt"
	math "math"

	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
	proto.RegisterFile("assets/transactions/mutate/service.proto", fileDescriptor_39a94abb379d69cb)
}

var fileDescriptor_39a94abb379d69cb = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xd2, 0x31, 0x4b, 0xf3, 0x40,
	0x18, 0x07, 0xf0, 0x26, 0xf0, 0xf6, 0x85, 0xe0, 0x54, 0x10, 0x21, 0x68, 0x86, 0x4e, 0x4e, 0x77,
	0x50, 0x11, 0xe1, 0x9c, 0x52, 0x07, 0x5d, 0x0e, 0x8a, 0x76, 0x92, 0x80, 0x5c, 0xd3, 0x23, 0x06,
	0x92, 0xbb, 0x92, 0xe7, 0x2a, 0xce, 0x7e, 0x02, 0xc1, 0x6f, 0x20, 0xb8, 0xb8, 0xf9, 0x2d, 0xc4,
	0xa9, 0xe0, 0xe2, 0x28, 0xa9, 0x93, 0xbb, 0xbb, 0x34, 0xcf, 0x81, 0xe7, 0x60, 0x21, 0xeb, 0x93,
	0xfb, 0xfd, 0xef, 0xff, 0x1c, 0x09, 0x76, 0x05, 0x80, 0x34, 0x40, 0x4d, 0x25, 0x14, 0x88, 0xd4,
	0xe4, 0x5a, 0x01, 0x2d, 0xe7, 0x46, 0x18, 0x49, 0x41, 0x56, 0x57, 0x79, 0x2a, 0xc9, 0xac, 0xd2,
	0x46, 0xf7, 0x48, 0x73, 0xb2, 0x14, 0xca, 0x14, 0x92, 0x94, 0x7a, 0x3a, 0x2f, 0x24, 0xe0, 0x0c,
	0x88, 0xab, 0x09, 0xea, 0x70, 0x3b, 0xd3, 0x3a, 0x2b, 0x24, 0x15, 0xb3, 0x9c, 0x0a, 0xa5, 0xb4,
	0x11, 0xf8, 0xb9, 0x49, 0x0b, 0xd7, 0xdd, 0x5b, 0x4a, 0x00, 0x91, 0xd9, 0x7b, 0xc3, 0xfd, 0x35,
	0x27, 0x9d, 0xd9, 0x45, 0x25, 0x61, 0xa6, 0x15, 0x58, 0x36, 0x78, 0xf2, 0x82, 0xff, 0x67, 0xb8,
	0x40, 0xef, 0xc1, 0x0b, 0xba, 0x27, 0x42, 0x4d, 0x0b, 0xd9, 0x3b, 0x68, 0xb9, 0x06, 0xe1, 0x58,
	0x26, 0x3c, 0x6a, 0x0b, 0xc7, 0x3f, 0xb3, 0x53, 0x5b, 0xad, 0xbf, 0x73, 0xf3, 0xfa, 0x71, 0xe7,
	0x6f, 0xf5, 0x37, 0x29, 0xe6, 0x50, 0xbb, 0x1b, 0x92, 0xe1, 0x97, 0xff, 0x5c, 0x47, 0xde, 0xa2,
	0x8e, 0xbc, 0xf7, 0x3a, 0xf2, 0x6e, 0x97, 0x51, 0x67, 0xb1, 0x8c, 0x3a, 0x6f, 0xcb, 0xa8, 0x13,
	0x0c, 0x52, 0x5d, 0xb6, 0x6c, 0x30, 0xdc, 0xb0, 0xfb, 0x8f, 0x56, 0x0f, 0x32, 0xf2, 0xce, 0x0f,
	0xb3, 0xdc, 0x5c, 0xce, 0x27, 0x24, 0xd5, 0x25, 0x8d, 0x57, 0x8c, 0x63, 0x09, 0x1b, 0x45, 0xaf,
	0xe9, 0xdf, 0x4f, 0x7d, 0xef, 0xff, 0x8b, 0x79, 0x3c, 0xe6, 0x8f, 0x3e, 0x89, 0x9d, 0x16, 0xdc,
	0xb6, 0x88, 0xb1, 0xc5, 0xd8, 0x6d, 0xc1, 0x1b, 0xf8, 0xf2, 0x0b, 0x24, 0x16, 0x24, 0x08, 0x12,
	0x17, 0x24, 0x08, 0x6a, 0x9f, 0xb5, 0x03, 0xc9, 0xf1, 0x68, 0xc8, 0xa5, 0x11, 0x53, 0x61, 0xc4,
	0xa7, 0x3f, 0x70, 0x30, 0x63, 0x56, 0x33, 0x8c, 0x04, 0xc6, 0x5c, 0xcf, 0x18, 0x06, 0x4c, 0xba,
	0xcd, 0x2f, 0xb3, 0xf7, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x19, 0xf9, 0x1a, 0xf7, 0x0d, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	Handle(ctx context.Context, in *Message, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type serviceClient struct {
	cc grpc1.ClientConn
}

func NewServiceClient(cc grpc1.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Handle(ctx context.Context, in *Message, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/assetmantle.modules.assets.transactions.mutate.Service/Handle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	Handle(context.Context, *Message) (*TransactionResponse, error)
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) Handle(ctx context.Context, req *Message) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handle not implemented")
}

func RegisterServiceServer(s grpc1.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_Handle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assetmantle.modules.assets.transactions.mutate.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assetmantle.modules.assets.transactions.mutate.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "assets/transactions/mutate/service.proto",
}