// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/assets/internal/transactions/burn/service.proto

package burn

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
	proto.RegisterFile("x/assets/internal/transactions/burn/service.proto", fileDescriptor_5612498c9ebe80bc)
}

var fileDescriptor_5612498c9ebe80bc = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xb1, 0x4a, 0xf4, 0x40,
	0x14, 0x85, 0x37, 0x81, 0x7f, 0x7f, 0x08, 0x56, 0x2b, 0x36, 0x51, 0x52, 0x6c, 0xbd, 0xcc, 0xac,
	0x8a, 0x20, 0x53, 0x08, 0x49, 0xa3, 0x4d, 0x60, 0xd1, 0xad, 0x24, 0x20, 0x77, 0xb3, 0x43, 0x0c,
	0x24, 0x33, 0xcb, 0xcc, 0x44, 0xac, 0x7d, 0x02, 0xc1, 0x07, 0x10, 0x2c, 0xf5, 0x45, 0xc4, 0x6a,
	0xc1, 0xc6, 0x52, 0xb2, 0x56, 0x96, 0x3e, 0x81, 0x64, 0x67, 0xc4, 0xb1, 0x4b, 0xda, 0x3b, 0xe7,
	0x3b, 0xf7, 0x9c, 0xcb, 0x78, 0xbb, 0xd7, 0x18, 0xa4, 0xa4, 0x4a, 0xe2, 0x9c, 0x29, 0x2a, 0x18,
	0x14, 0x58, 0x09, 0x60, 0x12, 0x52, 0x95, 0x73, 0x26, 0xf1, 0xac, 0x12, 0x0c, 0x4b, 0x2a, 0xae,
	0xf2, 0x94, 0xa2, 0x85, 0xe0, 0x8a, 0x0f, 0x46, 0x6b, 0xa0, 0x04, 0xa6, 0x0a, 0x8a, 0x4a, 0x3e,
	0xaf, 0x0a, 0x2a, 0x91, 0x36, 0x41, 0x36, 0x8b, 0x1a, 0xd6, 0xdf, 0xc9, 0x38, 0xcf, 0x0a, 0x8a,
	0x61, 0x91, 0x63, 0x60, 0x8c, 0x2b, 0xd0, 0x8f, 0x6b, 0x2f, 0xbf, 0xd5, 0xfa, 0x92, 0x4a, 0x09,
	0x99, 0x59, 0xef, 0x1f, 0xb5, 0x41, 0xac, 0xc9, 0x85, 0xa0, 0x72, 0xc1, 0x99, 0x34, 0xfc, 0xde,
	0x93, 0xe3, 0xfd, 0x3f, 0xd3, 0x85, 0x06, 0xf7, 0x8e, 0xd7, 0x3f, 0x01, 0x36, 0x2f, 0xe8, 0xe0,
	0x00, 0x75, 0xa9, 0x85, 0x62, 0x9d, 0xc9, 0x0f, 0xbb, 0x61, 0xd3, 0xdf, 0xc9, 0xa9, 0x89, 0x35,
	0xdc, 0xbe, 0x79, 0xfd, 0xb8, 0x73, 0xb7, 0x86, 0x9b, 0x58, 0xbb, 0xfc, 0xd4, 0x6b, 0x80, 0xe8,
	0xcb, 0x7d, 0xae, 0x03, 0x67, 0x59, 0x07, 0xce, 0x7b, 0x1d, 0x38, 0xb7, 0xab, 0xa0, 0xb7, 0x5c,
	0x05, 0xbd, 0xb7, 0x55, 0xd0, 0xf3, 0xc6, 0x29, 0x2f, 0x3b, 0x6d, 0x8f, 0x36, 0x4c, 0xef, 0x49,
	0x73, 0x88, 0x89, 0x73, 0x1e, 0x65, 0xb9, 0xba, 0xac, 0x66, 0x28, 0xe5, 0x25, 0x0e, 0x1b, 0x28,
	0xd6, 0x01, 0x8c, 0x11, 0x6e, 0x71, 0xe9, 0x07, 0xf7, 0x5f, 0x18, 0x87, 0xd3, 0xe8, 0xd1, 0x1d,
	0x85, 0x56, 0x94, 0xd8, 0x44, 0x09, 0x75, 0x94, 0xa9, 0x1d, 0x25, 0xaa, 0x04, 0x7b, 0xf9, 0x23,
	0x4f, 0x8c, 0x3c, 0xd1, 0xf2, 0xc4, 0x96, 0x27, 0x8d, 0xbc, 0x76, 0x0f, 0xbb, 0xc8, 0x93, 0xe3,
	0x49, 0x14, 0x53, 0x05, 0x73, 0x50, 0xf0, 0xe9, 0x8e, 0x2d, 0x94, 0x10, 0xc3, 0x12, 0xa2, 0x61,
	0x42, 0x6c, 0x9a, 0x90, 0x06, 0x9f, 0xf5, 0xd7, 0x3f, 0x65, 0xff, 0x3b, 0x00, 0x00, 0xff, 0xff,
	0xc3, 0xf0, 0xad, 0x3e, 0x1d, 0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/assetmantle.modules.assets.transactions.burn.Service/Handle", in, out, opts...)
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
		FullMethod: "/assetmantle.modules.assets.transactions.burn.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assetmantle.modules.assets.transactions.burn.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "x/assets/internal/transactions/burn/service.proto",
}
