// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/orders/internal/queries/orders/service.proto

package orders

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
	proto.RegisterFile("x/orders/internal/queries/orders/service.proto", fileDescriptor_268887c36b1e4b8a)
}

var fileDescriptor_268887c36b1e4b8a = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x31, 0x4b, 0xfb, 0x40,
	0x18, 0xc6, 0x9b, 0xc0, 0xbf, 0x7f, 0x08, 0x4e, 0x19, 0x8b, 0x46, 0x10, 0x1c, 0x3a, 0xf4, 0x8e,
	0x5a, 0x45, 0xc9, 0x22, 0xed, 0xa2, 0x4b, 0x68, 0xab, 0x9b, 0x64, 0xb9, 0x36, 0x2f, 0xf5, 0x20,
	0xb9, 0x6b, 0x73, 0x97, 0xa2, 0xab, 0x9f, 0x40, 0xf0, 0x1b, 0x74, 0x74, 0xf0, 0x73, 0x88, 0x53,
	0xc1, 0xc5, 0xb1, 0xa4, 0x4e, 0x7e, 0x0a, 0xa9, 0xf7, 0x56, 0xe3, 0xd4, 0x38, 0x3e, 0x97, 0xe7,
	0xf7, 0xbe, 0xcf, 0xf3, 0x12, 0x87, 0xdc, 0x50, 0x99, 0x46, 0x90, 0x2a, 0xca, 0x85, 0x86, 0x54,
	0xb0, 0x98, 0x4e, 0x32, 0x48, 0x39, 0xa8, 0xf5, 0xbb, 0x82, 0x74, 0xca, 0x87, 0x40, 0xc6, 0xa9,
	0xd4, 0xd2, 0xad, 0x33, 0xa5, 0x40, 0x27, 0x4c, 0xe8, 0x18, 0x48, 0x22, 0xa3, 0x2c, 0x06, 0x45,
	0x8c, 0x93, 0x20, 0x88, 0xb2, 0xb6, 0x3d, 0x92, 0x72, 0x14, 0x03, 0x65, 0x63, 0x4e, 0x99, 0x10,
	0x52, 0x33, 0xcd, 0xa5, 0x50, 0x66, 0x50, 0xad, 0xb5, 0x71, 0xf1, 0x4a, 0xde, 0x5e, 0xc0, 0x24,
	0x03, 0xa5, 0x11, 0x3a, 0x2c, 0x0b, 0xa9, 0xb1, 0x14, 0x0a, 0x33, 0x1f, 0x3c, 0x59, 0xce, 0xff,
	0x4b, 0xd3, 0xc2, 0x9d, 0x59, 0x4e, 0xf5, 0x9c, 0x89, 0x28, 0x06, 0xf7, 0x98, 0x94, 0xee, 0x42,
	0xfa, 0x85, 0x2c, 0xb5, 0x93, 0xbf, 0x83, 0x26, 0xcf, 0xde, 0xfe, 0xdd, 0xeb, 0xfb, 0x83, 0xbd,
	0xeb, 0xee, 0x50, 0x03, 0xaf, 0x53, 0x4f, 0x9b, 0x03, 0xd0, 0xac, 0x89, 0xb2, 0xb3, 0xb0, 0x9f,
	0x73, 0xcf, 0x9a, 0xe7, 0x9e, 0xb5, 0xc8, 0x3d, 0xeb, 0x7e, 0xe9, 0x55, 0xe6, 0x4b, 0xaf, 0xf2,
	0xb6, 0xf4, 0x2a, 0x4e, 0x63, 0x28, 0x93, 0xf2, 0xeb, 0x3b, 0x5b, 0xd8, 0xbb, 0xb7, 0x3a, 0x44,
	0xcf, 0xba, 0x3a, 0x1d, 0x71, 0x7d, 0x9d, 0x0d, 0xc8, 0x50, 0x26, 0xb4, 0xbd, 0x9a, 0x12, 0x98,
	0x1c, 0x38, 0x85, 0x6e, 0xba, 0xef, 0xcc, 0xfe, 0xd7, 0x0e, 0xba, 0xfd, 0xee, 0xa3, 0x5d, 0x6f,
	0x17, 0x42, 0x04, 0x18, 0xa2, 0xfb, 0x53, 0x9a, 0x7f, 0xcb, 0x97, 0x5f, 0xde, 0x10, 0xbd, 0xa1,
	0xf9, 0x18, 0xa2, 0x17, 0x65, 0x6e, 0x1f, 0x95, 0xf6, 0x86, 0x67, 0xbd, 0x4e, 0x00, 0x9a, 0x45,
	0x4c, 0xb3, 0x0f, 0xbb, 0x51, 0xe0, 0x7c, 0x1f, 0x41, 0xdf, 0x37, 0x56, 0xdf, 0x47, 0x74, 0xfd,
	0x30, 0xa8, 0x7e, 0xfd, 0x1a, 0xad, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x50, 0x38, 0x42,
	0x00, 0x03, 0x00, 0x00,
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
	Handle(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
}

type serviceClient struct {
	cc grpc1.ClientConn
}

func NewServiceClient(cc grpc1.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Handle(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/assetmantle.modules.orders.queries.orders.Service/Handle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	Handle(context.Context, *QueryRequest) (*QueryResponse, error)
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) Handle(ctx context.Context, req *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handle not implemented")
}

func RegisterServiceServer(s grpc1.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_Handle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assetmantle.modules.orders.queries.orders.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assetmantle.modules.orders.queries.orders.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "x/orders/internal/queries/orders/service.proto",
}
