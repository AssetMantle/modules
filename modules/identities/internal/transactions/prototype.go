// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transactions

import (
	"github.com/AssetMantle/modules/modules/identities/internal/transactions/define"
	"github.com/AssetMantle/modules/modules/identities/internal/transactions/deputize"
	"github.com/AssetMantle/modules/modules/identities/internal/transactions/issue"
	"github.com/AssetMantle/modules/modules/identities/internal/transactions/mutate"
	"github.com/AssetMantle/modules/modules/identities/internal/transactions/nub"
	"github.com/AssetMantle/modules/modules/identities/internal/transactions/provision"
	"github.com/AssetMantle/modules/modules/identities/internal/transactions/quash"
	"github.com/AssetMantle/modules/modules/identities/internal/transactions/revoke"
	"github.com/AssetMantle/modules/modules/identities/internal/transactions/unprovision"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/helpers/base"
)

func Prototype() helpers.Transactions {
	return base.NewTransactions(
		define.Transaction,
		deputize.Transaction,
		issue.Transaction,
		mutate.Transaction,
		nub.Transaction,
		provision.Transaction,
		quash.Transaction,
		revoke.Transaction,
		unprovision.Transaction,
	)
}
