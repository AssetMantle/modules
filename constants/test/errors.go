// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/persistenceOne/persistenceSDK/constants"
)

var (
	MockError = errors.Register(constants.ProjectRoute, 999, "MockError")
)
