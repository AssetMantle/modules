// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"strconv"

	"github.com/AssetMantle/modules/schema/data"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
)

type booleanData struct {
	Value bool `json:"value"`
}

var _ data.BooleanData = (*booleanData)(nil)

func (booleanData booleanData) GetID() ids.DataID {
	return baseIDs.NewDataID(booleanData)
}
func (booleanData booleanData) Compare(listable traits.Listable) int {
	compareBooleanData, err := booleanDataFromInterface(listable)
	if err != nil {
		panic(err)
	}

	if booleanData.Value == compareBooleanData.Value {
		return 0
	} else if booleanData.Value == true {
		return 1
	}

	return -1
}
func (booleanData booleanData) String() string {
	return strconv.FormatBool(booleanData.Value)
}
func (booleanData booleanData) GetType() ids.ID {
	return dataConstants.BooleanDataID
}
func (booleanData booleanData) ZeroValue() data.Data {
	return NewBooleanData(false)
}
func (booleanData booleanData) GenerateHash() ids.ID {
	if booleanData.Compare(booleanData.ZeroValue()) == 0 {
		return baseIDs.NewID(strconv.FormatBool(false))
	}

	return baseIDs.NewID(strconv.FormatBool(true))
}
func (booleanData booleanData) Get() bool {
	return booleanData.Value
}

func booleanDataFromInterface(listable traits.Listable) (booleanData, error) {
	switch value := listable.(type) {
	case booleanData:
		return value, nil
	default:
		return booleanData{}, constants.MetaDataError
	}
}

func NewBooleanData(value bool) data.Data {
	return booleanData{
		Value: value,
	}
}
