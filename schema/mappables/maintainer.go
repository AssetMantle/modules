/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappables

import (
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Maintainer interface {
	GetIdentityID() types.ID
	GetMaintainedID() types.ID

	CanAddMaintainer() bool
	CanRemoveMaintainer() bool
	CanMutateMaintainer() bool

	MaintainsTrait(types.ID) bool
	traits.Mappable
}
