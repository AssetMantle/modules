// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package burn

import (
	"github.com/AssetMantle/modules/schema/helpers/base"
)

var Auxiliary = base.NewAuxiliary(
	"burn",
	keeperPrototype,
)

var AuxiliaryMock = base.NewAuxiliary(
	"burn",
	keeperPrototypeMock,
)
