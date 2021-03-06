/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package traits

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Document interface {
	GetID() types.ID
	GetClassificationID() types.ID
	// GetProperty returns property from a document searching in both Mutables and Immutables
	// * Returns nil if property is not found
	GetProperty(types.ID) types.Property

	HasImmutables
	HasMutables
}
