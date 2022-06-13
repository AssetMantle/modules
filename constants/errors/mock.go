// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package errors

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

// TODO move to errors
var (
	MockError = errors.Register(projectRoute, 999, "MockError")
)
