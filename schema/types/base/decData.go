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

const DecType = "D"

type decData struct {
	Value sdkTypes.Dec `json:"value"`
}

var _ types.Data = (*decData)(nil)

func (decData decData) GenerateHash() string {
	return meta.Hash(decData.Value.String())
}

func (decData decData) AsString() (string, error) {
	return "", errors.EntityNotFound
}

func (decData decData) AsDec() (sdkTypes.Dec, error) {
	return decData.Value, nil
}

func (decData decData) AsHeight() (types.Height, error) {
	return height{}, errors.EntityNotFound
}

func (decData decData) AsID() (types.ID, error) {
	return id{}, errors.EntityNotFound
}

func (decData decData) Get() interface{} {
	return decData.Value
}

func NewDecData(value sdkTypes.Dec) types.Data {
	return decData{
		Value: value,
	}
}

func ReadDecData(dataString string) types.Data {
	dec, Error := sdkTypes.NewDecFromStr(dataString)
	if Error != nil {
		return nil
	}
	return NewDecData(dec)
}
