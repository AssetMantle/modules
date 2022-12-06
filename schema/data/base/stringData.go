// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"strings"

	"github.com/AssetMantle/modules/schema/data"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ data.StringData = (*StringDataI_StringData)(nil)

func (stringData *StringDataI_StringData) GetID() ids.DataID {
	return baseIDs.GenerateDataID(stringData)
}
func (stringData *StringDataI_StringData) Compare(listable traits.Listable) int {
	compareStringData, err := stringDataFromInterface(listable)
	if err != nil {
		panic(err)
	}

	return strings.Compare(stringData.Get(), compareStringData.Get())
}
func (stringData *StringDataI_StringData) String() string {
	return stringData.StringData.String()
}
func (stringData *StringDataI_StringData) Bytes() []byte {
	return []byte(stringData.String())
}
func (stringData *StringDataI_StringData) GetType() ids.StringID {
	return dataConstants.StringDataID
}
func (stringData *StringDataI_StringData) ZeroValue() data.Data {
	return NewStringData("")
}
func (stringData *StringDataI_StringData) GenerateHashID() ids.HashID {
	return baseIDs.GenerateHashID(stringData.Bytes())
}
func (stringData *StringDataI_StringData) Get() string {
	return stringData.StringData.Value
}

func stringDataFromInterface(listable traits.Listable) (*StringDataI, error) {
	switch value := listable.(type) {
	case *StringDataI:
		return value, nil
	default:
		return nil, constants.MetaDataError
	}
}

func StringDataPrototype() data.StringData {
	return NewStringData("")
}

func NewStringData(value string) data.StringData {
	return &StringDataI{
		Impl: &StringDataI_StringData{
			StringData: &StringData{
				Value: value,
			},
		},
	}
}
