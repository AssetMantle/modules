// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package constants

import (
	"github.com/AssetMantle/modules/schema/errors/base"
)

const projectRoute = "/AssetMantle"

var (
	EntityAlreadyExists  = base.NewError(projectRoute, 102, "EntityAlreadyExists")
	EntityNotFound       = base.NewError(projectRoute, 103, "EntityNotFound")
	IncorrectFormat      = base.NewError(projectRoute, 104, "IncorrectFormat")
	IncorrectMessage     = base.NewError(projectRoute, 105, "IncorrectMessage")
	InsufficientBalance  = base.NewError(projectRoute, 106, "InsufficientBalance")
	InvalidParameter     = base.NewError(projectRoute, 107, "InvalidParameter")
	InvalidRequest       = base.NewError(projectRoute, 108, "InvalidRequest")
	MetaDataError        = base.NewError(projectRoute, 109, "MetaDataError")
	MockError            = base.NewError(projectRoute, 999, "MockError")
	NotAuthorized        = base.NewError(projectRoute, 110, "NotAuthorized")
	UninitializedUsage   = base.NewError(projectRoute, 111, "UninitializedUsage")
	UnsupportedParameter = base.NewError(projectRoute, 112, "UnsupportedParameter")
)
