package constants

import "github.com/cosmos/cosmos-sdk/types/errors"

var BurnNotAllowed = errors.Register(ProjectRoute, 101, "BurnNotAllowed")
var EntityAlreadyExists = errors.Register(ProjectRoute, 102, "EntityAlreadyExists")
var EntityNotFound = errors.Register(ProjectRoute, 103, "EntityNotFound")
var IncorrectMessage = errors.Register(ProjectRoute, 104, "IncorrectMessage")
var MutableNotFound = errors.Register(ProjectRoute, 105, "MutableNotFound")
