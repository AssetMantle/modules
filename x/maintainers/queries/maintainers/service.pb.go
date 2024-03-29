// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: maintainers/queries/maintainers/service.proto

package maintainers

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
	proto.RegisterFile("maintainers/queries/maintainers/service.proto", fileDescriptor_6e9809c6d4ffccad)
}

var fileDescriptor_6e9809c6d4ffccad = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x31, 0x4b, 0xf3, 0x40,
	0x18, 0xc7, 0x9b, 0xbc, 0xb4, 0x43, 0x78, 0xa7, 0x8c, 0x45, 0x32, 0xb8, 0x88, 0x83, 0x77, 0xd4,
	0x08, 0x85, 0x2c, 0x9a, 0x2c, 0xba, 0x1c, 0xb4, 0xba, 0x49, 0x40, 0xae, 0xed, 0x43, 0x3d, 0x48,
	0xee, 0xda, 0xdc, 0xa5, 0xe8, 0xea, 0x27, 0x10, 0xfa, 0x0d, 0x1c, 0xfd, 0x14, 0x8e, 0xd2, 0xa9,
	0xe0, 0xe2, 0x28, 0xa9, 0x93, 0x9f, 0x42, 0xd2, 0xdc, 0x70, 0x71, 0xa9, 0xc4, 0xf5, 0x9f, 0xe7,
	0xff, 0xe3, 0x77, 0xcf, 0x13, 0xe7, 0x28, 0xa5, 0x8c, 0x2b, 0xca, 0x38, 0x64, 0x12, 0xcf, 0x73,
	0xc8, 0x18, 0x48, 0x6c, 0x66, 0x12, 0xb2, 0x05, 0x1b, 0x03, 0x9a, 0x65, 0x42, 0x09, 0xd7, 0xa7,
	0x52, 0x82, 0x4a, 0x29, 0x57, 0x09, 0xa0, 0x54, 0x4c, 0xf2, 0x04, 0x24, 0x32, 0xc6, 0x91, 0x46,
	0x98, 0x59, 0x77, 0x6f, 0x2a, 0xc4, 0x34, 0x01, 0x4c, 0x67, 0x0c, 0x53, 0xce, 0x85, 0xa2, 0x8a,
	0x09, 0x2e, 0x2b, 0x64, 0xd7, 0xdf, 0x65, 0x50, 0x66, 0xf7, 0x37, 0x19, 0xcc, 0x73, 0x90, 0x4a,
	0x97, 0x4e, 0x7e, 0x5b, 0x92, 0x33, 0xc1, 0xa5, 0xb6, 0x3f, 0x5e, 0x59, 0x4e, 0x7b, 0x58, 0x7e,
	0x70, 0x5f, 0x2c, 0xa7, 0x73, 0x41, 0xf9, 0x24, 0x01, 0x37, 0x44, 0x0d, 0xde, 0x84, 0xb6, 0x98,
	0xcb, 0xca, 0xa9, 0x1b, 0xfd, 0x05, 0x51, 0x19, 0xee, 0xe3, 0x87, 0xb7, 0xcf, 0xa5, 0x7d, 0xe8,
	0x1e, 0xe0, 0x0a, 0x53, 0x7b, 0xcd, 0xa2, 0x37, 0x02, 0x45, 0x7b, 0x66, 0x16, 0x2d, 0xff, 0xbd,
	0x16, 0x9e, 0xb5, 0x2e, 0x3c, 0xeb, 0xa3, 0xf0, 0xac, 0xc7, 0x8d, 0xd7, 0x5a, 0x6f, 0xbc, 0xd6,
	0xfb, 0xc6, 0x6b, 0x39, 0xfd, 0xb1, 0x48, 0x9b, 0x28, 0x45, 0xff, 0xaf, 0xaa, 0x6b, 0x0f, 0xca,
	0x75, 0x0d, 0xac, 0xeb, 0xb3, 0x29, 0x53, 0xb7, 0xf9, 0x08, 0x8d, 0x45, 0x8a, 0xc3, 0x92, 0x47,
	0xb4, 0x5b, 0xc5, 0xc3, 0x77, 0x78, 0xc7, 0x1d, 0x9e, 0xec, 0x76, 0x48, 0xc8, 0x90, 0x3c, 0xdb,
	0x7e, 0x68, 0xf8, 0x10, 0xed, 0x43, 0x7e, 0xac, 0x83, 0xd5, 0xb3, 0x55, 0xad, 0x15, 0xeb, 0x56,
	0x6c, 0x4c, 0xc4, 0xba, 0x65, 0x66, 0x85, 0x7d, 0xda, 0xa0, 0x15, 0x9f, 0x0f, 0x22, 0x02, 0x8a,
	0x4e, 0xa8, 0xa2, 0x5f, 0x76, 0xdf, 0x20, 0x04, 0x81, 0x46, 0x04, 0x81, 0x31, 0x1f, 0x04, 0x1a,
	0x52, 0x4b, 0x47, 0x9d, 0xed, 0x9f, 0xe6, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x04, 0x18, 0xdf,
	0x7a, 0x58, 0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/assetmantle.modules.maintainers.queries.maintainers.Query/Handle", in, out, opts...)
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
		FullMethod: "/assetmantle.modules.maintainers.queries.maintainers.Query/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Handle(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assetmantle.modules.maintainers.queries.maintainers.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Query_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "maintainers/queries/maintainers/service.proto",
}
