package base

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

type getter interface {
	GetID() ids.ID
}

var _ ids.ID = (*AnyID)(nil)
var _ getter = (*AnyID_ClassificationID)(nil)
var _ getter = (*AnyID_AssetID)(nil)
var _ getter = (*AnyID_DataID)(nil)
var _ getter = (*AnyID_HashID)(nil)
var _ getter = (*AnyID_IdentityID)(nil)
var _ getter = (*AnyID_MaintainerID)(nil)
var _ getter = (*AnyID_OrderID)(nil)
var _ getter = (*AnyID_OwnableID)(nil)
var _ getter = (*AnyID_PropertyID)(nil)
var _ getter = (*AnyID_SplitID)(nil)
var _ getter = (*AnyID_StringID)(nil)

func (m *AnyID_ClassificationID) GetID() ids.ID {
	return m.ClassificationID
}
func (m *AnyID_AssetID) GetID() ids.ID {
	return m.AssetID
}
func (m *AnyID_DataID) GetID() ids.ID {
	return m.DataID
}
func (m *AnyID_HashID) GetID() ids.ID {
	return m.HashID
}
func (m *AnyID_IdentityID) GetID() ids.ID {
	return m.IdentityID
}
func (m *AnyID_MaintainerID) GetID() ids.ID {
	return m.MaintainerID
}
func (m *AnyID_OrderID) GetID() ids.ID {
	return m.OrderID
}
func (m *AnyID_OwnableID) GetID() ids.ID {
	return m.OwnableID
}
func (m *AnyID_PropertyID) GetID() ids.ID {
	return m.PropertyID
}
func (m *AnyID_SplitID) GetID() ids.ID {
	return m.SplitID
}
func (m *AnyID_StringID) GetID() ids.ID {
	return m.StringID
}

func (m *AnyID) Compare(listable traits.Listable) int {
	return m.Impl.(ids.ID).Compare(listable)
}

func (m *AnyID) Bytes() []byte {
	return m.Impl.(ids.ID).Bytes()
}

func (m *AnyID) ToAnyID() ids.AnyID {
	return m.Impl.(ids.ID).ToAnyID()
}
