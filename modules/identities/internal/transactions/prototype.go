/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package transactions

import (
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/transactions/define"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/transactions/issue"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/transactions/nub"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/transactions/provision"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/transactions/unprovision"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

func Prototype() helpers.Transactions {
	return base.NewTransactions(
		define.Transaction,
		issue.Transaction,
		nub.Transaction,
		provision.Transaction,
		unprovision.Transaction,
	)
}
