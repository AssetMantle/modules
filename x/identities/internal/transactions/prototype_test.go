// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transactions

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/identities/internal/transactions/define"
	"github.com/AssetMantle/modules/x/identities/internal/transactions/deputize"
	"github.com/AssetMantle/modules/x/identities/internal/transactions/issue"
	"github.com/AssetMantle/modules/x/identities/internal/transactions/nub"
	"github.com/AssetMantle/modules/x/identities/internal/transactions/provision"
	"github.com/AssetMantle/modules/x/identities/internal/transactions/revoke"
	"github.com/AssetMantle/modules/x/identities/internal/transactions/unprovision"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, Prototype().Get("unprovision").GetName(), baseHelpers.NewTransactions(
		define.Transaction,
		deputize.Transaction,
		issue.Transaction,
		nub.Transaction,
		provision.Transaction,
		revoke.Transaction,
		unprovision.Transaction,
	).Get("unprovision").GetName())
}
