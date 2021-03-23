/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"strings"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/utilities/meta"
)

type accAddressListData struct {
	Value sortedAccAddresses `json:"value"`
}

var _ types.ListData = (*accAddressListData)(nil)

func (accAddressListData accAddressListData) String() string {
	accAddressDataStringList := make([]string, len(accAddressListData.Value))

	for i, accAddress := range accAddressListData.Value {
		accAddressDataStringList[i] = accAddress.String()
	}

	return strings.Join(accAddressDataStringList, constants.ListDataStringSeparator)
}
func (accAddressListData accAddressListData) GetTypeID() types.ID {
	return NewID("AL")
}
func (accAddressListData accAddressListData) ZeroValue() types.Data {
	return NewAccAddressListData([]sdkTypes.AccAddress{}...)
}
func (accAddressListData accAddressListData) GenerateHashID() types.ID {
	if len(accAddressListData.Value) == 0 {
		return NewID("")
	}

	return NewID(meta.Hash(accAddressListData.String()))
}
func (accAddressListData accAddressListData) AsAccAddress() (sdkTypes.AccAddress, error) {
	zeroValue, _ := accAddressData{}.ZeroValue().AsAccAddress()
	return zeroValue, errors.IncorrectFormat
}
func (accAddressListData accAddressListData) AsAccAddressList() ([]sdkTypes.AccAddress, error) {
	return accAddressListData.Value, nil
}
func (accAddressListData accAddressListData) AsString() (string, error) {
	zeroValue, _ := stringData{}.ZeroValue().AsString()
	return zeroValue, errors.IncorrectFormat
}
func (accAddressListData accAddressListData) AsDec() (sdkTypes.Dec, error) {
	zeroValue, _ := decData{}.ZeroValue().AsDec()
	return zeroValue, errors.IncorrectFormat
}
func (accAddressListData accAddressListData) AsHeight() (types.Height, error) {
	zeroValue, _ := heightData{}.ZeroValue().AsHeight()
	return zeroValue, errors.IncorrectFormat
}
func (accAddressListData accAddressListData) AsID() (types.ID, error) {
	zeroValue, _ := idData{}.ZeroValue().AsID()
	return zeroValue, errors.IncorrectFormat
}
func (accAddressListData accAddressListData) Get() interface{} {
	return accAddressListData.Value
}
func (accAddressListData accAddressListData) Equal(data types.Data) bool {
	compareAccAddressListData, Error := accAddressListDataFromInterface(data)
	if Error != nil {
		return false
	}

	if len(accAddressListData.Value) != len(compareAccAddressListData.Value) {
		return false
	}

	return accAddressListData.GenerateHashID().Equals(compareAccAddressListData.GenerateHashID())
}
func (accAddressListData accAddressListData) Add(dataList ...types.Data) types.ListData {
	for _, data := range dataList {
		accAddressData, Error := accAddressDataFromInterface(data)
		if Error != nil {
			panic(Error)
		}

		accAddressListData.Value = accAddressListData.Value.Insert(accAddressData.Value).(sortedAccAddresses)
	}

	return accAddressListData
}
func (accAddressListData accAddressListData) Remove(dataList ...types.Data) types.ListData {
	for _, data := range dataList {
		accAddressData, Error := accAddressDataFromInterface(data)
		if Error != nil {
			panic(Error)
		}

		accAddressListData.Value = accAddressListData.Value.Delete(accAddressData.Value).(sortedAccAddresses)
	}

	return accAddressListData
}
func (accAddressListData accAddressListData) IsPresent(data types.Data) bool {
	accAddressData, Error := accAddressDataFromInterface(data)
	if Error != nil {
		panic(Error)
	}

	return accAddressListData.Value.Search(accAddressData.Value) != accAddressListData.Value.Len()
}
func accAddressListDataFromInterface(data types.Data) (accAddressListData, error) {
	switch value := data.(type) {
	case accAddressListData:
		return value, nil
	default:
		return accAddressListData{}, errors.MetaDataError
	}
}

func NewAccAddressListData(value ...sdkTypes.AccAddress) types.Data {
	return accAddressListData{
		Value: sortedAccAddresses(value).Sort().(sortedAccAddresses),
	}
}

func ReadAccAddressListData(dataString string) (types.Data, error) {
	if dataString == "" {
		return accAddressListData{}.ZeroValue(), nil
	}

	dataStringList := strings.Split(dataString, constants.ListDataStringSeparator)
	accAddressList := make([]sdkTypes.AccAddress, len(dataStringList))

	for i, accAddressString := range dataStringList {
		accAddress, Error := sdkTypes.AccAddressFromBech32(accAddressString)
		if Error != nil {
			return accAddressListData{}.ZeroValue(), Error
		}

		accAddressList[i] = accAddress
	}

	return NewAccAddressListData(accAddressList...), nil
}
