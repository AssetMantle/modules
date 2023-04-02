// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"
	"encoding/binary"
	"strconv"
	"strings"

	"github.com/AssetMantle/modules/schema/data"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
	"github.com/AssetMantle/modules/schema/types"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

var _ data.HeightData = (*HeightData)(nil)

func (heightData *HeightData) ValidateBasic() error {
	return heightData.Value.ValidateBasic()
}
func (heightData *HeightData) GetID() ids.DataID {
	return baseIDs.GenerateDataID(heightData)
}
func (heightData *HeightData) GetBondWeight() int64 {
	return dataConstants.HeightDataWeight
}
func (heightData *HeightData) AsString() string {
	return strconv.FormatInt(heightData.Value.Get(), 10)
}
func (heightData *HeightData) FromString(dataString string) (data.Data, error) {
	dataString = strings.TrimSpace(dataString)
	if dataString == "" {
		return PrototypeHeightData(), nil
	}

	Height, err := strconv.ParseInt(dataString, 10, 64)
	if err != nil {
		return PrototypeHeightData(), err
	}

	return NewHeightData(baseTypes.NewHeight(Height)), nil
}
func (heightData *HeightData) Compare(listable traits.Listable) int {
	compareHeightData, err := dataFromListable(listable)
	if err != nil {
		panic(err)
	}

	return bytes.Compare(heightData.Bytes(), compareHeightData.Bytes())
}
func (heightData *HeightData) Bytes() []byte {
	Bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(Bytes, uint64(heightData.Get().Get()))

	return Bytes
}
func (heightData *HeightData) GetTypeID() ids.StringID {
	return dataConstants.HeightDataTypeID
}
func (heightData *HeightData) ZeroValue() data.Data {
	return NewHeightData(baseTypes.NewHeight(-1))
}
func (heightData *HeightData) GenerateHashID() ids.HashID {
	if heightData.Compare(heightData.ZeroValue()) == 0 {
		return baseIDs.GenerateHashID()
	}
	return baseIDs.GenerateHashID(heightData.Bytes())
}
func (heightData *HeightData) Get() types.Height {
	return heightData.Value
}
func (heightData *HeightData) ToAnyData() data.AnyData {
	return &AnyData{
		Impl: &AnyData_HeightData{
			HeightData: heightData,
		},
	}
}

func PrototypeHeightData() data.HeightData {
	return NewHeightData(baseTypes.NewHeight(0)).ZeroValue().(data.HeightData)
}

func NewHeightData(value types.Height) data.HeightData {
	return &HeightData{
		Value: value.(*baseTypes.Height),
	}
}
