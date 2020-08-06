/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappers

import (
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Classifications interface {
	GetID() types.ID

	Get(types.ID) mappables.Classification
	GetList() []mappables.Classification

	Fetch(types.ID) Classifications
	Add(mappables.Classification) Classifications
	Remove(mappables.Classification) Classifications
	Mutate(mappables.Classification) Classifications
}
