// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"strconv"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/lists"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type booleanData struct {
	Value bool `json:"value"`
}

var _ types.Data = (*booleanData)(nil)

func (booleanData booleanData) GetID() types.ID {
	return dataID{
		TypeID: booleanData.GetTypeID(),
		HashID: booleanData.GenerateHashID(),
	}
}
func (booleanData booleanData) Compare(data types.Data) int {
	compareBooleanData, Error := booleanDataFromInterface(data)
	if Error != nil {
		panic(Error)
	}

	if booleanData.Value == compareBooleanData.Value {
		return 0
	} else if booleanData.Value == true { //nolint:gosimple
		return 1
	}

	return -1
}
func (booleanData booleanData) String() string {
	return strconv.FormatBool(booleanData.Value)
}
func (booleanData booleanData) GetTypeID() types.ID {
	return booleanDataID
}
func (booleanData booleanData) ZeroValue() types.Data {
	return NewBooleanData(false)
}
func (booleanData booleanData) GenerateHashID() types.ID {
	if booleanData.Compare(booleanData.ZeroValue()) == 0 {
		return NewID(strconv.FormatBool(false))
	}

	return NewID(strconv.FormatBool(true))
}
func (booleanData booleanData) AsAccAddress() (sdkTypes.AccAddress, error) {
	zeroValue, _ := accAddressData{}.ZeroValue().AsAccAddress()
	return zeroValue, errors.IncorrectFormat
}
func (booleanData booleanData) AsDataList() (lists.DataList, error) {
	zeroValue, _ := listData{}.ZeroValue().AsDataList()
	return zeroValue, errors.IncorrectFormat
}
func (booleanData booleanData) AsString() (string, error) {
	zeroValue, _ := stringData{}.ZeroValue().AsString()
	return zeroValue, errors.IncorrectFormat
}
func (booleanData booleanData) AsDec() (sdkTypes.Dec, error) {
	zeroValue, _ := decData{}.ZeroValue().AsDec()
	return zeroValue, errors.IncorrectFormat
}
func (booleanData booleanData) AsHeight() (types.Height, error) {
	zeroValue, _ := heightData{}.ZeroValue().AsHeight()
	return zeroValue, errors.IncorrectFormat
}
func (booleanData booleanData) AsID() (types.ID, error) {
	zeroValue, _ := idData{}.ZeroValue().AsID()
	return zeroValue, errors.IncorrectFormat
}
func (booleanData booleanData) Get() interface{} {
	return booleanData.Value
}
func booleanDataFromInterface(data types.Data) (booleanData, error) {
	switch value := data.(type) {
	case booleanData:
		return value, nil
	default:
		return booleanData{}, errors.MetaDataError
	}
}

func NewBooleanData(value bool) types.Data {
	return booleanData{
		Value: value,
	}
}

func ReadBooleanData(dataString string) (types.Data, error) {
	if dataString == "" {
		return booleanData{}.ZeroValue(), nil
	}

	Bool, Error := strconv.ParseBool(dataString)
	if Error != nil {
		return booleanData{}.ZeroValue(), Error
	}

	return NewBooleanData(Bool), nil
}
