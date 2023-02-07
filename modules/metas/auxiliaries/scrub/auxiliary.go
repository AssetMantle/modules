// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package scrub

import (
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
)

var Auxiliary = baseHelpers.NewAuxiliary(
	"scrub",
	keeperPrototype,
)

var AuxiliaryMock = baseHelpers.NewAuxiliary(
	"scrub",
	keeperPrototypeMock,
)
