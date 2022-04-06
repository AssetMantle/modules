// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"strconv"

	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/types"
	"github.com/AssetMantle/modules/schema/types/base"
)

type booleanData struct {
	Value bool `json:"value"`
}

var _ data.BooleanData = (*booleanData)(nil)

func (booleanData booleanData) GetID() types.ID {
	return base.NewDataID(booleanData)
}
func (booleanData booleanData) Compare(data types.Data) int {
	compareBooleanData, Error := booleanDataFromInterface(data)
	if Error != nil {
		panic(Error)
	}

	if booleanData.Value == compareBooleanData.Value {
		return 0
	} else if booleanData.Value == true { //nolint:gosimple
		return 1
	}

	return -1
}
func (booleanData booleanData) String() string {
	return strconv.FormatBool(booleanData.Value)
}
func (booleanData booleanData) GetTypeID() types.ID {
	return BooleanDataID
}
func (booleanData booleanData) ZeroValue() types.Data {
	return NewBooleanData(false)
}
func (booleanData booleanData) GenerateHashID() types.ID {
	if booleanData.Compare(booleanData.ZeroValue()) == 0 {
		return base.NewID(strconv.FormatBool(false))
	}

	return base.NewID(strconv.FormatBool(true))
}
func (booleanData booleanData) Get() bool {
	return booleanData.Value
}

func booleanDataFromInterface(data types.Data) (booleanData, error) {
	switch value := data.(type) {
	case booleanData:
		return value, nil
	default:
		return booleanData{}, errors.MetaDataError
	}
}

func NewBooleanData(value bool) types.Data {
	return booleanData{
		Value: value,
	}
}

func ReadBooleanData(dataString string) (types.Data, error) {
	if dataString == "" {
		return booleanData{}.ZeroValue(), nil
	}

	Bool, Error := strconv.ParseBool(dataString)
	if Error != nil {
		return booleanData{}.ZeroValue(), Error
	}

	return NewBooleanData(Bool), nil
}
