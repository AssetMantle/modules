// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"strconv"

	"github.com/AssetMantle/modules/constants/errors"
	idsConstants "github.com/AssetMantle/modules/constants/ids"
	"github.com/AssetMantle/modules/schema/capabilities"
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/types"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
	"github.com/AssetMantle/modules/utilities/meta"
)

type heightData struct {
	Value types.Height `json:"value"`
}

var _ data.HeightData = (*heightData)(nil)

func (heightData heightData) GetID() ids.DataID {
	return baseIDs.NewDataID(heightData)
}
func (heightData heightData) Compare(listable capabilities.Listable) int {
	compareHeightData, err := heightDataFromInterface(listable)
	if err != nil {
		panic(err)
	}

	return heightData.Value.Compare(compareHeightData.Value)
}
func (heightData heightData) String() string {
	return strconv.FormatInt(heightData.Value.Get(), 10)
}
func (heightData heightData) GetType() types.ID {
	return idsConstants.HeightDataID
}
func (heightData heightData) ZeroValue() types.Data {
	return NewHeightData(baseTypes.NewHeight(0))
}
func (heightData heightData) GenerateHash() types.ID {
	if heightData.Compare(heightData.ZeroValue()) == 0 {
		return baseIDs.NewID("")
	}

	return baseIDs.NewID(meta.Hash(strconv.FormatInt(heightData.Value.Get(), 10)))
}
func (heightData heightData) Get() types.Height {
	return heightData.Value
}

func heightDataFromInterface(listable capabilities.Listable) (heightData, error) {
	switch value := listable.(type) {
	case heightData:
		return value, nil
	default:
		return heightData{}, errors.MetaDataError
	}
}

func NewHeightData(value types.Height) types.Data {
	return heightData{
		Value: value,
	}
}

func ReadHeightData(dataString string) (types.Data, error) {
	if dataString == "" {
		return heightData{}.ZeroValue(), nil
	}

	height, err := strconv.ParseInt(dataString, 10, 64)
	if err != nil {
		return nil, err
	}

	return NewHeightData(baseTypes.NewHeight(height)), nil
}
