package mapper

import "github.com/persistenceOne/persistenceSDK/constants"

const ModuleName = "splits"
const ModuleRoute = constants.ProjectRoute + "/" + ModuleName
const DefaultParamspace = "splits"
const StoreKey = "splits"
const QueryRoute = ModuleRoute
const TransactionRoute = ModuleRoute

var StoreKeyPrefix = []byte{0x13}
