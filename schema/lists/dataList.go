// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package lists

import (
	"github.com/AssetMantle/modules/schema/data"
)

type DataList interface {
	Size() int
	GetList() []data.Data
	Search(data data.Data) (int, bool)
}
