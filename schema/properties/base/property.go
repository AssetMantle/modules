package base

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ properties.Property = (*Property)(nil)

func (m *Property) GetID() ids.ID {
	return m.Impl.(properties.Property).GetID()
}

func (m *Property) GetDataID() ids.ID {
	return m.Impl.(properties.Property).GetDataID()
}

func (m *Property) GetKey() ids.ID {
	return m.Impl.(properties.Property).GetKey()
}

func (m *Property) GetType() ids.ID {
	return m.Impl.(properties.Property).GetType()
}

func (m *Property) IsMeta() bool {
	return m.Impl.(properties.Property).IsMeta()
}

func (m *Property) Compare(listable traits.Listable) int {
	return m.Impl.(properties.Property).Compare(listable)
}
