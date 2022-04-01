// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transfer

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Auxiliary = base.NewAuxiliary(
	"transfer",
	keeperPrototype,
)

var AuxiliaryMock = base.NewAuxiliary(
	"transfer",
	keeperPrototypeMock,
)
