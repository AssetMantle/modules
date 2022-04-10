// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package super

import (
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
)

var Auxiliary = baseHelpers.NewAuxiliary(
	"super",
	keeperPrototype,
)

var AuxiliaryMock = baseHelpers.NewAuxiliary(
	"super",
	keeperPrototypeMock,
)
