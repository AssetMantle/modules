// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package verify

import (
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
)

var Auxiliary = baseHelpers.NewAuxiliary(
	"maintainersVerify",
	keeperPrototype,
)

var AuxiliaryMock = baseHelpers.NewAuxiliary(
	"maintainersVerify",
	keeperPrototypeMock,
)
