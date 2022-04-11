// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package lists

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/types"
)

type PropertyList interface {
	GetProperty(ids.PropertyID) types.Property
	GetList() []types.Property

	Add(...types.Property) PropertyList
	Remove(...types.Property) PropertyList
	Mutate(...types.Property) PropertyList
}
