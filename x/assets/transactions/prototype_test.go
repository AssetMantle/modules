// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transactions

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/assets/transactions/burn"
	"github.com/AssetMantle/modules/x/assets/transactions/define"
	"github.com/AssetMantle/modules/x/assets/transactions/deputize"
	"github.com/AssetMantle/modules/x/assets/transactions/mint"
	"github.com/AssetMantle/modules/x/assets/transactions/mutate"
	"github.com/AssetMantle/modules/x/assets/transactions/renumerate"
	"github.com/AssetMantle/modules/x/assets/transactions/revoke"
	"github.com/AssetMantle/modules/x/assets/transactions/send"
	"github.com/AssetMantle/modules/x/assets/transactions/unwrap"
	"github.com/AssetMantle/modules/x/assets/transactions/wrap"
)

func TestPrototype(t *testing.T) {
	want := baseHelpers.NewTransactions(burn.Transaction,
		define.Transaction,
		deputize.Transaction,
		mint.Transaction,
		mutate.Transaction,
		renumerate.Transaction,
		revoke.Transaction,
		send.Transaction,
		unwrap.Transaction,
		wrap.Transaction,
	)

	require.Equal(t, Prototype().GetTransaction(""), want.GetTransaction(""))

}
