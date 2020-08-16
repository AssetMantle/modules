/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"strings"
)

type properties struct {
	PropertyList []types.Property `json:"propertyList"`
}

var _ types.Properties = (*properties)(nil)

func (properties properties) Get(id types.ID) types.Property {
	for _, property := range properties.GetList() {
		if property.GetID().Compare(id) == 0 {
			return property
		}
	}
	return nil
}
func (properties properties) GetList() []types.Property {
	var propertyList []types.Property
	for _, baseProperty := range properties.PropertyList {
		propertyList = append(propertyList, baseProperty)
	}
	return propertyList
}
func (properties properties) Add(property types.Property) types.Properties {
	propertyList := properties.GetList()
	propertyList = append(propertyList, property)
	return NewProperties(propertyList)
}
func (properties properties) Remove(property types.Property) types.Properties {
	propertyList := properties.GetList()
	for i, oldProperty := range propertyList {
		if oldProperty.GetID().Compare(property.GetID()) == 0 {
			propertyList = append(propertyList[:i], propertyList[i+1:]...)
		}
	}
	return NewProperties(propertyList)
}
func (properties properties) Mutate(property types.Property) types.Properties {
	propertyList := properties.GetList()
	for i, oldProperty := range propertyList {
		if oldProperty.GetID().Compare(property.GetID()) == 0 {
			propertyList[i] = property
		}
	}
	return NewProperties(propertyList)
}
func NewProperties(propertyList []types.Property) types.Properties {
	return properties{
		PropertyList: propertyList,
	}
}
func ReadProperties(Properties string) types.Properties {
	var propertyList []types.Property
	properties := strings.Split(Properties, constants.PropertiesSeparator)
	for _, property := range properties {
		propertyList = append(propertyList, ReadProperty(property))
	}
	return NewProperties(propertyList)
}
