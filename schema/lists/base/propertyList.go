// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/capabilities"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/types"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

type propertyList struct {
	types.List
}

var _ lists.PropertyList = (*propertyList)(nil)

// TODO write test
func (propertyList propertyList) GetProperty(propertyID ids.PropertyID) types.Property {
	if i, found := propertyList.List.Search(baseTypes.NewEmptyPropertyFromID(propertyID)); found {
		return propertyList.GetList()[i]
	}
	return nil
}
func (propertyList propertyList) GetList() []types.Property {
	properties := make([]types.Property, propertyList.Size())
	for i, listable := range propertyList.List.Get() {
		properties[i] = listable.(types.Property)
	}

	return properties
}
func (propertyList propertyList) Add(properties ...types.Property) lists.PropertyList {
	propertyList.List = propertyList.List.Add(propertiesToListables(properties...)...)
	return propertyList
}
func (propertyList propertyList) Remove(properties ...types.Property) lists.PropertyList {
	propertyList.List = propertyList.List.Remove(propertiesToListables(properties...)...)
	return propertyList
}
func (propertyList propertyList) Mutate(properties ...types.Property) lists.PropertyList {
	propertyList.List = propertyList.List.Mutate(propertiesToListables(properties...)...)
	return propertyList
}

func propertiesToListables(properties ...types.Property) []capabilities.Listable {
	listables := make([]capabilities.Listable, len(properties))
	for i, property := range properties {
		listables[i] = property
	}
	return listables
}
func NewPropertyList(properties ...types.Property) lists.PropertyList {
	return propertyList{List: baseTypes.NewList(propertiesToListables(properties...)...)}
}
func ReadProperties(propertiesString string) (lists.PropertyList, error) {
	properties, err := ReadMetaProperties(propertiesString)
	if err != nil {
		return nil, err
	}

	return properties.ToPropertyList(), nil
}
