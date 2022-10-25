// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/ids"

	"github.com/AssetMantle/modules/schema/traits"
)

type booleanData struct {
	Value bool `json:"value"`
}

var _ data.BooleanData = (*booleanData)(nil)

func (booleanData booleanData) GetID() ids.DataID {
	panic("")
	//return NewDataID(booleanData)
}
func (booleanData booleanData) Compare(listable traits.Listable) int {
	return -1
}
func (booleanData booleanData) String() string {
	panic("implement me")
}

// TODO test
func (booleanData booleanData) Bytes() []byte {
	if booleanData.Get() {
		return []byte{0x1}
	}
	return []byte{0x0}
}
func (booleanData booleanData) GetType() ids.StringID {
	return NewStringID("B")
}
func (booleanData booleanData) ZeroValue() data.Data {
	return NewBooleanData(false)
}
func (booleanData booleanData) GenerateHashID() ids.HashID {
	if booleanData.Compare(booleanData.ZeroValue()) == 0 {
		return GenerateHashID()
	}

	return GenerateHashID(booleanData.Bytes())
}
func (booleanData booleanData) Get() bool {
	return booleanData.Value
}

func NewBooleanData(value bool) data.Data {
	return booleanData{
		Value: value,
	}
}
