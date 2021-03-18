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

type accAddressData struct {
	Value sdkTypes.AccAddress `json:"value"`
}

var _ types.Data = (*accAddressData)(nil)

func (accAddressData accAddressData) String() string {
	return accAddressData.Value.String()
}
func (accAddressData accAddressData) GetTypeID() types.ID {
	return NewID("A")
}
func (accAddressData accAddressData) ZeroValue() types.Data {
	return NewAccAddressData(sdkTypes.AccAddress{})
}
func (accAddressData accAddressData) GenerateHashID() types.ID {
	if accAddressData.Equal(accAddressData.ZeroValue()) {
		return NewID("")
	}

	return NewID(meta.Hash(accAddressData.Value.String()))
}
func (accAddressData accAddressData) AsAccAddressData() (sdkTypes.AccAddress, error) {
	return accAddressData.Value, nil
}
func (accAddressData accAddressData) AsString() (string, error) {
	zeroValue, _ := stringData{}.ZeroValue().AsString()
	return zeroValue, errors.IncorrectFormat
}
func (accAddressData accAddressData) AsDec() (sdkTypes.Dec, error) {
	zeroValue, _ := decData{}.ZeroValue().AsDec()
	return zeroValue, errors.IncorrectFormat
}
func (accAddressData accAddressData) AsHeight() (types.Height, error) {
	zeroValue, _ := heightData{}.ZeroValue().AsHeight()
	return zeroValue, errors.IncorrectFormat
}
func (accAddressData accAddressData) AsID() (types.ID, error) {
	zeroValue, _ := idData{}.ZeroValue().AsID()
	return zeroValue, errors.IncorrectFormat
}
func (accAddressData accAddressData) Get() interface{} {
	return accAddressData.Value
}
func (accAddressData accAddressData) Equal(data types.Data) bool {
	compareAccAddressData, Error := accAddressDataFromInterface(data)
	if Error != nil {
		return false
	}

	return accAddressData.Value.Equals(compareAccAddressData.Value)
}
func accAddressDataFromInterface(data types.Data) (accAddressData, error) {
	switch value := data.(type) {
	case accAddressData:
		return value, nil
	default:
		return accAddressData{}, errors.MetaDataError
	}
}

func NewAccAddressData(value sdkTypes.AccAddress) types.Data {
	return accAddressData{
		Value: value,
	}
}

func ReadAccAddressData(dataString string) (types.Data, error) {
	if dataString == "" {
		return accAddressData{}.ZeroValue(), nil
	}

	accAddress, Error := sdkTypes.AccAddressFromBech32(dataString)
	if Error != nil {
		return accAddressData{}.ZeroValue(), Error
	}

	return NewAccAddressData(accAddress), nil
}
