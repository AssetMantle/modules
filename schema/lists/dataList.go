// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package lists

import (
	"github.com/AssetMantle/modules/schema/data"
)

type AnyDataList interface {
	Size() int
	GetList() []data.AnyData
	Search(data.AnyData) (int, bool)
	Add(...data.AnyData) AnyDataList
	Remove(...data.AnyData) AnyDataList
}
