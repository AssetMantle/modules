// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"strings"

	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/types"
	"github.com/AssetMantle/modules/schema/types/base"
	"github.com/AssetMantle/modules/utilities/meta"
)

type stringData struct {
	Value string `json:"value"`
}

var _ data.StringData = (*stringData)(nil)

func (stringData stringData) GetID() types.ID {
	return base.NewDataID(stringData)
}
func (stringData stringData) Compare(data types.Data) int {
	compareStringData, err := stringDataFromInterface(data)
	if err != nil {
		panic(err)
	}

	return strings.Compare(stringData.Value, compareStringData.Value)
}
func (stringData stringData) String() string {
	return stringData.Value
}
func (stringData stringData) GetTypeID() types.ID {
	return StringDataID
}
func (stringData stringData) ZeroValue() types.Data {
	return NewStringData("")
}
func (stringData stringData) GenerateHashID() types.ID {
	return base.NewID(meta.Hash(stringData.Value))
}
func (stringData stringData) Get() string {
	return stringData.Value
}

func stringDataFromInterface(data types.Data) (stringData, error) {
	switch value := data.(type) {
	case stringData:
		return value, nil
	default:
		return stringData{}, errors.MetaDataError
	}
}

func NewStringData(value string) types.Data {
	return stringData{
		Value: value,
	}
}

func ReadStringData(stringData string) (types.Data, error) {
	return NewStringData(stringData), nil
}
