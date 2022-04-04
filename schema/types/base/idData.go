// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/types"
	"github.com/AssetMantle/modules/utilities/meta"
)

type idData struct {
	Value types.ID `json:"value"`
}

var _ types.Data = (*idData)(nil)

func (idData idData) GetID() types.ID {
	return dataID{
		TypeID: idData.GetTypeID(),
		HashID: idData.GenerateHashID(),
	}
}
func (idData idData) Compare(data types.Data) int {
	compareIDData, err := idDataFromInterface(data)
	if err != nil {
		panic(err)
	}

	return bytes.Compare(idData.Value.Bytes(), compareIDData.Value.Bytes())
}
func (idData idData) String() string {
	return idData.Value.String()
}
func (idData idData) ZeroValue() types.Data {
	return NewIDData(NewID(""))
}
func (idData idData) GetTypeID() types.ID {
	return idDataID
}
func (idData idData) GenerateHashID() types.ID {
	return NewID(meta.Hash(idData.Value.String()))
}
func (idData idData) AsAccAddress() (sdkTypes.AccAddress, error) {
	zeroValue, _ := accAddressData{}.ZeroValue().AsAccAddress()
	return zeroValue, errors.EntityNotFound
}
func (idData idData) AsDataList() (lists.DataList, error) {
	zeroValue, _ := listData{}.ZeroValue().AsDataList()
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
