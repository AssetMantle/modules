package mapper

import "github.com/persistenceOne/persistenceSDK/constants"

const ModuleName = "assets"
const ModuleRoute = constants.ProjectRoute + "/" + ModuleName
const DefaultParamspace = "assets"
const StoreKey = "assets"
const QueryRoute = ModuleRoute
const TransactionRoute = ModuleRoute

var StoreKeyPrefix = []byte{0x11}
