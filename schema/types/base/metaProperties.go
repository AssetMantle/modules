/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"strings"

	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type metaProperties struct {
	MetaPropertyList []types.MetaProperty `json:"metaPropertyList,omitempty"`
}

var _ types.MetaProperties = (*metaProperties)(nil)

func (metaProperties metaProperties) Get(id types.ID) types.MetaProperty {
	for _, metaProperty := range metaProperties.GetList() {
		if metaProperty.GetID().Compare(id) == 0 {
			return metaProperty
		}
	}

	return nil
}

func (metaProperties metaProperties) GetList() []types.MetaProperty {
	return metaProperties.MetaPropertyList
}
func (metaProperties metaProperties) Add(metaPropertyList ...types.MetaProperty) types.MetaProperties {
	newMetaPropertyList := metaProperties.GetList()

	for _, addMetaProperty := range metaPropertyList {
		if metaProperties.Get(addMetaProperty.GetID()) == nil {
			newMetaPropertyList = append(newMetaPropertyList, addMetaProperty)
		}
	}

	return NewMetaProperties(newMetaPropertyList...)
}
func (metaProperties metaProperties) Remove(metaPropertyList ...types.MetaProperty) types.MetaProperties {
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
func (metaProperties metaProperties) Mutate(metaPropertyList ...types.MetaProperty) types.MetaProperties {
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
func (metaProperties metaProperties) RemoveData() types.Properties {
	propertyList := make([]types.Property, len(metaProperties.GetList()))
	for i, oldMetaProperty := range metaProperties.GetList() {
		propertyList[i] = oldMetaProperty.RemoveData()
	}

	return NewProperties(propertyList...)
}

func NewMetaProperties(metaPropertyList ...types.MetaProperty) types.MetaProperties {
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

	return NewMetaProperties(metaPropertyList...), nil
}

func ReadMetaProperty(metaPropertyString string) (types.MetaProperty, error) {
	propertyIDAndData := strings.Split(metaPropertyString, constants.PropertyIDAndDataSeparator)
	if len(propertyIDAndData) == 2 && propertyIDAndData[0] != "" {
		data, Error := ReadData(propertyIDAndData[1])
		if Error != nil {
			return nil, Error
		}

		return NewMetaProperty(NewID(propertyIDAndData[0]), data), nil
	}

	return nil, errors.IncorrectFormat
}

func ReadData(dataString string) (types.Data, error) {
	dataTypeAndString := strings.SplitN(dataString, constants.DataTypeAndValueSeparator, 2)
	if len(dataTypeAndString) == 2 {
		dataType, dataString := dataTypeAndString[0], dataTypeAndString[1]

		var data types.Data

		var Error error

		switch NewID(dataType) {
		case decData{}.GetTypeID():
			data, Error = ReadDecData(dataString)
		case idData{}.GetTypeID():
			data, Error = ReadIDData(dataString)
		case heightData{}.GetTypeID():
			data, Error = ReadHeightData(dataString)
		case stringData{}.GetTypeID():
			data, Error = ReadStringData(dataString)
		case accAddressData{}.GetTypeID():
			data, Error = ReadAccAddressData(dataString)
		case listData{}.GetTypeID():
			data, Error = ReadAccAddressListData(dataString)
		default:
			data, Error = nil, errors.UnsupportedParameter
		}

		if Error != nil {
			return nil, Error
		}

		return data, nil
	}

	return nil, errors.IncorrectFormat
}
