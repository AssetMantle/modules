package identity

import (
	"github.com/persistenceOne/persistenceSDK/modules/identities/constants"
	"github.com/persistenceOne/persistenceSDK/types"
)

var Query = types.NewQuery(
	constants.ModuleName,
	constants.IdentityQuery,
	constants.IdentityQueryShort,
	constants.IdentityQueryLong,
	packageCodec,
	registerCodec,
	initializeQueryKeeper,
	queryRequestPrototype,
	queryResponsePrototype,
	[]types.CLIFlag{constants.IdentityID},
)
