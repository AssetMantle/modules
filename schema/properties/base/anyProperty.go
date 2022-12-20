package base

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ properties.Property = (*AnyProperty)(nil)

func (m *AnyProperty) GetID() ids.PropertyID {
	return m.Impl.(properties.Property).GetID()
}
func (m *AnyProperty) ScrubData() properties.Property {
	return m.Impl.(properties.Property).ScrubData()
}
func (m *AnyProperty) GetDataID() ids.DataID {
	return m.Impl.(properties.Property).GetDataID()
}
func (m *AnyProperty) GetKey() ids.StringID {
	return m.Impl.(properties.Property).GetKey()
}
func (m *AnyProperty) GetData() data.AnyData {
	return m.Impl.(properties.Property).GetData()
}
func (m *AnyProperty) GetType() ids.StringID {
	return m.Impl.(properties.Property).GetType()
}
func (m *AnyProperty) IsMeta() bool {
	return m.Impl.(properties.Property).IsMeta()
}
func (m *AnyProperty) ToAnyProperty() properties.AnyProperty {
	return m.Impl.(properties.Property).ToAnyProperty()
}
func (m *AnyProperty) Compare(listable traits.Listable) int {
	return m.Impl.(properties.Property).Compare(listable)
}
