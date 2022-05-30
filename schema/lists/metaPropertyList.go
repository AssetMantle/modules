// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package lists

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/properties"
)

type MetaPropertyList interface {
	GetList() []properties.MetaProperty
	GetMetaProperty(ids.PropertyID) properties.MetaProperty
	ToPropertyList() PropertyList

	Add(...properties.MetaProperty) MetaPropertyList
}
