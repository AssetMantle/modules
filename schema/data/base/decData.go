// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/types"
	"github.com/AssetMantle/modules/schema/types/base"
	"github.com/AssetMantle/modules/utilities/meta"
)

type decData struct {
	Value sdkTypes.Dec `json:"value"`
}

var _ types.Data = (*decData)(nil)

func (decData decData) GetID() types.ID {
	return base.NewDataID(decData)
}
func (decData decData) Compare(data types.Data) int {
	compareDecData, err := decDataFromInterface(data)
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
func (decData decData) GetTypeID() types.ID {
	return DecDataID
}
func (decData decData) ZeroValue() types.Data {
	return NewDecData(sdkTypes.ZeroDec())
}
func (decData decData) GenerateHashID() types.ID {
	if decData.Compare(decData.ZeroValue()) == 0 {
		return base.NewID("")
	}

	return base.NewID(meta.Hash(decData.Value.String()))
}
func (decData decData) AsAccAddress() (sdkTypes.AccAddress, error) {
	zeroValue, _ := accAddressData{}.ZeroValue().AsAccAddress()
	return zeroValue, errors.IncorrectFormat
}
func (decData decData) AsDataList() (lists.DataList, error) {
	zeroValue, _ := listData{}.ZeroValue().AsDataList()
	return zeroValue, errors.IncorrectFormat
}
func (decData decData) AsString() (string, error) {
	zeroValue, _ := stringData{}.ZeroValue().AsString()
	return zeroValue, errors.IncorrectFormat
}
func (decData decData) AsDec() (sdkTypes.Dec, error) {
	return decData.Value, nil
}
func (decData decData) AsHeight() (types.Height, error) {
	zeroValue, _ := heightData{}.ZeroValue().AsHeight()
	return zeroValue, errors.IncorrectFormat
}
func (decData decData) AsID() (types.ID, error) {
	zeroValue, _ := idData{}.ZeroValue().AsID()
	return zeroValue, errors.IncorrectFormat
}
func (decData decData) Get() interface{} {
	return decData.Value
}
func decDataFromInterface(data types.Data) (decData, error) {
	switch value := data.(type) {
	case decData:
		return value, nil
	default:
		return decData{}, errors.MetaDataError
	}
}

func NewDecData(value sdkTypes.Dec) types.Data {
	return decData{
		Value: value,
	}
}

func ReadDecData(dataString string) (types.Data, error) {
	if dataString == "" {
		return decData{}.ZeroValue(), nil
	}

	dec, err := sdkTypes.NewDecFromStr(dataString)
	if err != nil {
		return decData{}.ZeroValue(), err
	}

	return NewDecData(dec), nil
}
