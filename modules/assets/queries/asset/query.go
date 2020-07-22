package asset

import (
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/assets/mapper"
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

var Query = utility.NewQuery(
	mapper.ModuleName,
	QueryName,
	QueryRoute,
	QueryShort,
	QueryLong,
	packageCodec,
	registerCodec,
	initializeQueryKeeper,
	queryRequestPrototype,
	queryResponsePrototype,
	[]utility.CLIFlag{constants.AssetID},
)
