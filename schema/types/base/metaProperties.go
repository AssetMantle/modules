/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"strings"

	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type metaProperties struct {
	MetaPropertyList []types.MetaProperty `json:"metaPropertyList,omitempty"`
}

var _ types.MetaProperties = (*metaProperties)(nil)

func (metaProperties metaProperties) GetMetaProperty(id types.ID) types.MetaProperty {
	for _, metaProperty := range metaProperties.GetMetaPropertyList() {
		if metaProperty.GetID().Equals(id) {
			return metaProperty
		}
	}

	return nil
}

func (metaProperties metaProperties) GetMetaPropertyList() []types.MetaProperty {
	return metaProperties.MetaPropertyList
}

func (metaProperties metaProperties) AddMetaProperty(metaProperty types.MetaProperty) types.MetaProperties {
	metaPropertyList := metaProperties.GetMetaPropertyList()
	metaPropertyList = append(metaPropertyList, metaProperty)

	return NewMetaProperties(metaPropertyList)
}

func (metaProperties metaProperties) RemoveMetaProperty(metaProperty types.MetaProperty) types.MetaProperties {
	metaPropertyList := metaProperties.GetMetaPropertyList()
	for i, oldMetaProperty := range metaPropertyList {
		if oldMetaProperty.GetID().Equals(metaProperty.GetID()) {
			metaPropertyList = append(metaPropertyList[:i], metaPropertyList[i+1:]...)
		}
	}

	return NewMetaProperties(metaPropertyList)
}

func (metaProperties metaProperties) MutateMetaProperty(metaProperty types.MetaProperty) types.MetaProperties {
	metaPropertyList := metaProperties.GetMetaPropertyList()
	for i, oldProperty := range metaPropertyList {
		if oldProperty.GetID().Equals(metaProperty.GetID()) {
			metaPropertyList[i] = metaProperty
		}
	}

	return NewMetaProperties(metaPropertyList)
}

func (metaProperties metaProperties) Get(id types.ID) types.Property {
	if metaProperty := metaProperties.GetMetaProperty(id); metaProperty != nil {
		return metaProperty.RemoveData()
	}

	return nil
}

func (metaProperties metaProperties) GetList() []types.Property {
	propertyList := make([]types.Property, len(metaProperties.MetaPropertyList))
	for i, metaProperty := range metaProperties.MetaPropertyList {
		propertyList[i] = metaProperty
	}

	return propertyList
}

func (metaProperties metaProperties) Add(property types.Property) types.Properties {
	propertyList := metaProperties.GetList()
	propertyList = append(propertyList, property)

	return NewProperties(propertyList...)
}

func (metaProperties metaProperties) Remove(property types.Property) types.Properties {
	propertyList := metaProperties.GetList()
	for i, oldProperty := range propertyList {
		if oldProperty.GetID().Equals(property.GetID()) {
			propertyList = append(propertyList[:i], propertyList[i+1:]...)
		}
	}

	return NewProperties(propertyList...)
}

func (metaProperties metaProperties) Mutate(property types.Property) types.Properties {
	propertyList := metaProperties.GetList()
	for i, oldProperty := range propertyList {
		if oldProperty.GetID().Equals(property.GetID()) {
			propertyList[i] = property
		}
	}

	return NewProperties(propertyList...)
}

func (metaProperties metaProperties) RemoveData() types.Properties {
	propertyList := make([]types.Property, len(metaProperties.GetMetaPropertyList()))
	for i, oldMetaProperty := range metaProperties.GetMetaPropertyList() {
		propertyList[i] = oldMetaProperty.RemoveData()
	}

	return NewProperties(propertyList...)
}

func NewMetaProperties(metaPropertyList []types.MetaProperty) types.MetaProperties {
	return metaProperties{
		MetaPropertyList: metaPropertyList,
	}
}

func ReadMetaProperties(metaPropertiesString string) (types.MetaProperties, error) {
	var metaPropertyList []types.MetaProperty

	metaProperties := strings.Split(metaPropertiesString, constants.PropertiesSeparator)
	for _, metaPropertyString := range metaProperties {
		if metaPropertyString != "" {
			metaProperty, Error := ReadMetaProperty(metaPropertyString)
			if Error != nil {
				return nil, Error
			}

			metaPropertyList = append(metaPropertyList, metaProperty)
		}
	}

	return NewMetaProperties(metaPropertyList), nil
}
