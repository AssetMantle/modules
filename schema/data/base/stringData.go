// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"strings"

	"github.com/AssetMantle/modules/schema/data"
	idsConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
	stringUtilities "github.com/AssetMantle/modules/utilities/string"
)

type stringData struct {
	Value string `json:"value"`
}

var _ data.StringData = (*stringData)(nil)

func (stringData stringData) GetID() ids.DataID {
	return baseIDs.NewDataID(stringData)
}
func (stringData stringData) Compare(listable traits.Listable) int {
	compareStringData, err := stringDataFromInterface(listable)
	if err != nil {
		panic(err)
	}

	return strings.Compare(stringData.Value, compareStringData.Value)
}
func (stringData stringData) String() string {
	return stringData.Value
}
func (stringData stringData) GetType() ids.ID {
	return idsConstants.StringDataID
}
func (stringData stringData) ZeroValue() data.Data {
	return NewStringData("")
}
func (stringData stringData) GenerateHash() ids.ID {
	return baseIDs.NewID(stringUtilities.Hash(stringData.Value))
}
func (stringData stringData) Get() string {
	return stringData.Value
}

func stringDataFromInterface(listable traits.Listable) (stringData, error) {
	switch value := listable.(type) {
	case stringData:
		return value, nil
	default:
		return stringData{}, constants.MetaDataError
	}
}

func NewStringData(value string) data.Data {
	return stringData{
		Value: value,
	}
}

func ReadStringData(stringData string) (data.Data, error) {
	return NewStringData(stringData), nil
}
