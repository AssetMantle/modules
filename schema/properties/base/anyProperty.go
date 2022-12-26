package base

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/traits"
)

type getter interface {
	getProperty() properties.Property
}

var _ getter = (*AnyProperty_MetaProperty)(nil)
var _ getter = (*AnyProperty_MesaProperty)(nil)

func (m *AnyProperty_MetaProperty) getProperty() properties.Property {
	return m.MetaProperty
}
func (m *AnyProperty_MesaProperty) getProperty() properties.Property {
	return m.MesaProperty
}

func (m *AnyProperty) GetID() ids.PropertyID {
	return m.Impl.(getter).getProperty().GetID()
}
func (m *AnyProperty) ScrubData() properties.Property {
	return m.Impl.(getter).getProperty().ScrubData()
}
func (m *AnyProperty) GetDataID() ids.DataID {
	return m.Impl.(getter).getProperty().GetDataID()
}
func (m *AnyProperty) GetKey() ids.StringID {
	return m.Impl.(getter).getProperty().GetKey()
}
func (m *AnyProperty) GetData() data.AnyData {
	return m.Impl.(getter).getProperty().GetData()
}
func (m *AnyProperty) GetType() ids.StringID {
	return m.Impl.(getter).getProperty().GetType()
}
func (m *AnyProperty) IsMeta() bool {
	return m.Impl.(getter).getProperty().IsMeta()
}
func (m *AnyProperty) ToAnyProperty() properties.AnyProperty {
	return m.Impl.(getter).getProperty().ToAnyProperty()
}
func (m *AnyProperty) Compare(listable traits.Listable) int {
	return m.Impl.(getter).getProperty().Compare(listable)
}
