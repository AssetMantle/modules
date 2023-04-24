// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/assets/internal/queries/asset/queryRequest.proto

package asset

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	base "github.com/AssetMantle/schema/go/ids/base"
	proto "github.com/gogo/protobuf/proto"
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

type QueryRequest struct {
	AssetID *base.AssetID `protobuf:"bytes,1,opt,name=asset_i_d,json=assetID,proto3" json:"asset_i_d,omitempty"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f47128a2354eb24, []int{0}
}
func (m *QueryRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(m, src)
}
func (m *QueryRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

func (m *QueryRequest) GetAssetID() *base.AssetID {
	if m != nil {
		return m.AssetID
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryRequest)(nil), "assetmantle.modules.assets.queries.asset.QueryRequest")
}

func init() {
	proto.RegisterFile("x/assets/internal/queries/asset/queryRequest.proto", fileDescriptor_4f47128a2354eb24)
}

var fileDescriptor_4f47128a2354eb24 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0xaa, 0xd0, 0x4f, 0x2c,
	0x2e, 0x4e, 0x2d, 0x29, 0xd6, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x2f, 0x2c,
	0x4d, 0x2d, 0xca, 0x4c, 0x2d, 0x86, 0x88, 0x83, 0x79, 0x95, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5,
	0x25, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x1a, 0x60, 0x99, 0xdc, 0xc4, 0xbc, 0x92, 0x9c,
	0x54, 0xbd, 0xdc, 0xfc, 0x94, 0xd2, 0x9c, 0xd4, 0x62, 0x3d, 0x88, 0x29, 0x7a, 0x50, 0xcd, 0x10,
	0xae, 0x94, 0x78, 0x66, 0x4a, 0xb1, 0x7e, 0x52, 0x62, 0x71, 0x2a, 0xc4, 0xb0, 0xf8, 0xcc, 0x14,
	0x88, 0x11, 0x4a, 0x01, 0x5c, 0x3c, 0x81, 0x48, 0x06, 0x0b, 0x39, 0x70, 0x71, 0x42, 0x55, 0xc4,
	0xa7, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0xa9, 0xe8, 0x21, 0x5b, 0x53, 0x9c, 0x9c, 0x91,
	0x9a, 0x9b, 0xa8, 0x97, 0x99, 0x52, 0xac, 0x07, 0x32, 0x4f, 0xcf, 0x11, 0x24, 0xe7, 0xe9, 0x12,
	0xc4, 0x9e, 0x08, 0x61, 0x38, 0xdd, 0x67, 0x3a, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6,
	0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39,
	0x06, 0x2e, 0x9d, 0xe4, 0xfc, 0x5c, 0x3d, 0x62, 0xdd, 0xec, 0x24, 0x88, 0xec, 0xb0, 0x00, 0x90,
	0x6b, 0x03, 0x18, 0xa3, 0xec, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5,
	0xc1, 0x56, 0xfb, 0x82, 0x4d, 0xd2, 0x87, 0x9a, 0xa4, 0x4f, 0x20, 0x14, 0x17, 0x31, 0xb1, 0x3a,
	0xfa, 0x3a, 0x06, 0x3a, 0xae, 0x62, 0xd2, 0x70, 0x44, 0x72, 0x87, 0x2f, 0xd4, 0x1d, 0x8e, 0x10,
	0x77, 0x04, 0x42, 0xdd, 0x01, 0xe6, 0x9e, 0x42, 0x51, 0x1a, 0x03, 0x55, 0x1a, 0x03, 0x51, 0x1a,
	0x03, 0x55, 0x0a, 0xe1, 0x3e, 0x62, 0x32, 0x21, 0x56, 0x69, 0x8c, 0x7b, 0x80, 0x93, 0x6f, 0x6a,
	0x49, 0x62, 0x4a, 0x62, 0x49, 0xe2, 0x2b, 0x26, 0x1d, 0x24, 0x6d, 0x56, 0x56, 0x50, 0x7d, 0x56,
	0x56, 0x10, 0x8d, 0x56, 0x56, 0x50, 0x9d, 0x50, 0x81, 0x24, 0x36, 0x70, 0xd4, 0x19, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x16, 0x62, 0xbb, 0x04, 0x33, 0x02, 0x00, 0x00,
}

func (m *QueryRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AssetID != nil {
		{
			size, err := m.AssetID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQueryRequest(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQueryRequest(dAtA []byte, offset int, v uint64) int {
	offset -= sovQueryRequest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AssetID != nil {
		l = m.AssetID.Size()
		n += 1 + l + sovQueryRequest(uint64(l))
	}
	return n
}

func sovQueryRequest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQueryRequest(x uint64) (n int) {
	return sovQueryRequest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryRequest
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQueryRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AssetID == nil {
				m.AssetID = &base.AssetID{}
			}
			if err := m.AssetID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryRequest
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQueryRequest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQueryRequest
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQueryRequest
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQueryRequest
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQueryRequest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQueryRequest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQueryRequest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQueryRequest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQueryRequest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQueryRequest = fmt.Errorf("proto: unexpected end of group")
)
