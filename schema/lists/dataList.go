// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package lists

import (
	"github.com/AssetMantle/modules/schema/data"
)

type DataList interface {
	Size() int
	GetList() []data.Data
	Search(data.Data) (int, bool)
	Add(...data.Data) DataList
	Delete(...data.Data) DataList
}
