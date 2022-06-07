// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/AssetMantle/modules/constants"
)

// TODO move to errors
var (
	MockError = errors.Register(constants.ProjectRoute, 999, "MockError")
)
