// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/lists"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type dataList struct {
	types.List
}

func (dataList dataList) Search(data types.Data) int {
	// TODO implement me
	panic("implement me")
}

var _ lists.DataList = (*dataList)(nil)

func (dataList dataList) GetList() []types.Data {
	// TODO implement me
	panic("implement me")
}
