/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappers

import (
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Splits interface {
	GetID() types.ID

	Get(types.ID) mappables.Split
	GetList() []mappables.Split

	Fetch(types.ID) Splits
	Add(mappables.Split) Splits
	Remove(mappables.Split) Splits
	Mutate(mappables.Split) Splits
}
