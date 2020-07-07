package burn

import (
	"github.com/persistenceOne/persistenceSDK/modules/assets/constants"
	"github.com/persistenceOne/persistenceSDK/types"
)

var Transaction = types.NewTransaction(
	constants.ModuleName,
	constants.BurnTransaction,
	NewTransactionKeeper,
	constants.BurnTransactionShort,
	constants.BurnTransactionLong,
	requestPrototype,
	registerCodec,
	[]types.CLIFlag{
		constants.AssetID,
	},
)
