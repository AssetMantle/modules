/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mapper

import "github.com/persistenceOne/persistenceSDK/constants"

const ModuleName = "orders"
const ModuleRoute = constants.ProjectRoute + "/" + ModuleName
const DefaultParamspace = "orders"
const StoreKey = "orders"
const QueryRoute = ModuleRoute
const TransactionRoute = ModuleRoute

var StoreKeyPrefix = []byte{0x14}
