// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: AssetMantle/modules/x/metas/queries/meta/service.proto

package meta

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
	proto.RegisterFile("AssetMantle/modules/x/metas/queries/meta/service.proto", fileDescriptor_aef7a6cab4cafd46)
}

var fileDescriptor_aef7a6cab4cafd46 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x73, 0x2c, 0x2e, 0x4e,
	0x2d, 0xf1, 0x4d, 0xcc, 0x2b, 0xc9, 0x49, 0xd5, 0xcf, 0xcd, 0x4f, 0x29, 0xcd, 0x49, 0x2d, 0xd6,
	0xaf, 0xd0, 0xcf, 0x4d, 0x2d, 0x49, 0x2c, 0xd6, 0x2f, 0x2c, 0x4d, 0x2d, 0xca, 0x4c, 0x2d, 0x06,
	0xf3, 0xf4, 0x8b, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85,
	0x34, 0x90, 0xf4, 0xe9, 0x41, 0xf5, 0xe9, 0x55, 0xe8, 0x81, 0xf5, 0xe9, 0x41, 0xf5, 0x81, 0x79,
	0x52, 0x36, 0x44, 0xdb, 0x00, 0xe2, 0x54, 0xc6, 0x17, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x40,
	0xec, 0x91, 0xb2, 0x25, 0x59, 0x77, 0x71, 0x41, 0x7e, 0x5e, 0x31, 0xd4, 0x99, 0x46, 0x2d, 0x8c,
	0x5c, 0xac, 0x81, 0x20, 0x09, 0xa1, 0x6a, 0x2e, 0x36, 0x8f, 0xc4, 0xbc, 0x94, 0x9c, 0x54, 0x21,
	0x33, 0x3d, 0x62, 0xdd, 0xae, 0x07, 0xd6, 0x1a, 0x04, 0x71, 0x90, 0x94, 0x39, 0xc9, 0xfa, 0x20,
	0x4e, 0x51, 0x62, 0x70, 0xba, 0xc8, 0x74, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f,
	0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c,
	0x5c, 0x3a, 0xc9, 0xf9, 0xb9, 0x44, 0x1b, 0xec, 0xc4, 0x13, 0x0c, 0x89, 0x85, 0x00, 0x90, 0xef,
	0x02, 0x18, 0xa3, 0x8c, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x89,
	0x0d, 0xa9, 0x45, 0x4c, 0x6c, 0x8e, 0xbe, 0x11, 0xbe, 0x81, 0xbe, 0xab, 0x98, 0x50, 0x62, 0xd1,
	0x17, 0x6a, 0x71, 0x84, 0x9e, 0x2f, 0xd8, 0xe2, 0x40, 0xa8, 0xc5, 0x20, 0xde, 0x29, 0x14, 0xa5,
	0x31, 0x50, 0xa5, 0x31, 0x11, 0x31, 0x60, 0xa5, 0x31, 0x50, 0xa5, 0x60, 0xde, 0x23, 0x26, 0x13,
	0x62, 0x95, 0xc6, 0xb8, 0x07, 0x38, 0x81, 0xe8, 0x94, 0xc4, 0x92, 0xc4, 0x57, 0x4c, 0xba, 0x48,
	0xda, 0xac, 0xac, 0xa0, 0xfa, 0xac, 0xac, 0x22, 0xac, 0xac, 0xc0, 0x3a, 0xad, 0xac, 0xa0, 0x5a,
	0x21, 0xfc, 0x24, 0x36, 0x70, 0x0c, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc7, 0xd1, 0x6c,
	0xa3, 0xc2, 0x02, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/AssetMantle.modules.x.metas.queries.meta.Query/Handle", in, out, opts...)
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
		FullMethod: "/AssetMantle.modules.x.metas.queries.meta.Query/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Handle(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var Query_serviceDesc = _Query_serviceDesc
var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AssetMantle.modules.x.metas.queries.meta.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Query_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "AssetMantle/modules/x/metas/queries/meta/service.proto",
}
