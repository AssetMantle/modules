// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/data"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ data.DecData = (*DecData)(nil)

func (decData *DecData) GetID() ids.DataID {
	return baseIDs.GenerateDataID(decData)
}
func (decData *DecData) GetBondWeight() sdkTypes.Dec {
	return dataConstants.DecDataWeight
}
func (decData *DecData) Compare(listable traits.Listable) int {
	compareDecData, err := dataFromListable(listable)
	if err != nil {
		panic(err)
	}

	return bytes.Compare(decData.Bytes(), compareDecData.Bytes())
}
func (decData *DecData) Bytes() []byte {
	return sdkTypes.SortableDecBytes(decData.Value)
}
func (decData *DecData) GetType() ids.StringID {
	return dataConstants.DecDataID
}
func (decData *DecData) ZeroValue() data.Data {
	return NewDecData(sdkTypes.ZeroDec())
}
func (decData *DecData) GenerateHashID() ids.HashID {
	if decData.Compare(decData.ZeroValue()) == 0 {
		return baseIDs.GenerateHashID()
	}

	return baseIDs.GenerateHashID(decData.Bytes())
}
func (decData *DecData) AsString() string {
	return joinDataTypeAndValueStrings(decData.GetType().AsString(), decData.Value.String())
}
func (decData *DecData) FromString(dataTypeAndValueString string) (data.Data, error) {
	dataTypeString, dataString := splitDataTypeAndValueStrings(dataTypeAndValueString)

	if dataTypeString != decData.GetType().AsString() {
		return PrototypeDecData(), constants.IncorrectFormat
	}

	if dataString == "" {
		return PrototypeDecData(), nil
	}

	dec, err := sdkTypes.NewDecFromStr(dataString)
	if err != nil {
		return PrototypeDecData(), err
	}

	return NewDecData(dec), nil
}
func (decData *DecData) Get() sdkTypes.Dec {
	return decData.Value
}
func (decData *DecData) ToAnyData() data.AnyData {
	return &AnyData{
		Impl: &AnyData_DecData{
			DecData: decData,
		}}
}

func PrototypeDecData() data.DecData {
	return NewDecData(sdkTypes.ZeroDec()).ZeroValue().(data.DecData)
}

func NewDecData(value sdkTypes.Dec) data.DecData {
	return &DecData{
		Value: value,
	}
}
