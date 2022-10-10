// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"strconv"

	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/schema/data"
	idsConstants "github.com/AssetMantle/modules/schema/data/constants"
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
	return idsConstants.HeightDataID
}
func (heightData heightData) ZeroValue() data.Data {
	return NewHeightData(baseTypes.NewHeight(-1))
}
func (heightData heightData) GenerateHash() ids.ID {
	if heightData.Compare(heightData.ZeroValue()) == 0 {
		return baseIDs.NewID("")
	}

	return baseIDs.NewID(stringUtilities.Hash(strconv.FormatInt(heightData.Value.Get(), 10)))
}
func (heightData heightData) Get() types.Height {
	return heightData.Value
}

func heightDataFromInterface(listable traits.Listable) (heightData, error) {
	switch value := listable.(type) {
	case heightData:
		return value, nil
	default:
		return heightData{}, errors.MetaDataError
	}
}

func NewHeightData(value types.Height) data.Data {
	return heightData{
		Value: value,
	}
}

func ReadHeightData(dataString string) (data.Data, error) {
	if dataString == "" {
		return heightData{}.ZeroValue(), nil
	}

	height, err := strconv.ParseInt(dataString, 10, 64)
	if err != nil {
		return nil, err
	}

	return NewHeightData(baseTypes.NewHeight(height)), nil
}
