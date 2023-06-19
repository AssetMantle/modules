// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transactions

import (
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/identities/transactions/define"
	"github.com/AssetMantle/modules/x/identities/transactions/deputize"
	"github.com/AssetMantle/modules/x/identities/transactions/issue"
	"github.com/AssetMantle/modules/x/identities/transactions/nub"
	"github.com/AssetMantle/modules/x/identities/transactions/provision"
	"github.com/AssetMantle/modules/x/identities/transactions/quash"
	"github.com/AssetMantle/modules/x/identities/transactions/revoke"
	"github.com/AssetMantle/modules/x/identities/transactions/unprovision"
	"github.com/AssetMantle/modules/x/identities/transactions/update"
)

func Prototype() helpers.Transactions {
	return baseHelpers.NewTransactions(
		define.Transaction,
		deputize.Transaction,
		issue.Transaction,
		update.Transaction,
		nub.Transaction,
		provision.Transaction,
		quash.Transaction,
		revoke.Transaction,
		unprovision.Transaction,
	)
}
