// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/lists"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type idList struct {
	types.List
}

func (i idList) GetList() []types.ID {
	// TODO implement me
	panic("implement me")
}

var _ lists.IDList = (*idList)(nil)
