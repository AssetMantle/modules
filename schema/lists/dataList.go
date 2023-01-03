// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package lists

import (
	"github.com/AssetMantle/modules/schema/data"
)

type AnyDataList interface {
	GetList() []data.AnyData
	Search(data.Data) (int, bool)
	Add(...data.Data) AnyDataList
	Remove(...data.Data) AnyDataList
}
