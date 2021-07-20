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

type idData struct {
	Value types.ID `json:"value"`
}

var _ types.Data = (*idData)(nil)

func (idData idData) Compare(data types.Data) int {
	compareIDData, Error := idDataFromInterface(data)
	if Error != nil {
		panic(Error)
	}

	return bytes.Compare(idData.Value.Bytes(), compareIDData.Value.Bytes()) % 2
}
func (idData idData) String() string {
	return idData.Value.String()
}
func (idData idData) ZeroValue() types.Data {
	return NewIDData(NewID(""))
}
func (idData idData) GetTypeID() types.ID {
	return NewID("I")
}
func (idData idData) GenerateHashID() types.ID {
	return NewID(meta.Hash(idData.Value.String()))
}
func (idData idData) AsAccAddress() (sdkTypes.AccAddress, error) {
	zeroValue, _ := accAddressData{}.ZeroValue().AsAccAddress()
	return zeroValue, errors.EntityNotFound
}
func (idData idData) AsListData() (types.ListData, error) {
	zeroValue, _ := listData{}.ZeroValue().AsListData()
	return zeroValue, errors.IncorrectFormat
}
func (idData idData) AsString() (string, error) {
	zeroValue, _ := stringData{}.ZeroValue().AsString()
	return zeroValue, errors.IncorrectFormat
}
func (idData idData) AsDec() (sdkTypes.Dec, error) {
	zeroValue, _ := decData{}.ZeroValue().AsDec()
	return zeroValue, errors.IncorrectFormat
}
func (idData idData) AsHeight() (types.Height, error) {
	zeroValue, _ := heightData{}.ZeroValue().AsHeight()
	return zeroValue, errors.IncorrectFormat
}
func (idData idData) AsID() (types.ID, error) {
	return idData.Value, nil
}
func (idData idData) Get() interface{} {
	return idData.Value
}
func idDataFromInterface(data types.Data) (idData, error) {
	switch value := data.(type) {
	case idData:
		return value, nil
	default:
		return idData{}, errors.MetaDataError
	}
}

func NewIDData(value types.ID) types.Data {
	return idData{
		Value: value,
	}
}

func ReadIDData(idData string) (types.Data, error) {
	return NewIDData(NewID(idData)), nil
}
