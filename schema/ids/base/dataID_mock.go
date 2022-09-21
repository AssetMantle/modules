// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
	"strconv"
)

type booleanData struct {
	Value bool `json:"value"`
}

func (booleanData booleanData) String() string {
	panic("implement me")
}

func (booleanData booleanData) GetType() ids.ID {
	return NewID("B")
}

func (booleanData booleanData) ZeroValue() data.Data {
	panic("implement me")
}

func (booleanData booleanData) GenerateHash() ids.ID {
	if booleanData.Value {
		return NewID(strconv.FormatBool(true))
	} else {
		return NewID(strconv.FormatBool(false))
	}

}

func (booleanData booleanData) Compare(listable traits.Listable) int {
	panic("implement me")
}

func (booleanData booleanData) GetID() ids.DataID {
	return NewDataID(booleanData)
}

func NewBooleanData(value bool) data.Data {
	return booleanData{
		Value: value,
	}
}
