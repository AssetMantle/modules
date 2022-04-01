// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transactions

import (
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/transactions/burn"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/transactions/define"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/transactions/deputize"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/transactions/mint"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/transactions/mutate"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/transactions/renumerate"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/transactions/revoke"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

func Prototype() helpers.Transactions {
	return base.NewTransactions(
		burn.Transaction,
		define.Transaction,
		deputize.Transaction,
		mint.Transaction,
		mutate.Transaction,
		renumerate.Transaction,
		revoke.Transaction,
	)
}
