package mapper

import "github.com/persistenceOne/persistenceSDK/constants"

const ModuleName = "orders"
const ModuleRoute = constants.ProjectRoute + "/" + ModuleName
const DefaultParamspace = "orders"
const StoreKey = "orders"
const QueryRoute = ModuleRoute
const TransactionRoute = ModuleRoute

var StoreKeyPrefix = []byte{0x14}
