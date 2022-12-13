// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"

	"github.com/AssetMantle/modules/schema/data"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ data.StringData = (*Data_StringData)(nil)

func (stringData *Data_StringData) Unmarshal(bytes []byte) error {
	// TODO implement me
	panic("implement me")
}
func (stringData *Data_StringData) GetID() ids.ID {
	return baseIDs.GenerateDataID(stringData)
}
func (stringData *Data_StringData) Compare(listable traits.Listable) int {
	compareStringData, err := dataFromInterface(listable)
	if err != nil {
		panic(err)
	}

	return bytes.Compare(stringData.Bytes(), compareStringData.Bytes())
}
func (stringData *Data_StringData) String() string {
	return stringData.StringData.String()
}
func (stringData *Data_StringData) Bytes() []byte {
	return []byte(stringData.String())
}
func (stringData *Data_StringData) GetType() ids.ID {
	return dataConstants.StringDataID
}
func (stringData *Data_StringData) ZeroValue() data.Data {
	return NewStringData("")
}
func (stringData *Data_StringData) GenerateHashID() ids.ID {
	return baseIDs.GenerateHashID(stringData.Bytes())
}
func (stringData *Data_StringData) Get() string {
	return stringData.StringData.Value
}

func StringDataPrototype() data.Data {
	return NewStringData("")
}

func NewStringData(value string) data.Data {
	return &Data{
		Impl: &Data_StringData{
			StringData: &StringData{
				Value: value,
			},
		},
	}
}
