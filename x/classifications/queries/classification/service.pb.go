// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: AssetMantle/modules/x/classifications/queries/classification/service.proto

package classification

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
	proto.RegisterFile("AssetMantle/modules/x/classifications/queries/classification/service.proto", fileDescriptor_e5438836dfa39dff)
}

var fileDescriptor_e5438836dfa39dff = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xb1, 0x4a, 0xc3, 0x40,
	0x18, 0xc7, 0x9b, 0x88, 0x1d, 0x82, 0x53, 0xc6, 0x22, 0x19, 0x7c, 0x80, 0x3b, 0xab, 0xe0, 0x70,
	0x38, 0xd8, 0x76, 0x68, 0xa9, 0x04, 0x12, 0x5d, 0x82, 0x04, 0xe4, 0x9a, 0x7e, 0xd6, 0x83, 0x24,
	0xd7, 0xe6, 0x2e, 0xa5, 0xae, 0x3e, 0x81, 0xe0, 0x1b, 0x38, 0xfa, 0x24, 0xea, 0x54, 0x70, 0xd1,
	0x4d, 0x12, 0x27, 0x9f, 0x42, 0xda, 0xdc, 0xd0, 0xde, 0x9a, 0xae, 0xff, 0x2f, 0xff, 0x1f, 0xbf,
	0x7c, 0xf7, 0x59, 0xc3, 0x8e, 0x10, 0x20, 0x5d, 0x9a, 0xca, 0x18, 0x70, 0xc2, 0xc7, 0x79, 0x0c,
	0x02, 0x2f, 0x70, 0x14, 0x53, 0x21, 0xd8, 0x1d, 0x8b, 0xa8, 0x64, 0x3c, 0x15, 0x78, 0x96, 0x43,
	0xc6, 0x40, 0x68, 0x39, 0x16, 0x90, 0xcd, 0x59, 0x04, 0x68, 0x9a, 0x71, 0xc9, 0xed, 0xf3, 0x0d,
	0x16, 0x52, 0x2c, 0xb4, 0x40, 0x1a, 0x0b, 0x29, 0x96, 0x96, 0xb7, 0x0e, 0x27, 0x9c, 0x4f, 0x62,
	0xc0, 0x74, 0xca, 0x30, 0x4d, 0x53, 0x2e, 0xd5, 0xe7, 0x6b, 0x76, 0xcb, 0xab, 0xe5, 0xb9, 0x8a,
	0x1f, 0x6e, 0x33, 0x98, 0xe5, 0x20, 0xa4, 0x22, 0xfa, 0x3b, 0x21, 0x8a, 0x29, 0x4f, 0x85, 0x5a,
	0xc0, 0x49, 0x69, 0x58, 0xfb, 0xfe, 0x6a, 0x60, 0x7f, 0x1b, 0x56, 0x73, 0x40, 0xd3, 0x71, 0x0c,
	0xf6, 0x10, 0xd5, 0x59, 0x0b, 0x5a, 0xf3, 0xae, 0x2a, 0xf3, 0xd6, 0xe5, 0x4e, 0x58, 0x95, 0xf3,
	0xd1, 0xd9, 0xe3, 0xe7, 0xef, 0xb3, 0x79, 0x6c, 0x23, 0x9c, 0x54, 0x4b, 0xd0, 0x7f, 0x7d, 0xde,
	0x1e, 0x81, 0xa4, 0x6d, 0x2d, 0xef, 0xbe, 0xef, 0xbd, 0x15, 0x8e, 0xb1, 0x2c, 0x1c, 0xe3, 0xa7,
	0x70, 0x8c, 0xa7, 0xd2, 0x69, 0x2c, 0x4b, 0xa7, 0xf1, 0x55, 0x3a, 0x0d, 0xeb, 0x22, 0xe2, 0x49,
	0x2d, 0xc5, 0xee, 0xc1, 0x75, 0x75, 0x52, 0xde, 0x6a, 0xa1, 0x9e, 0x71, 0xd3, 0x9f, 0x30, 0x79,
	0x9f, 0x8f, 0x50, 0xc4, 0x13, 0x5c, 0xe7, 0xc1, 0x5e, 0xcc, 0x66, 0xc7, 0x0d, 0x7a, 0x7e, 0xef,
	0xd5, 0xdc, 0x3a, 0x53, 0x57, 0x09, 0x06, 0xa8, 0xa7, 0x09, 0xfa, 0x4a, 0x70, 0x3b, 0xff, 0xd8,
	0xaa, 0x87, 0xaa, 0x1e, 0x06, 0xa1, 0x56, 0x0f, 0x55, 0x5d, 0xcb, 0x0b, 0x73, 0x50, 0xa7, 0x1e,
	0xf6, 0xbd, 0xae, 0x0b, 0x92, 0x8e, 0xa9, 0xa4, 0x7f, 0x66, 0x67, 0x03, 0x45, 0x88, 0x62, 0x11,
	0x12, 0x10, 0xa2, 0xd1, 0x08, 0x51, 0x38, 0x7d, 0x32, 0x6a, 0xae, 0x0f, 0xf7, 0xf4, 0x3f, 0x00,
	0x00, 0xff, 0xff, 0xbb, 0x34, 0xaa, 0x80, 0x07, 0x04, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/AssetMantle.modules.x.classifications.queries.classification.Query/Handle", in, out, opts...)
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
		FullMethod: "/AssetMantle.modules.x.classifications.queries.classification.Query/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Handle(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AssetMantle.modules.x.classifications.queries.classification.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Query_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "AssetMantle/modules/x/classifications/queries/classification/service.proto",
}
