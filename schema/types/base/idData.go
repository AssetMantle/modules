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

func (idData idData) GenerateHash() string {
	if idData.Value.String() == "" {
		return ""
	}
	return meta.Hash(idData.Value.String())
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
	return idData.Value, errors.EntityNotFound
}

func (idData idData) Get() interface{} {
	return idData.Value
}

func NewIDData(value types.ID) types.Data {
	return idData{
		Value: value,
	}
}

func ReadIDData(idData string) types.Data {
	return NewIDData(NewID(idData))
}
