// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"strings"

	"github.com/AssetMantle/modules/constants"
	"github.com/AssetMantle/modules/constants/errors"
	idsConstants "github.com/AssetMantle/modules/constants/ids"
	"github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/traits"
	"github.com/AssetMantle/modules/schema/types"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

type metaPropertyList struct {
	types.List
}

var _ lists.MetaPropertyList = (*metaPropertyList)(nil)

func (metaPropertyList metaPropertyList) GetList() []types.MetaProperty {
	metaProperties := make([]types.MetaProperty, metaPropertyList.Size())
	for i, listable := range metaPropertyList.Get() {
		metaProperties[i] = listable.(types.MetaProperty)
	}
	return metaProperties
}
func (metaPropertyList metaPropertyList) GetMetaProperty(propertyID ids.PropertyID) types.MetaProperty {
	if i, found := metaPropertyList.List.Search(baseTypes.NewMetaPropertyFromID(propertyID)); found {
		return metaPropertyList.GetList()[i]
	}
	return nil
}
func (metaPropertyList metaPropertyList) ToPropertyList() lists.PropertyList {
	propertyList := make([]types.Property, len(metaPropertyList.GetList()))
	for i, oldMetaProperty := range metaPropertyList.GetList() {
		propertyList[i] = oldMetaProperty.RemoveData()
	}

	return NewPropertyList(propertyList...)
}
func (metaPropertyList metaPropertyList) Add(metaProperties ...types.MetaProperty) lists.MetaPropertyList {
	metaPropertyList.List = metaPropertyList.List.Add(metaPropertiesToListables(metaProperties...)...)
	return metaPropertyList
}
func metaPropertiesToListables(metaProperties ...types.MetaProperty) []traits.Listable {
	listables := make([]traits.Listable, len(metaProperties))
	for i, property := range metaProperties {
		listables[i] = property
	}
	return listables
}
func NewMetaProperties(metaProperties ...types.MetaProperty) lists.MetaPropertyList {
	return metaPropertyList{List: baseTypes.NewList(metaPropertiesToListables(metaProperties...)...)}
}

func ReadMetaProperties(metaPropertiesString string) (lists.MetaPropertyList, error) {
	var metaPropertyList []types.MetaProperty

	metaProperties := strings.Split(metaPropertiesString, constants.PropertiesSeparator)
	for _, metaPropertyString := range metaProperties {
		if metaPropertyString != "" {
			metaProperty, err := baseTypes.ReadMetaProperty(metaPropertyString)
			if err != nil {
				return nil, err
			}

			metaPropertyList = append(metaPropertyList, metaProperty)
		}
	}

	return NewMetaProperties(metaPropertyList...), nil
}

// ReadData
// CHECK-TODO if data type added see if added here
func ReadData(dataString string) (types.Data, error) {
	dataTypeAndString := strings.SplitN(dataString, constants.DataTypeAndValueSeparator, 2)
	if len(dataTypeAndString) == 2 {
		dataTypeID, dataString := dataTypeAndString[0], dataTypeAndString[1]

		var data types.Data

		var Error error

		switch baseIDs.NewID(dataTypeID) {
		case idsConstants.AccAddressDataID:
			data, Error = base.ReadAccAddressData(dataString)
		case idsConstants.BooleanDataID:
			data, Error = base.ReadBooleanData(dataString)
		case idsConstants.DecDataID:
			data, Error = base.ReadDecData(dataString)
		case idsConstants.HeightDataID:
			data, Error = base.ReadHeightData(dataString)
		case idsConstants.IDDataID:
			data, Error = base.ReadIDData(dataString)
		case idsConstants.ListDataID:
			data, Error = base.ReadListData(dataString)
		case idsConstants.StringDataID:
			data, Error = base.ReadStringData(dataString)
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
