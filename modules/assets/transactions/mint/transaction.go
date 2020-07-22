package mint

import (
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/assets/mapper"
	"github.com/persistenceOne/persistenceSDK/types/utility"
	"github.com/persistenceOne/persistenceSDK/types/utility/base"
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
	[]utility.CLIFlag{constants.ClassificationID, constants.MaintainersID, constants.Properties, constants.Lock, constants.Burn},
)
