// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transactions

import (
	"github.com/AssetMantle/modules/modules/splits/internal/transactions/send"
	"github.com/AssetMantle/modules/modules/splits/internal/transactions/unwrap"
	"github.com/AssetMantle/modules/modules/splits/internal/transactions/wrap"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/helpers/base"
)

func Prototype() helpers.Transactions {
	return base.NewTransactions(
		send.Transaction,
		unwrap.Transaction,
		wrap.Transaction,
	)
}
