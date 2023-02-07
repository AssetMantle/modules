// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transactions

import (
	"github.com/AssetMantle/modules/modules/orders/internal/transactions/cancel"
	"github.com/AssetMantle/modules/modules/orders/internal/transactions/define"
	"github.com/AssetMantle/modules/modules/orders/internal/transactions/make"
	"github.com/AssetMantle/modules/modules/orders/internal/transactions/take"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
)

func Prototype() helpers.Transactions {
	return baseHelpers.NewTransactions(
		// TODO ***** revisit business logic
		// deputize.Transaction,
		// immediate.Transaction,
		// modify.Transaction,
		// revoke.Transaction,
		take.Transaction,
		cancel.Transaction,
		define.Transaction,
		make.Transaction,
	)
}
