// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: orders/queries/order/service.proto

package order

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
	proto.RegisterFile("orders/queries/order/service.proto", fileDescriptor_d39626a640b3db69)
}

var fileDescriptor_d39626a640b3db69 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0x9b, 0x40, 0xbb, 0x08, 0xdf, 0x2a, 0xcb, 0x52, 0x66, 0x51, 0xbe, 0x45, 0x85, 0x32,
	0x43, 0xfd, 0x0b, 0xd9, 0xb5, 0x1b, 0xdd, 0x84, 0xb6, 0xba, 0x93, 0x80, 0x4c, 0x9b, 0x4b, 0x0d,
	0x24, 0x99, 0x76, 0x66, 0x52, 0x74, 0xeb, 0x13, 0x08, 0x3e, 0x80, 0xe0, 0x4e, 0x9f, 0x44, 0x5c,
	0x55, 0xdc, 0xb8, 0x94, 0xd4, 0x95, 0x4f, 0x21, 0xc9, 0x9d, 0x45, 0x85, 0x22, 0x71, 0x79, 0x92,
	0xf3, 0xbb, 0xe7, 0xcc, 0x9d, 0x71, 0xda, 0x42, 0x86, 0x20, 0x15, 0x5b, 0x64, 0x20, 0x23, 0x50,
	0xac, 0x94, 0x4c, 0x81, 0x5c, 0x46, 0x53, 0xa0, 0x73, 0x29, 0xb4, 0x70, 0x3b, 0x5c, 0x29, 0xd0,
	0x09, 0x4f, 0x75, 0x0c, 0x34, 0x11, 0x61, 0x16, 0x83, 0xa2, 0xc8, 0x51, 0xc3, 0xa1, 0x6c, 0xb6,
	0x66, 0x42, 0xcc, 0x62, 0x60, 0x7c, 0x1e, 0x31, 0x9e, 0xa6, 0x42, 0x73, 0x1d, 0x89, 0x54, 0xe1,
	0x9c, 0x66, 0x67, 0x6b, 0x56, 0xa1, 0xae, 0x2f, 0x24, 0x2c, 0x32, 0x50, 0xda, 0x38, 0x77, 0x7e,
	0x75, 0xaa, 0xb9, 0x48, 0x95, 0x29, 0xb7, 0xfb, 0x68, 0x39, 0xf5, 0x71, 0xf1, 0xc3, 0xbd, 0xb7,
	0x9c, 0xc6, 0x09, 0x4f, 0xc3, 0x18, 0xdc, 0x43, 0x5a, 0xb5, 0x32, 0x2d, 0xd9, 0x53, 0x4c, 0x6f,
	0x1e, 0xfd, 0x99, 0xc3, 0x2e, 0xed, 0xff, 0x37, 0x6f, 0x9f, 0x77, 0x36, 0x71, 0x5b, 0x0c, 0x59,
	0x66, 0x0e, 0xb0, 0xec, 0x4d, 0x40, 0xf3, 0x1e, 0xca, 0xc1, 0xab, 0xfd, 0x9c, 0x13, 0x6b, 0x95,
	0x13, 0xeb, 0x23, 0x27, 0xd6, 0xed, 0x9a, 0xd4, 0x56, 0x6b, 0x52, 0x7b, 0x5f, 0x93, 0x9a, 0xd3,
	0x9d, 0x8a, 0xa4, 0x72, 0xf8, 0xe0, 0xdf, 0x19, 0x5e, 0xd0, 0xa8, 0x58, 0xc1, 0xc8, 0x3a, 0x3f,
	0x98, 0x45, 0xfa, 0x32, 0x9b, 0xd0, 0xa9, 0x48, 0x58, 0xbf, 0x18, 0xe2, 0x63, 0x0b, 0x33, 0x84,
	0x5d, 0xb1, 0x6d, 0x0b, 0x7d, 0xb0, 0xeb, 0x7d, 0x7f, 0x38, 0x1e, 0x3e, 0xd9, 0x9d, 0xfe, 0x46,
	0xb2, 0x6f, 0x92, 0x87, 0x98, 0x3c, 0x36, 0xc9, 0xa5, 0x7c, 0xf9, 0x61, 0x0d, 0x8c, 0x35, 0x40,
	0x6b, 0x60, 0xac, 0x28, 0x73, 0x7b, 0xbf, 0xaa, 0x35, 0x38, 0x1e, 0x0d, 0x7c, 0xd0, 0x3c, 0xe4,
	0x9a, 0x7f, 0xd9, 0xdd, 0x0d, 0xcc, 0xf3, 0x0c, 0xe7, 0x79, 0x08, 0x7a, 0x9e, 0x21, 0xcd, 0x87,
	0x49, 0xa3, 0x7c, 0x06, 0x7b, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x73, 0x0b, 0x99, 0x85, 0xc9,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	Handle(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Handle(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/assetmantle.modules.orders.queries.order.Query/Handle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	Handle(context.Context, *QueryRequest) (*QueryResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Handle(ctx context.Context, req *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handle not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Handle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/assetmantle.modules.orders.queries.order.Query/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Handle(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assetmantle.modules.orders.queries.order.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Query_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orders/queries/order/service.proto",
}
