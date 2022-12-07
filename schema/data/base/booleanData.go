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

var _ data.BooleanData = (*Data_BooleanData)(nil)

func (booleanData *Data_BooleanData) GetID() ids.ID {
	return baseIDs.GenerateDataID(booleanData)
}
func (booleanData *Data_BooleanData) Compare(listable traits.Listable) int {
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
func (booleanData *Data_BooleanData) String() string {
	return booleanData.BooleanData.String()
}
func (booleanData *Data_BooleanData) Bytes() []byte {
	if booleanData.Get() {
		return []byte{0x1}
	}
	return []byte{0x0}
}
func (booleanData *Data_BooleanData) GetType() ids.ID {
	return dataConstants.BooleanDataID
}
func (booleanData *Data_BooleanData) ZeroValue() data.Data {
	return NewBooleanData(false)
}
func (booleanData *Data_BooleanData) GenerateHashID() ids.ID {
	if booleanData.Compare(booleanData.ZeroValue()) == 0 {
		return baseIDs.GenerateHashID()
	}

	return baseIDs.GenerateHashID(booleanData.Bytes())
}
func (booleanData *Data_BooleanData) Get() bool {
	return booleanData.BooleanData.Value
}

func BooleanDataPrototype() data.Data {
	return (&Data_BooleanData{}).ZeroValue()
}

func NewBooleanData(value bool) data.Data {
	return &Data{
		Impl: &Data_BooleanData{
			BooleanData: &BooleanData{
				Value: value,
			},
		},
	}
}
