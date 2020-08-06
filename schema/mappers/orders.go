/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappers

import (
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Orders interface {
	GetID() types.ID

	Get(types.ID) mappables.Order
	GetList() []mappables.Order

	Fetch(types.ID) Orders
	Add(mappables.Order) Orders
	Remove(mappables.Order) Orders
	Mutate(mappables.Order) Orders
}
