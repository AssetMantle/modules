// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transactions

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/identities/transactions/define"
	"github.com/AssetMantle/modules/x/identities/transactions/deputize"
	"github.com/AssetMantle/modules/x/identities/transactions/issue"
	"github.com/AssetMantle/modules/x/identities/transactions/name"
	"github.com/AssetMantle/modules/x/identities/transactions/provision"
	"github.com/AssetMantle/modules/x/identities/transactions/revoke"
	"github.com/AssetMantle/modules/x/identities/transactions/unprovision"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, Prototype().GetTransaction("unprovision").GetServicePath(), baseHelpers.NewTransactions(
		define.Transaction,
		deputize.Transaction,
		issue.Transaction,
		name.Transaction,
		provision.Transaction,
		revoke.Transaction,
		unprovision.Transaction,
	).GetTransaction("unprovision").GetServicePath())
}
