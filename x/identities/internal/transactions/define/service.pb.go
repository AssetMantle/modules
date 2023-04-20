// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/identities/internal/transactions/define/service.proto

package define

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
	proto.RegisterFile("x/identities/internal/transactions/define/service.proto", fileDescriptor_6483863e7b166365)
}

var fileDescriptor_6483863e7b166365 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x31, 0x4b, 0xfb, 0x40,
	0x18, 0xc6, 0x9b, 0xfc, 0xa1, 0x7f, 0x08, 0x4e, 0xdd, 0x8c, 0x72, 0x68, 0x17, 0xc1, 0xe1, 0x0e,
	0xec, 0x20, 0x64, 0xb3, 0x16, 0x6c, 0x87, 0x42, 0xd1, 0x4e, 0xd2, 0xe5, 0x6d, 0xf2, 0x1a, 0x0f,
	0x92, 0xbb, 0x92, 0x7b, 0x2b, 0xce, 0x6e, 0x6e, 0x82, 0xf8, 0x05, 0x1c, 0x1d, 0xfc, 0x1c, 0xe2,
	0x54, 0x70, 0x71, 0x94, 0xd4, 0xc9, 0x4f, 0x21, 0xf6, 0x0a, 0x3d, 0x10, 0xda, 0xba, 0x26, 0xbf,
	0xe7, 0x9e, 0xe7, 0x77, 0x17, 0x1c, 0x5e, 0x0b, 0x99, 0xa0, 0x22, 0x49, 0x12, 0x8d, 0x90, 0x8a,
	0xb0, 0x50, 0x90, 0x09, 0x2a, 0x40, 0x19, 0x88, 0x49, 0x6a, 0x65, 0x44, 0x82, 0x17, 0x52, 0xa1,
	0x30, 0x58, 0x5c, 0xc9, 0x18, 0xf9, 0xa8, 0xd0, 0xa4, 0x6b, 0x6c, 0x11, 0xe3, 0x2e, 0xcd, 0x2d,
	0x1d, 0x6e, 0xa7, 0x5a, 0xa7, 0x19, 0x0a, 0x18, 0x49, 0x01, 0x4a, 0x69, 0x02, 0xfb, 0x7b, 0x96,
	0x0e, 0xff, 0x50, 0x9b, 0xa3, 0x31, 0x90, 0xce, 0x6b, 0xc3, 0xe3, 0xf5, 0x83, 0xce, 0xb7, 0x53,
	0x34, 0x23, 0xad, 0xcc, 0xfc, 0x90, 0x83, 0x07, 0x2f, 0xf8, 0x7f, 0x66, 0x6d, 0x6a, 0xb7, 0x5e,
	0x50, 0x6d, 0x83, 0x4a, 0x32, 0xac, 0xed, 0xf1, 0xe5, 0x4e, 0xbc, 0x6b, 0xa7, 0x84, 0x8d, 0x55,
	0x60, 0xff, 0x77, 0x75, 0x7d, 0xf7, 0xe6, 0xed, 0xf3, 0xde, 0xdf, 0xaa, 0x6f, 0x8a, 0x1c, 0x14,
	0x65, 0xe8, 0x7a, 0xd8, 0x58, 0xf3, 0xd9, 0x7f, 0x29, 0x99, 0x37, 0x29, 0x99, 0xf7, 0x51, 0x32,
	0xef, 0x6e, 0xca, 0x2a, 0x93, 0x29, 0xab, 0xbc, 0x4f, 0x59, 0x25, 0xa8, 0xc7, 0x3a, 0x5f, 0xd1,
	0xda, 0xdc, 0x98, 0x3b, 0xf5, 0x7e, 0x24, 0x7b, 0xde, 0x79, 0x3b, 0x95, 0x74, 0x39, 0x1e, 0xf2,
	0x58, 0xe7, 0xe2, 0xc8, 0x18, 0xa4, 0xae, 0x2d, 0xce, 0x75, 0x32, 0xce, 0xd0, 0x88, 0xb5, 0xaf,
	0xf2, 0xd1, 0xff, 0xd7, 0xe9, 0xb7, 0x9e, 0x7c, 0xd6, 0x59, 0x0c, 0xe8, 0xbb, 0x03, 0x5a, 0x33,
	0xec, 0xd5, 0x05, 0x06, 0x2e, 0x30, 0xb0, 0x40, 0xe9, 0xef, 0x2f, 0x07, 0x06, 0x27, 0xbd, 0x66,
	0x17, 0x09, 0x12, 0x20, 0xf8, 0xf2, 0x77, 0x16, 0x70, 0x14, 0xb9, 0x74, 0x14, 0x59, 0x7c, 0x58,
	0x9d, 0xbd, 0x67, 0xe3, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xad, 0x33, 0x8f, 0xc6, 0x02, 0x00,
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
	err := c.cc.Invoke(ctx, "/identities.transactions.define.Service/Handle", in, out, opts...)
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
		FullMethod: "/identities.transactions.define.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "identities.transactions.define.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "x/identities/internal/transactions/define/service.proto",
}