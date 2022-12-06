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

var _ data.DecData = (*DecDataI_DecData)(nil)

func (decData *DecDataI_DecData) GetID() ids.DataID {
	return baseIDs.GenerateDataID(decData)
}
func (decData *DecDataI_DecData) Compare(listable traits.Listable) int {
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
func (decData *DecDataI_DecData) String() string {
	return decData.DecData.Value
}
func (decData *DecDataI_DecData) Bytes() []byte {
	dec, _ := sdkTypes.NewDecFromStr(decData.DecData.Value)

	return sdkTypes.SortableDecBytes(dec)
}
func (decData *DecDataI_DecData) GetType() ids.StringID {
	return dataConstants.DecDataID
}
func (decData *DecDataI_DecData) ZeroValue() data.Data {
	return NewDecData(sdkTypes.ZeroDec())
}
func (decData *DecDataI_DecData) GenerateHashID() ids.HashID {
	if decData.Compare(decData.ZeroValue()) == 0 {
		return baseIDs.GenerateHashID()
	}

	return baseIDs.GenerateHashID(decData.Bytes())
}
func (decData *DecDataI_DecData) Get() sdkTypes.Dec {
	dec, _ := sdkTypes.NewDecFromStr(decData.DecData.Value)
	return dec
}

func decDataFromInterface(listable traits.Listable) (*DecDataI_DecData, error) {
	switch value := listable.(type) {
	case *DecDataI_DecData:
		return value, nil
	default:
		panic(constants.MetaDataError)
	}
}

func DecDataPrototype() data.DecData {
	return (&DecDataI_DecData{}).ZeroValue().(data.DecData)
}

func NewDecData(value sdkTypes.Dec) data.DecData {
	return &DecDataI_DecData{
		DecData: &DecData{
			Value: value.String(),
		},
	}
}
