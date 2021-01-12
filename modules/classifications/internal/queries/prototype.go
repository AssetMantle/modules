/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package queries

import (
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/queries/classification"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

func Prototype() helpers.Queries {
	return base.NewQueries(
		classification.Query,
	)
}
