// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package lists

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/types"
)

type MetaPropertyList interface {
	GetList() []types.MetaProperty
	GetMetaProperty(ids.PropertyID) types.MetaProperty
	ToPropertyList() PropertyList

	Add(...types.MetaProperty) MetaPropertyList
}
