// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transfer

import (
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
)

var Auxiliary = baseHelpers.NewAuxiliary(
	"transfer",
	keeperPrototype,
)

var AuxiliaryMock = baseHelpers.NewAuxiliary(
	"transfer",
	keeperPrototypeMock,
)
