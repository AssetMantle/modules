// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transactions

import (
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/orders/transactions/cancel"
	"github.com/AssetMantle/modules/x/orders/transactions/get"
	"github.com/AssetMantle/modules/x/orders/transactions/put"
)

func Prototype() helpers.Transactions {
	return baseHelpers.NewTransactions(
		// TODO ***** revisit business logic
		// deputize.Transaction,
		// immediate.Transaction,
		// modify.Transaction,
		// revoke.Transaction,
		// define.Transaction,
		// make.Transaction,
		// take.Transaction,
		cancel.Transaction,
		get.Transaction,
		put.Transaction,
	)
}
