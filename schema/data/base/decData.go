// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/schema/data"
	idsConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
	stringUtilities "github.com/AssetMantle/modules/utilities/string"
)

type decData struct {
	Value sdkTypes.Dec `json:"value"`
}

var _ data.DecData = (*decData)(nil)

func (decData decData) GetID() ids.DataID {
	return baseIDs.NewDataID(decData)
}
func (decData decData) Compare(listable traits.Listable) int {
	compareDecData, err := decDataFromInterface(listable)
	if err != nil {
		panic(err)
	}

	if decData.Value.GT(compareDecData.Value) {
		return 1
	} else if decData.Value.LT(compareDecData.Value) {
		return -1
	}

	return 0
}
func (decData decData) String() string {
	return decData.Value.String()
}
func (decData decData) GetType() ids.ID {
	return idsConstants.DecDataID
}
func (decData decData) ZeroValue() data.Data {
	return NewDecData(sdkTypes.ZeroDec())
}
func (decData decData) GenerateHash() ids.ID {
	if decData.Compare(decData.ZeroValue()) == 0 {
		return baseIDs.NewID("")
	}

	return baseIDs.NewID(stringUtilities.Hash(decData.Value.String()))
}
func (decData decData) Get() sdkTypes.Dec {
	return decData.Value
}

func decDataFromInterface(listable traits.Listable) (decData, error) {
	switch value := listable.(type) {
	case decData:
		return value, nil
	default:
		return decData{}, errors.MetaDataError
	}
}

func NewDecData(value sdkTypes.Dec) data.Data {
	return decData{
		Value: value,
	}
}

func ReadDecData(dataString string) (data.Data, error) {
	if dataString == "" {
		return decData{}.ZeroValue(), nil
	}

	dec, err := sdkTypes.NewDecFromStr(dataString)
	if err != nil {
		return decData{}.ZeroValue(), err
	}

	return NewDecData(dec), nil
}
