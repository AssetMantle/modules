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

type idData struct {
	Value types.ID `json:"value"`
}

func (idData idData) MarshalAmino() (string, error) {
	return idData.Value.String(), nil
}

func (idData *idData) UnmarshalAmino(text string) (err error) {
	idData.Value = NewID(text)
	return nil
}

func (idData idData) MarshalJSON() ([]byte, error) {
	return NewCodec().MarshalJSON(idData.Value)
	//return []byte("mi"), nil
}

func (i *idData) UnmarshalJSON(bz []byte) error {
	//var text string
	//
	//err := NewCodec().UnmarshalJSON(bz, &text)
	//idData.Value = NewID(text)
	//return err

	i = &idData{NewID("dummy")}
	return nil
}

var _ types.Data = (*idData)(nil)

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
	if idData.Value.String() == "" {
		return NewID("")
	}

	return NewID(meta.Hash(idData.Value.String()))
}
func (idData idData) AsString() (string, error) {
	return "", errors.EntityNotFound
}
func (idData idData) AsDec() (sdkTypes.Dec, error) {
	return sdkTypes.Dec{}, errors.EntityNotFound
}
func (idData idData) AsHeight() (types.Height, error) {
	return height{}, errors.EntityNotFound
}
func (idData idData) AsID() (types.ID, error) {
	return idData.Value, nil
}
func (idData idData) Get() interface{} {
	return idData.Value
}
func (idData idData) Equal(data types.Data) bool {
	compareIDData, Error := idDataFromInterface(data)
	if Error != nil {
		return false
	}

	return idData.Value.Equals(compareIDData.Value)
}
func idDataFromInterface(data types.Data) (*idData, error) {
	switch value := data.(type) {
	case *idData:
		return value, nil
	default:
		return &idData{}, errors.MetaDataError
	}
}

func NewIDData(value types.ID) types.Data {
	return &idData{
		Value: value,
	}
}

func ReadIDData(idData string) (types.Data, error) {
	return NewIDData(NewID(idData)), nil
}
