// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package properties

import (
	"github.com/AssetMantle/modules/schema/data"
)

type MetaProperty interface {
	GetData() data.Data
	RemoveData() Property
	Property
}
