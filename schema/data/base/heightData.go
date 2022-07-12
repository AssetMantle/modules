// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"strconv"

	"github.com/AssetMantle/modules/schema/data"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
	"github.com/AssetMantle/modules/schema/types"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
	stringUtilities "github.com/AssetMantle/modules/utilities/string"
)

type heightData struct {
	Value types.Height `json:"value"`
}

var _ data.HeightData = (*heightData)(nil)

func (heightData heightData) GetID() ids.DataID {
	return baseIDs.NewDataID(heightData)
}
func (heightData heightData) Compare(listable traits.Listable) int {
	compareHeightData, err := heightDataFromInterface(listable)
	if err != nil {
		panic(err)
	}

	return heightData.Value.Compare(compareHeightData.Value)
}
func (heightData heightData) String() string {
	return strconv.FormatInt(heightData.Value.Get(), 10)
}
func (heightData heightData) GetType() ids.ID {
	return dataConstants.HeightDataID
}
func (heightData heightData) ZeroValue() data.Data {
	return NewHeightData(baseTypes.NewHeight(0))
}
func (heightData heightData) GenerateHash() ids.ID {
	if heightData.Compare(heightData.ZeroValue()) == 0 {
		return baseIDs.NewStringID("")
	}

	return baseIDs.NewStringID(stringUtilities.Hash(strconv.FormatInt(heightData.Value.Get(), 10)))
}
func (heightData heightData) Get() types.Height {
	return heightData.Value
}

func heightDataFromInterface(listable traits.Listable) (heightData, error) {
	switch value := listable.(type) {
	case heightData:
		return value, nil
	default:
		return heightData{}, constants.MetaDataError
	}
}

func NewHeightData(value types.Height) data.Data {
	return heightData{
		Value: value,
	}
}
