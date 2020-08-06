/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappers

import (
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Maintainers interface {
	GetID() types.ID

	Get(types.ID) mappables.Maintainer
	GetList() []mappables.Maintainer

	Fetch(types.ID) Maintainers
	Add(mappables.Maintainer) Maintainers
	Remove(mappables.Maintainer) Maintainers
	Mutate(mappables.Maintainer) Maintainers
}
