// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transactions

import (
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/orders/transactions/cancel"
	"github.com/AssetMantle/modules/x/orders/transactions/define"
	"github.com/AssetMantle/modules/x/orders/transactions/deputize"
	"github.com/AssetMantle/modules/x/orders/transactions/get"
	"github.com/AssetMantle/modules/x/orders/transactions/immediate"
	"github.com/AssetMantle/modules/x/orders/transactions/make"
	"github.com/AssetMantle/modules/x/orders/transactions/modify"
	"github.com/AssetMantle/modules/x/orders/transactions/put"
	"github.com/AssetMantle/modules/x/orders/transactions/revoke"
	"github.com/AssetMantle/modules/x/orders/transactions/take"
)

func Prototype() helpers.Transactions {
	return baseHelpers.NewTransactions(
		cancel.Transaction,
		define.Transaction,
		deputize.Transaction,
		get.Transaction,
		immediate.Transaction,
		make.Transaction,
		modify.Transaction,
		put.Transaction,
		revoke.Transaction,
		take.Transaction,
	)
}
