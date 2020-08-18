/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package metas

import (
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/scrub"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/supplement"
	"github.com/persistenceOne/persistenceSDK/modules/metas/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/metas/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/metas/queries/meta"
	"github.com/persistenceOne/persistenceSDK/modules/metas/transactions/reveal"
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
	[]helpers.Auxiliary{scrub.Auxiliary, supplement.Auxiliary},
	[]helpers.Query{meta.Query},
	[]helpers.Transaction{reveal.Transaction},
)
