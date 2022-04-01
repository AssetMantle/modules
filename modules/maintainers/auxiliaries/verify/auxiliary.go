// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package verify

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Auxiliary = base.NewAuxiliary(
	"maintainersVerify",
	keeperPrototype,
)

var AuxiliaryMock = base.NewAuxiliary(
	"maintainersVerify",
	keeperPrototypeMock,
)
