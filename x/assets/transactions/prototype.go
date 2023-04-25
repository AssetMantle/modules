// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transactions

import (
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/assets/transactions/burn"
	"github.com/AssetMantle/modules/x/assets/transactions/define"
	"github.com/AssetMantle/modules/x/assets/transactions/deputize"
	"github.com/AssetMantle/modules/x/assets/transactions/mint"
	"github.com/AssetMantle/modules/x/assets/transactions/mutate"
	"github.com/AssetMantle/modules/x/assets/transactions/renumerate"
	"github.com/AssetMantle/modules/x/assets/transactions/revoke"
)

func Prototype() helpers.Transactions {
	return baseHelpers.NewTransactions(
		burn.Transaction,
		define.Transaction,
		deputize.Transaction,
		mint.Transaction,
		mutate.Transaction,
		renumerate.Transaction,
		revoke.Transaction,
	)
}
