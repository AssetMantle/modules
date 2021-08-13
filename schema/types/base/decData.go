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

var _ types.Data = (*DecData)(nil)

func (decData DecData) Compare(data types.Data) int {
	compareDecData, Error := decDataFromInterface(data)
	if Error != nil {
		panic(Error)
	}

	if decData.Value.Dec.GT(compareDecData.Value.Dec) {
		return 1
	} else if decData.Value.Dec.LT(compareDecData.Value.Dec) {
		return -1
	}

	return 0
}
func (decData DecData) String() string {
	return decData.Value.Dec.String()
}
func (decData DecData) GetTypeID() types.ID {
	return NewID("D")
}
func (decData DecData) ZeroValue() types.Data {
	return NewDecData(sdkTypes.ZeroDec())
}
func (decData DecData) GenerateHashID() types.ID {
	if decData.Compare(decData.ZeroValue()) == 0 {
		return NewID("")
	}

	return NewID(meta.Hash(decData.Value.Dec.String()))
}
func (decData DecData) AsAccAddress() (sdkTypes.AccAddress, error) {
	zeroValue, _ := AccAddressData{}.ZeroValue().AsAccAddress()
	return zeroValue, errors.IncorrectFormat
}
func (decData DecData) AsListData() (types.ListData, error) {
	zeroValue, _ := ListData{}.ZeroValue().AsListData()
	return zeroValue, errors.IncorrectFormat
}
func (decData DecData) AsString() (string, error) {
	zeroValue, _ := StringData{}.ZeroValue().AsString()
	return zeroValue, errors.IncorrectFormat
}
func (decData DecData) AsDec() (sdkTypes.Dec, error) {
	return decData.Value.Dec, nil
}
func (decData DecData) AsHeight() (types.Height, error) {
	zeroValue, _ := HeightData{}.ZeroValue().AsHeight()
	return zeroValue, errors.IncorrectFormat
}
func (decData DecData) AsID() (types.ID, error) {
	zeroValue, _ := IDData{}.ZeroValue().AsID()
	return zeroValue, errors.IncorrectFormat
}
func (decData DecData) Get() interface{} {
	return decData.Value.Dec
}
func decDataFromInterface(data types.Data) (DecData, error) {
	switch value := data.(type) {
	case *DecData:
		return *value, nil
	default:
		return DecData{}, errors.MetaDataError
	}
}

func NewDecData(value sdkTypes.Dec) types.Data {
	return &DecData{
		Value: sdkTypes.DecProto{Dec: value},
	}
}

func ReadDecData(dataString string) (types.Data, error) {
	if dataString == "" {
		return DecData{}.ZeroValue(), nil
	}

	dec, Error := sdkTypes.NewDecFromStr(dataString)
	if Error != nil {
		return DecData{}.ZeroValue(), Error
	}

	return NewDecData(dec), nil
}
