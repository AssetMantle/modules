/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package transactions

import (
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/transactions/cancel"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/transactions/define"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/transactions/make"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/transactions/take"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Transactions = base.NewTransactions(cancel.Transaction, define.Transaction, make.Transaction, take.Transaction)
