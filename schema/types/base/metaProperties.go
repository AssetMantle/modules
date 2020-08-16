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
		if metaProperty.GetID().Compare(id) == 0 {
			return metaProperty
		}
	}
	return nil
}

func (metaProperties metaProperties) GetMetaPropertyList() []types.MetaProperty {
	var metaPropertyList []types.MetaProperty
	for _, metaProperty := range metaProperties.MetaPropertyList {
		metaPropertyList = append(metaPropertyList, metaProperty)
	}
	return metaPropertyList
}

func (metaProperties metaProperties) AddMetaProperty(metaProperty types.MetaProperty) types.MetaProperties {
	metaPropertyList := metaProperties.GetMetaPropertyList()
	metaPropertyList = append(metaPropertyList, metaProperty)
	return NewMetaProperties(metaPropertyList)
}

func (metaProperties metaProperties) RemoveMetaProperty(metaProperty types.MetaProperty) types.MetaProperties {
	metaPropertyList := metaProperties.GetMetaPropertyList()
	for i, oldMetaProperty := range metaPropertyList {
		if oldMetaProperty.GetID().Compare(metaProperty.GetID()) == 0 {
			metaPropertyList = append(metaPropertyList[:i], metaPropertyList[i+1:]...)
		}
	}
	return NewMetaProperties(metaPropertyList)
}

func (metaProperties metaProperties) MutateMetaProperty(metaProperty types.MetaProperty) types.MetaProperties {
	metaPropertyList := metaProperties.GetMetaPropertyList()
	for i, oldProperty := range metaPropertyList {
		if oldProperty.GetID().Compare(metaProperty.GetID()) == 0 {
			metaPropertyList[i] = metaProperty
		}
	}
	return NewMetaProperties(metaPropertyList)
}

func (metaProperties metaProperties) Get(id types.ID) types.Property {
	for _, property := range metaProperties.GetList() {
		if property.GetID().Compare(id) == 0 {
			return property
		}
	}
	return nil
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
	return NewProperties(propertyList)
}

func (metaProperties metaProperties) Remove(property types.Property) types.Properties {
	propertyList := metaProperties.GetList()
	for i, oldProperty := range propertyList {
		if oldProperty.GetID().Compare(property.GetID()) == 0 {
			propertyList = append(propertyList[:i], propertyList[i+1:]...)
		}
	}
	return NewProperties(propertyList)
}

func (metaProperties metaProperties) Mutate(property types.Property) types.Properties {
	propertyList := metaProperties.GetList()
	for i, oldProperty := range propertyList {
		if oldProperty.GetID().Compare(property.GetID()) == 0 {
			propertyList[i] = property
		}
	}
	return NewProperties(propertyList)
}

func (metaProperties metaProperties) RemoveData() types.Properties {
	var propertyList []types.Property
	for _, oldMetaProperty := range metaProperties.GetMetaPropertyList() {
		propertyList = append(propertyList, oldMetaProperty.RemoveData())
	}
	return NewProperties(propertyList)
}

func NewMetaProperties(metaPropertyList []types.MetaProperty) types.MetaProperties {
	return metaProperties{
		MetaPropertyList: metaPropertyList,
	}
}

func ReadMetaProperties(MetaProperties string) types.MetaProperties {
	var metaPropertyList []types.MetaProperty
	metaProperties := strings.Split(MetaProperties, constants.MetaPropertiesSeparator)
	for _, metaProperty := range metaProperties {
		metaPropertyList = append(metaPropertyList, ReadMetaProperty(metaProperty))
	}
	return NewMetaProperties(metaPropertyList)
}
