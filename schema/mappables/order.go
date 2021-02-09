/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappables

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Order interface {
	GetID() types.ID
	GetClassificationID() types.ID
	GetExchangeRate() types.Data
	GetCreationHeight() types.Data
	GetMakerOwnableID() types.ID
	GetTakerOwnableID() types.ID
	GetMakerID() types.ID
	GetKey() helpers.Key

	GetTakerID() types.Property
	GetCreation() types.Property
	GetExpiry() types.Property
	GetMakerOwnableSplit() types.Property

	traits.HasMutables
	traits.HasImmutables
	helpers.Mappable
}
