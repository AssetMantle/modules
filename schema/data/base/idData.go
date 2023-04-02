// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"
	"strings"

	"github.com/AssetMantle/modules/schema/data"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ data.IDData = (*IDData)(nil)

func (idData *IDData) ValidateBasic() error {
	return idData.Value.ValidateBasic()
}
func (idData *IDData) GetID() ids.DataID {
	return baseIDs.GenerateDataID(idData)
}
func (idData *IDData) GetBondWeight() int64 {
	return dataConstants.IDDataWeight
}
func (idData *IDData) AsString() string {
	return idData.Value.AsString()
}
func (idData *IDData) FromString(dataString string) (data.Data, error) {
	dataString = strings.TrimSpace(dataString)
	if dataString == "" {
		return PrototypeIDData(), nil
	}

	id, err := baseIDs.PrototypeAnyID().FromString(dataString)
	if err != nil {
		return PrototypeIDData(), err
	}

	return NewIDData(id), nil
}
func (idData *IDData) Compare(listable traits.Listable) int {
	compareIDData, err := dataFromListable(listable)
	if err != nil {
		panic(err)
	}

	return bytes.Compare(idData.Bytes(), compareIDData.Bytes())
}
func (idData *IDData) Bytes() []byte {
	return idData.Value.Bytes()
}
func (idData *IDData) GetTypeID() ids.StringID {
	return dataConstants.IDDataTypeID
}
func (idData *IDData) ZeroValue() data.Data {
	return PrototypeIDData()
}
func (idData *IDData) GenerateHashID() ids.HashID {
	return baseIDs.GenerateHashID(idData.Bytes())
}
func (idData *IDData) Get() ids.AnyID {
	return idData.Value
}
func (idData *IDData) ToAnyData() data.AnyData {
	return &AnyData{
		Impl: &AnyData_IDData{
			IDData: idData,
		},
	}
}
func PrototypeIDData() data.IDData {
	return NewIDData(baseIDs.NewStringID(""))
}

func NewIDData(value ids.ID) data.IDData {
	return &IDData{
		Value: value.ToAnyID().(*baseIDs.AnyID),
	}
}
