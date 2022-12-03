// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package lists

import (
	"github.com/AssetMantle/modules/schema/data"
)

type DataList interface {
	Size() int
	GetList() []data.DataI
	Search(data.DataI) (int, bool)
	Add(...data.DataI) DataList
	Remove(...data.DataI) DataList
}
