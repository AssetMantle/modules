// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queries

import (
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/splits/internal/queries/ownable"
	"github.com/AssetMantle/modules/x/splits/internal/queries/split"
)

func Prototype() helpers.Queries {
	return baseHelpers.NewQueries(
		split.Query,
		ownable.Query,
	)
}
