package mapper

import "github.com/persistenceOne/persistenceSDK/constants"

const ModuleName = "identities"
const ModuleRoute = constants.ProjectRoute + "/" + ModuleName
const DefaultParamspace = "identities"
const StoreKey = "identities"
const QueryRoute = ModuleRoute
const TransactionRoute = ModuleRoute

var StoreKeyPrefix = []byte{0x12}
