/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package module

import (
	"github.com/persistenceOne/persistenceSDK/modules/assets/auxiliaries"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/queries"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/simulator"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/transactions"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Prototype = base.NewModule(
	Name,
	simulator.Prototype,
	parameters.Prototype,
	auxiliaries.Prototype,
	queries.Prototype,
	transactions.Prototype,
)
