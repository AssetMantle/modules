// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"strings"

	"github.com/AssetMantle/modules/constants"
	"github.com/AssetMantle/modules/schema/capabilities"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/types"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

type metaPropertyList struct {
	types.List
}

var _ lists.MetaPropertyList = (*metaPropertyList)(nil)

func (metaPropertyList metaPropertyList) GetList() []types.MetaProperty {
	metaProperties := make([]types.MetaProperty, metaPropertyList.Size())
	for i, listable := range metaPropertyList.Get() {
		metaProperties[i] = listable.(types.MetaProperty)
	}
	return metaProperties
}
func (metaPropertyList metaPropertyList) GetMetaProperty(propertyID ids.PropertyID) types.MetaProperty {
	if i, found := metaPropertyList.List.Search(baseTypes.NewEmptyMetaPropertyFromID(propertyID)); found {
		return metaPropertyList.GetList()[i]
	}
	return nil
}
func (metaPropertyList metaPropertyList) ToPropertyList() lists.PropertyList {
	propertyList := make([]types.Property, len(metaPropertyList.GetList()))
	for i, oldMetaProperty := range metaPropertyList.GetList() {
		propertyList[i] = oldMetaProperty.RemoveData()
	}

	return NewPropertyList(propertyList...)
}
func (metaPropertyList metaPropertyList) Add(metaProperties ...types.MetaProperty) lists.MetaPropertyList {
	metaPropertyList.List = metaPropertyList.List.Add(metaPropertiesToListables(metaProperties...)...)
	return metaPropertyList
}
func metaPropertiesToListables(metaProperties ...types.MetaProperty) []capabilities.Listable {
	listables := make([]capabilities.Listable, len(metaProperties))
	for i, property := range metaProperties {
		listables[i] = property
	}
	return listables
}
func NewMetaProperties(metaProperties ...types.MetaProperty) lists.MetaPropertyList {
	return metaPropertyList{List: baseTypes.NewList(metaPropertiesToListables(metaProperties...)...)}
}

func ReadMetaProperties(metaPropertiesString string) (lists.MetaPropertyList, error) {
	var metaPropertyList []types.MetaProperty

	metaProperties := strings.Split(metaPropertiesString, constants.PropertiesSeparator)
	for _, metaPropertyString := range metaProperties {
		if metaPropertyString != "" {
			metaProperty, err := baseTypes.ReadMetaProperty(metaPropertyString)
			if err != nil {
				return nil, err
			}

			metaPropertyList = append(metaPropertyList, metaProperty)
		}
	}

	return NewMetaProperties(metaPropertyList...), nil
}
