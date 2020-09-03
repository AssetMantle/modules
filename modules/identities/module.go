/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package identities

import (
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/queries/identity"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/transactions/define"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/transactions/issue"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/transactions/nub"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/transactions/provision"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/transactions/unprovision"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Module = base.NewModule(
	mapper.ModuleName,
	mapper.DefaultParamspace,
	mapper.QueryRoute,
	mapper.TransactionRoute,
	genesis.State,
	mapper.Mapper,
	[]helpers.Auxiliary{verify.Auxiliary},
	[]helpers.Query{identity.Query},
	[]helpers.Transaction{define.Transaction, issue.Transaction, nub.Transaction, provision.Transaction, unprovision.Transaction},
)
