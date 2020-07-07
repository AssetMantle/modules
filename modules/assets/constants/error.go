package constants

import "github.com/cosmos/cosmos-sdk/types/errors"

var UnknownMessage = errors.Register(ModuleName, 001, "UnknownMessage")
var IncorrectMessage = errors.Register(ModuleName, 002, "IncorrectMessage")
var UnknownQuery = errors.Register(ModuleName, 101, "UnknownQuery")
var IncorrectQuery = errors.Register(ModuleName, 102, "IncorrectQuery")
var EntityNotFound = errors.Register(ModuleName, 103, "EntityNotFound")
var EntityAlreadyExists = errors.Register(ModuleName, 104, "EntityAlreadyExists")
var BurnNotAllowed = errors.Register(ModuleName, 105, "BurnNotAllowed")
var MutableNotFound = errors.Register(ModuleName, 105, "MutableNotFound")
