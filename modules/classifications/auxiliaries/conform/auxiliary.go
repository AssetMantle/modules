// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package conform

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Auxiliary = base.NewAuxiliary(
	"conform",
	keeperPrototype,
)

var AuxiliaryMock = base.NewAuxiliary(
	"conform",
	keeperPrototypeMock,
)
