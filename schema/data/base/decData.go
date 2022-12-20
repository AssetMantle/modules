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

var _ data.DecData = (*Data_DecData)(nil)

func (decData *Data_DecData) Unmarshal(bytes []byte) error {
	// TODO implement me
	panic("implement me")
}
func (decData *Data_DecData) GetID() ids.ID {
	return baseIDs.GenerateDataID(decData)
}
func (decData *Data_DecData) Compare(listable traits.Listable) int {
	compareDecData, err := dataFromInterface(listable)
	if err != nil {
		panic(err)
	}

	return bytes.Compare(decData.Bytes(), compareDecData.Bytes())
}
func (decData *Data_DecData) String() string {
	return decData.DecData.Value.String()
}
func (decData *Data_DecData) Bytes() []byte {
	return sdkTypes.SortableDecBytes(decData.DecData.Value)
}
func (decData *Data_DecData) GetType() ids.ID {
	return dataConstants.DecDataID
}
func (decData *Data_DecData) ZeroValue() data.Data {
	return NewDecData(sdkTypes.ZeroDec())
}
func (decData *Data_DecData) GenerateHashID() ids.ID {
	if decData.Compare(decData.ZeroValue()) == 0 {
		return baseIDs.GenerateHashID()
	}

	return baseIDs.GenerateHashID(decData.Bytes())
}
func (decData *Data_DecData) Get() sdkTypes.Dec {
	return decData.DecData.Value
}

func DecDataPrototype() data.Data {
	return NewDecData(sdkTypes.ZeroDec()).ZeroValue().(data.DecData)
}

func NewDecData(value sdkTypes.Dec) data.Data {
	return &Data{
		Impl: &Data_DecData{
			DecData: &DecData{
				Value: value,
			},
		},
	}
}
