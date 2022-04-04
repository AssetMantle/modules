// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package scrub

import (
	"github.com/AssetMantle/modules/schema/helpers/base"
)

var Auxiliary = base.NewAuxiliary(
	"scrub",
	keeperPrototype,
)

var AuxiliaryMock = base.NewAuxiliary(
	"scrub",
	keeperPrototypeMock,
)
