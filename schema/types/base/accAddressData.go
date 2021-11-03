/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"bytes"
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/utilities/meta"
)

var _ types.Data = (*AccAddressData)(nil)

func accAddressDataFromInterface(data types.Data) (AccAddressData, error) {
	switch value := data.(type) {
	case *AccAddressData:
		return *value, nil
	default:
		return AccAddressData{}, errors.MetaDataError
	}
}

func NewAccAddressData(address sdkTypes.AccAddress) types.Data {
	return &AccAddressData{
		Value: address.Bytes(),
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

func (accAddressData AccAddressData) Compare(compareData types.Data) int {
	compareAccAddressData, err := compareData.AsAccAddress()
	if err != nil {
		panic(err)
	}
	return bytes.Compare(accAddressData.Value, compareAccAddressData.Bytes())
}
func (accAddressData AccAddressData) String() string {
	return sdkTypes.AccAddress(accAddressData.Value).String()
}
func (accAddressData AccAddressData) GetTypeID() types.ID {
	id := NewID("A")
	return &id
}
func (accAddressData AccAddressData) ZeroValue() types.Data {
	return NewAccAddressData(sdkTypes.AccAddress{})
}
func (accAddressData *AccAddressData) AsAny() (*codecTypes.Any, error) {
	return codecTypes.NewAnyWithValue(accAddressData)
}
func (accAddressData AccAddressData) GenerateHashID() types.ID {
	if accAddressData.Compare(accAddressData.ZeroValue()) == 0 {
		id := NewID("")
		return &id
	}

	id := NewID(meta.Hash(accAddressData.String()))
	return &id
}
func (accAddressData AccAddressData) AsAccAddress() (sdkTypes.AccAddress, error) {
	return accAddressData.Value, nil
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
