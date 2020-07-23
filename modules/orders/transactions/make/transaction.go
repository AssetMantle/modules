package make

import (
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/orders/mapper"
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

var Transaction = utility.NewTransaction(
	mapper.ModuleName,
	TransactionName,
	TransactionRoute,
	TransactionShort,
	TransactionLong,
	registerCodec,
	initializeTransactionKeeper,
	requestPrototype,
	[]utility.CLIFlag{constants.ClassificationID, constants.MaintainersID, constants.Properties, constants.Lock, constants.Burn,
		constants.TakerAddress, constants.SenderAddress, constants.FeeRecipientAddress, constants.MakerAssetAmount, constants.MakerAssetData,
		constants.MakerFee, constants.MakerFeeAssetData, constants.TakerAssetAmount, constants.TakerAssetData, constants.TakerFee,
		constants.TakerFeeAssetData, constants.ExpirationTime, constants.Salt},
)
