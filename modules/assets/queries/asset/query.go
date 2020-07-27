package asset

import (
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/assets/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Query = base.NewQuery(
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
	[]helpers.CLIFlag{constants.AssetID},
)
