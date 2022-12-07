// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
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
func (decData *DecData) Compare(listable traits.Listable) int {
	compareDecData, err := decDataFromInterface(listable)
	if err != nil {
		panic(err)
	}

	if decData.Get().GT(compareDecData.Get()) {
		return 1
	} else if decData.Get().LT(compareDecData.Get()) {
		return -1
	}

	return 0
}
func (decData *DecData) Bytes() []byte {
	dec, _ := sdkTypes.NewDecFromStr(decData.Value)

	return sdkTypes.SortableDecBytes(dec)
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
	dec, _ := sdkTypes.NewDecFromStr(decData.Value)
	return dec
}

func decDataFromInterface(listable traits.Listable) (*DecDataI, error) {
	switch value := listable.(type) {
	case *DecDataI:
		return value, nil
	default:
		return nil, constants.MetaDataError
	}
}

func DecDataPrototype() data.DecData {
	return NewDecData(sdkTypes.ZeroDec()).ZeroValue().(data.DecData)
}

func NewDecData(value sdkTypes.Dec) data.DecData {
	return &DecData{
		Value: value,
	}
}
