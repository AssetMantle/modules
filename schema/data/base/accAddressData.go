// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"

	"buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/data/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/data"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
)

type accAddressData base.AccAddressData

var _ data.AccAddressData = (*accAddressData)(nil)

func (accAddressData *accAddressData) GetID() ids.DataID {
	return baseIDs.NewDataID(accAddressData)
}
func (accAddressData *accAddressData) Compare(listable traits.Listable) int {
	compareAccAddressData, err := accAddressDataFromInterface(listable)
	if err != nil {
		panic(err)
	}

	return bytes.Compare(accAddressData.Value, compareAccAddressData.Value)
}
func (accAddressData *accAddressData) String() string {
	return sdkTypes.AccAddress(accAddressData.Value).String()
}
func (accAddressData *accAddressData) Bytes() []byte {
	return sdkTypes.AccAddress(accAddressData.Value).Bytes()
}
func (accAddressData *accAddressData) GetType() ids.StringID {
	return dataConstants.AccAddressDataID
}
func (accAddressData *accAddressData) ZeroValue() data.DataI {
	return NewAccAddressData(sdkTypes.AccAddress{})
}
func (accAddressData *accAddressData) GenerateHashID() ids.HashID {
	if accAddressData.Compare(accAddressData.ZeroValue()) == 0 {
		// TODO test
		return baseIDs.GenerateHashID()
	}

	return baseIDs.GenerateHashID(accAddressData.Bytes())
}
func (accAddressData *accAddressData) Get() sdkTypes.AccAddress {
	return accAddressData.Value
}

func accAddressDataFromInterface(listable traits.Listable) (*accAddressData, error) {
	switch value := listable.(type) {
	case *accAddressData:
		return value, nil
	default:
		return &accAddressData{}, errorConstants.MetaDataError
	}
}

func AccAddressDataPrototype() data.AccAddressData {
	return (&accAddressData{}).ZeroValue().(data.AccAddressData)
}

func NewAccAddressData(value sdkTypes.AccAddress) data.AccAddressData {
	return &accAddressData{
		Value: value,
	}
}
