/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package transactions

import (
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/transactions/deputize"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Transactions = base.NewTransactions(deputize.Transaction)
