// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/classifications/internal/queries/classification/queryRequest.proto

package classification

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
	ClassificationID *base.ClassificationID `protobuf:"bytes,1,opt,name=classification_i_d,json=classificationID,proto3" json:"classification_i_d,omitempty"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ee91ed9aa903bda, []int{0}
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

func (m *QueryRequest) GetClassificationID() *base.ClassificationID {
	if m != nil {
		return m.ClassificationID
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryRequest)(nil), "classifications.queries.classification.QueryRequest")
}

func init() {
	proto.RegisterFile("x/classifications/internal/queries/classification/queryRequest.proto", fileDescriptor_3ee91ed9aa903bda)
}

var fileDescriptor_3ee91ed9aa903bda = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0xa9, 0xd0, 0x4f, 0xce,
	0x49, 0x2c, 0x2e, 0xce, 0x4c, 0xcb, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6, 0xcf, 0xcc,
	0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x2f, 0x2c, 0x4d, 0x2d, 0xca, 0x4c, 0x2d, 0x46, 0x53,
	0x00, 0x16, 0xae, 0x0c, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x52, 0x43, 0x33, 0x43, 0x0f, 0xaa, 0x55, 0x0f, 0x55, 0x5c, 0x4a, 0xb1, 0x42, 0x3f, 0x33,
	0xa5, 0x58, 0x3f, 0x29, 0xb1, 0x38, 0x15, 0xcd, 0x54, 0x4f, 0x17, 0x88, 0x51, 0x4a, 0xc1, 0x5c,
	0x3c, 0x81, 0x48, 0x16, 0x08, 0x39, 0x73, 0x09, 0xa1, 0xaa, 0x8c, 0xcf, 0x8c, 0x4f, 0x91, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x36, 0x12, 0xd5, 0xcb, 0x4c, 0x29, 0xd6, 0x73, 0x46, 0x33, 0x28, 0x48,
	0x00, 0xdd, 0x68, 0xa7, 0xc7, 0x4c, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0,
	0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0xc0,
	0xa5, 0x95, 0x9c, 0x9f, 0xab, 0x47, 0x9c, 0xf3, 0x9d, 0x04, 0x91, 0x5d, 0x16, 0x00, 0x72, 0x6e,
	0x00, 0x63, 0x54, 0x40, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0xbe, 0x63,
	0x71, 0x71, 0x6a, 0x89, 0x6f, 0x62, 0x5e, 0x49, 0x4e, 0xaa, 0x7e, 0x6e, 0x7e, 0x4a, 0x69, 0x4e,
	0x6a, 0xb1, 0x3e, 0xc9, 0x01, 0xbc, 0x88, 0x89, 0xd9, 0x39, 0xd0, 0x79, 0x15, 0x93, 0x9a, 0x33,
	0x9a, 0xab, 0x02, 0xa1, 0xae, 0x42, 0x15, 0x3f, 0x85, 0xa1, 0x30, 0x06, 0xaa, 0x30, 0x06, 0x55,
	0xfc, 0x11, 0x93, 0x11, 0x71, 0x0a, 0x63, 0xdc, 0x03, 0x9c, 0x7c, 0x53, 0x4b, 0x12, 0x53, 0x12,
	0x4b, 0x12, 0x5f, 0x31, 0x69, 0xa0, 0x69, 0xb2, 0xb2, 0x82, 0xea, 0xb2, 0xb2, 0x42, 0x95, 0x49,
	0x62, 0x03, 0xc7, 0xa0, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x63, 0x27, 0xba, 0x54, 0x02,
	0x00, 0x00,
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
	if m.ClassificationID != nil {
		{
			size, err := m.ClassificationID.MarshalToSizedBuffer(dAtA[:i])
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
	if m.ClassificationID != nil {
		l = m.ClassificationID.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field ClassificationID", wireType)
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
			if m.ClassificationID == nil {
				m.ClassificationID = &base.ClassificationID{}
			}
			if err := m.ClassificationID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
