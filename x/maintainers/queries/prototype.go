// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queries

import (
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/maintainers/queries/maintainer"
	"github.com/AssetMantle/modules/x/maintainers/queries/maintainers"
	"github.com/AssetMantle/modules/x/maintainers/queries/parameters"
)

func Prototype() helpers.Queries {
	return baseHelpers.NewQueries(
		maintainer.Query,
		maintainers.Query,
		parameters.Query,
	)
}
