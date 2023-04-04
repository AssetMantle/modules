// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transactions

import (
	"github.com/AssetMantle/schema/x/helpers"
	baseHelpers "github.com/AssetMantle/schema/x/helpers/base"

	"github.com/AssetMantle/modules/modules/assets/internal/transactions/burn"
	"github.com/AssetMantle/modules/modules/assets/internal/transactions/define"
	"github.com/AssetMantle/modules/modules/assets/internal/transactions/deputize"
	"github.com/AssetMantle/modules/modules/assets/internal/transactions/mint"
	"github.com/AssetMantle/modules/modules/assets/internal/transactions/mutate"
	"github.com/AssetMantle/modules/modules/assets/internal/transactions/renumerate"
	"github.com/AssetMantle/modules/modules/assets/internal/transactions/revoke"
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
