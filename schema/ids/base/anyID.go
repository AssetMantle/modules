package base

import (
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

type idGetter interface {
	get() ids.ID
}

var _ idGetter = (*AnyID_ClassificationID)(nil)
var _ idGetter = (*AnyID_AssetID)(nil)
var _ idGetter = (*AnyID_DataID)(nil)
var _ idGetter = (*AnyID_HashID)(nil)
var _ idGetter = (*AnyID_IdentityID)(nil)
var _ idGetter = (*AnyID_MaintainerID)(nil)
var _ idGetter = (*AnyID_OrderID)(nil)
var _ idGetter = (*AnyID_AnyOwnableID)(nil)
var _ idGetter = (*AnyID_PropertyID)(nil)
var _ idGetter = (*AnyID_SplitID)(nil)
var _ idGetter = (*AnyID_StringID)(nil)

func (m *AnyID_ClassificationID) get() ids.ID {
	return m.ClassificationID
}
func (m *AnyID_AssetID) get() ids.ID {
	return m.AssetID
}
func (m *AnyID_DataID) get() ids.ID {
	return m.DataID
}
func (m *AnyID_HashID) get() ids.ID {
	return m.HashID
}
func (m *AnyID_IdentityID) get() ids.ID {
	return m.IdentityID
}
func (m *AnyID_MaintainerID) get() ids.ID {
	return m.MaintainerID
}
func (m *AnyID_OrderID) get() ids.ID {
	return m.OrderID
}
func (m *AnyID_AnyOwnableID) get() ids.ID {
	return m.AnyOwnableID
}
func (m *AnyID_PropertyID) get() ids.ID {
	return m.PropertyID
}
func (m *AnyID_SplitID) get() ids.ID {
	return m.SplitID
}
func (m *AnyID_StringID) get() ids.ID {
	return m.StringID
}

var _ ids.AnyID = (*AnyID)(nil)

func (m *AnyID) AsString() string {
	return m.Impl.(idGetter).get().AsString()
}
func (m *AnyID) Get() ids.ID {
	return m.Impl.(idGetter).get()
}
func (m *AnyID) Compare(listable traits.Listable) int {
	return m.Impl.(idGetter).get().Compare(listable)
}
func (m *AnyID) Bytes() []byte {
	return m.Impl.(idGetter).get().Bytes()
}
func (m *AnyID) ToAnyID() ids.AnyID {
	return m.Impl.(idGetter).get().ToAnyID()
}

func idFromListable(listable traits.Listable) (ids.ID, error) {
	switch listable.(type) {
	case ids.ID:
		return listable.(ids.ID), nil

	default:
		return nil, errorConstants.MetaDataError
	}
}
