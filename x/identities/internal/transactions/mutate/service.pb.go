// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/identities/internal/transactions/mutate/service.proto

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
	proto.RegisterFile("x/identities/internal/transactions/mutate/service.proto", fileDescriptor_a3f5029c61eea546)
}

var fileDescriptor_a3f5029c61eea546 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x31, 0x4b, 0xc3, 0x40,
	0x18, 0x86, 0x9b, 0x08, 0x15, 0x82, 0x53, 0x37, 0xa3, 0x1c, 0xda, 0x45, 0x70, 0xc8, 0x81, 0x1d,
	0x84, 0x6c, 0xd6, 0xc1, 0x76, 0x08, 0x14, 0xed, 0x24, 0x5d, 0xbe, 0x26, 0x1f, 0xf1, 0x20, 0x77,
	0x57, 0x72, 0x5f, 0xc5, 0xd9, 0xcd, 0x4d, 0x10, 0xff, 0x80, 0xa3, 0x83, 0xbf, 0x43, 0x9c, 0x0a,
	0x2e, 0x8e, 0x92, 0x3a, 0xf9, 0x2b, 0xc4, 0x5e, 0xa1, 0x07, 0x42, 0x5b, 0xd7, 0xe4, 0x79, 0xef,
	0x7d, 0x9f, 0xbb, 0xe0, 0xf8, 0x86, 0x8b, 0x0c, 0x15, 0x09, 0x12, 0x68, 0xb8, 0x50, 0x84, 0xa5,
	0x82, 0x82, 0x53, 0x09, 0xca, 0x40, 0x4a, 0x42, 0x2b, 0xc3, 0xe5, 0x98, 0x80, 0x90, 0x1b, 0x2c,
	0xaf, 0x45, 0x8a, 0xd1, 0xa8, 0xd4, 0xa4, 0x1b, 0x6c, 0x11, 0x8b, 0x5c, 0x3a, 0xb2, 0x74, 0xb8,
	0x9b, 0x6b, 0x9d, 0x17, 0xc8, 0x61, 0x24, 0x38, 0x28, 0xa5, 0x09, 0xec, 0xef, 0x59, 0x3a, 0xfc,
	0x47, 0xad, 0x44, 0x63, 0x20, 0x9f, 0xd7, 0x86, 0xa7, 0xeb, 0x07, 0x9d, 0x6f, 0xe7, 0x68, 0x46,
	0x5a, 0x99, 0xf9, 0x21, 0x47, 0x8f, 0x5e, 0xb0, 0x79, 0x61, 0x6d, 0x1a, 0x77, 0x5e, 0x50, 0xef,
	0x80, 0xca, 0x0a, 0x6c, 0x1c, 0x44, 0xcb, 0x9d, 0xa2, 0xc4, 0x4e, 0x09, 0x5b, 0xab, 0xc0, 0xfe,
	0xdf, 0xea, 0xe6, 0xfe, 0xed, 0xfb, 0xd7, 0x83, 0xbf, 0xd3, 0xdc, 0xe6, 0x12, 0x14, 0x15, 0xe8,
	0x7a, 0xd8, 0x58, 0xfb, 0xc5, 0x7f, 0xad, 0x98, 0x37, 0xa9, 0x98, 0xf7, 0x59, 0x31, 0xef, 0x7e,
	0xca, 0x6a, 0x93, 0x29, 0xab, 0x7d, 0x4c, 0x59, 0x2d, 0x68, 0xa6, 0x5a, 0xae, 0x68, 0x6d, 0x6f,
	0xcd, 0x9d, 0x7a, 0xbf, 0x92, 0x3d, 0xef, 0xb2, 0x93, 0x0b, 0xba, 0x1a, 0x0f, 0xa3, 0x54, 0x4b,
	0x7e, 0x62, 0x0c, 0x52, 0x62, 0x8b, 0xa5, 0xce, 0xc6, 0x05, 0x1a, 0xbe, 0xf6, 0x55, 0x3e, 0xf9,
	0x1b, 0xdd, 0x7e, 0xf2, 0xec, 0xb3, 0xee, 0x62, 0x40, 0xdf, 0x1d, 0x90, 0xcc, 0xb0, 0x37, 0x17,
	0x18, 0xb8, 0xc0, 0xc0, 0x02, 0x95, 0x7f, 0xb8, 0x1c, 0x18, 0x9c, 0xf5, 0xda, 0x09, 0x12, 0x64,
	0x40, 0xf0, 0xed, 0xef, 0x2d, 0xe0, 0x38, 0x76, 0xe9, 0x38, 0xb6, 0xf8, 0xb0, 0x3e, 0x7b, 0xcf,
	0xd6, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x62, 0x2b, 0xdf, 0xd0, 0xc6, 0x02, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/identities.transactions.mutate.Service/Handle", in, out, opts...)
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
		FullMethod: "/identities.transactions.mutate.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "identities.transactions.mutate.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "x/identities/internal/transactions/mutate/service.proto",
}