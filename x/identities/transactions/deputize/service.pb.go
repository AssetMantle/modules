// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: identities/transactions/deputize/service.proto

package deputize

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
	proto.RegisterFile("identities/transactions/deputize/service.proto", fileDescriptor_5ad62bbbb70c71f9)
}

var fileDescriptor_5ad62bbbb70c71f9 = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xb1, 0x4a, 0xe4, 0x40,
	0x18, 0x80, 0x37, 0x39, 0x6e, 0x0f, 0xc2, 0x55, 0x5b, 0xe6, 0xee, 0x52, 0xec, 0xf5, 0x33, 0xa0,
	0x5b, 0x2c, 0x23, 0x82, 0x59, 0x16, 0x74, 0x8b, 0xc0, 0xa2, 0x5b, 0x49, 0x40, 0x66, 0x93, 0x9f,
	0x38, 0x90, 0xcc, 0x84, 0xcc, 0xac, 0x88, 0xa5, 0x4f, 0x20, 0x58, 0xd9, 0x5a, 0xfa, 0x02, 0x3e,
	0x80, 0x8d, 0x58, 0x2d, 0xd8, 0x58, 0x4a, 0xd6, 0xca, 0xa7, 0x90, 0xdd, 0x4c, 0xc8, 0x58, 0x2d,
	0xa4, 0x1d, 0xfe, 0xef, 0xe3, 0xfb, 0xff, 0xc4, 0x41, 0x2c, 0x06, 0xae, 0x98, 0x62, 0x20, 0xb1,
	0x2a, 0x28, 0x97, 0x34, 0x52, 0x4c, 0x70, 0x89, 0x63, 0xc8, 0x17, 0x8a, 0x5d, 0x01, 0x96, 0x50,
	0x5c, 0xb0, 0x08, 0x50, 0x5e, 0x08, 0x25, 0x7a, 0x03, 0x2a, 0x25, 0xa8, 0x8c, 0x72, 0x95, 0x02,
	0xca, 0x44, 0xbc, 0x48, 0x41, 0x1a, 0x0e, 0x64, 0x3a, 0x50, 0xed, 0x70, 0xff, 0x26, 0x42, 0x24,
	0x29, 0x60, 0x9a, 0x33, 0x4c, 0x39, 0x17, 0x8a, 0x56, 0x03, 0x1b, 0xa7, 0xbb, 0xbd, 0x21, 0x03,
	0x29, 0x69, 0xa2, 0x1b, 0xdc, 0xbd, 0xad, 0xf3, 0xc6, 0xeb, 0x59, 0x01, 0x32, 0x17, 0x5c, 0x6a,
	0x78, 0xe7, 0xc9, 0x72, 0x7e, 0x9d, 0x54, 0x2b, 0xf5, 0x1e, 0x2d, 0xa7, 0x7b, 0x44, 0x79, 0x9c,
	0x42, 0x6f, 0x1f, 0xb5, 0x59, 0x0c, 0x05, 0x55, 0x98, 0x3b, 0x69, 0x87, 0xcf, 0x9a, 0xd7, 0x63,
	0x9d, 0xd9, 0xff, 0x7f, 0xfd, 0xfa, 0x71, 0x6b, 0xff, 0xeb, 0xff, 0xc1, 0x95, 0x0d, 0x1b, 0xdb,
	0xd6, 0xe0, 0xe8, 0xee, 0xc7, 0x73, 0xe9, 0x59, 0xcb, 0xd2, 0xb3, 0xde, 0x4b, 0xcf, 0xba, 0x59,
	0x79, 0x9d, 0xe5, 0xca, 0xeb, 0xbc, 0xad, 0xbc, 0x8e, 0x33, 0x8c, 0x44, 0xd6, 0xaa, 0x66, 0xf4,
	0x5b, 0xdf, 0x65, 0xba, 0x3e, 0xd4, 0xd4, 0x3a, 0xf5, 0x13, 0xa6, 0xce, 0x17, 0x73, 0x14, 0x89,
	0x0c, 0xfb, 0x6b, 0x61, 0x50, 0x05, 0x69, 0x21, 0xbe, 0xc4, 0xdb, 0x3e, 0xc4, 0xbd, 0xfd, 0xd3,
	0x0f, 0x26, 0xb3, 0xf1, 0x83, 0x3d, 0xf0, 0x8d, 0xa2, 0x40, 0x17, 0x4d, 0x9a, 0xa2, 0x99, 0x59,
	0x34, 0xd6, 0xf8, 0xcb, 0x37, 0x2c, 0xd4, 0x58, 0xd8, 0x60, 0xa1, 0x89, 0x85, 0x35, 0x56, 0xda,
	0x07, 0x6d, 0xb0, 0xf0, 0x70, 0x3a, 0x0a, 0x40, 0xd1, 0x98, 0x2a, 0xfa, 0x69, 0x0f, 0x0d, 0x05,
	0x21, 0xda, 0x41, 0x48, 0x23, 0x21, 0xc4, 0xb4, 0x10, 0x52, 0x6b, 0xe6, 0xdd, 0xcd, 0x8f, 0xb6,
	0xfb, 0x15, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x78, 0xc5, 0x03, 0x5b, 0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/assetmantle.modules.identities.transactions.deputize.Service/Handle", in, out, opts...)
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
		FullMethod: "/assetmantle.modules.identities.transactions.deputize.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assetmantle.modules.identities.transactions.deputize.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "identities/transactions/deputize/service.proto",
}