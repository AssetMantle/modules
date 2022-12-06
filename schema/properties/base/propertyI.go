package base

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/traits"
)

type propertyI PropertyI

func (p propertyI) GetID() ids.PropertyID {
	return p.Impl.(properties.Property).GetID()
}

func (p propertyI) GetDataID() ids.DataID {
	return p.Impl.(properties.Property).GetDataID()
}

func (p propertyI) GetKey() ids.StringID {
	return p.Impl.(properties.Property).GetKey()
}

func (p propertyI) GetType() ids.StringID {
	return p.Impl.(properties.Property).GetType()
}

func (p propertyI) IsMeta() bool {
	return p.Impl.(properties.Property).IsMeta()
}

func (p propertyI) Compare(listable traits.Listable) int {
	return p.Impl.(properties.Property).Compare(listable)
}

var _ properties.Property = (*propertyI)(nil)
