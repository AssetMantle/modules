// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"
	"strings"

	"github.com/AssetMantle/modules/schema/data/utilities"

	"github.com/AssetMantle/modules/schema/data"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ data.StringData = (*StringData)(nil)

func (stringData *StringData) ValidateBasic() error {
	if !utilities.IsValidStringData(stringData.Value) {
		return errorConstants.IncorrectFormat
	}

	return nil
}

func (stringData *StringData) GetID() ids.DataID {
	return baseIDs.GenerateDataID(stringData)
}
func (stringData *StringData) GetBondWeight() int64 {
	return dataConstants.StringDataWeight
}
func (stringData *StringData) Compare(listable traits.Listable) int {
	compareStringData, err := dataFromListable(listable)
	if err != nil {
		panic(err)
	}

	return bytes.Compare(stringData.Bytes(), compareStringData.Bytes())
}
func (stringData *StringData) Bytes() []byte {
	return []byte(stringData.Value)
}
func (stringData *StringData) GetTypeID() ids.StringID {
	return dataConstants.StringDataTypeID
}
func (stringData *StringData) ZeroValue() data.Data {
	return NewStringData("")
}
func (stringData *StringData) GenerateHashID() ids.HashID {
	return baseIDs.GenerateHashID(stringData.Bytes())
}
func (stringData *StringData) AsString() string {
	return stringData.Value
}
func (stringData *StringData) FromString(dataString string) (data.Data, error) {
	dataString = strings.TrimSpace(dataString)
	if dataString == "" {
		return PrototypeStringData(), nil
	}

	return NewStringData(dataString), nil
}
func (stringData *StringData) ToAnyData() data.AnyData {
	return &AnyData{
		Impl: &AnyData_StringData{
			StringData: stringData,
		},
	}
}

func PrototypeStringData() data.StringData {
	return NewStringData("")
}

func NewStringData(value string) data.StringData {
	return &StringData{
		Value: value,
	}
}
