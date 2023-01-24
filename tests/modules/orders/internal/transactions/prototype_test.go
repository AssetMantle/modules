// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transactions

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/modules/orders/internal/transactions/cancel"
	"github.com/AssetMantle/modules/modules/orders/internal/transactions/define"
	"github.com/AssetMantle/modules/modules/orders/internal/transactions/make"
	"github.com/AssetMantle/modules/modules/orders/internal/transactions/take"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, Prototype().Get("cancel").GetName(), baseHelpers.NewTransactions(
		take.Transaction,
		cancel.Transaction,
		define.Transaction,
		make.Transaction,
	).Get("cancel").GetName())
	require.Equal(t, Prototype().Get("define").GetName(), baseHelpers.NewTransactions(
		take.Transaction,
		cancel.Transaction,
		define.Transaction,
		make.Transaction,
	).Get("define").GetName())
	require.Equal(t, Prototype().Get("make").GetName(), baseHelpers.NewTransactions(
		take.Transaction,
		cancel.Transaction,
		define.Transaction,
		make.Transaction,
	).Get("make").GetName())
	require.Equal(t, Prototype().Get("take").GetName(), baseHelpers.NewTransactions(
		take.Transaction,
		cancel.Transaction,
		define.Transaction,
		make.Transaction,
	).Get("take").GetName())

}
