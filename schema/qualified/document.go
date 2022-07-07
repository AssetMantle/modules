// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package qualified

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/properties"
)

type Document interface {
	GetID() ids.ID
	GetClassificationID() ids.ID
	// GetProperty returns property from a document searching in both Mutables and Immutables
	// * Returns nil if property is not found
	GetProperty(ids.PropertyID) properties.Property
	GetImmutablePropertyList() lists.PropertyList
	GetMutablePropertyList() lists.PropertyList

	GenerateHashID() ids.ID
	Mutate(...properties.Property) Document
}
