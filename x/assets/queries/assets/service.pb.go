// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: assets/queries/assets/service.proto

package assets

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
	proto.RegisterFile("assets/queries/assets/service.proto", fileDescriptor_2342b589a413aef8)
}

var fileDescriptor_2342b589a413aef8 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0x2c, 0x2e, 0x4e,
	0x2d, 0x29, 0xd6, 0x2f, 0x2c, 0x4d, 0x2d, 0xca, 0x4c, 0x2d, 0xd6, 0x87, 0x72, 0x8b, 0x53, 0x8b,
	0xca, 0x32, 0x93, 0x53, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x34, 0xc1, 0xa2, 0xb9, 0x89,
	0x79, 0x25, 0x39, 0xa9, 0x7a, 0xb9, 0xf9, 0x29, 0xa5, 0x39, 0xa9, 0xc5, 0x7a, 0x10, 0x95, 0x7a,
	0x50, 0x8d, 0x50, 0xae, 0x94, 0x26, 0x76, 0xf3, 0x40, 0xdc, 0xca, 0xf8, 0xa2, 0xd4, 0xc2, 0xd2,
	0xd4, 0xe2, 0x12, 0x88, 0xa9, 0x52, 0x5a, 0xf8, 0x95, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0x43, 0x5d,
	0x20, 0x25, 0x93, 0x9e, 0x9f, 0x9f, 0x9e, 0x93, 0xaa, 0x9f, 0x58, 0x90, 0xa9, 0x9f, 0x98, 0x97,
	0x97, 0x5f, 0x92, 0x58, 0x92, 0x99, 0x9f, 0x57, 0x0c, 0x91, 0x35, 0x5a, 0xc3, 0xc8, 0xc5, 0x1a,
	0x08, 0xd2, 0x26, 0xb4, 0x88, 0x91, 0x8b, 0xcd, 0x23, 0x31, 0x2f, 0x25, 0x27, 0x55, 0xc8, 0x5c,
	0x8f, 0x68, 0x57, 0xeb, 0x81, 0x35, 0x07, 0x41, 0x5c, 0x27, 0x65, 0x41, 0xba, 0x46, 0x88, 0x5b,
	0x95, 0x54, 0x9b, 0x2e, 0x3f, 0x99, 0xcc, 0x24, 0x2f, 0x24, 0xab, 0x0f, 0xd1, 0x0c, 0xf3, 0x52,
	0x99, 0x61, 0x52, 0x6a, 0x49, 0xa2, 0x21, 0x94, 0xeb, 0x74, 0x9d, 0xe9, 0xc4, 0x23, 0x39, 0xc6,
	0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39,
	0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xb8, 0x74, 0x93, 0xf3, 0x73, 0x89, 0xb7, 0xde, 0x89, 0x27, 0x18,
	0x12, 0x4f, 0x01, 0xa0, 0x60, 0x08, 0x60, 0x8c, 0x32, 0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2,
	0x4b, 0xce, 0xcf, 0xd5, 0x77, 0x04, 0x29, 0xf1, 0x85, 0xb8, 0x03, 0x6a, 0x8a, 0x7e, 0x85, 0x3e,
	0xd6, 0x30, 0x5f, 0xc4, 0xc4, 0xea, 0xe8, 0xeb, 0x18, 0xe8, 0xb8, 0x8a, 0x49, 0xd3, 0x11, 0xc9,
	0x6e, 0x5f, 0xa8, 0xdd, 0x8e, 0x08, 0xbf, 0x66, 0xc2, 0xb9, 0xa7, 0x50, 0xd4, 0xc6, 0x40, 0xd5,
	0xc6, 0x40, 0x24, 0x63, 0xa0, 0x6a, 0xa1, 0xdc, 0x47, 0x4c, 0xa6, 0x44, 0xab, 0x8d, 0x71, 0x0f,
	0x70, 0xf2, 0x4d, 0x2d, 0x49, 0x4c, 0x49, 0x2c, 0x49, 0x7c, 0xc5, 0xa4, 0x8b, 0xa4, 0xcf, 0xca,
	0x0a, 0xaa, 0xd1, 0xca, 0x0a, 0xa2, 0xd4, 0xca, 0x0a, 0xaa, 0x15, 0x26, 0x90, 0xc4, 0x06, 0x4e,
	0x0f, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x39, 0x0d, 0xb4, 0xbd, 0xd6, 0x02, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/assetmantle.modules.assets.queries.assets.Query/Handle", in, out, opts...)
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
		FullMethod: "/assetmantle.modules.assets.queries.assets.Query/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Handle(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assetmantle.modules.assets.queries.assets.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Query_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "assets/queries/assets/service.proto",
}
