// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transactions

import (
	"testing"

	"github.com/AssetMantle/modules/x/orders/transactions/cancel"
	"github.com/AssetMantle/modules/x/orders/transactions/define"
	make2 "github.com/AssetMantle/modules/x/orders/transactions/make"
	"github.com/AssetMantle/modules/x/orders/transactions/take"

	"github.com/stretchr/testify/require"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, Prototype().Get("cancel").GetName(), baseHelpers.NewTransactions(
		take.Transaction,
		cancel.Transaction,
		define.Transaction,
		make2.Transaction,
	).Get("cancel").GetName())
	require.Equal(t, Prototype().Get("define").GetName(), baseHelpers.NewTransactions(
		take.Transaction,
		cancel.Transaction,
		define.Transaction,
		make2.Transaction,
	).Get("define").GetName())
	require.Equal(t, Prototype().Get("make").GetName(), baseHelpers.NewTransactions(
		take.Transaction,
		cancel.Transaction,
		define.Transaction,
		make2.Transaction,
	).Get("make").GetName())
	require.Equal(t, Prototype().Get("take").GetName(), baseHelpers.NewTransactions(
		take.Transaction,
		cancel.Transaction,
		define.Transaction,
		make2.Transaction,
	).Get("take").GetName())

}
