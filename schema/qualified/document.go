// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package qualified

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/types"
)

// TODO retire or send to base
type Document interface {
	GetID() types.ID
	GetClassificationID() types.ID
	// GetProperty returns property from a document searching in both Mutables and Immutables
	// * Returns nil if property is not found
	GetProperty(ids.PropertyID) types.Property
	GetImmutablePropertyList() lists.PropertyList
	GetMutablePropertyList() lists.PropertyList

	GenerateHashID() types.ID
	Mutate(propertyList ...types.Property) Document
}
