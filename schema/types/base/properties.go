/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type properties struct {
	PropertyList []types.Property `json:"propertyList,omitempty"`
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
func (properties properties) Add(propertyList ...types.Property) types.Properties {
	newPropertyList := properties.GetList()

	for _, addProperty := range propertyList {
		if properties.Get(addProperty.GetID()) == nil {
			newPropertyList = append(newPropertyList, addProperty)
		}
	}

	return NewProperties(newPropertyList...)
}
func (properties properties) Remove(propertyList ...types.Property) types.Properties {
	newPropertyList := properties.GetList()

	for _, removeProperty := range propertyList {
		for i, oldProperty := range newPropertyList {
			if oldProperty.GetID().Equals(removeProperty.GetID()) {
				newPropertyList = append(newPropertyList[:i], newPropertyList[i+1:]...)
				break
			}
		}
	}

	return NewProperties(newPropertyList...)
}
func (properties properties) Mutate(propertyList ...types.Property) types.Properties {
	newPropertyList := properties.GetList()

	for _, mutateProperty := range propertyList {
		for i, oldProperty := range newPropertyList {
			if oldProperty.GetID().Equals(mutateProperty.GetID()) {
				newPropertyList[i] = mutateProperty
				break
			}
		}
	}

	return NewProperties(newPropertyList...)
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
