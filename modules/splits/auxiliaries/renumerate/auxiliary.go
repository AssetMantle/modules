// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package renumerate

import (
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
)

var Auxiliary = baseHelpers.NewAuxiliary(
	"renumerate",
	keeperPrototype,
)

var AuxiliaryMock = baseHelpers.NewAuxiliary(
	"renumerate",
	keeperPrototypeMock,
)
