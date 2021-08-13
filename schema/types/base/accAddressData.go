/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"bytes"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/utilities/meta"
)

var _ types.Data = (*AccAddressData)(nil)

func (accAddressData AccAddressData) Compare(sortable types.Data) int {
	compareAccAddressData, Error := accAddressDataFromInterface(sortable)
	if Error != nil {
		panic(Error)
	}

	return bytes.Compare(accAddressData.Value.GetBytes(), compareAccAddressData.Value.GetBytes())
}
func (accAddressData AccAddressData) String() string {
	return accAddressData.Value.String()
}
func (accAddressData AccAddressData) GetTypeID() types.ID {
	return NewID("A")
}
func (accAddressData AccAddressData) ZeroValue() types.Data {
	return NewAccAddressData(sdkTypes.AccAddress{})
}
func (accAddressData AccAddressData) GenerateHashID() types.ID {
	if accAddressData.Compare(accAddressData.ZeroValue()) == 0 {
		return NewID("")
	}

	return NewID(meta.Hash(accAddressData.Value.String()))
}
func (accAddressData AccAddressData) AsAccAddress() (sdkTypes.AccAddress, error) {
	return accAddressData.Value.AsSDKTypesAccAddress(), nil
}
func (accAddressData AccAddressData) AsListData() (types.ListData, error) {
	zeroValue, _ := ListData{}.ZeroValue().AsListData()
	return zeroValue, errors.IncorrectFormat
}
func (accAddressData AccAddressData) AsString() (string, error) {
	zeroValue, _ := StringData{}.ZeroValue().AsString()
	return zeroValue, errors.IncorrectFormat
}
func (accAddressData AccAddressData) AsDec() (sdkTypes.Dec, error) {
	zeroValue, _ := DecData{}.ZeroValue().AsDec()
	return zeroValue, errors.IncorrectFormat
}
func (accAddressData AccAddressData) AsHeight() (types.Height, error) {
	zeroValue, _ := HeightData{}.ZeroValue().AsHeight()
	return zeroValue, errors.IncorrectFormat
}
func (accAddressData AccAddressData) AsID() (types.ID, error) {
	zeroValue, _ := IDData{}.ZeroValue().AsID()
	return zeroValue, errors.IncorrectFormat
}
func (accAddressData AccAddressData) Get() interface{} {
	return accAddressData.Value
}
func accAddressDataFromInterface(data types.Data) (AccAddressData, error) {
	switch value := data.(type) {
	case *AccAddressData:
		return *value, nil
	default:
		return AccAddressData{}, errors.MetaDataError
	}
}

func NewAccAddressData(value sdkTypes.AccAddress) types.Data {
	return &AccAddressData{
		Value: NewAccAddressFromSDKTypesAccAddress(value),
	}
}

func ReadAccAddressData(dataString string) (types.Data, error) {
	if dataString == "" {
		return AccAddressData{}.ZeroValue(), nil
	}

	accAddress, Error := sdkTypes.AccAddressFromBech32(dataString)
	if Error != nil {
		return AccAddressData{}.ZeroValue(), Error
	}

	return NewAccAddressData(accAddress), nil
}
