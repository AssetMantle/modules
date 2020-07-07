package mutate

import (
	"github.com/persistenceOne/persistenceSDK/modules/assets/constants"
	"github.com/persistenceOne/persistenceSDK/types"
)

var Transaction = types.NewTransaction(
	constants.ModuleName,
	constants.MutateTransaction,
	NewTransactionKeeper,
	constants.MutateTransactionShort,
	constants.MutateTransactionLong,
	requestPrototype,
	registerCodec,
	[]types.CLIFlag{
		constants.AssetID,
		constants.Properties,
		constants.Lock,
		constants.Burn},
)
