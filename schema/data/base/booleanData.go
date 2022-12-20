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

var _ data.BooleanData = (*BooleanData)(nil)

func (booleanData *BooleanData) GetID() ids.DataID {
	return baseIDs.GenerateDataID(booleanData)
}
func (booleanData *BooleanData) Compare(listable traits.Listable) int {
	compareBooleanData, err := dataFromInterface(listable)
	if err != nil {
		panic(err)
	}

	if value := bytes.Compare(booleanData.Bytes(), compareBooleanData.Bytes()); value == 0 {
		return 0
	} else if value > 0 {
		return 1
	} else {
		return -1
	}
}
func (booleanData *BooleanData) Bytes() []byte {
	if booleanData.Get() {
		return []byte{0x1}
	}
	return []byte{0x0}
}
func (booleanData *BooleanData) GetType() ids.StringID {
	return dataConstants.BooleanDataID
}
func (booleanData *BooleanData) ZeroValue() data.Data {
	return NewBooleanData(false)
}
func (booleanData *BooleanData) GenerateHashID() ids.HashID {
	if booleanData.Compare(booleanData.ZeroValue()) == 0 {
		return baseIDs.GenerateHashID()
	}

	return baseIDs.GenerateHashID(booleanData.Bytes())
}
func (booleanData *BooleanData) Get() bool {
	return booleanData.Value
}
func (booleanData *BooleanData) ToAnyData() data.AnyData {
	return &AnyData{
		Impl: &AnyData_BooleanData{
			BooleanData: booleanData,
		},
	}
}

func BooleanDataPrototype() data.Data {
	return &BooleanData{}
}

func NewBooleanData(value bool) data.Data {
	return &AnyData{
		Impl: &AnyData_BooleanData{
			BooleanData: &BooleanData{
				Value: value,
			},
		},
	}
}
