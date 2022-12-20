// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package properties

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

type Property interface {
	GetID() ids.PropertyID
	GetDataID() ids.DataID
	GetKey() ids.StringID
	GetType() ids.StringID
	GetData() data.AnyData

	IsMeta() bool
	ScrubData() Property

	ToAnyProperty() AnyProperty
	traits.Listable
}
