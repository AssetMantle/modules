// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package data

import (
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/types"
)

type ListData interface {
	types.Data
	Get() lists.DataList
}
