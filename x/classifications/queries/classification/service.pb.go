// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: classifications/queries/classification/service.proto

package classification

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
	proto.RegisterFile("classifications/queries/classification/service.proto", fileDescriptor_59f874cbf61dd532)
}

var fileDescriptor_59f874cbf61dd532 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xb1, 0x4a, 0xfb, 0x40,
	0x1c, 0xc7, 0x9b, 0xfc, 0xf9, 0x57, 0x08, 0x4e, 0x19, 0x8b, 0x64, 0xf0, 0x01, 0xee, 0xac, 0x8a,
	0xc3, 0x09, 0x85, 0xb6, 0x60, 0xeb, 0x10, 0x68, 0x75, 0x93, 0x80, 0x5c, 0xd3, 0x9f, 0xf5, 0x20,
	0xb9, 0x6b, 0x73, 0x97, 0xa2, 0xab, 0x4f, 0x20, 0xf8, 0x06, 0x8e, 0xbe, 0x83, 0xbb, 0x38, 0x15,
	0x44, 0x70, 0xd4, 0xd4, 0xc9, 0xa7, 0x90, 0x34, 0x37, 0x34, 0x37, 0xd5, 0xea, 0xfa, 0x4d, 0xbe,
	0x9f, 0x7c, 0xbe, 0x97, 0x73, 0xf6, 0xc3, 0x88, 0x4a, 0xc9, 0x2e, 0x58, 0x48, 0x15, 0x13, 0x5c,
	0xe2, 0x49, 0x0a, 0x09, 0x03, 0x89, 0xcb, 0x39, 0x96, 0x90, 0x4c, 0x59, 0x08, 0x68, 0x9c, 0x08,
	0x25, 0x5c, 0x42, 0xa5, 0x04, 0x15, 0x53, 0xae, 0x22, 0x40, 0xb1, 0x18, 0xa6, 0x11, 0x48, 0x64,
	0x90, 0x90, 0x26, 0x19, 0x79, 0x6d, 0x6b, 0x24, 0xc4, 0x28, 0x02, 0x4c, 0xc7, 0x0c, 0x53, 0xce,
	0x85, 0xd2, 0xaf, 0x2f, 0xc8, 0x35, 0xb2, 0xa2, 0x4f, 0x1e, 0x5f, 0x9f, 0x27, 0x30, 0x49, 0x41,
	0x2a, 0xdd, 0x3d, 0xfc, 0x61, 0x57, 0x8e, 0x05, 0x97, 0x7a, 0xd2, 0xee, 0x87, 0xe5, 0x6c, 0x9c,
	0x16, 0x23, 0xdd, 0x57, 0xcb, 0xa9, 0x76, 0x29, 0x1f, 0x46, 0xe0, 0x76, 0xd1, 0xfa, 0x53, 0x51,
	0x3f, 0xff, 0xd0, 0x49, 0xe1, 0x58, 0x3b, 0xfe, 0x03, 0x52, 0x61, 0xbc, 0x7d, 0x70, 0xf3, 0xf2,
	0x79, 0x67, 0xef, 0xb8, 0x08, 0x17, 0x34, 0x6c, 0x0e, 0x9f, 0xd6, 0x07, 0xa0, 0x68, 0xdd, 0xc8,
	0x5b, 0x8f, 0xff, 0x9e, 0x32, 0xcf, 0x9a, 0x65, 0x9e, 0xf5, 0x9e, 0x79, 0xd6, 0xed, 0xdc, 0xab,
	0xcc, 0xe6, 0x5e, 0xe5, 0x6d, 0xee, 0x55, 0x9c, 0x46, 0x28, 0xe2, 0x5f, 0x08, 0xb6, 0x36, 0xf5,
	0xd9, 0xf5, 0xf2, 0xc3, 0xec, 0x59, 0x67, 0x9d, 0x11, 0x53, 0x97, 0xe9, 0x00, 0x85, 0x22, 0xc6,
	0xcd, 0x1c, 0xeb, 0x17, 0xa6, 0x1a, 0x8b, 0xaf, 0xf0, 0x6a, 0x3f, 0xeb, 0xde, 0xfe, 0xdf, 0xf4,
	0xdb, 0xfd, 0xf6, 0x83, 0x4d, 0x9a, 0x4b, 0x76, 0xbe, 0xb6, 0x6b, 0x1b, 0x76, 0x7d, 0x6d, 0x57,
	0xce, 0x9f, 0x4b, 0xe5, 0x40, 0x97, 0x03, 0xa3, 0x1c, 0xe8, 0xb2, 0x91, 0x67, 0xf6, 0xd1, 0xfa,
	0xe5, 0xa0, 0xd3, 0x6b, 0xf9, 0xa0, 0xe8, 0x90, 0x2a, 0xfa, 0x65, 0x37, 0x96, 0x40, 0x84, 0x68,
	0x12, 0x21, 0x06, 0x8a, 0x10, 0xcd, 0x32, 0x9f, 0x0c, 0xaa, 0x8b, 0xab, 0xba, 0xf7, 0x1d, 0x00,
	0x00, 0xff, 0xff, 0x3d, 0x3a, 0x55, 0x4f, 0xb5, 0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/assetmantle.modules.classifications.queries.classification.Service/Handle", in, out, opts...)
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
		FullMethod: "/assetmantle.modules.classifications.queries.classification.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assetmantle.modules.classifications.queries.classification.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "classifications/queries/classification/service.proto",
}