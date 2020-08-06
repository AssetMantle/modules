/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappers

import (
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type InterIdentities interface {
	GetID() types.ID

	Get(types.ID) mappables.InterIdentity
	GetList() []mappables.InterIdentity

	Fetch(types.ID) InterIdentities
	Add(mappables.InterIdentity) InterIdentities
	Remove(mappables.InterIdentity) InterIdentities
	Mutate(mappables.InterIdentity) InterIdentities
}
