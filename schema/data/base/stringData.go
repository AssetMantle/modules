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

var _ data.StringData = (*StringData)(nil)

func (stringData *StringData) GetID() ids.DataID {
	return baseIDs.GenerateDataID(stringData)
}
func (stringData *StringData) Compare(listable traits.Listable) int {
	compareStringData, err := dataFromInterface(listable)
	if err != nil {
		panic(err)
	}

	return bytes.Compare(stringData.Bytes(), compareStringData.Bytes())
}
func (stringData *StringData) Bytes() []byte {
	return []byte(stringData.String())
}
func (stringData *StringData) GetType() ids.StringID {
	return dataConstants.StringDataID
}
func (stringData *StringData) ZeroValue() data.Data {
	return NewStringData("")
}
func (stringData *StringData) GenerateHashID() ids.HashID {
	return baseIDs.GenerateHashID(stringData.Bytes())
}
func (stringData *StringData) Get() string {
	return stringData.Value
}
func (stringData *StringData) ToAnyData() data.AnyData {
	return &AnyData{
		Impl: &AnyData_StringData{
			StringData: stringData,
		},
	}
}

func StringDataPrototype() data.Data {
	return NewStringData("")
}

func NewStringData(value string) data.Data {
	return &AnyData{
		Impl: &AnyData_StringData{
			StringData: &StringData{
				Value: value,
			},
		},
	}
}
