// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package properties

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

type Property interface {
	GetID() ids.ID
	GetDataID() ids.ID
	GetKey() ids.ID
	GetType() ids.ID
	GetData() data.Data

	IsMeta() bool
	ScrubData() Property

	traits.Listable
}
