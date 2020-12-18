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

type decData struct {
	Value sdkTypes.Dec `json:"value"`
}

var _ types.Data = (*decData)(nil)

func (DecData decData) String() string {
	return DecData.Value.String()
}
func (DecData decData) GetTypeID() types.ID {
	return NewID("D")
}
func (DecData decData) ZeroValue() types.Data {
	return NewDecData(sdkTypes.ZeroDec())
}
func (DecData decData) GenerateHashID() types.ID {
	if DecData.Equal(DecData.ZeroValue()) {
		return NewID("")
	}
	return NewID(meta.Hash(DecData.Value.String()))
}
func (DecData decData) AsString() (string, error) {
	return "", errors.EntityNotFound
}
func (DecData decData) AsDec() (sdkTypes.Dec, error) {
	return DecData.Value, nil
}
func (DecData decData) AsHeight() (types.Height, error) {
	return height{}, errors.EntityNotFound
}
func (DecData decData) AsID() (types.ID, error) {
	return id{}, errors.EntityNotFound
}
func (DecData decData) Get() interface{} {
	return DecData.Value
}
func (DecData decData) Equal(data types.Data) bool {
	switch value := data.(type) {
	case decData:
		return value.Value.Equal(DecData.Value)
	default:
		return false
	}
}

func NewDecData(value sdkTypes.Dec) types.Data {
	return decData{
		Value: value,
	}
}

func ReadDecData(dataString string) (types.Data, error) {
	if dataString == "" {
		return decData{}.ZeroValue(), nil
	}
	dec, Error := sdkTypes.NewDecFromStr(dataString)
	if Error != nil {
		return decData{}.ZeroValue(), Error
	}

	return NewDecData(dec), nil
}
