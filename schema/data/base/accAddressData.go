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

var _ data.AccAddressData = (*AccAddressDataI_AccAddressData)(nil)

func (accAddressData *AccAddressDataI_AccAddressData) GetID() ids.DataID {
	return baseIDs.GenerateDataID(accAddressData)
}
func (accAddressData *AccAddressDataI_AccAddressData) Compare(listable traits.Listable) int {
	compareAccAddressData, err := accAddressDataFromInterface(listable)
	if err != nil {
		panic(err)
	}
	return bytes.Compare(accAddressData.AccAddressData.Value, compareAccAddressData.Bytes())
}
func (accAddressData *AccAddressDataI_AccAddressData) String() string {
	return sdkTypes.AccAddress(accAddressData.AccAddressData.Value).String()
}
func (accAddressData *AccAddressDataI_AccAddressData) Bytes() []byte {
	return sdkTypes.AccAddress(accAddressData.AccAddressData.Value).Bytes()
}
func (accAddressData *AccAddressDataI_AccAddressData) GetType() ids.StringID {
	return dataConstants.AccAddressDataID
}
func (accAddressData *AccAddressDataI_AccAddressData) ZeroValue() data.Data {
	return AccAddressDataPrototype()
}
func (accAddressData *AccAddressDataI_AccAddressData) GenerateHashID() ids.HashID {
	if accAddressData.Compare(accAddressData.ZeroValue()) == 0 {
		// TODO test
		return baseIDs.GenerateHashID()
	}

	return baseIDs.GenerateHashID(accAddressData.Bytes())
}
func (accAddressData *AccAddressDataI_AccAddressData) Get() sdkTypes.AccAddress {
	return accAddressData.AccAddressData.Value
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

func NewAccAddressData(value sdkTypes.AccAddress) data.AccAddressData {
	return &AccAddressDataI{
		Impl: &AccAddressDataI_AccAddressData{
			AccAddressData: &AccAddressData{
				Value: value,
			},
		}}
}
