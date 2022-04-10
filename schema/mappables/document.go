// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappables

import (
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/types"
)

type Document interface {
	GetID() types.ID
	GetClassificationID() types.ID
	// GetProperty returns property from a document searching in both Mutables and Immutables
	// * Returns nil if property is not found
	GetProperty(types.ID) types.Property

	// TODO see how to adjust in direct data
	qualified.HasImmutables
	qualified.HasMutables
}
