// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queries

import (
	"github.com/AssetMantle/modules/modules/classifications/internal/queries/classification"
	"github.com/AssetMantle/modules/modules/classifications/internal/queries/classifications"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
)

func Prototype() helpers.Queries {
	return baseHelpers.NewQueries(
		classification.Query,
		classifications.Query,
	)
}
