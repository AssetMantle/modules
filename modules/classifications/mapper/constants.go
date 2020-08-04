package mapper

import "github.com/persistenceOne/persistenceSDK/constants"

const ModuleName = "classifications"
const ModuleRoute = constants.ProjectRoute + "/" + ModuleName
const DefaultParamspace = "classifications"
const StoreKey = "classifications"
const QueryRoute = ModuleRoute
const TransactionRoute = ModuleRoute

var StoreKeyPrefix = []byte{0x16}
