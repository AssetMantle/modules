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

type metaProperties struct {
	MetaPropertyList []types.MetaProperty `json:"metaPropertyList"`
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
	return metaProperties.GetMetaProperty(id).RemoveData()
}

func (metaProperties metaProperties) GetList() []types.Property {
	var propertyList []types.Property
	for _, metaProperty := range metaProperties.MetaPropertyList {
		propertyList = append(propertyList, metaProperty)
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
	var propertyList []types.Property
	for _, oldMetaProperty := range metaProperties.GetMetaPropertyList() {
		propertyList = append(propertyList, oldMetaProperty.RemoveData())
	}
	return NewProperties(propertyList...)
}

func NewMetaProperties(metaPropertyList []types.MetaProperty) types.MetaProperties {
	return metaProperties{
		MetaPropertyList: metaPropertyList,
	}
}

func ReadMetaProperties(MetaProperties string) (types.MetaProperties, error) {
	var metaPropertyList []types.MetaProperty
	metaProperties := strings.Split(MetaProperties, constants.PropertiesSeparator)
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
