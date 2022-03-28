/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package traits

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type HasImmutables interface {
	// GetImmutableProperties return the immutable properties object
	// does not return nil
	GetImmutableProperties() types.Properties

	GenerateHashID() types.ID
	types.Proto
}
