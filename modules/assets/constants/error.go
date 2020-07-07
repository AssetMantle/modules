package constants

import "github.com/cosmos/cosmos-sdk/types/errors"

var IncorrectMessage = errors.Register(ModuleName, 002, "IncorrectMessage")
var EntityNotFound = errors.Register(ModuleName, 103, "EntityNotFound")
var EntityAlreadyExists = errors.Register(ModuleName, 104, "EntityAlreadyExists")
var BurnNotAllowed = errors.Register(ModuleName, 105, "BurnNotAllowed")
var MutableNotFound = errors.Register(ModuleName, 106, "MutableNotFound")
