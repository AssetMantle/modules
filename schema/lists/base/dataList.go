// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/types"
)

type dataList struct {
	types.List
}

var _ lists.DataList = (*dataList)(nil)

func (dataList dataList) GetList() []types.Data {
	// TODO implement me
	panic("implement me")
}

func (dataList dataList) Search(data types.Data) (bool, int) {
	// TODO implement me
	panic("implement me")
}

func NewDataList(dataList ...types.Data) lists.DataList {
	// TODO implement me
	panic("implement me")
}
