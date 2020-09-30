/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package super

import (
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Auxiliary = base.NewAuxiliary(
	mapper.ModuleName,
	AuxiliaryName,
	AuxiliaryRoute,
	initializeAuxiliaryKeeper,
)

var AuxiliaryMock = base.NewAuxiliary(
	mapper.ModuleName,
	AuxiliaryName,
	AuxiliaryRoute,
	initializeAuxiliaryKeeperMock,
)
