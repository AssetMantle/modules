// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package qualified

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/properties"
)

type Document interface {
	GenerateHashID() ids.HashID
	GetClassificationID() ids.ClassificationID
	// GetProperty returns property from a document searching in both Mutables and Immutables
	// * Returns nil if property is not found
	GetProperty(ids.PropertyID) properties.Property
	GetImmutables() Immutables
	GetMutables() Mutables

	Mutate(...properties.Property) Document
}
