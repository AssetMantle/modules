// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package data

import (
	"github.com/AssetMantle/modules/schema/lists"
)

type ListData interface {
	Get() lists.DataList

	Search(Data) (int, bool)
	Add(...Data) ListData
	Remove(...Data) ListData

	Data
}
