// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

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
		if property.GetID().Compare(id) == 0 {
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
			if oldProperty.GetID().Compare(removeProperty.GetID()) == 0 {
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
			if oldProperty.GetID().Compare(mutateProperty.GetID()) == 0 {
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
	properties, err := ReadMetaProperties(propertiesString)
	if err != nil {
		return nil, err
	}

	return properties.RemoveData(), nil
}
