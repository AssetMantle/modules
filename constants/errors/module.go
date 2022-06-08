// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package errors

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

const projectRoute = "/AssetMantle"

// TODO create error types
var (
	DeletionNotAllowed   = errors.Register(projectRoute, 101, "DeletionNotAllowed")
	EntityAlreadyExists  = errors.Register(projectRoute, 102, "EntityAlreadyExists")
	EntityNotFound       = errors.Register(projectRoute, 103, "EntityNotFound")
	IncorrectFormat      = errors.Register(projectRoute, 104, "IncorrectFormat")
	IncorrectMessage     = errors.Register(projectRoute, 105, "IncorrectMessage")
	InsufficientBalance  = errors.Register(projectRoute, 106, "InsufficientBalance")
	InvalidParameter     = errors.Register(projectRoute, 107, "InvalidParameter")
	InvalidRequest       = errors.Register(projectRoute, 108, "InvalidRequest")
	MetaDataError        = errors.Register(projectRoute, 109, "MetaDataError")
	NotAuthorized        = errors.Register(projectRoute, 110, "NotAuthorized")
	UninitializedUsage   = errors.Register(projectRoute, 111, "UninitializedUsage")
	UnsupportedParameter = errors.Register(projectRoute, 112, "UnsupportedParameter")
)
