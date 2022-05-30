// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/schema/data"
	idsConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
	stringUtilities "github.com/AssetMantle/modules/utilities/string"
)

type accAddressData struct {
	Value sdkTypes.AccAddress `json:"value"`
}

var _ data.AccAddressData = (*accAddressData)(nil)

func (accAddressData accAddressData) GetID() ids.DataID {
	return baseIDs.NewDataID(accAddressData)
}
func (accAddressData accAddressData) Compare(listable traits.Listable) int {
	compareAccAddressData, err := accAddressDataFromInterface(listable)
	if err != nil {
		panic(err)
	}

	return bytes.Compare(accAddressData.Value.Bytes(), compareAccAddressData.Value.Bytes())
}
func (accAddressData accAddressData) String() string {
	return accAddressData.Value.String()
}
func (accAddressData accAddressData) GetType() ids.ID {
	return idsConstants.AccAddressDataID
}
func (accAddressData accAddressData) ZeroValue() data.Data {
	return NewAccAddressData(sdkTypes.AccAddress{})
}
func (accAddressData accAddressData) GenerateHash() ids.ID {
	if accAddressData.Compare(accAddressData.ZeroValue()) == 0 {
		return baseIDs.NewID("")
	}

	return baseIDs.NewID(stringUtilities.Hash(accAddressData.Value.String()))
}
func (accAddressData accAddressData) Get() sdkTypes.AccAddress {
	return accAddressData.Value
}

func accAddressDataFromInterface(listable traits.Listable) (accAddressData, error) {
	switch value := listable.(type) {
	case accAddressData:
		return value, nil
	default:
		return accAddressData{}, errors.MetaDataError
	}
}

func NewAccAddressData(value sdkTypes.AccAddress) data.Data {
	return accAddressData{
		Value: value,
	}
}

func ReadAccAddressData(dataString string) (data.Data, error) {
	if dataString == "" {
		return accAddressData{}.ZeroValue(), nil
	}

	accAddress, err := sdkTypes.AccAddressFromBech32(dataString)
	if err != nil {
		return accAddressData{}.ZeroValue(), err
	}

	return NewAccAddressData(accAddress), nil
}
