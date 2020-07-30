package mapper

import "github.com/persistenceOne/persistenceSDK/constants"

const ModuleName = "exchanges"
const ModuleRoute = constants.ProjectRoute + "/" + ModuleName
const DefaultParamspace = "exchanges"
const StoreKey = "exchanges"
const QueryRoute = ModuleRoute
const TransactionRoute = ModuleRoute

var StoreKeyPrefix = []byte{0x15}
