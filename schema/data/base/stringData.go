// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"strings"

	"github.com/AssetMantle/modules/constants/errors"
	idsConstants "github.com/AssetMantle/modules/constants/ids"
	"github.com/AssetMantle/modules/schema/capabilities"
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/types"
	"github.com/AssetMantle/modules/utilities/meta"
)

type stringData struct {
	Value string `json:"value"`
}

var _ data.StringData = (*stringData)(nil)

func (stringData stringData) GetID() ids.DataID {
	return baseIDs.NewDataID(stringData)
}
func (stringData stringData) Compare(listable capabilities.Listable) int {
	compareStringData, err := stringDataFromInterface(listable)
	if err != nil {
		panic(err)
	}

	return strings.Compare(stringData.Value, compareStringData.Value)
}
func (stringData stringData) String() string {
	return stringData.Value
}
func (stringData stringData) GetType() types.ID {
	return idsConstants.StringDataID
}
func (stringData stringData) ZeroValue() types.Data {
	return NewStringData("")
}
func (stringData stringData) GenerateHash() types.ID {
	return baseIDs.NewID(meta.Hash(stringData.Value))
}
func (stringData stringData) Get() string {
	return stringData.Value
}

func stringDataFromInterface(listable capabilities.Listable) (stringData, error) {
	switch value := listable.(type) {
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
