/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappers

import (
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Metas interface {
	GetID() types.ID

	Get(types.ID) mappables.Meta
	GetList() []mappables.Meta

	Fetch(types.ID) Metas
	Add(mappables.Meta) Metas
	Remove(mappables.Meta) Metas
	Mutate(mappables.Meta) Metas
}
