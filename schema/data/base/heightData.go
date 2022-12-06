// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"encoding/binary"
	"strconv"

	"buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/data/base"

	"github.com/AssetMantle/modules/schema/data"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
	"github.com/AssetMantle/modules/schema/types"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

type heightData base.HeightData

var _ data.HeightData = (*heightData)(nil)

func (heightData *heightData) GetID() ids.DataID {
	return baseIDs.NewDataID(heightData)
}
func (heightData *heightData) Compare(listable traits.Listable) int {
	compareHeightData, err := heightDataFromInterface(listable)
	if err != nil {
		panic(err)
	}

	return heightData.Get().Compare(compareHeightData.Get())
}
func (heightData *heightData) String() string {
	return strconv.FormatInt(heightData.Get().Get(), 10)
}
func (heightData *heightData) Bytes() []byte {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, uint64(heightData.Get().Get()))
	return bytes
}
func (heightData *heightData) GetType() ids.StringID {
	return dataConstants.HeightDataID
}
func (heightData *heightData) ZeroValue() data.Data {
	return NewHeightData(baseTypes.NewHeight(-1))
}
func (heightData *heightData) GenerateHashID() ids.HashID {
	if heightData.Compare(heightData.ZeroValue()) == 0 {
		return baseIDs.GenerateHashID()
	}
	return baseIDs.GenerateHashID(heightData.Bytes())
}
func (heightData *heightData) Get() types.Height {
	return baseTypes.NewHeight(heightData.Value)
}

func heightDataFromInterface(listable traits.Listable) (*heightData, error) {
	switch value := listable.(type) {
	case *heightData:
		return value, nil
	default:
		return &heightData{}, constants.MetaDataError
	}
}

func HeightDataPrototype() data.HeightData {
	return (&heightData{}).ZeroValue().(data.HeightData)
}

func NewHeightData(value types.Height) data.HeightData {
	return &heightData{
		Value: value.Get(),
	}
}
