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

var _ types.Data = (*IDData)(nil)

func idDataFromInterface(data types.Data) (IDData, error) {
	switch value := data.(type) {
	case *IDData:
		return *value, nil
	default:
		return IDData{}, errors.MetaDataError
	}
}

func NewIDData(value types.ID) types.Data {
	id := NewID(value.String())
	return &IDData{
		Value: id,
	}
}

func ReadIDData(idData string) (types.Data, error) {
	return NewIDData(NewTypeID(idData)), nil
}

func (idData IDData) Compare(data types.Data) int {
	compareIDData, Error := dummyIDDataFromInterface(data)
	if Error != nil {
		panic(Error)
	}

	return bytes.Compare(idData.Value.Bytes(), compareIDData.Value.Bytes())
}
func (idData IDData) String() string {
	return idData.Value.String()
}
func (idData IDData) ZeroValue() types.Data {
	return NewIDData(NewTypeID(""))
}
func (idData IDData) GetTypeID() types.ID {
	return NewTypeID("I")
}
func (idData IDData) GenerateHashID() types.ID {
	return NewTypeID(meta.Hash(idData.Value.String()))
}
func (idData IDData) AsAccAddress() (sdkTypes.AccAddress, error) {
	zeroValue, _ := AccAddressData{}.ZeroValue().AsAccAddress()
	return zeroValue, errors.EntityNotFound
}
func (idData IDData) AsListData() (types.ListData, error) {
	zeroValue, _ := ListData{}.ZeroValue().AsListData()
	return zeroValue, errors.IncorrectFormat
}
func (idData IDData) AsString() (string, error) {
	zeroValue, _ := StringData{}.ZeroValue().AsString()
	return zeroValue, errors.IncorrectFormat
}
func (idData IDData) AsDec() (sdkTypes.Dec, error) {
	zeroValue, _ := DecData{}.ZeroValue().AsDec()
	return zeroValue, errors.IncorrectFormat
}
func (idData IDData) AsHeight() (types.Height, error) {
	zeroValue, _ := HeightData{}.ZeroValue().AsHeight()
	return zeroValue, errors.IncorrectFormat
}
func (idData IDData) AsID() (types.ID, error) {
	return &idData.Value, nil
}
func (idData IDData) Get() interface{} {
	return idData.Value
}
func (idData *IDData) AsAny() (*codecTypes.Any, error) {
	return codecTypes.NewAnyWithValue(idData)
}
func dummyIDDataFromInterface(data types.Data) (IDData, error) {
	switch value := data.(type) {
	case *IDData:
		return *value, nil
	default:
		return IDData{}, errors.MetaDataError
	}
}
