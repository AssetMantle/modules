// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package lists

import (
	"github.com/AssetMantle/modules/schema/types"
)

type DataList interface {
	Size() int
	GetList() []types.Data
	Search(data types.Data) (bool, int)
}
