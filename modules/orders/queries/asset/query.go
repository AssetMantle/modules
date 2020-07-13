package asset

import (
	"github.com/persistenceOne/persistenceSDK/modules/orders/constants"
	"github.com/persistenceOne/persistenceSDK/types"
)

var Query = types.NewQuery(
	constants.ModuleName,
	constants.AssetQuery,
	constants.AssetQueryShort,
	constants.AssetQueryLong,
	packageCodec,
	registerCodec,
	initializeQueryKeeper,
	queryRequestPrototype,
	queryResponsePrototype,
	[]types.CLIFlag{constants.AssetID},
)
