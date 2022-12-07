// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/data"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ data.BooleanData = (*BooleanData)(nil)

func (booleanData *BooleanData) GetID() ids.DataID {
	return baseIDs.GenerateDataID(booleanData)
}
func (booleanData *BooleanData) Compare(listable traits.Listable) int {
	compareBooleanData, err := booleanDataFromInterface(listable)
	if err != nil {
		panic(err)
	}

	if booleanData.Value == compareBooleanData.GetBooleanData().Value {
		return 0
	} else if booleanData.Value == true {
		return 1
	}

	return -1
}

// TODO test
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

func booleanDataFromInterface(listable traits.Listable) (*BooleanDataI, error) {
	switch value := listable.(type) {
	case *BooleanDataI:
		return value, nil
	default:
		panic(constants.MetaDataError)
	}
}

func BooleanDataPrototype() data.BooleanData {
	return (&BooleanData{}).ZeroValue().(data.BooleanData)
}

func NewBooleanData(value bool) data.BooleanData {
	return &BooleanData{
		Value: value,
	}
}
