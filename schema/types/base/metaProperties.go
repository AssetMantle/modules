// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"strings"

	"github.com/AssetMantle/modules/constants"
	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/types"
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
			metaProperty, err := ReadMetaProperty(metaPropertyString)
			if err != nil {
				return nil, err
			}

			metaPropertyList = append(metaPropertyList, metaProperty)
		}
	}

	return NewMetaProperties(metaPropertyList...), nil
}

func ReadData(dataString string) (types.Data, error) {
	dataTypeAndString := strings.SplitN(dataString, constants.DataTypeAndValueSeparator, 2)
	if len(dataTypeAndString) == 2 {
		dataTypeID, dataString := dataTypeAndString[0], dataTypeAndString[1]

		var data types.Data

		var Error error

		switch baseIDs.NewID(dataTypeID) {
		case base.DecDataID:
			data, Error = base.ReadDecData(dataString)
		case base.IDDataID:
			data, Error = base.ReadIDData(dataString)
		case base.HeightDataID:
			data, Error = base.ReadHeightData(dataString)
		case base.StringDataID:
			data, Error = base.ReadStringData(dataString)
		case base.AccAddressDataID:
			data, Error = base.ReadAccAddressData(dataString)
			// TODO Check
		case base.ListDataID:
			data, Error = base.ReadAccAddressListData(dataString)
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
