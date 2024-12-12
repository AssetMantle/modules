// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
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
	require.Equal(t, Prototype().GetTransaction("cancel").GetServicePath(), baseHelpers.NewTransactions(
		take.Transaction,
		cancel.Transaction,
		define.Transaction,
		make2.Transaction,
	).GetTransaction("cancel").GetServicePath())
	require.Equal(t, Prototype().GetTransaction("define").GetServicePath(), baseHelpers.NewTransactions(
		take.Transaction,
		cancel.Transaction,
		define.Transaction,
		make2.Transaction,
	).GetTransaction("define").GetServicePath())
	require.Equal(t, Prototype().GetTransaction("make").GetServicePath(), baseHelpers.NewTransactions(
		take.Transaction,
		cancel.Transaction,
		define.Transaction,
		make2.Transaction,
	).GetTransaction("make").GetServicePath())
	require.Equal(t, Prototype().GetTransaction("take").GetServicePath(), baseHelpers.NewTransactions(
		take.Transaction,
		cancel.Transaction,
		define.Transaction,
		make2.Transaction,
	).GetTransaction("take").GetServicePath())

}
