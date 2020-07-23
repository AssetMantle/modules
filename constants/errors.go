package constants

import "github.com/cosmos/cosmos-sdk/types/errors"

var DeletionNotAllowed = errors.Register(ProjectRoute, 101, "DeletionNotAllowed")
var EntityAlreadyExists = errors.Register(ProjectRoute, 102, "EntityAlreadyExists")
var EntityNotFound = errors.Register(ProjectRoute, 103, "EntityNotFound")
var IncorrectMessage = errors.Register(ProjectRoute, 104, "IncorrectMessage")
var MutationNotAllowed = errors.Register(ProjectRoute, 105, "MutationNotAllowed")
var NotAuthorized = errors.Register(ProjectRoute, 106, "NotAuthorized")
