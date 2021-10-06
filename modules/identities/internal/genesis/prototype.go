/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package genesis

import (
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/parameters/dummy"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

func Prototype() helpers.Genesis {
	return NewGenesis([]mappable.Identity{}, []dummy.DummyParameter{})
}
