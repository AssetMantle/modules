/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package traits

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type HasImmutables interface {
	GetImmutableProperties() types.Properties

	GenerateHashID() types.ID
}
