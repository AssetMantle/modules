package asset

import (
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/constants"
	"github.com/persistenceOne/persistenceSDK/types"
)

var Query = types.NewQuery(
	constants.ModuleName,
	constants.AssetQuery,
	constants.AssetQueryShort,
	constants.AssetQueryLong,
	queryRequestPrototype,
	queryResponsePrototype,
	registerCodec,
	[]types.CLIFlag{constants.AssetID},
)
