// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package genesis

import (
	"github.com/AssetMantle/modules/modules/maintainers/internal/key"
	"github.com/AssetMantle/modules/modules/maintainers/internal/mappable"
	"github.com/AssetMantle/modules/modules/maintainers/internal/module"
	"github.com/AssetMantle/modules/modules/maintainers/internal/parameters"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
)

func Prototype() helpers.Genesis {
	return baseHelpers.NewGenesis(key.Prototype, mappable.Prototype, []helpers.Mappable{mappable.NewMappable(module.GenesisMaintainer)}, parameters.Prototype().GetList())
}
