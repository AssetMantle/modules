// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package constants

import (
	"github.com/AssetMantle/modules/constants"
	"github.com/AssetMantle/modules/schema/errors/base"
)

var (
	DeletionNotAllowed   = base.NewError(constants.ProjectRoute, 101, "DeletionNotAllowed")
	EntityAlreadyExists  = base.NewError(constants.ProjectRoute, 102, "EntityAlreadyExists")
	EntityNotFound       = base.NewError(constants.ProjectRoute, 103, "EntityNotFound")
	IncorrectFormat      = base.NewError(constants.ProjectRoute, 104, "IncorrectFormat")
	IncorrectMessage     = base.NewError(constants.ProjectRoute, 105, "IncorrectMessage")
	InsufficientBalance  = base.NewError(constants.ProjectRoute, 106, "InsufficientBalance")
	InvalidParameter     = base.NewError(constants.ProjectRoute, 107, "InvalidParameter")
	InvalidRequest       = base.NewError(constants.ProjectRoute, 108, "InvalidRequest")
	MetaDataError        = base.NewError(constants.ProjectRoute, 109, "MetaDataError")
	MockError            = base.NewError(constants.ProjectRoute, 999, "MockError")
	NotAuthorized        = base.NewError(constants.ProjectRoute, 110, "NotAuthorized")
	UninitializedUsage   = base.NewError(constants.ProjectRoute, 111, "UninitializedUsage")
	UnsupportedParameter = base.NewError(constants.ProjectRoute, 112, "UnsupportedParameter")
)
