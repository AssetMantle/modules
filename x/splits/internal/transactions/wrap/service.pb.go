// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/splits/internal/transactions/wrap/service.proto

package wrap

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
	proto.RegisterFile("x/splits/internal/transactions/wrap/service.proto", fileDescriptor_d49a5b68eac05057)
}

var fileDescriptor_d49a5b68eac05057 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x4b, 0xfb, 0x40,
	0x18, 0xc6, 0x9b, 0xfc, 0xa1, 0x7f, 0x08, 0x4e, 0x15, 0xa1, 0xc4, 0x12, 0xb0, 0xb8, 0x7a, 0x87,
	0xba, 0x05, 0x1c, 0xcc, 0xa2, 0x4b, 0xa1, 0xd8, 0x42, 0x41, 0xba, 0xbc, 0x4d, 0x8f, 0x78, 0x70,
	0xb9, 0x3b, 0xee, 0xbd, 0x5a, 0xe7, 0x8e, 0x4e, 0x82, 0xdf, 0xc0, 0xd1, 0x4f, 0x22, 0x4e, 0x05,
	0x17, 0x47, 0x49, 0x9d, 0xfc, 0x14, 0x92, 0x5e, 0x0b, 0x01, 0x09, 0x74, 0xcd, 0xf3, 0x7b, 0xf8,
	0x3d, 0x6f, 0x2e, 0x38, 0x7d, 0xa0, 0xa8, 0x05, 0xb7, 0x48, 0xb9, 0xb4, 0xcc, 0x48, 0x10, 0xd4,
	0x1a, 0x90, 0x08, 0xa9, 0xe5, 0x4a, 0x22, 0x9d, 0x1b, 0xd0, 0x14, 0x99, 0xb9, 0xe7, 0x29, 0x23,
	0xda, 0x28, 0xab, 0x5a, 0x6d, 0x57, 0x20, 0x55, 0x8e, 0x94, 0x5c, 0xd8, 0xc9, 0x94, 0xca, 0x04,
	0xa3, 0xa0, 0x39, 0x05, 0x29, 0x95, 0x05, 0x17, 0xae, 0x7b, 0xe1, 0x4e, 0xaa, 0x9c, 0x21, 0x42,
	0xb6, 0x51, 0x85, 0x17, 0xbb, 0x54, 0x2a, 0x5f, 0x6e, 0x18, 0x6a, 0x25, 0x71, 0x53, 0x3f, 0x5b,
	0x78, 0xc1, 0xff, 0x81, 0xdb, 0xde, 0x9a, 0x07, 0xcd, 0x6b, 0x90, 0x53, 0xc1, 0x5a, 0x47, 0xa4,
	0xee, 0x00, 0xd2, 0x73, 0xf6, 0xf0, 0xa4, 0x1e, 0x19, 0xfe, 0xb5, 0x75, 0x0f, 0x17, 0x1f, 0xdf,
	0xcf, 0xfe, 0x41, 0x77, 0x9f, 0xe6, 0x20, 0xad, 0x60, 0xdb, 0xd1, 0x65, 0x21, 0x79, 0xf4, 0xdf,
	0x8a, 0xc8, 0x5b, 0x16, 0x91, 0xf7, 0x55, 0x44, 0xde, 0xd3, 0x2a, 0x6a, 0x2c, 0x57, 0x51, 0xe3,
	0x73, 0x15, 0x35, 0x82, 0x4e, 0xaa, 0xf2, 0x5a, 0x53, 0xb2, 0xb7, 0x99, 0xde, 0x2f, 0x6f, 0xe9,
	0x7b, 0xb7, 0x49, 0xc6, 0xed, 0xdd, 0x6c, 0x42, 0x52, 0x95, 0xd3, 0x4b, 0x44, 0x66, 0x7b, 0x4e,
	0x96, 0xab, 0xe9, 0x4c, 0x30, 0xa4, 0x3b, 0xfc, 0xab, 0x17, 0xff, 0xdf, 0x60, 0x38, 0x7a, 0xf5,
	0xdb, 0x03, 0x27, 0x1d, 0x56, 0xa5, 0x23, 0x03, 0xfa, 0x7d, 0x1b, 0x8d, 0xab, 0xd1, 0xb8, 0x8c,
	0x0a, 0xff, 0xb8, 0x2e, 0x1a, 0x5f, 0xf5, 0x93, 0x1e, 0xb3, 0x30, 0x05, 0x0b, 0x3f, 0x7e, 0xe8,
	0xb0, 0x38, 0xae, 0x72, 0x71, 0x5c, 0x82, 0x93, 0xe6, 0xfa, 0x61, 0xce, 0x7f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0xc2, 0x20, 0x9b, 0x57, 0x77, 0x02, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/splits.transactions.wrap.Service/Handle", in, out, opts...)
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
		FullMethod: "/splits.transactions.wrap.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "splits.transactions.wrap.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "x/splits/internal/transactions/wrap/service.proto",
}