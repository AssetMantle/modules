// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transactions

import (
	"testing"

	baseHelpers "github.com/AssetMantle/schema/x/helpers/base"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/x/assets/internal/transactions/burn"
	"github.com/AssetMantle/modules/x/assets/internal/transactions/define"
	"github.com/AssetMantle/modules/x/assets/internal/transactions/deputize"
	"github.com/AssetMantle/modules/x/assets/internal/transactions/mint"
	"github.com/AssetMantle/modules/x/assets/internal/transactions/mutate"
	"github.com/AssetMantle/modules/x/assets/internal/transactions/renumerate"
	"github.com/AssetMantle/modules/x/assets/internal/transactions/revoke"
)

func TestPrototype(t *testing.T) {
	want := baseHelpers.NewTransactions(burn.Transaction,
		define.Transaction,
		deputize.Transaction,
		mint.Transaction,
		mutate.Transaction,
		renumerate.Transaction,
		revoke.Transaction)

	require.Equal(t, Prototype().Get(""), want.Get(""))

}
