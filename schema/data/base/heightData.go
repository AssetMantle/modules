// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"encoding/binary"
	"strconv"

	"github.com/AssetMantle/modules/schema/data"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
	"github.com/AssetMantle/modules/schema/types"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

var _ data.HeightData = (*HeightDataI_HeightData)(nil)

func (heightData *HeightDataI_HeightData) GetID() ids.DataID {
	return baseIDs.GenerateDataID(heightData)
}
func (heightData *HeightDataI_HeightData) Compare(listable traits.Listable) int {
	compareHeightData, err := heightDataFromInterface(listable)
	if err != nil {
		panic(err)
	}

	return heightData.Get().Compare(compareHeightData.Get())
}
func (heightData *HeightDataI_HeightData) String() string {
	return strconv.FormatInt(heightData.Get().Get(), 10)
}
func (heightData *HeightDataI_HeightData) Bytes() []byte {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, uint64(heightData.Get().Get()))
	return bytes
}
func (heightData *HeightDataI_HeightData) GetType() ids.StringID {
	return dataConstants.HeightDataID
}
func (heightData *HeightDataI_HeightData) ZeroValue() data.Data {
	return NewHeightData(baseTypes.NewHeight(-1))
}
func (heightData *HeightDataI_HeightData) GenerateHashID() ids.HashID {
	if heightData.Compare(heightData.ZeroValue()) == 0 {
		return baseIDs.GenerateHashID()
	}
	return baseIDs.GenerateHashID(heightData.Bytes())
}
func (heightData *HeightDataI_HeightData) Get() types.Height {
	return baseTypes.NewHeight(heightData.HeightData.Value)
}

func heightDataFromInterface(listable traits.Listable) (*HeightDataI_HeightData, error) {
	switch value := listable.(type) {
	case *HeightDataI_HeightData:
		return value, nil
	default:
		return nil, constants.MetaDataError
	}
}

func HeightDataPrototype() data.HeightData {
	return NewHeightData(baseTypes.NewHeight(0)).ZeroValue().(data.HeightData)
}

func NewHeightData(value types.Height) data.HeightData {
	return &HeightDataI{
		Impl: &HeightDataI_HeightData{
			HeightData: &HeightData{
				Value: value.Get(),
			},
		},
	}
}
