// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"
	"encoding/binary"
	"strconv"

	"github.com/AssetMantle/modules/schema/data"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
	"github.com/AssetMantle/modules/schema/types"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

var _ data.HeightData = (*Data_HeightData)(nil)

func (heightData *Data_HeightData) Unmarshal(bytes []byte) error {
	// TODO implement me
	panic("implement me")
}
func (heightData *Data_HeightData) GetID() ids.ID {
	return baseIDs.GenerateDataID(heightData)
}
func (heightData *Data_HeightData) Compare(listable traits.Listable) int {
	compareHeightData, err := dataFromInterface(listable)
	if err != nil {
		panic(err)
	}

	return bytes.Compare(heightData.Bytes(), compareHeightData.Bytes())
}
func (heightData *Data_HeightData) String() string {
	return strconv.FormatInt(heightData.Get().Get(), 10)
}
func (heightData *Data_HeightData) Bytes() []byte {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, uint64(heightData.Get().Get()))
	return bytes
}
func (heightData *Data_HeightData) GetType() ids.ID {
	return dataConstants.HeightDataID
}
func (heightData *Data_HeightData) ZeroValue() data.Data {
	return NewHeightData(baseTypes.NewHeight(-1))
}
func (heightData *Data_HeightData) GenerateHashID() ids.ID {
	if heightData.Compare(heightData.ZeroValue()) == 0 {
		return baseIDs.GenerateHashID()
	}
	return baseIDs.GenerateHashID(heightData.Bytes())
}
func (heightData *Data_HeightData) Get() types.Height {
	return baseTypes.NewHeight(heightData.HeightData.Value)
}

func HeightDataPrototype() data.HeightData {
	return NewHeightData(baseTypes.NewHeight(0)).ZeroValue().(data.HeightData)
}

func NewHeightData(value types.Height) data.Data {
	return &Data{
		Impl: &Data_HeightData{
			HeightData: &HeightData{
				Value: value.Get(),
			},
		},
	}
}
