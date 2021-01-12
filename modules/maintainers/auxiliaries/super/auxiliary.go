/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package super

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Auxiliary = base.NewAuxiliary(
	"super",
	keeperPrototype,
)

var AuxiliaryMock = base.NewAuxiliary(
	"super",
	keeperPrototypeMock,
)
