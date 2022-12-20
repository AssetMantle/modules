// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/data"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ data.DecData = (*DecData)(nil)

func (decData *DecData) GetID() ids.DataID {
	return baseIDs.GenerateDataID(decData)
}
func (decData *DecData) Compare(listable traits.Listable) int {
	compareDecData, err := dataFromInterface(listable)
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
func (decData *DecData) Get() sdkTypes.Dec {
	return decData.Value
}
func (decData *DecData) ToAnyData() data.AnyData {
	return &AnyData{
		Impl: &AnyData_DecData{
			DecData: decData,
		}}
}

func DecDataPrototype() data.Data {
	return NewDecData(sdkTypes.ZeroDec()).ZeroValue().(data.DecData)
}

func NewDecData(value sdkTypes.Dec) data.Data {
	return &AnyData{
		Impl: &AnyData_DecData{
			DecData: &DecData{
				Value: value,
			},
		},
	}
}
