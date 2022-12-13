// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package authenticate

import (
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
)

var Auxiliary = baseHelpers.NewAuxiliary(
	"authenticate",
	keeperPrototype,
)

var AuxiliaryMock = baseHelpers.NewAuxiliary(
	"authenticate",
	keeperPrototypeMock,
)
