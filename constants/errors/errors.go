/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package errors

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/persistenceOne/persistenceSDK/constants"
)

var (
	DeletionNotAllowed   = errors.Register(constants.ProjectRoute, 101, "DeletionNotAllowed")
	EntityAlreadyExists  = errors.Register(constants.ProjectRoute, 102, "EntityAlreadyExists")
	EntityNotFound       = errors.Register(constants.ProjectRoute, 103, "EntityNotFound")
	IncorrectFormat      = errors.Register(constants.ProjectRoute, 104, "IncorrectFormat")
	IncorrectMessage     = errors.Register(constants.ProjectRoute, 105, "IncorrectMessage")
	InsufficientBalance  = errors.Register(constants.ProjectRoute, 106, "InsufficientBalance")
	InvalidParameter     = errors.Register(constants.ProjectRoute, 107, "InvalidParameter")
	InvalidRequest       = errors.Register(constants.ProjectRoute, 108, "InvalidRequest")
	MetaDataError        = errors.Register(constants.ProjectRoute, 109, "MetaDataError")
	NotAuthorized        = errors.Register(constants.ProjectRoute, 110, "NotAuthorized")
	UninitializedUsage   = errors.Register(constants.ProjectRoute, 111, "UninitializedUsage")
	UnsupportedParameter = errors.Register(constants.ProjectRoute, 112, "UnsupportedParameter")
)
