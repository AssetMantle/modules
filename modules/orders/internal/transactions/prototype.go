// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transactions

import (
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/transactions/cancel"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/transactions/define"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/transactions/deputize"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/transactions/immediate"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/transactions/make"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/transactions/modify"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/transactions/revoke"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/transactions/take"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

func Prototype() helpers.Transactions {
	return base.NewTransactions(
		cancel.Transaction,
		define.Transaction,
		deputize.Transaction,
		immediate.Transaction,
		make.Transaction,
		modify.Transaction,
		revoke.Transaction,
		take.Transaction,
	)
}
