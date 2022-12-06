// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package metas

import (
	"github.com/AssetMantle/modules/modules/metas/auxiliaries"
	"github.com/AssetMantle/modules/modules/metas/module/block"
	"github.com/AssetMantle/modules/modules/metas/module/genesis"
	"github.com/AssetMantle/modules/modules/metas/module/invariants"
	"github.com/AssetMantle/modules/modules/metas/module/mapper"
	"github.com/AssetMantle/modules/modules/metas/module/module"
	"github.com/AssetMantle/modules/modules/metas/module/parameters"
	"github.com/AssetMantle/modules/modules/metas/module/queries"
	"github.com/AssetMantle/modules/modules/metas/module/transactions"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
)

func Prototype() helpers.Module {
	return baseHelpers.NewModule(
		module.Name,
		1,
		auxiliaries.Prototype,
		block.Prototype,
		genesis.Prototype,
		invariants.Prototype,
		mapper.Prototype,
		nil,
		parameters.Prototype,
		queries.Prototype,
		nil,
		transactions.Prototype,
	)
}
