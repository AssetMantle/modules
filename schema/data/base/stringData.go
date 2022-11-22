// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"strings"

	"github.com/AssetMantle/modules/schema/data"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
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
func (stringData stringData) Bytes() []byte {
	return []byte(stringData.Value)
}
func (stringData stringData) GetType() ids.StringID {
	return dataConstants.StringDataID
}
func (stringData stringData) ZeroValue() data.Data {
	return NewStringData("")
}
func (stringData stringData) GenerateHashID() ids.HashID {
	return baseIDs.GenerateHashID(stringData.Bytes())
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

func StringDataPrototype() data.StringData {
	return stringData{}.ZeroValue().(data.StringData)
}

func NewStringData(value string) data.StringData {
	return stringData{
		Value: value,
	}
}
