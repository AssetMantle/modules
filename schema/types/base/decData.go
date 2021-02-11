/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"encoding/json"
	"math/big"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/utilities/meta"
)

var (
	nilAmino string
	nilJSON  []byte
)

func init() {
	empty := new(big.Int)
	bz, err := empty.MarshalText()
	if err != nil {
		panic("bad nil amino init")
	}
	nilAmino = string(bz)

	nilJSON, err = json.Marshal(string(bz))
	if err != nil {
		panic("bad nil json init")
	}
}

type decData struct {
	Value sdkTypes.Dec `json:"value"`
}

func (decData decData) MarshalAmino() (string, error) {
	if decData.Value.Int == nil {
		return nilAmino, nil
	}
	bz, err := decData.Value.Int.MarshalText()
	return string(bz), err
}

func (decData *decData) UnmarshalAmino(text string) (err error) {
	tempInt := new(big.Int)
	err = tempInt.UnmarshalText([]byte(text))
	if err != nil {
		return err
	}
	decData.Value.Int = tempInt
	return nil
}

func (decData decData) MarshalJSON() ([]byte, error) {
	//if decData.Value.Int == nil {
	//	return nilJSON, nil
	//}
	//
	//return NewCodec().MarshalJSON(decData.Value.String())

	return []byte("1"), nil
}

func (d *decData) UnmarshalJSON(bz []byte) error {
	//if d.Value.Int == nil {
	//	d.Value.Int = new(big.Int)
	//}
	//
	//var text string
	//err := NewCodec().UnmarshalJSON(bz, &text)
	//if err != nil {
	//	return err
	//}
	//// TODO: Reuse dec allocation
	//newDec, err := sdkTypes.NewDecFromStr(text)
	//if err != nil {
	//	return err
	//}
	//d.Value.Int = newDec.Int

	d = &decData{sdkTypes.NewDec(1)}
	return nil
}

var _ types.Data = (*decData)(nil)

func (decData decData) String() string {
	return decData.Value.String()
}
func (decData decData) GetTypeID() types.ID {
	return NewID("D")
}
func (decData decData) ZeroValue() types.Data {
	return NewDecData(sdkTypes.ZeroDec())
}
func (decData decData) GenerateHashID() types.ID {
	if decData.Equal(decData.ZeroValue()) {
		return NewID("")
	}

	return NewID(meta.Hash(decData.Value.String()))
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
func (decData decData) Equal(data types.Data) bool {
	compareDecData, Error := decDataFromInterface(data)
	if Error != nil {
		return false
	}

	return decData.Value.Equal(compareDecData.Value)
}
func decDataFromInterface(data types.Data) (*decData, error) {
	switch value := data.(type) {
	case *decData:
		return value, nil
	default:
		return &decData{}, errors.MetaDataError
	}
}

func NewDecData(value sdkTypes.Dec) types.Data {
	return &decData{
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
