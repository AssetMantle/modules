package base

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/traits"
)

type getter interface {
	isAnyProperty_Impl
	GetProperty() properties.Property
}

var _ properties.Property = (*AnyProperty)(nil)
var _ getter = (*AnyProperty_MetaProperty)(nil)
var _ getter = (*AnyProperty_MesaProperty)(nil)

func (m *AnyProperty_MetaProperty) GetProperty() properties.Property {
	return m.MetaProperty
}
func (m *AnyProperty_MesaProperty) GetProperty() properties.Property {
	return m.MesaProperty
}

func (m *AnyProperty) GetID() ids.PropertyID {
	return m.Impl.(getter).GetProperty().GetID()
}
func (m *AnyProperty) ScrubData() properties.Property {
	return m.Impl.(getter).GetProperty().ScrubData()
}
func (m *AnyProperty) GetDataID() ids.DataID {
	return m.Impl.(getter).GetProperty().GetDataID()
}
func (m *AnyProperty) GetKey() ids.StringID {
	return m.Impl.(getter).GetProperty().GetKey()
}
func (m *AnyProperty) GetData() data.AnyData {
	return m.Impl.(getter).GetProperty().GetData()
}
func (m *AnyProperty) GetType() ids.StringID {
	return m.Impl.(getter).GetProperty().GetType()
}
func (m *AnyProperty) IsMeta() bool {
	return m.Impl.(getter).GetProperty().IsMeta()
}
func (m *AnyProperty) ToAnyProperty() properties.AnyProperty {
	return m.Impl.(getter).GetProperty().ToAnyProperty()
}
func (m *AnyProperty) Compare(listable traits.Listable) int {
	return m.Impl.(getter).GetProperty().Compare(listable)
}
