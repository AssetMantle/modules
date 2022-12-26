package base

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

type getter interface {
	getID() ids.ID
}

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

func (m *AnyID_ClassificationID) getID() ids.ID {
	return m.ClassificationID
}
func (m *AnyID_AssetID) getID() ids.ID {
	return m.AssetID
}
func (m *AnyID_DataID) getID() ids.ID {
	return m.DataID
}
func (m *AnyID_HashID) getID() ids.ID {
	return m.HashID
}
func (m *AnyID_IdentityID) getID() ids.ID {
	return m.IdentityID
}
func (m *AnyID_MaintainerID) getID() ids.ID {
	return m.MaintainerID
}
func (m *AnyID_OrderID) getID() ids.ID {
	return m.OrderID
}
func (m *AnyID_OwnableID) getID() ids.ID {
	return m.OwnableID
}
func (m *AnyID_PropertyID) getID() ids.ID {
	return m.PropertyID
}
func (m *AnyID_SplitID) getID() ids.ID {
	return m.SplitID
}
func (m *AnyID_StringID) getID() ids.ID {
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
