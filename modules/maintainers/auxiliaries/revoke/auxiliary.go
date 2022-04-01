// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package revoke

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Auxiliary = base.NewAuxiliary(
	"revoke",
	keeperPrototype,
)

var AuxiliaryMock = base.NewAuxiliary(
	"revoke",
	keeperPrototypeMock,
)
