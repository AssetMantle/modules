package make

import (
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/orders/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/utilities"
	"github.com/persistenceOne/persistenceSDK/schema/utilities/base"
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
	[]utilities.CLIFlag{constants.ClassificationID, constants.MaintainersID, constants.Properties, constants.Lock, constants.Burn,
		constants.TakerAddress, constants.MakerAssetAmount, constants.MakerAssetData, constants.TakerAssetAmount,
		constants.TakerAssetData, constants.Salt},
)
