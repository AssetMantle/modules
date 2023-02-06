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

	IsMeta() bool
	ValidateBasic() error

	ToAnyProperty() AnyProperty

	Mutate(data.Data) Property
	traits.Listable
}
