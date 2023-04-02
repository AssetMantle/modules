// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"
	"strings"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/data"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ data.AccAddressData = (*AccAddressData)(nil)

func (accAddressData *AccAddressData) ValidateBasic() error {
	return sdkTypes.VerifyAddressFormat(accAddressData.Value)
}
func (accAddressData *AccAddressData) GetID() ids.DataID {
	return baseIDs.GenerateDataID(accAddressData)
}
func (accAddressData *AccAddressData) GetBondWeight() int64 {
	return dataConstants.AccAddressDataWeight
}
func (accAddressData *AccAddressData) Compare(listable traits.Listable) int {
	compareAccAddressData, err := dataFromListable(listable)
	if err != nil {
		panic(err)
	}

	return bytes.Compare(accAddressData.Bytes(), compareAccAddressData.Bytes())
}
func (accAddressData *AccAddressData) AsString() string {
	return sdkTypes.AccAddress(accAddressData.Value).String()
}
func (accAddressData *AccAddressData) FromString(dataString string) (data.Data, error) {
	dataString = strings.TrimSpace(dataString)
	if dataString == "" {
		return PrototypeAccAddressData(), nil
	}

	accAddress, err := sdkTypes.AccAddressFromBech32(dataString)
	if err != nil {
		return PrototypeAccAddressData(), err
	}

	return NewAccAddressData(accAddress), nil
}
func (accAddressData *AccAddressData) Bytes() []byte {
	return sdkTypes.AccAddress(accAddressData.Value).Bytes()
}
func (accAddressData *AccAddressData) GetTypeID() ids.StringID {
	return dataConstants.AccAddressDataTypeID
}
func (accAddressData *AccAddressData) ZeroValue() data.Data {
	return PrototypeAccAddressData()
}
func (accAddressData *AccAddressData) GenerateHashID() ids.HashID {
	if accAddressData.Compare(accAddressData.ZeroValue()) == 0 {
		// TODO test
		return baseIDs.GenerateHashID()
	}

	return baseIDs.GenerateHashID(accAddressData.Bytes())
}
func (accAddressData *AccAddressData) Get() sdkTypes.AccAddress {
	return accAddressData.Value
}
func (accAddressData *AccAddressData) ToAnyData() data.AnyData {
	return &AnyData{
		Impl: &AnyData_AccAddressData{
			AccAddressData: accAddressData,
		}}
}

func PrototypeAccAddressData() data.Data {
	return NewAccAddressData(sdkTypes.AccAddress{})
}

func NewAccAddressData(value sdkTypes.AccAddress) data.AccAddressData {
	return &AccAddressData{
		Value: value,
	}
}
