// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transactions

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/orders/transactions/cancel"
	"github.com/AssetMantle/modules/x/orders/transactions/define"
	make2 "github.com/AssetMantle/modules/x/orders/transactions/make"
	"github.com/AssetMantle/modules/x/orders/transactions/take"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, Prototype().GetTransaction("cancel").GetName(), baseHelpers.NewTransactions(
		take.Transaction,
		cancel.Transaction,
		define.Transaction,
		make2.Transaction,
	).GetTransaction("cancel").GetName())
	require.Equal(t, Prototype().GetTransaction("define").GetName(), baseHelpers.NewTransactions(
		take.Transaction,
		cancel.Transaction,
		define.Transaction,
		make2.Transaction,
	).GetTransaction("define").GetName())
	require.Equal(t, Prototype().GetTransaction("make").GetName(), baseHelpers.NewTransactions(
		take.Transaction,
		cancel.Transaction,
		define.Transaction,
		make2.Transaction,
	).GetTransaction("make").GetName())
	require.Equal(t, Prototype().GetTransaction("take").GetName(), baseHelpers.NewTransactions(
		take.Transaction,
		cancel.Transaction,
		define.Transaction,
		make2.Transaction,
	).GetTransaction("take").GetName())

}
