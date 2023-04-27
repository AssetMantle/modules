// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queries

import (
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/splits/queries/ownable"
	"github.com/AssetMantle/modules/x/splits/queries/split"
	"github.com/AssetMantle/modules/x/splits/queries/splits"
)

func Prototype() helpers.Queries {
	return baseHelpers.NewQueries(
		split.Query,
		splits.Query,
		ownable.Query,
	)
}