// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Auxiliary = base.NewAuxiliary(
	"define",
	keeperPrototype,
)

var AuxiliaryMock = base.NewAuxiliary(
	"define",
	keeperPrototypeMock,
)
