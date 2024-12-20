// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: AssetMantle/modules/x/splits/queries/splits/service.proto

package splits

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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
	proto.RegisterFile("AssetMantle/modules/x/splits/queries/splits/service.proto", fileDescriptor_99bae0ab2f4f64e9)
}

var fileDescriptor_99bae0ab2f4f64e9 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x74, 0x2c, 0x2e, 0x4e,
	0x2d, 0xf1, 0x4d, 0xcc, 0x2b, 0xc9, 0x49, 0xd5, 0xcf, 0xcd, 0x4f, 0x29, 0xcd, 0x49, 0x2d, 0xd6,
	0xaf, 0xd0, 0x2f, 0x2e, 0xc8, 0xc9, 0x2c, 0x29, 0xd6, 0x2f, 0x2c, 0x4d, 0x2d, 0xca, 0x4c, 0x2d,
	0x86, 0x71, 0x8b, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85,
	0xb4, 0x91, 0xb4, 0xea, 0x41, 0xb5, 0xea, 0x55, 0xe8, 0x41, 0xd4, 0xea, 0x41, 0xb5, 0x42, 0xb9,
	0x52, 0xf6, 0xa4, 0xd8, 0x03, 0xe2, 0x56, 0xc6, 0x17, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x40,
	0x6c, 0x93, 0x72, 0x20, 0xc7, 0x80, 0xe2, 0x82, 0xfc, 0xbc, 0x62, 0xa8, 0x7b, 0x8d, 0xba, 0x19,
	0xb9, 0x58, 0x03, 0x41, 0x12, 0x42, 0x8d, 0x8c, 0x5c, 0x6c, 0x1e, 0x89, 0x79, 0x29, 0x39, 0xa9,
	0x42, 0x96, 0x7a, 0x24, 0xf8, 0x42, 0x0f, 0xac, 0x3d, 0x08, 0xe2, 0x2e, 0x29, 0x2b, 0x72, 0xb4,
	0x42, 0x5c, 0xa4, 0xc4, 0xe0, 0xf4, 0x98, 0xe9, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18,
	0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5,
	0x18, 0xb8, 0xf4, 0x93, 0xf3, 0x73, 0x49, 0x31, 0xdb, 0x89, 0x27, 0x18, 0x12, 0x31, 0x01, 0x20,
	0x7f, 0x06, 0x30, 0x46, 0x99, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea,
	0x93, 0x10, 0x6c, 0x8b, 0x98, 0xd8, 0x1c, 0x7d, 0x23, 0x82, 0x03, 0x83, 0x57, 0x31, 0xa1, 0xc4,
	0xad, 0x2f, 0xd4, 0xfa, 0x08, 0xbd, 0x60, 0x84, 0x5f, 0x40, 0xd6, 0x43, 0xb8, 0xa7, 0x50, 0x54,
	0xc7, 0x40, 0x55, 0xc7, 0x44, 0xc4, 0x40, 0xa4, 0x63, 0xa0, 0xaa, 0xa1, 0xdc, 0x47, 0x4c, 0xe6,
	0x24, 0xa8, 0x8e, 0x71, 0x0f, 0x70, 0xf2, 0x4d, 0x2d, 0x49, 0x4c, 0x49, 0x2c, 0x49, 0x7c, 0xc5,
	0x64, 0x80, 0xa4, 0xd3, 0xca, 0x0a, 0xaa, 0xd5, 0xca, 0x2a, 0xc2, 0xca, 0x0a, 0xa2, 0xda, 0xca,
	0x0a, 0xaa, 0x1b, 0x26, 0x90, 0xc4, 0x06, 0x8e, 0x7a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x95, 0xb4, 0xd6, 0xbc, 0xe7, 0x02, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/AssetMantle.modules.x.splits.queries.splits.Query/Handle", in, out, opts...)
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
		FullMethod: "/AssetMantle.modules.x.splits.queries.splits.Query/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Handle(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var Query_serviceDesc = _Query_serviceDesc
var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AssetMantle.modules.x.splits.queries.splits.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Query_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "AssetMantle/modules/x/splits/queries/splits/service.proto",
}
