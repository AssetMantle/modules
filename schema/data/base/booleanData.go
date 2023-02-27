// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"
	"strconv"

	"github.com/AssetMantle/modules/schema/data"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ data.BooleanData = (*BooleanData)(nil)

func (booleanData *BooleanData) GetID() ids.DataID {
	return baseIDs.GenerateDataID(booleanData)
}
func (booleanData *BooleanData) GetBondWeight() int64 {
	return dataConstants.BooleanDataWeight
}
func (booleanData *BooleanData) Compare(listable traits.Listable) int {
	compareBooleanData, err := dataFromListable(listable)
	if err != nil {
		panic(err)
	}

	if value := bytes.Compare(booleanData.Bytes(), compareBooleanData.Bytes()); value == 0 {
		return 0
	} else if value > 0 {
		return 1
	} else {
		return -1
	}
}
func (booleanData *BooleanData) AsString() string {
	return joinDataTypeAndValueStrings(booleanData.GetType().AsString(), strconv.FormatBool(booleanData.Value))
}
func (booleanData *BooleanData) FromString(dataTypeAndValueString string) (data.Data, error) {
	dataTypeString, dataString := splitDataTypeAndValueStrings(dataTypeAndValueString)

	if dataTypeString != booleanData.GetType().AsString() {
		return PrototypeBooleanData(), errorConstants.IncorrectFormat.Wrapf("incorrect format for BooleanData, expected type identifier %s, got %s", booleanData.GetType().AsString(), dataTypeString)
	}

	if dataString == "" {
		return PrototypeBooleanData(), nil
	}

	Bool, err := strconv.ParseBool(dataString)
	if err != nil {
		return PrototypeBooleanData(), err
	}

	return NewBooleanData(Bool), nil
}
func (booleanData *BooleanData) Bytes() []byte {
	if booleanData.Get() {
		return []byte{0x1}
	}
	return []byte{0x0}
}
func (booleanData *BooleanData) GetType() ids.StringID {
	return dataConstants.BooleanDataID
}
func (booleanData *BooleanData) ZeroValue() data.Data {
	return NewBooleanData(false)
}
func (booleanData *BooleanData) GenerateHashID() ids.HashID {
	if booleanData.Compare(booleanData.ZeroValue()) == 0 {
		return baseIDs.GenerateHashID()
	}

	return baseIDs.GenerateHashID(booleanData.Bytes())
}
func (booleanData *BooleanData) Get() bool {
	return booleanData.Value
}
func (booleanData *BooleanData) ToAnyData() data.AnyData {
	return &AnyData{
		Impl: &AnyData_BooleanData{
			BooleanData: booleanData,
		},
	}
}

func PrototypeBooleanData() data.BooleanData {
	return &BooleanData{}
}

func NewBooleanData(value bool) data.BooleanData {
	return &BooleanData{
		Value: value,
	}
}
