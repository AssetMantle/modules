/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package transactions

import (
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/transactions/burn"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/transactions/define"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/transactions/mint"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/transactions/mutate"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

func Prototype() helpers.Transactions {
	return base.NewTransactions(
		burn.Transaction,
		define.Transaction,
		mint.Transaction,
		mutate.Transaction,
	)
}
