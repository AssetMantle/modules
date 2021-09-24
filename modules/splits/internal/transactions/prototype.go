/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package transactions

import (
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/transactions/send"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/transactions/unwrap"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/transactions/wrap"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

func Prototype() helpers.Transactions {
	return base.NewTransactions(
		send.Transaction,
		unwrap.Transaction,
		wrap.Transaction,
	)
}
