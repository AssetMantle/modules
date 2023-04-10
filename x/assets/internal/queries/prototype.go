// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queries

import (
	"github.com/AssetMantle/modules/x/assets/internal/queries/assets"
	"github.com/AssetMantle/schema/x/helpers"
	baseHelpers "github.com/AssetMantle/schema/x/helpers/base"

	"github.com/AssetMantle/modules/x/assets/internal/queries/asset"
)

func Prototype() helpers.Queries {
	return baseHelpers.NewQueries(
		asset.Query,
		assets.Query,
	)
}
