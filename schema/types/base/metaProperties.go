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

var _ types.MetaProperties = (*MetaProperties)(nil)

func (metaProperties MetaProperties) Get(id types.ID) types.MetaProperty {
	for _, metaProperty := range metaProperties.GetList() {
		if metaProperty.GetID().Compare(id) == 0 {
			return metaProperty
		}
	}

	return nil
}
func (metaProperties MetaProperties) GetList() []types.MetaProperty {
	return metaProperties.MetaPropertyList
}
func (metaProperties MetaProperties) Add(metaPropertyList ...types.MetaProperty) types.MetaProperties {
	newMetaPropertyList := metaProperties.GetList()

	for _, addMetaProperty := range metaPropertyList {
		if metaProperties.Get(addMetaProperty.GetID()) == nil {
			newMetaPropertyList = append(newMetaPropertyList, addMetaProperty)
		}
	}

	return NewMetaProperties(newMetaPropertyList...)
}
func (metaProperties MetaProperties) Remove(metaPropertyList ...types.MetaProperty) types.MetaProperties {
	newMetaPropertyList := metaProperties.GetList()

	for _, removeMetaProperty := range metaPropertyList {
		for i, oldMetaProperty := range newMetaPropertyList {
			if oldMetaProperty.GetID().Compare(removeMetaProperty.GetID()) == 0 {
				newMetaPropertyList = append(newMetaPropertyList[:i], newMetaPropertyList[i+1:]...)
				break
			}
		}
	}

	return NewMetaProperties(newMetaPropertyList...)
}
func (metaProperties MetaProperties) Mutate(metaPropertyList ...types.MetaProperty) types.MetaProperties {
	newMetaPropertyList := metaProperties.GetList()

	for _, mutateMetaProperty := range metaPropertyList {
		for i, oldMetaProperty := range newMetaPropertyList {
			if oldMetaProperty.GetID().Compare(mutateMetaProperty.GetID()) == 0 {
				newMetaPropertyList[i] = mutateMetaProperty
				break
			}
		}
	}

	return NewMetaProperties(newMetaPropertyList...)
}
func (metaProperties MetaProperties) RemoveData() types.Properties {
	propertyList := make([]types.Property, len(metaProperties.GetList()))
	for i, oldMetaProperty := range metaProperties.GetList() {
		propertyList[i] = oldMetaProperty.RemoveData()
	}

	return NewProperties(propertyList...)
}

func NewMetaProperties(metaPropertyList ...types.MetaProperty) types.MetaProperties {
	return MetaProperties{
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

	return NewMetaProperties(metaPropertyList...), nil
}
