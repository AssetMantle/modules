/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package transactions

import (
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/transactions/reveal"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Transactions = base.NewTransactions(reveal.Transaction)
