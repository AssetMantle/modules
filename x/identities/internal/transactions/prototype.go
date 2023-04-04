// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transactions

import (
	"github.com/AssetMantle/schema/x/helpers"
	baseHelpers "github.com/AssetMantle/schema/x/helpers/base"

	"github.com/AssetMantle/modules/x/identities/internal/transactions/define"
	"github.com/AssetMantle/modules/x/identities/internal/transactions/deputize"
	"github.com/AssetMantle/modules/x/identities/internal/transactions/issue"
	"github.com/AssetMantle/modules/x/identities/internal/transactions/mutate"
	"github.com/AssetMantle/modules/x/identities/internal/transactions/nub"
	"github.com/AssetMantle/modules/x/identities/internal/transactions/provision"
	"github.com/AssetMantle/modules/x/identities/internal/transactions/quash"
	"github.com/AssetMantle/modules/x/identities/internal/transactions/revoke"
	"github.com/AssetMantle/modules/x/identities/internal/transactions/unprovision"
)

func Prototype() helpers.Transactions {
	return baseHelpers.NewTransactions(
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
