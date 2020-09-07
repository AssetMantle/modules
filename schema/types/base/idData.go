/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"

	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/utilities/meta"
)

const IDType = "I"

type idData struct {
	Value types.ID `json:"value"`
}

var _ types.Data = (*idData)(nil)

func (IDData idData) String() string {
	return IDData.Value.String()
}

func (IDData idData) GenerateHash() string {
	if IDData.Value.String() == "" {
		return ""
	}
	return meta.Hash(IDData.Value.String())
}

func (IDData idData) AsString() (string, error) {
	return "", errors.EntityNotFound
}

func (IDData idData) AsDec() (sdkTypes.Dec, error) {
	return sdkTypes.Dec{}, errors.EntityNotFound
}

func (IDData idData) AsHeight() (types.Height, error) {
	return height{}, errors.EntityNotFound
}

func (IDData idData) AsID() (types.ID, error) {
	return IDData.Value, errors.EntityNotFound
}

func (IDData idData) Get() interface{} {
	return IDData.Value
}

func (IDData idData) Equal(data types.Data) bool {
	switch value := data.(type) {
	case decData:
		return value.Equal(IDData)
	default:
		return false
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
