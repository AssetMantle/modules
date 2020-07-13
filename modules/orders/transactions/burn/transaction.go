package burn

import (
	"github.com/persistenceOne/persistenceSDK/modules/orders/constants"
	"github.com/persistenceOne/persistenceSDK/types"
)

var Transaction = types.NewTransaction(
	constants.ModuleName,
	constants.BurnTransaction,
	constants.BurnTransactionShort,
	constants.BurnTransactionLong,
	registerCodec,
	initializeTransactionKeeper,
	requestPrototype,
	[]types.CLIFlag{constants.OrderID},
)
