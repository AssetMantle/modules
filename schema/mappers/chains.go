/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappers

import (
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Chains interface {
	GetID() types.ID

	Get(types.ID) mappables.Chain
	GetList() []mappables.Chain

	Fetch(types.ID) Chains
	Add(mappables.Chain) Chains
	Remove(mappables.Chain) Chains
	Mutate(mappables.Chain) Chains
}
