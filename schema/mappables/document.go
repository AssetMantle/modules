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

type Document interface {
	GetID() types.ID
	GetClassificationID() types.ID
	GetProperty(types.ID) types.Property

	traits.HasImmutables
	traits.HasMutables

	helpers.Mappable
}
