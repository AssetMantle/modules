// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/data"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ data.AccAddressData = (*Data_AccAddressData)(nil)

func (accAddressData *Data_AccAddressData) String() string {
	return sdkTypes.AccAddress(accAddressData.AccAddressData.Value).String()
}
func (accAddressData *Data_AccAddressData) GetID() ids.ID {
	return baseIDs.GenerateDataID(accAddressData)
}
func (accAddressData *Data_AccAddressData) Compare(listable traits.Listable) int {
	compareAccAddressData, err := accAddressDataFromInterface(listable)
	if err != nil {
		panic(err)
	}
	return bytes.Compare(accAddressData.Value, compareAccAddressData.Bytes())
}
func (accAddressData *Data_AccAddressData) Bytes() []byte {
	return sdkTypes.AccAddress(accAddressData.Value).Bytes()
}
func (accAddressData *Data_AccAddressData) GetType() ids.StringID {
	return dataConstants.AccAddressDataID
}
func (accAddressData *Data_AccAddressData) ZeroValue() data.Data {
	return AccAddressDataPrototype()
}
func (accAddressData *Data_AccAddressData) GenerateHashID() ids.HashID {
	if accAddressData.Compare(accAddressData.ZeroValue()) == 0 {
		// TODO test
		return baseIDs.GenerateHashID()
	}

	return baseIDs.GenerateHashID(accAddressData.Bytes())
}
func (accAddressData *Data_AccAddressData) Get() sdkTypes.AccAddress {
	return accAddressData.Value
}

func accAddressDataFromInterface(listable traits.Listable) (*AccAddressDataI, error) {
	switch value := listable.(type) {
	case *AccAddressDataI:
		return value, nil
	default:
		panic(errorConstants.MetaDataError)
	}
}

func AccAddressDataPrototype() data.AccAddressData {
	return NewAccAddressData(sdkTypes.AccAddress{})
}

func GenerateAccAddressData(value sdkTypes.AccAddress) data.AccAddressData {
	return NewAccAddressData(value)
}

func NewAccAddressData(value sdkTypes.AccAddress) data.Data {
	return &Data{Impl: &Data_AccAddressData{AccAddressData: &AccAddressData{Value: value.Bytes()}}}
}
