// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/types"
)

type idList struct {
	types.List
}

var _ lists.IDList = (*idList)(nil)

func (idList idList) Get() []types.ID {
	// TODO implement me
	panic("implement me")
}
