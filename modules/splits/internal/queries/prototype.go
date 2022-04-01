// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queries

import (
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/queries/ownable"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/queries/split"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

func Prototype() helpers.Queries {
	return base.NewQueries(
		split.Query,
		ownable.Query,
	)
}
