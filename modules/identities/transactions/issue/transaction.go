package issue

import (
	"github.com/persistenceOne/persistenceSDK/modules/identities/constants"
	"github.com/persistenceOne/persistenceSDK/types"
)

var Transaction = types.NewTransaction(
	constants.ModuleName,
	constants.IssueTransaction,
	constants.IssueTransactionShort,
	constants.IssueTransactionLong,
	registerCodec,
	initializeTransactionKeeper,
	requestPrototype,
	[]types.CLIFlag{constants.IdentityID},
)
