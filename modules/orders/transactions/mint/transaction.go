package mint

import (
	"github.com/persistenceOne/persistenceSDK/modules/orders/constants"
	"github.com/persistenceOne/persistenceSDK/types"
)

var Transaction = types.NewTransaction(
	constants.ModuleName,
	constants.MintTransaction,
	constants.MintTransaction,
	constants.MintTransactionShort,
	constants.MintTransactionLong,
	registerCodec,
	initializeTransactionKeeper,
	requestPrototype,
	[]types.CLIFlag{constants.BuyCoinDenom, constants.BuyCoinAmount, constants.SellCoinDenom, constants.SellCoinAmount, constants.Properties},
)
