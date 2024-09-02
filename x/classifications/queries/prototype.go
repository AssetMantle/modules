// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queries

import (
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/classifications/queries/classification"
	"github.com/AssetMantle/modules/x/classifications/queries/classifications"
	"github.com/AssetMantle/modules/x/classifications/queries/parameters"
)

func Prototype() helpers.Queries {
	return baseHelpers.NewQueries(
		classification.Query,
		classifications.Query,
		parameters.Query,
	)
}
