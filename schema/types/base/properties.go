/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

var _ types.Properties = (*Properties)(nil)

func (properties Properties) Get(id types.ID) types.Property {
	for _, property := range properties.GetList() {
		if property.GetID().Compare(id) == 0 {
			return property
		}
	}

	return nil
}
func (properties Properties) GetList() []types.Property {
	newPropertyList := make([]types.Property, len(properties.PropertyList))
	for i, element := range properties.PropertyList {
		newPropertyList[i] = &element
	}
	return newPropertyList
}
func (properties Properties) Add(propertyList ...types.Property) types.Properties {
	newPropertyList := properties.GetList()

	for _, addProperty := range propertyList {
		if properties.Get(addProperty.GetID()) == nil {
			newPropertyList = append(newPropertyList, addProperty)
		}
	}

	return NewProperties(newPropertyList...)
}
func (properties Properties) Remove(propertyList ...types.Property) types.Properties {
	newPropertyList := properties.GetList()

	for _, removeProperty := range propertyList {
		for i, oldProperty := range newPropertyList {
			if oldProperty.GetID().Compare(removeProperty.GetID()) == 0 {
				newPropertyList = append(newPropertyList[:i], newPropertyList[i+1:]...)
				break
			}
		}
	}

	return NewProperties(newPropertyList...)
}
func (properties Properties) Mutate(propertyList ...types.Property) types.Properties {
	newPropertyList := properties.GetList()

	for _, mutateProperty := range propertyList {
		for i, oldProperty := range newPropertyList {
			if oldProperty.GetID().Compare(mutateProperty.GetID()) == 0 {
				newPropertyList[i] = mutateProperty
				break
			}
		}
	}

	return NewProperties(newPropertyList...)
}
func NewProperties(propertyList ...types.Property) *Properties {
	newPropertyList := make([]Property, len(propertyList))
	for i, element := range propertyList {
		newPropertyList[i] = *NewProperty(element.GetID(), element.GetFact())
	}
	return &Properties{
		PropertyList: newPropertyList,
	}
}

func ReadProperties(propertiesString string) (types.Properties, error) {
	properties, Error := ReadMetaProperties(propertiesString)
	if Error != nil {
		return nil, Error
	}

	return properties.RemoveData(), nil
}
