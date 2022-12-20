// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: schema/ids/base/id.v1.proto

package base

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type ID struct {
	// Types that are valid to be assigned to Impl:
	//	*ID_AssetID
	//	*ID_ClassificationID
	//	*ID_DataID
	//	*ID_HashID
	//	*ID_IdentityID
	//	*ID_MaintainerID
	//	*ID_OrderID
	//	*ID_OwnableID
	//	*ID_PropertyID
	//	*ID_SplitID
	//	*ID_StringID
	Impl isID_Impl `protobuf_oneof:"impl"`
}

func (m *ID) Reset()         { *m = ID{} }
func (m *ID) String() string { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()    {}
func (*ID) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3244aad97cc7d61, []int{0}
}
func (m *ID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ID.Merge(m, src)
}
func (m *ID) XXX_Size() int {
	return m.Size()
}
func (m *ID) XXX_DiscardUnknown() {
	xxx_messageInfo_ID.DiscardUnknown(m)
}

var xxx_messageInfo_ID proto.InternalMessageInfo

type isID_Impl interface {
	isID_Impl()
	MarshalTo([]byte) (int, error)
	Size() int
}

type ID_AssetID struct {
	AssetID *AssetID `protobuf:"bytes,1,opt,name=AssetID,proto3,oneof" json:"AssetID,omitempty"`
}
type ID_ClassificationID struct {
	ClassificationID *ClassificationID `protobuf:"bytes,2,opt,name=ClassificationID,proto3,oneof" json:"ClassificationID,omitempty"`
}
type ID_DataID struct {
	DataID *DataID `protobuf:"bytes,3,opt,name=DataID,proto3,oneof" json:"DataID,omitempty"`
}
type ID_HashID struct {
	HashID *HashID `protobuf:"bytes,4,opt,name=HashID,proto3,oneof" json:"HashID,omitempty"`
}
type ID_IdentityID struct {
	IdentityID *IdentityID `protobuf:"bytes,5,opt,name=IdentityID,proto3,oneof" json:"IdentityID,omitempty"`
}
type ID_MaintainerID struct {
	MaintainerID *MaintainerID `protobuf:"bytes,6,opt,name=MaintainerID,proto3,oneof" json:"MaintainerID,omitempty"`
}
type ID_OrderID struct {
	OrderID *OrderID `protobuf:"bytes,7,opt,name=OrderID,proto3,oneof" json:"OrderID,omitempty"`
}
type ID_OwnableID struct {
	OwnableID *OwnableID `protobuf:"bytes,8,opt,name=OwnableID,proto3,oneof" json:"OwnableID,omitempty"`
}
type ID_PropertyID struct {
	PropertyID *PropertyID `protobuf:"bytes,9,opt,name=PropertyID,proto3,oneof" json:"PropertyID,omitempty"`
}
type ID_SplitID struct {
	SplitID *SplitID `protobuf:"bytes,10,opt,name=SplitID,proto3,oneof" json:"SplitID,omitempty"`
}
type ID_StringID struct {
	StringID *StringID `protobuf:"bytes,11,opt,name=StringID,proto3,oneof" json:"StringID,omitempty"`
}

func (*ID_AssetID) isID_Impl()          {}
func (*ID_ClassificationID) isID_Impl() {}
func (*ID_DataID) isID_Impl()           {}
func (*ID_HashID) isID_Impl()           {}
func (*ID_IdentityID) isID_Impl()       {}
func (*ID_MaintainerID) isID_Impl()     {}
func (*ID_OrderID) isID_Impl()          {}
func (*ID_OwnableID) isID_Impl()        {}
func (*ID_PropertyID) isID_Impl()       {}
func (*ID_SplitID) isID_Impl()          {}
func (*ID_StringID) isID_Impl()         {}

func (m *ID) GetImpl() isID_Impl {
	if m != nil {
		return m.Impl
	}
	return nil
}

func (m *ID) GetAssetID() *AssetID {
	if x, ok := m.GetImpl().(*ID_AssetID); ok {
		return x.AssetID
	}
	return nil
}

func (m *ID) GetClassificationID() *ClassificationID {
	if x, ok := m.GetImpl().(*ID_ClassificationID); ok {
		return x.ClassificationID
	}
	return nil
}

func (m *ID) GetDataID() *DataID {
	if x, ok := m.GetImpl().(*ID_DataID); ok {
		return x.DataID
	}
	return nil
}

func (m *ID) GetHashID() *HashID {
	if x, ok := m.GetImpl().(*ID_HashID); ok {
		return x.HashID
	}
	return nil
}

func (m *ID) GetIdentityID() *IdentityID {
	if x, ok := m.GetImpl().(*ID_IdentityID); ok {
		return x.IdentityID
	}
	return nil
}

func (m *ID) GetMaintainerID() *MaintainerID {
	if x, ok := m.GetImpl().(*ID_MaintainerID); ok {
		return x.MaintainerID
	}
	return nil
}

func (m *ID) GetOrderID() *OrderID {
	if x, ok := m.GetImpl().(*ID_OrderID); ok {
		return x.OrderID
	}
	return nil
}

func (m *ID) GetOwnableID() *OwnableID {
	if x, ok := m.GetImpl().(*ID_OwnableID); ok {
		return x.OwnableID
	}
	return nil
}

func (m *ID) GetPropertyID() *PropertyID {
	if x, ok := m.GetImpl().(*ID_PropertyID); ok {
		return x.PropertyID
	}
	return nil
}

func (m *ID) GetSplitID() *SplitID {
	if x, ok := m.GetImpl().(*ID_SplitID); ok {
		return x.SplitID
	}
	return nil
}

func (m *ID) GetStringID() *StringID {
	if x, ok := m.GetImpl().(*ID_StringID); ok {
		return x.StringID
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ID) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ID_AssetID)(nil),
		(*ID_ClassificationID)(nil),
		(*ID_DataID)(nil),
		(*ID_HashID)(nil),
		(*ID_IdentityID)(nil),
		(*ID_MaintainerID)(nil),
		(*ID_OrderID)(nil),
		(*ID_OwnableID)(nil),
		(*ID_PropertyID)(nil),
		(*ID_SplitID)(nil),
		(*ID_StringID)(nil),
	}
}

func init() {
	proto.RegisterType((*ID)(nil), "ids.ID")
}

func init() { proto.RegisterFile("schema/ids/base/id.v1.proto", fileDescriptor_a3244aad97cc7d61) }

var fileDescriptor_a3244aad97cc7d61 = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x4f, 0x8b, 0xd3, 0x40,
	0x18, 0xc6, 0x27, 0x69, 0x6d, 0xb7, 0xd3, 0xd5, 0xd5, 0x80, 0x30, 0xac, 0x10, 0xd7, 0x95, 0x85,
	0x15, 0x21, 0xa1, 0x7a, 0xf0, 0x6c, 0xb7, 0x60, 0x72, 0x28, 0x5b, 0xb2, 0x20, 0x8b, 0x78, 0x99,
	0x66, 0xc6, 0xed, 0x40, 0xfe, 0x91, 0x99, 0x55, 0xfc, 0x06, 0x1e, 0xbd, 0x7a, 0xf5, 0xe8, 0x27,
	0x11, 0x4f, 0x7b, 0xf4, 0x28, 0xed, 0xcd, 0x4f, 0xb1, 0xcc, 0x9b, 0xb4, 0x9d, 0x26, 0xb0, 0x97,
	0x32, 0x7d, 0xe6, 0xf7, 0x24, 0xef, 0xfb, 0x3e, 0x6f, 0xf0, 0x13, 0x19, 0x2f, 0x78, 0x4a, 0x7d,
	0xc1, 0xa4, 0x3f, 0xa7, 0x92, 0xfb, 0x82, 0x79, 0x9f, 0x47, 0x5e, 0x51, 0xe6, 0x2a, 0x77, 0x3a,
	0x82, 0xc9, 0xc3, 0xa3, 0x26, 0x41, 0xa5, 0xe4, 0x2a, 0x9c, 0x6c, 0xb0, 0xc3, 0x17, 0x4d, 0x22,
	0x4e, 0xa8, 0x94, 0xe2, 0x93, 0x88, 0xa9, 0x12, 0x79, 0x66, 0xa2, 0x4f, 0x9b, 0x28, 0xa3, 0x8a,
	0xde, 0x09, 0x2c, 0xa8, 0x5c, 0x98, 0xc0, 0xf3, 0x76, 0xc1, 0x3c, 0x53, 0x42, 0x7d, 0x35, 0xa1,
	0x93, 0x26, 0x94, 0x52, 0x91, 0x29, 0x2a, 0x32, 0x5e, 0x9a, 0x58, 0xab, 0xb5, 0xbc, 0x64, 0xbb,
	0xc4, 0x71, 0x8b, 0xf8, 0x92, 0xd1, 0x79, 0xc2, 0xef, 0xac, 0xa8, 0x28, 0xf3, 0x82, 0x97, 0xbb,
	0x15, 0xb5, 0x5e, 0x25, 0x8b, 0x44, 0xec, 0x4c, 0xf1, 0x59, 0x8b, 0x50, 0xa5, 0xc8, 0xae, 0x0c,
	0xe4, 0xf8, 0x47, 0x17, 0xdb, 0xe1, 0xc4, 0x39, 0xc5, 0xfd, 0xb7, 0x55, 0x06, 0xc4, 0x3a, 0xb2,
	0x4e, 0x87, 0xaf, 0xf6, 0x3d, 0xc1, 0xa4, 0x57, 0x6b, 0x01, 0x8a, 0xd6, 0xd7, 0xce, 0x19, 0x7e,
	0x78, 0xd6, 0xc8, 0x82, 0xd8, 0x60, 0x79, 0x0c, 0x96, 0xe6, 0x65, 0x80, 0xa2, 0x96, 0xc1, 0x39,
	0xc1, 0xbd, 0x09, 0xa4, 0x44, 0x3a, 0x60, 0x1d, 0x82, 0xb5, 0x92, 0x02, 0x14, 0xd5, 0x97, 0x1a,
	0x0b, 0x20, 0x2b, 0xd2, 0x35, 0xb0, 0x4a, 0xd2, 0x58, 0x75, 0x72, 0x46, 0x18, 0x87, 0x9b, 0xc4,
	0xc8, 0x3d, 0x40, 0x0f, 0x00, 0xdd, 0xca, 0x01, 0x8a, 0x0c, 0xc8, 0x79, 0x83, 0xf7, 0xa7, 0x46,
	0x7e, 0xa4, 0x07, 0xa6, 0x47, 0x60, 0x32, 0x2f, 0x02, 0x14, 0xed, 0x80, 0x7a, 0x50, 0xe7, 0x55,
	0xa2, 0xa4, 0x6f, 0x0c, 0xaa, 0xd6, 0xf4, 0xa0, 0xea, 0xa3, 0xe3, 0xe1, 0xc1, 0xf9, 0x3a, 0x59,
	0xb2, 0x07, 0xec, 0x83, 0x8a, 0x5d, 0xab, 0x01, 0x8a, 0xb6, 0x88, 0xee, 0x62, 0xb6, 0x49, 0x99,
	0x0c, 0x8c, 0x2e, 0xb6, 0xb2, 0xee, 0x62, 0xfb, 0x4f, 0x17, 0x73, 0x51, 0x65, 0x4e, 0xb0, 0x51,
	0x4c, 0xad, 0xe9, 0x62, 0xea, 0xa3, 0xf3, 0x12, 0xef, 0x5d, 0xd4, 0xd9, 0x93, 0x21, 0xa0, 0xf7,
	0x2b, 0xb4, 0x16, 0x03, 0x14, 0x6d, 0x80, 0x71, 0x0f, 0x77, 0x45, 0x5a, 0x24, 0xe3, 0x6f, 0xd6,
	0xef, 0xa5, 0x6b, 0xdd, 0x2c, 0x5d, 0xeb, 0xdf, 0xd2, 0xb5, 0xbe, 0xaf, 0x5c, 0x74, 0xb3, 0x72,
	0xd1, 0xdf, 0x95, 0x8b, 0x70, 0x3f, 0xce, 0x53, 0xfd, 0x80, 0xf1, 0x20, 0x64, 0xef, 0x47, 0x33,
	0xbd, 0x4a, 0x33, 0xeb, 0x83, 0x77, 0x25, 0xd4, 0xe2, 0x7a, 0xee, 0xc5, 0x79, 0xea, 0xc3, 0xbe,
	0x4c, 0x69, 0xa6, 0x12, 0xee, 0xa7, 0x39, 0xbb, 0x4e, 0xb8, 0xf4, 0x1b, 0xeb, 0xf8, 0xd3, 0xee,
	0x84, 0x97, 0x97, 0xbf, 0xec, 0x4e, 0xc8, 0xe4, 0x1f, 0xf8, 0x5d, 0xda, 0x07, 0x21, 0x93, 0x1f,
	0xdf, 0xcd, 0xc6, 0x53, 0xae, 0xa8, 0xfe, 0x90, 0xff, 0x83, 0x3e, 0xef, 0xc1, 0xb6, 0xbe, 0xbe,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x88, 0x44, 0x87, 0x5c, 0x04, 0x00, 0x00,
}

func (m *ID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Impl != nil {
		{
			size := m.Impl.Size()
			i -= size
			if _, err := m.Impl.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *ID_AssetID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ID_AssetID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.AssetID != nil {
		{
			size, err := m.AssetID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIdV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *ID_ClassificationID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ID_ClassificationID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ClassificationID != nil {
		{
			size, err := m.ClassificationID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIdV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *ID_DataID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ID_DataID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.DataID != nil {
		{
			size, err := m.DataID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIdV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *ID_HashID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ID_HashID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.HashID != nil {
		{
			size, err := m.HashID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIdV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *ID_IdentityID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ID_IdentityID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.IdentityID != nil {
		{
			size, err := m.IdentityID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIdV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	return len(dAtA) - i, nil
}
func (m *ID_MaintainerID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ID_MaintainerID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.MaintainerID != nil {
		{
			size, err := m.MaintainerID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIdV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	return len(dAtA) - i, nil
}
func (m *ID_OrderID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ID_OrderID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.OrderID != nil {
		{
			size, err := m.OrderID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIdV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	return len(dAtA) - i, nil
}
func (m *ID_OwnableID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ID_OwnableID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.OwnableID != nil {
		{
			size, err := m.OwnableID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIdV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	return len(dAtA) - i, nil
}
func (m *ID_PropertyID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ID_PropertyID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PropertyID != nil {
		{
			size, err := m.PropertyID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIdV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	return len(dAtA) - i, nil
}
func (m *ID_SplitID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ID_SplitID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SplitID != nil {
		{
			size, err := m.SplitID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIdV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	return len(dAtA) - i, nil
}
func (m *ID_StringID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ID_StringID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.StringID != nil {
		{
			size, err := m.StringID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIdV1(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x5a
	}
	return len(dAtA) - i, nil
}
func encodeVarintIdV1(dAtA []byte, offset int, v uint64) int {
	offset -= sovIdV1(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Impl != nil {
		n += m.Impl.Size()
	}
	return n
}

func (m *ID_AssetID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AssetID != nil {
		l = m.AssetID.Size()
		n += 1 + l + sovIdV1(uint64(l))
	}
	return n
}
func (m *ID_ClassificationID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ClassificationID != nil {
		l = m.ClassificationID.Size()
		n += 1 + l + sovIdV1(uint64(l))
	}
	return n
}
func (m *ID_DataID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DataID != nil {
		l = m.DataID.Size()
		n += 1 + l + sovIdV1(uint64(l))
	}
	return n
}
func (m *ID_HashID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HashID != nil {
		l = m.HashID.Size()
		n += 1 + l + sovIdV1(uint64(l))
	}
	return n
}
func (m *ID_IdentityID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IdentityID != nil {
		l = m.IdentityID.Size()
		n += 1 + l + sovIdV1(uint64(l))
	}
	return n
}
func (m *ID_MaintainerID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaintainerID != nil {
		l = m.MaintainerID.Size()
		n += 1 + l + sovIdV1(uint64(l))
	}
	return n
}
func (m *ID_OrderID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OrderID != nil {
		l = m.OrderID.Size()
		n += 1 + l + sovIdV1(uint64(l))
	}
	return n
}
func (m *ID_OwnableID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OwnableID != nil {
		l = m.OwnableID.Size()
		n += 1 + l + sovIdV1(uint64(l))
	}
	return n
}
func (m *ID_PropertyID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PropertyID != nil {
		l = m.PropertyID.Size()
		n += 1 + l + sovIdV1(uint64(l))
	}
	return n
}
func (m *ID_SplitID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SplitID != nil {
		l = m.SplitID.Size()
		n += 1 + l + sovIdV1(uint64(l))
	}
	return n
}
func (m *ID_StringID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StringID != nil {
		l = m.StringID.Size()
		n += 1 + l + sovIdV1(uint64(l))
	}
	return n
}

func sovIdV1(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIdV1(x uint64) (n int) {
	return sovIdV1(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIdV1
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
			return fmt.Errorf("proto: ID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdV1
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
				return ErrInvalidLengthIdV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIdV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &AssetID{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &ID_AssetID{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassificationID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdV1
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
				return ErrInvalidLengthIdV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIdV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ClassificationID{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &ID_ClassificationID{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdV1
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
				return ErrInvalidLengthIdV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIdV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &DataID{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &ID_DataID{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HashID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdV1
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
				return ErrInvalidLengthIdV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIdV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &HashID{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &ID_HashID{v}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdentityID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdV1
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
				return ErrInvalidLengthIdV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIdV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &IdentityID{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &ID_IdentityID{v}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaintainerID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdV1
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
				return ErrInvalidLengthIdV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIdV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &MaintainerID{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &ID_MaintainerID{v}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdV1
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
				return ErrInvalidLengthIdV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIdV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &OrderID{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &ID_OrderID{v}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnableID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdV1
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
				return ErrInvalidLengthIdV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIdV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &OwnableID{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &ID_OwnableID{v}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PropertyID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdV1
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
				return ErrInvalidLengthIdV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIdV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &PropertyID{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &ID_PropertyID{v}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SplitID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdV1
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
				return ErrInvalidLengthIdV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIdV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &SplitID{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &ID_SplitID{v}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StringID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdV1
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
				return ErrInvalidLengthIdV1
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIdV1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &StringID{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Impl = &ID_StringID{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIdV1(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIdV1
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
func skipIdV1(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIdV1
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
					return 0, ErrIntOverflowIdV1
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
					return 0, ErrIntOverflowIdV1
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
				return 0, ErrInvalidLengthIdV1
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIdV1
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIdV1
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIdV1        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIdV1          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIdV1 = fmt.Errorf("proto: unexpected end of group")
)