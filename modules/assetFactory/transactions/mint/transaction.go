package mint

import (
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/constants"
	"github.com/persistenceOne/persistenceSDK/types"
)

var Transaction = types.NewTransaction(
	constants.ModuleName,
	constants.MintTransaction,
	constants.MintTransactionShort,
	constants.MintTransactionLong,
	requestPrototype,
	registerCodec,
	[]types.CLIFlag{
		constants.Properties,
		constants.ChainID,
		constants.MaintainersID,
		constants.ClassificationID,
		constants.Lock,
		constants.Burn},
)
