/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mint

import (
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Auxiliary = base.NewAuxiliary(
	Name,
	initializeAuxiliaryKeeper,
)
var AuxiliaryMock = base.NewAuxiliary(
	Name,
	AuxiliaryRoute,
	initializeAuxiliaryKeeperMock,
)
