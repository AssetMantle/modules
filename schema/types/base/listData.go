/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	"strings"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/utilities/meta"
)

var _ types.ListData = (*ListData)(nil)

func (listData ListData) UnpackInterfaces(unpacker codecTypes.AnyUnpacker) error {
	for _, any := range listData.GetValue() {
		var data types.Data
		err := unpacker.UnpackAny(any, &data)
		if err != nil {
			return err
		}
	}
	return nil
}

func listDataFromData(data types.Data) (ListData, error) {
	switch value := data.(type) {
	case *ListData:
		return *value, nil
	default:
		return ListData{}, errors.MetaDataError
	}
}

func NewListData(value ...types.Data) ListData {
	dataAnyList := make([]*codecTypes.Any, len(value))
	for i, data := range value {
		dataAny, err := data.AsAny()
		if err != nil {
			panic(err)
		}
		dataAnyList[i] = dataAny
	}
	return ListData{Value: dataAnyList}
}

func ReadAccAddressListData(dataString string) (types.Data, error) {
	panic("implement me")
}

// TODO: find a better impl for types.DummyListData
func (listData ListData) Compare(data types.Data) int {
	panic("implement me")
}
func (listData ListData) String() string {
	dataStringList := make([]string, len(listData.Value))

	for i, data := range listData.Value {
		dataStringList[i] = data.String()
	}

	return strings.Join(dataStringList, constants.ListDataStringSeparator)
}
func (listData ListData) GetTypeID() types.ID {
	return NewTypeID("LD")
}
func (listData ListData) ZeroValue() types.Data {
	ld := NewListData([]types.Data{}...)
	return &ld
}
func (listData ListData) GenerateHashID() types.ID {
	if len(listData.Value) == 0 {
		return NewTypeID("")
	}

	return NewTypeID(meta.Hash(listData.String()))
}
func (listData ListData) AsAccAddress() (sdkTypes.AccAddress, error) {
	zeroValue, _ := AccAddressData{}.ZeroValue().AsAccAddress()
	return zeroValue, errors.IncorrectFormat
}
func (listData ListData) AsListData() (types.ListData, error) {
	//return &listData, nil
	return nil, nil
}
func (listData ListData) AsString() (string, error) {
	zeroValue, _ := StringData{}.ZeroValue().AsString()
	return zeroValue, errors.IncorrectFormat
}
func (listData ListData) AsDec() (sdkTypes.Dec, error) {
	zeroValue, _ := DecData{}.ZeroValue().AsDec()
	return zeroValue, errors.IncorrectFormat
}
func (listData ListData) AsHeight() (types.Height, error) {
	zeroValue, _ := HeightData{}.ZeroValue().AsHeight()
	return zeroValue, errors.IncorrectFormat
}
func (listData ListData) AsID() (types.ID, error) {
	zeroValue, _ := IDData{}.ZeroValue().AsID()
	return zeroValue, errors.IncorrectFormat
}
func (listData ListData) Get() interface{} {
	return listData.Value
}
func (listData *ListData) AsAny() (*codecTypes.Any, error) {
	return codecTypes.NewAnyWithValue(listData)
}
