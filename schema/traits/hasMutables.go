/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package traits

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type HasMutables interface {
	// GetMutableProperties return the mutable properties object
	// does not return nil
	GetMutableProperties() types.Properties

	Mutate(propertyList ...types.Property) HasMutables
}
