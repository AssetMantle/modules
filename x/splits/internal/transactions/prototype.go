// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transactions

import (
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/splits/internal/transactions/send"
	"github.com/AssetMantle/modules/x/splits/internal/transactions/unwrap"
	"github.com/AssetMantle/modules/x/splits/internal/transactions/wrap"
)

func Prototype() helpers.Transactions {
	return baseHelpers.NewTransactions(
		send.Transaction,
		unwrap.Transaction,
		wrap.Transaction,
	)
}
