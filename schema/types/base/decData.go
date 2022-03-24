/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/utilities/meta"
)

type decData struct {
	Value sdkTypes.Dec `json:"value"`
}

var _ types.Data = (*decData)(nil)

func (decData decData) GetID() types.ID {
	return dataID{
		TypeID: decData.GetTypeID(),
		HashID: decData.GenerateHashID(),
	}
}
func (decData decData) Compare(data types.Data) int {
	compareDecData, Error := decDataFromInterface(data)
	if Error != nil {
		panic(Error)
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
	return decDataID
}
func (decData decData) ZeroValue() types.Data {
	return NewDecData(sdkTypes.ZeroDec())
}
func (decData decData) GenerateHashID() types.ID {
	if decData.Compare(decData.ZeroValue()) == 0 {
		return NewID("")
	}

	return NewID(meta.Hash(decData.Value.String()))
}
func (decData decData) AsAccAddress() (sdkTypes.AccAddress, error) {
	zeroValue, _ := accAddressData{}.ZeroValue().AsAccAddress()
	return zeroValue, errors.IncorrectFormat
}
func (decData decData) AsListData() (types.ListData, error) {
	zeroValue, _ := listData{}.ZeroValue().AsListData()
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

	dec, Error := sdkTypes.NewDecFromStr(dataString)
	if Error != nil {
		return decData{}.ZeroValue(), Error
	}

	return NewDecData(dec), nil
}
