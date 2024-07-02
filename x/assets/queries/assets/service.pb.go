// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: AssetMantle/modules/x/assets/queries/assets/service.proto

package assets

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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
	proto.RegisterFile("AssetMantle/modules/x/assets/queries/assets/service.proto", fileDescriptor_c8b1b0efe74062b4)
}

var fileDescriptor_c8b1b0efe74062b4 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x31, 0x4b, 0xc3, 0x40,
	0x18, 0x86, 0x9b, 0x80, 0x1d, 0x82, 0x53, 0xc6, 0xa0, 0x11, 0x04, 0x27, 0xe1, 0xce, 0x2a, 0x28,
	0xbd, 0x45, 0xd3, 0x45, 0x97, 0x40, 0xab, 0x4b, 0x90, 0x80, 0x5c, 0xdb, 0x8f, 0x1a, 0x48, 0x72,
	0x6d, 0xee, 0x52, 0xea, 0xea, 0x2f, 0x10, 0xfc, 0x07, 0x0e, 0x0e, 0x4e, 0xfe, 0x0c, 0x71, 0x2a,
	0xb8, 0x38, 0x4a, 0xe2, 0xe4, 0xaf, 0x90, 0xf4, 0xbe, 0x42, 0x3b, 0xa6, 0xe3, 0x7b, 0xf7, 0x3e,
	0xf9, 0xde, 0xf7, 0xf2, 0x59, 0x6d, 0x4f, 0x4a, 0x50, 0x3e, 0x4f, 0x55, 0x0c, 0x34, 0x11, 0xc3,
	0x3c, 0x06, 0x49, 0x67, 0x94, 0x57, 0xa7, 0x92, 0x4e, 0x72, 0xc8, 0x22, 0x90, 0x4b, 0x29, 0x21,
	0x9b, 0x46, 0x03, 0x20, 0xe3, 0x4c, 0x28, 0x61, 0x1f, 0xae, 0xa0, 0x04, 0x51, 0x32, 0x23, 0xda,
	0x4b, 0x10, 0x45, 0xe9, 0x9c, 0xd7, 0x99, 0x53, 0xc9, 0x87, 0xbb, 0x0c, 0x26, 0x39, 0x48, 0xa5,
	0xa7, 0x39, 0x17, 0x9b, 0x7c, 0x40, 0x8e, 0x45, 0x2a, 0x31, 0xaf, 0xb3, 0x33, 0x12, 0x62, 0x14,
	0x03, 0xe5, 0xe3, 0x88, 0xf2, 0x34, 0x15, 0x8a, 0xab, 0x48, 0xa4, 0x52, 0xdf, 0x1e, 0xbf, 0x1b,
	0xd6, 0x56, 0xaf, 0xc2, 0xec, 0x57, 0xc3, 0x6a, 0x5e, 0xf1, 0x74, 0x18, 0x83, 0xdd, 0x26, 0x35,
	0x3a, 0x92, 0x05, 0x7e, 0xad, 0x53, 0x3b, 0x6c, 0x13, 0x54, 0xe7, 0xdd, 0x3f, 0x78, 0xfc, 0xfa,
	0x7d, 0x36, 0xf7, 0xec, 0x5d, 0x9a, 0xe8, 0xbe, 0x58, 0x6b, 0xda, 0xea, 0x83, 0xe2, 0x2d, 0x94,
	0x9d, 0xd2, 0xfc, 0x28, 0x5c, 0x63, 0x5e, 0xb8, 0xc6, 0x4f, 0xe1, 0x1a, 0x4f, 0xa5, 0xdb, 0x98,
	0x97, 0x6e, 0xe3, 0xbb, 0x74, 0x1b, 0x16, 0x1d, 0x88, 0xa4, 0x4e, 0x80, 0xce, 0xf6, 0x8d, 0xfe,
	0xb7, 0xdd, 0xea, 0x31, 0xba, 0xc6, 0xed, 0xe9, 0x28, 0x52, 0xf7, 0x79, 0x9f, 0x0c, 0x44, 0x42,
	0x6b, 0xbc, 0xfc, 0x8b, 0xd9, 0xf4, 0xfc, 0xc0, 0xeb, 0x79, 0x6f, 0xe6, 0xda, 0x7a, 0xf8, 0x38,
	0x3e, 0xd0, 0xa1, 0x74, 0xe1, 0x6a, 0xbc, 0x96, 0x9f, 0x6b, 0xee, 0x10, 0xdd, 0x61, 0x10, 0xea,
	0xeb, 0x10, 0xdd, 0x28, 0x0b, 0xf3, 0xac, 0x86, 0x3b, 0xbc, 0xec, 0x76, 0x7c, 0x50, 0x7c, 0xc8,
	0x15, 0xff, 0x33, 0x8f, 0x56, 0x48, 0xc6, 0x10, 0x65, 0x2c, 0x60, 0x4c, 0xbb, 0x19, 0x43, 0x7a,
	0x79, 0xd0, 0x6f, 0x2e, 0xf6, 0xe3, 0xe4, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x56, 0x5a, 0xf8, 0x16,
	0x2a, 0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/AssetMantle.modules.x.assets.queries.assets.Query/Handle", in, out, opts...)
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
		FullMethod: "/AssetMantle.modules.x.assets.queries.assets.Query/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Handle(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AssetMantle.modules.x.assets.queries.assets.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Query_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "AssetMantle/modules/x/assets/queries/assets/service.proto",
}
