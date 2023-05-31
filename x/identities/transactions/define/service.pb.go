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
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x31, 0x4b, 0xeb, 0x40,
	0x1c, 0xc0, 0x9b, 0x3c, 0x5e, 0x1f, 0x84, 0x37, 0x75, 0x7b, 0x79, 0x12, 0xb0, 0xb3, 0xdc, 0x41,
	0x85, 0x82, 0x27, 0x28, 0x2d, 0x85, 0xda, 0x21, 0x50, 0xb4, 0x93, 0x04, 0xe4, 0x9a, 0xfc, 0x8d,
	0x07, 0xc9, 0x5d, 0xc9, 0x5d, 0xc5, 0xd9, 0x5d, 0x10, 0xfc, 0x06, 0x8e, 0x6e, 0x7e, 0x00, 0x77,
	0x71, 0x2a, 0xb8, 0x38, 0x4a, 0xea, 0xe4, 0xa7, 0x90, 0xf6, 0x0e, 0x72, 0x2e, 0x96, 0x76, 0xbd,
	0xe4, 0xf7, 0xbb, 0xdf, 0xff, 0xcf, 0x79, 0x3b, 0x2c, 0x01, 0xae, 0x98, 0x62, 0x20, 0xb1, 0x2a,
	0x28, 0x97, 0x34, 0x56, 0x4c, 0x70, 0x89, 0x13, 0x38, 0x67, 0x1c, 0xb0, 0x84, 0xe2, 0x92, 0xc5,
	0x80, 0x26, 0x85, 0x50, 0xa2, 0xd1, 0xa2, 0x52, 0x82, 0xca, 0x29, 0x57, 0x19, 0xa0, 0x5c, 0x24,
	0xd3, 0x0c, 0x24, 0xaa, 0x0c, 0xc8, 0x36, 0x20, 0x6d, 0xf0, 0xb7, 0x52, 0x21, 0xd2, 0x0c, 0x30,
	0x9d, 0x30, 0x4c, 0x39, 0x17, 0x8a, 0xea, 0xcf, 0x4b, 0xa3, 0xbf, 0xea, 0xfe, 0x1c, 0xa4, 0xa4,
	0xa9, 0xb9, 0xdf, 0xdf, 0x5b, 0xf1, 0xb7, 0x75, 0x76, 0x56, 0x80, 0x9c, 0x08, 0x2e, 0x0d, 0xda,
	0x7a, 0x72, 0xbc, 0x3f, 0x27, 0x7a, 0x98, 0xc6, 0xa3, 0xe3, 0xd5, 0x8f, 0x28, 0x4f, 0x32, 0x68,
	0xec, 0xa3, 0xf5, 0x47, 0x42, 0xa1, 0x8e, 0xf2, 0xfb, 0x9b, 0xc0, 0xa3, 0xea, 0xec, 0xd8, 0x24,
	0x36, 0xb7, 0xaf, 0x5f, 0x3f, 0xee, 0xdc, 0xff, 0xcd, 0x7f, 0x58, 0xbb, 0xb0, 0x35, 0xa7, 0xc6,
	0xba, 0x37, 0xbf, 0x9e, 0xcb, 0xc0, 0x99, 0x95, 0x81, 0xf3, 0x5e, 0x06, 0xce, 0xed, 0x3c, 0xa8,
	0xcd, 0xe6, 0x41, 0xed, 0x6d, 0x1e, 0xd4, 0xbc, 0x76, 0x2c, 0xf2, 0x0d, 0x4a, 0xba, 0x7f, 0xcd,
	0x3e, 0x86, 0x8b, 0x05, 0x0d, 0x9d, 0xd3, 0xc3, 0x94, 0xa9, 0x8b, 0xe9, 0x18, 0xc5, 0x22, 0xc7,
	0x9d, 0x85, 0x2e, 0xd4, 0x31, 0x46, 0x87, 0xaf, 0xf0, 0xcf, 0xeb, 0xbf, 0x77, 0x7f, 0x77, 0xc2,
	0xc1, 0xa8, 0xf7, 0xe0, 0xb6, 0x3a, 0x56, 0x4d, 0x68, 0x6a, 0x06, 0x55, 0xcd, 0xc8, 0xae, 0xe9,
	0x2d, 0xe1, 0x97, 0x6f, 0x50, 0x64, 0xa0, 0xa8, 0x82, 0x22, 0x1b, 0x8a, 0x34, 0x54, 0xba, 0x07,
	0xeb, 0x43, 0x51, 0x7f, 0xd8, 0x0d, 0x41, 0xd1, 0x84, 0x2a, 0xfa, 0xe9, 0xb6, 0x2d, 0x01, 0x21,
	0xc6, 0x40, 0x48, 0xa5, 0x20, 0xc4, 0x76, 0x10, 0xa2, 0x25, 0xe3, 0xfa, 0xf2, 0x59, 0xed, 0x7e,
	0x05, 0x00, 0x00, 0xff, 0xff, 0x80, 0xa4, 0x4a, 0x4e, 0x41, 0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/assetmantle.modules.identities.transactions.define.Service/Handle", in, out, opts...)
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
		FullMethod: "/assetmantle.modules.identities.transactions.define.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assetmantle.modules.identities.transactions.define.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "identities/transactions/define/service.proto",
}
