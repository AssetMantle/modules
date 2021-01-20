/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type properties struct {
	PropertyList []types.Property `json:"propertyList"`
}

var _ types.Properties = (*properties)(nil)

func (properties properties) Get(id types.ID) types.Property {
	for _, property := range properties.GetList() {
		if property.GetID().Equals(id) {
			return property
		}
	}

	return nil
}
func (properties properties) GetList() []types.Property {
	return properties.PropertyList
}
func (properties properties) Add(property types.Property) types.Properties {
	propertyList := properties.GetList()
	propertyList = append(propertyList, property)

	return NewProperties(propertyList...)
}
func (properties properties) Remove(property types.Property) types.Properties {
	propertyList := properties.GetList()
	for i, oldProperty := range propertyList {
		if oldProperty.GetID().Equals(property.GetID()) {
			propertyList = append(propertyList[:i], propertyList[i+1:]...)
		}
	}

	return NewProperties(propertyList...)
}
func (properties properties) Mutate(property types.Property) types.Properties {
	propertyList := properties.GetList()
	for i, oldProperty := range propertyList {
		if oldProperty.GetID().Equals(property.GetID()) {
			propertyList[i] = property
		}
	}

	return NewProperties(propertyList...)
}
func NewProperties(propertyList ...types.Property) types.Properties {
	return properties{
		PropertyList: propertyList,
	}
}
func ReadProperties(propertiesString string) (types.Properties, error) {
	properties, Error := ReadMetaProperties(propertiesString)
	if Error != nil {
		return nil, Error
	}

	return properties.RemoveData(), nil
}
