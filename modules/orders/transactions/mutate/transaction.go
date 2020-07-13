package mutate

import (
	"github.com/persistenceOne/persistenceSDK/modules/orders/constants"
	"github.com/persistenceOne/persistenceSDK/types"
)

var Transaction = types.NewTransaction(
	constants.ModuleName,
	constants.MutateTransaction,
	constants.MutateTransactionShort,
	constants.MutateTransactionLong,
	registerCodec,
	initializeTransactionKeeper,
	requestPrototype,
	[]types.CLIFlag{constants.AssetID, constants.Properties, constants.Lock, constants.Burn},
)
