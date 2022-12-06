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

var _ data.BooleanData = (*BooleanDataI_BooleanData)(nil)

func (booleanData *BooleanDataI_BooleanData) GetID() ids.DataID {
	return baseIDs.GenerateDataID(booleanData)
}
func (booleanData *BooleanDataI_BooleanData) Compare(listable traits.Listable) int {
	compareBooleanData, err := booleanDataFromInterface(listable)
	if err != nil {
		panic(err)
	}

	if booleanData.BooleanData.Value == compareBooleanData.BooleanData.Value {
		return 0
	} else if booleanData.BooleanData.Value == true {
		return 1
	}

	return -1
}

func (booleanData *BooleanDataI_BooleanData) String() string {
	return booleanData.BooleanData.String()
}

// TODO test
func (booleanData *BooleanDataI_BooleanData) Bytes() []byte {
	if booleanData.Get() {
		return []byte{0x1}
	}
	return []byte{0x0}
}
func (booleanData *BooleanDataI_BooleanData) GetType() ids.StringID {
	return dataConstants.BooleanDataID
}
func (booleanData *BooleanDataI_BooleanData) ZeroValue() data.Data {
	return NewBooleanData(false)
}
func (booleanData *BooleanDataI_BooleanData) GenerateHashID() ids.HashID {
	if booleanData.Compare(booleanData.ZeroValue()) == 0 {
		return baseIDs.GenerateHashID()
	}

	return baseIDs.GenerateHashID(booleanData.Bytes())
}
func (booleanData *BooleanDataI_BooleanData) Get() bool {
	return booleanData.BooleanData.Value
}

func booleanDataFromInterface(listable traits.Listable) (*BooleanDataI_BooleanData, error) {
	switch value := listable.(type) {
	case *BooleanDataI_BooleanData:
		return value, nil
	default:
		panic(constants.MetaDataError)
	}
}

func BooleanDataPrototype() data.BooleanData {
	return (&BooleanDataI_BooleanData{}).ZeroValue().(data.BooleanData)
}

func NewBooleanData(value bool) data.BooleanData {
	return &BooleanDataI{
		Impl: &BooleanDataI_BooleanData{
			BooleanData: &BooleanData{
				Value: value,
			},
		},
	}
}
