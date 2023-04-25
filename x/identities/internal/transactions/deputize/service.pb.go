// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/identities/internal/transactions/deputize/service.proto

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
	proto.RegisterFile("x/identities/internal/transactions/deputize/service.proto", fileDescriptor_779ae471960f97da)
}

var fileDescriptor_779ae471960f97da = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0x4a, 0xe3, 0x40,
	0x18, 0x80, 0x9b, 0x2c, 0xdb, 0x85, 0xb0, 0xa7, 0x1e, 0xb3, 0xbb, 0x39, 0x74, 0xef, 0x33, 0xa0,
	0x3d, 0xd4, 0x01, 0xc1, 0x96, 0xa2, 0x56, 0x08, 0x14, 0xed, 0x49, 0x02, 0x32, 0x4d, 0x7e, 0xe2,
	0x40, 0x32, 0x13, 0x32, 0x53, 0x11, 0x8f, 0x3e, 0x81, 0xe0, 0x1b, 0xf4, 0xe8, 0x0b, 0xf8, 0x00,
	0x5e, 0xc4, 0x53, 0xc1, 0x8b, 0x47, 0x49, 0x3d, 0xf9, 0x14, 0xd2, 0x66, 0x42, 0xc6, 0x63, 0x73,
	0xfd, 0xc9, 0xf7, 0xf1, 0xfd, 0x7f, 0xc6, 0xd9, 0xbb, 0xc6, 0x2c, 0x02, 0xae, 0x98, 0x62, 0x20,
	0x31, 0xe3, 0x0a, 0x72, 0x4e, 0x13, 0xac, 0x72, 0xca, 0x25, 0x0d, 0x15, 0x13, 0x5c, 0xe2, 0x08,
	0xb2, 0xb9, 0x62, 0x37, 0x80, 0x25, 0xe4, 0x57, 0x2c, 0x04, 0x94, 0xe5, 0x42, 0x89, 0x4e, 0x8f,
	0x4a, 0x09, 0x2a, 0xa5, 0x5c, 0x25, 0x80, 0x52, 0x11, 0xcd, 0x13, 0x90, 0xa8, 0x96, 0x21, 0xd3,
	0x81, 0x2a, 0x87, 0xfb, 0x37, 0x16, 0x22, 0x4e, 0x00, 0xd3, 0x8c, 0x61, 0xca, 0xb9, 0x50, 0xb4,
	0xfc, 0x60, 0xe3, 0x74, 0xb7, 0xca, 0x49, 0x41, 0x4a, 0x1a, 0xeb, 0x1c, 0xf7, 0x70, 0x1b, 0xd4,
	0x98, 0x5e, 0xe4, 0x20, 0x33, 0xc1, 0xa5, 0xf6, 0xec, 0x3c, 0x59, 0xce, 0xaf, 0xb3, 0x72, 0xd1,
	0xce, 0xa3, 0xe5, 0xb4, 0x8f, 0x29, 0x8f, 0x12, 0xe8, 0xec, 0xa3, 0x26, 0xeb, 0x22, 0xbf, 0x6c,
	0x74, 0xc7, 0xcd, 0xf0, 0x69, 0x3d, 0x3d, 0xd5, 0x99, 0xdd, 0xff, 0xb7, 0xaf, 0x1f, 0xf7, 0xf6,
	0xbf, 0xee, 0x1f, 0x5c, 0xda, 0xcc, 0xb5, 0x2b, 0x70, 0xb8, 0xf8, 0xf1, 0x5c, 0x78, 0xd6, 0xb2,
	0xf0, 0xac, 0xf7, 0xc2, 0xb3, 0xee, 0x56, 0x5e, 0x6b, 0xb9, 0xf2, 0x5a, 0x6f, 0x2b, 0xaf, 0xe5,
	0xf4, 0x43, 0x91, 0x36, 0xaa, 0x19, 0xfe, 0xd6, 0x77, 0x99, 0xac, 0x0f, 0x35, 0xb1, 0xce, 0x4f,
	0x62, 0xa6, 0x2e, 0xe7, 0x33, 0x14, 0x8a, 0x14, 0x0f, 0xd6, 0x42, 0xbf, 0x0c, 0xd2, 0x42, 0xbc,
	0xc5, 0x1f, 0x59, 0xd8, 0x3f, 0x07, 0xfe, 0x78, 0x3a, 0x7a, 0xb0, 0x7b, 0x03, 0x23, 0xcd, 0xd7,
	0x69, 0xe3, 0x3a, 0x6d, 0x6a, 0xa6, 0x8d, 0x34, 0xfe, 0xf2, 0x0d, 0x0b, 0x34, 0x16, 0xd4, 0x58,
	0x60, 0x62, 0x41, 0x85, 0x15, 0xf6, 0x41, 0x13, 0x2c, 0x38, 0x9a, 0x0c, 0x7d, 0x50, 0x34, 0xa2,
	0x8a, 0x7e, 0xda, 0x7d, 0x43, 0x41, 0x88, 0x76, 0x10, 0x52, 0x4b, 0x08, 0x31, 0x2d, 0x84, 0x54,
	0x9a, 0x59, 0x7b, 0xf3, 0xe2, 0x76, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x52, 0x40, 0x17, 0x42,
	0x85, 0x03, 0x00, 0x00,
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
	Metadata: "x/identities/internal/transactions/deputize/service.proto",
}
