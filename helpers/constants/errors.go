// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package constants

import (
	"github.com/AssetMantle/modules/helpers/base"
)

const projectRoute = "/AssetMantle"

var (
	MockError           = base.NewError(projectRoute, 100, "mock error")
	EntityAlreadyExists = base.NewError(projectRoute, 102, "entity already exists")
	EntityNotFound      = base.NewError(projectRoute, 103, "entity not found")
	IncorrectFormat     = base.NewError(projectRoute, 104, "incorrect format")
	InvalidMessage      = base.NewError(projectRoute, 105, "invalid message")
	InsufficientBalance = base.NewError(projectRoute, 106, "insufficient balance")
	InvalidParameter    = base.NewError(projectRoute, 107, "invalid parameter")
	InvalidRequest      = base.NewError(projectRoute, 108, "invalid request")
	MetaDataError       = base.NewError(projectRoute, 109, "meta data error")
	NotAuthorized       = base.NewError(projectRoute, 110, "not authorized")
	UninitializedUsage  = base.NewError(projectRoute, 111, "uninitialized usage")
)
