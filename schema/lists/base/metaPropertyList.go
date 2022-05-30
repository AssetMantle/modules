// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/traits"
)

type metaPropertyList struct {
	lists.List
}

var _ lists.MetaPropertyList = (*metaPropertyList)(nil)

func (metaPropertyList metaPropertyList) GetList() []properties.MetaProperty {
	metaProperties := make([]properties.MetaProperty, metaPropertyList.List.Size())
	for i, listable := range metaPropertyList.List.Get() {
		metaProperties[i] = listable.(properties.MetaProperty)
	}
	return metaProperties
}
func (metaPropertyList metaPropertyList) GetMetaProperty(propertyID ids.PropertyID) properties.MetaProperty {
	if i, found := metaPropertyList.List.Search(base.NewEmptyMetaPropertyFromID(propertyID)); found {
		return metaPropertyList.GetList()[i]
	}
	return nil
}
func (metaPropertyList metaPropertyList) ToPropertyList() lists.PropertyList {
	propertyList := make([]properties.Property, len(metaPropertyList.GetList()))
	for i, oldMetaProperty := range metaPropertyList.GetList() {
		propertyList[i] = oldMetaProperty.RemoveData()
	}

	return NewPropertyList(propertyList...)
}
func (metaPropertyList metaPropertyList) Add(metaProperties ...properties.MetaProperty) lists.MetaPropertyList {
	metaPropertyList.List = metaPropertyList.List.Add(metaPropertiesToListables(metaProperties...)...)
	return metaPropertyList
}
func metaPropertiesToListables(metaProperties ...properties.MetaProperty) []traits.Listable {
	listables := make([]traits.Listable, len(metaProperties))
	for i, property := range metaProperties {
		listables[i] = property
	}
	return listables
}
func NewMetaProperties(metaProperties ...properties.MetaProperty) lists.MetaPropertyList {
	return metaPropertyList{List: NewList(metaPropertiesToListables(metaProperties...)...)}
}
