/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappables

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Order interface {
	// TODO check if ID return type required
	GetRateID() types.ID
	GetCreationID() types.ID
	GetMakerOwnableID() types.ID
	GetTakerOwnableID() types.ID
	GetMakerID() types.ID

	GetCreation() types.MetaProperty
	GetExchangeRate() types.MetaProperty

	GetTakerID() types.Property
	GetExpiry() types.Property
	GetMakerOwnableSplit() types.Property

	Document
	helpers.Mappable
}
