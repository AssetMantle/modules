// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/identities/internal/queries/identity/service.proto

package identity

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
	proto.RegisterFile("x/identities/internal/queries/identity/service.proto", fileDescriptor_dd6933fe3afc0ee1)
}

var fileDescriptor_dd6933fe3afc0ee1 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xb1, 0x4a, 0xc3, 0x40,
	0x1c, 0xc6, 0x9b, 0x88, 0x15, 0x82, 0x53, 0xc6, 0x22, 0x19, 0x14, 0x27, 0xe1, 0x8e, 0x5a, 0xa7,
	0x13, 0x85, 0x16, 0x44, 0x3b, 0x04, 0x5a, 0xdd, 0x24, 0x20, 0xd7, 0xe6, 0x4f, 0x3d, 0x48, 0xee,
	0xda, 0xdc, 0xa5, 0x98, 0x55, 0x7c, 0x00, 0xc1, 0x37, 0x70, 0x74, 0xf6, 0x01, 0x1c, 0xc5, 0xa9,
	0xe0, 0xe2, 0x28, 0xa9, 0x93, 0x4f, 0x21, 0xed, 0x5d, 0x9b, 0xba, 0xb5, 0x5d, 0xbf, 0xe3, 0xf7,
	0xf1, 0xfb, 0xfe, 0x9c, 0x73, 0x74, 0x87, 0x59, 0x08, 0x5c, 0x31, 0xc5, 0x40, 0x62, 0xc6, 0x15,
	0x24, 0x9c, 0x46, 0x78, 0x90, 0x42, 0x32, 0x0d, 0xf4, 0x5b, 0x86, 0x25, 0x24, 0x43, 0xd6, 0x05,
	0xd4, 0x4f, 0x84, 0x12, 0x2e, 0xa6, 0x52, 0x82, 0x8a, 0x29, 0x57, 0x11, 0xa0, 0x58, 0x84, 0x69,
	0x04, 0x12, 0x15, 0x3d, 0xc8, 0xe0, 0xb3, 0x28, 0xab, 0xec, 0xf4, 0x84, 0xe8, 0x45, 0x80, 0x69,
	0x9f, 0x61, 0xca, 0xb9, 0x50, 0x54, 0x31, 0xc1, 0xa5, 0xae, 0xab, 0x90, 0x25, 0x25, 0x26, 0x41,
	0x76, 0x93, 0xc0, 0x20, 0x05, 0xa9, 0xd6, 0x61, 0x2f, 0x41, 0xf6, 0x05, 0x97, 0x66, 0xc6, 0xe1,
	0x9b, 0xe5, 0x6c, 0x5d, 0xe9, 0x61, 0xee, 0xab, 0xe5, 0x94, 0x2f, 0x28, 0x0f, 0x23, 0x70, 0x4f,
	0xd0, 0x8a, 0xf3, 0x50, 0x5b, 0x97, 0x4f, 0xbd, 0x2a, 0xa7, 0xeb, 0xe2, 0xda, 0x6d, 0xf7, 0xe0,
	0xfe, 0xf3, 0xe7, 0xc9, 0xde, 0x77, 0xf7, 0xb0, 0xae, 0x58, 0xdc, 0x37, 0xac, 0x76, 0x40, 0xd1,
	0xea, 0x7c, 0x56, 0xe3, 0x61, 0xe3, 0x3d, 0xf7, 0xac, 0x51, 0xee, 0x59, 0xdf, 0xb9, 0x67, 0x3d,
	0x8e, 0xbd, 0xd2, 0x68, 0xec, 0x95, 0xbe, 0xc6, 0x5e, 0xc9, 0xa9, 0x75, 0x45, 0xbc, 0xaa, 0x4a,
	0x63, 0xdb, 0xdc, 0xa3, 0x35, 0x39, 0x50, 0xcb, 0xba, 0x3e, 0xeb, 0x31, 0x75, 0x9b, 0x76, 0x50,
	0x57, 0xc4, 0xb8, 0x3e, 0xe9, 0xf2, 0xb5, 0x93, 0xe9, 0xc2, 0xcb, 0x5d, 0xff, 0xd9, 0xde, 0xac,
	0xfb, 0xcd, 0x76, 0xf3, 0xc5, 0xd6, 0x25, 0x46, 0xc8, 0x37, 0x42, 0xcd, 0x42, 0xa8, 0x6d, 0x84,
	0x4c, 0x94, 0x7d, 0xfc, 0x23, 0x02, 0x43, 0x04, 0x05, 0x11, 0x18, 0x62, 0x16, 0x65, 0xb9, 0x7d,
	0xbc, 0x22, 0x11, 0x9c, 0xb7, 0x1a, 0x3e, 0x28, 0x1a, 0x52, 0x45, 0x7f, 0xed, 0xda, 0x02, 0x4d,
	0x88, 0xc1, 0x09, 0x29, 0x78, 0x42, 0x4c, 0xc1, 0x3c, 0xcc, 0x3a, 0xe5, 0xe9, 0x87, 0xaa, 0xfd,
	0x05, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x07, 0x2d, 0xc2, 0x4f, 0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/assetmantle.modules.identities.queries.identity.Service/Handle", in, out, opts...)
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
		FullMethod: "/assetmantle.modules.identities.queries.identity.Service/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Handle(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assetmantle.modules.identities.queries.identity.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Service_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "x/identities/internal/queries/identity/service.proto",
}
