/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"sort"
	"strings"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/utilities/meta"
)

type accAddressListData struct {
	Value []sdkTypes.AccAddress `json:"value"`
}

var _ types.Data = (*accAddressListData)(nil)

func (accAddressListData accAddressListData) String() string {
	accAddressDataStringList := make([]string, len(accAddressListData.Value))

	for i, accAddress := range accAddressListData.Value {
		accAddressDataStringList[i] = accAddress.String()
	}

	sort.Strings(accAddressDataStringList)

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
func (accAddressListData accAddressListData) AsAccAddressData() (sdkTypes.AccAddress, error) {
	zeroValue, _ := accAddressData{}.ZeroValue().AsAccAddressData()
	return zeroValue, errors.IncorrectFormat
}
func (accAddressListData accAddressListData) AsAccAddressListData() ([]sdkTypes.AccAddress, error) {
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
		Value: value,
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
