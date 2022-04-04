// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

import (
	"github.com/AssetMantle/modules/schema/helpers/base"
)

var Auxiliary = base.NewAuxiliary(
	"deputize",
	keeperPrototype,
)

var AuxiliaryMock = base.NewAuxiliary(
	"deputize",
	keeperPrototypeMock,
)
