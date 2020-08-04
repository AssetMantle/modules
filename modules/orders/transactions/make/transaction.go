package make

import (
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/orders/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Transaction = base.NewTransaction(
	mapper.ModuleName,
	TransactionName,
	TransactionRoute,
	TransactionShort,
	TransactionLong,
	registerCodec,
	initializeTransactionKeeper,
	requestPrototype,
	[]helpers.CLIFlag{constants.Properties, constants.Lock, constants.Burn,
		constants.TakerAddress, constants.MakerAssetAmount, constants.MakerAssetData, constants.MakerAssetType, constants.TakerAssetAmount,
		constants.TakerAssetData, constants.TakerAssetType},
)
