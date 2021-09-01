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

func idDataFromInterface(data types.Data) (Data_IdData, error) {
	switch value := data.(type) {
	case *Data_IdData:
		return *value, nil
	default:
		return Data_IdData{}, errors.MetaDataError
	}
}

func NewIDData(value ID) types.Data {
	return &Data_IdData{
		IdData: &IDData{
			Value: value,
		},
	}
}

func ReadIDData(idData string) (types.Data, error) {
	return NewIDData(*NewID(idData)), nil
}

var _ types.Data = (*Data_IdData)(nil)

func (idData Data_IdData) Compare(data types.Data) int {
	compareIDData, Error := idDataFromInterface(data)
	if Error != nil {
		panic(Error)
	}

	return bytes.Compare(idData.IdData.Value.Bytes(), compareIDData.IdData.Value.Bytes())
}
func (idData Data_IdData) String() string {
	return idData.IdData.Value.String()
}
func (idData Data_IdData) ZeroValue() types.Data {
	return NewIDData(*NewID(""))
}
func (idData Data_IdData) GetTypeID() types.ID {
	return NewID("I")
}
func (idData Data_IdData) GenerateHashID() types.ID {
	return NewID(meta.Hash(idData.IdData.Value.String()))
}
func (idData Data_IdData) AsAccAddress() (sdkTypes.AccAddress, error) {
	zeroValue, _ := Data_AccAddressData{}.ZeroValue().AsAccAddress()
	return zeroValue, errors.EntityNotFound
}
func (idData Data_IdData) AsListData() (types.ListData, error) {
	zeroValue, _ := Data_ListData{}.ZeroValue().AsListData()
	return zeroValue, errors.IncorrectFormat
}
func (idData Data_IdData) AsString() (string, error) {
	zeroValue, _ := Data_StringData{}.ZeroValue().AsString()
	return zeroValue, errors.IncorrectFormat
}
func (idData Data_IdData) AsDec() (sdkTypes.Dec, error) {
	zeroValue, _ := Data_DecData{}.ZeroValue().AsDec()
	return zeroValue, errors.IncorrectFormat
}
func (idData Data_IdData) AsHeight() (types.Height, error) {
	zeroValue, _ := Data_HeightData{}.ZeroValue().AsHeight()
	return zeroValue, errors.IncorrectFormat
}
func (idData Data_IdData) AsID() (types.ID, error) {
	return &idData.IdData.Value, nil
}
func (idData Data_IdData) Get() interface{} {
	return idData.IdData.Value
}
func (idData Data_IdData) Unmarshal(dAtA []byte) error {
	return idData.IdData.Unmarshal(dAtA)
}
func (idData *Data_IdData) Reset() { *idData = Data_IdData{} }
func (*Data_IdData) ProtoMessage() {}
