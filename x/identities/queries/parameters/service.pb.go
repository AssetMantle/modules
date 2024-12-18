// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: AssetMantle/modules/x/identities/queries/parameters/service.proto

package parameters

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
	proto.RegisterFile("AssetMantle/modules/x/identities/queries/parameters/service.proto", fileDescriptor_899e37c996c98b07)
}

var fileDescriptor_899e37c996c98b07 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0x93, 0x20, 0x3a, 0x44, 0x4c, 0x1d, 0x3b, 0x78, 0xe0, 0x00, 0xb6, 0x44, 0x07, 0x90,
	0x07, 0x50, 0xb2, 0xb4, 0x1d, 0x22, 0xa5, 0xb0, 0x44, 0xc8, 0x12, 0x72, 0x9b, 0x27, 0x88, 0xd4,
	0xc4, 0xad, 0xed, 0xa0, 0x72, 0x0b, 0x10, 0x23, 0x1b, 0x23, 0x27, 0x41, 0x4c, 0x1d, 0x19, 0x51,
	0xb2, 0x71, 0x0a, 0x94, 0xc6, 0xc2, 0xed, 0x9a, 0x4e, 0x91, 0x62, 0x7d, 0xdf, 0xfb, 0x7f, 0xfb,
	0xf9, 0x41, 0xa0, 0x14, 0xe8, 0x88, 0x17, 0x7a, 0x01, 0x24, 0x17, 0x69, 0xb9, 0x00, 0x45, 0xd6,
	0x24, 0x4b, 0xa1, 0xd0, 0x99, 0xce, 0x40, 0x91, 0x55, 0x09, 0xb2, 0xf9, 0x2e, 0xb9, 0xe4, 0x39,
	0x68, 0x90, 0x8a, 0x28, 0x90, 0x8f, 0xd9, 0x1c, 0xf0, 0x52, 0x0a, 0x2d, 0xfa, 0xc3, 0x1d, 0x05,
	0x36, 0x0a, 0xbc, 0xc6, 0x56, 0x81, 0x8d, 0x02, 0x5b, 0xc5, 0x60, 0xd4, 0x65, 0x6e, 0xf3, 0xeb,
	0xe9, 0x4e, 0xc2, 0xaa, 0x04, 0xa5, 0xdb, 0xe9, 0x83, 0xf1, 0x21, 0x22, 0xb5, 0x14, 0x85, 0x32,
	0x3d, 0xce, 0xde, 0x5c, 0xff, 0x78, 0xda, 0x1c, 0xf4, 0x5f, 0x5c, 0xbf, 0x37, 0xe6, 0x45, 0xba,
	0x80, 0x7e, 0x80, 0x3b, 0xb4, 0xc3, 0x5b, 0xcd, 0x75, 0x9b, 0x73, 0x10, 0x1e, 0xa2, 0x68, 0x13,
	0x9e, 0x3a, 0xe1, 0xeb, 0xd1, 0x67, 0x85, 0xdc, 0x4d, 0x85, 0xdc, 0x9f, 0x0a, 0xb9, 0xcf, 0x35,
	0x72, 0x36, 0x35, 0x72, 0xbe, 0x6b, 0xe4, 0xf8, 0xe7, 0x73, 0x91, 0x77, 0x99, 0x11, 0x9e, 0xdc,
	0xb4, 0x0f, 0x19, 0x37, 0xfd, 0x63, 0xf7, 0xf6, 0xf2, 0x3e, 0xd3, 0x0f, 0xe5, 0x0c, 0xcf, 0x45,
	0x4e, 0x3a, 0x5c, 0xeb, 0xbb, 0xd7, 0x0b, 0xa2, 0x64, 0x32, 0x8d, 0x3f, 0xbc, 0xbd, 0x9d, 0x88,
	0x4c, 0x9c, 0x04, 0x4f, 0x6c, 0x9c, 0xa9, 0x89, 0x13, 0xff, 0xe3, 0x5f, 0x7b, 0x14, 0x33, 0x14,
	0x4b, 0x98, 0xa5, 0x98, 0xa1, 0x98, 0xa5, 0x2a, 0xef, 0xaa, 0x03, 0xc5, 0x46, 0x71, 0x18, 0x81,
	0xe6, 0x29, 0xd7, 0xfc, 0xd7, 0xbb, 0xd8, 0x31, 0x50, 0x6a, 0x14, 0x94, 0x26, 0x94, 0x5a, 0x09,
	0xa5, 0xc6, 0x42, 0xa9, 0xd5, 0xcc, 0x7a, 0xdb, 0xd5, 0x19, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff,
	0xab, 0x23, 0x88, 0xac, 0x47, 0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/AssetMantle.modules.x.identities.queries.parameters.Query/Handle", in, out, opts...)
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
		FullMethod: "/AssetMantle.modules.x.identities.queries.parameters.Query/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Handle(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var Query_serviceDesc = _Query_serviceDesc
var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AssetMantle.modules.x.identities.queries.parameters.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Query_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "AssetMantle/modules/x/identities/queries/parameters/service.proto",
}
