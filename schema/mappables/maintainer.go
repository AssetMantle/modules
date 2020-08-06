/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappables

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Maintainer interface {
	Sting() string

	GetAddress() sdkTypes.AccAddress
	GetID() types.ID

	CanMutateMaintainersProperty(types.ID) bool

	CanAddMaintainer() bool
	CanRemoveMaintainer() bool
	CanMutateMaintainer() bool

	CanMutateLock() bool
	CanMutateBurn() bool
	CanMutateTrait(types.ID) bool
	traits.Mappable
}
