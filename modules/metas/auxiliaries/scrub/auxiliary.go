/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package scrub

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Auxiliary = base.NewAuxiliary(
	"scrub",
	keeperPrototype,
)

var AuxiliaryMock = base.NewAuxiliary(
	"scrub",
	keeperPrototypeMock,
)
