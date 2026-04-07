// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transactions

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/orders/transactions/cancel"
	"github.com/AssetMantle/modules/x/orders/transactions/get"
	"github.com/AssetMantle/modules/x/orders/transactions/put"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, Prototype().GetTransaction(cancel.Transaction.GetServicePath()).GetServicePath(), baseHelpers.NewTransactions(
		cancel.Transaction,
		get.Transaction,
		put.Transaction,
	).GetTransaction(cancel.Transaction.GetServicePath()).GetServicePath())
	require.Equal(t, Prototype().GetTransaction(get.Transaction.GetServicePath()).GetServicePath(), baseHelpers.NewTransactions(
		cancel.Transaction,
		get.Transaction,
		put.Transaction,
	).GetTransaction(get.Transaction.GetServicePath()).GetServicePath())
	require.Equal(t, Prototype().GetTransaction(put.Transaction.GetServicePath()).GetServicePath(), baseHelpers.NewTransactions(
		cancel.Transaction,
		get.Transaction,
		put.Transaction,
	).GetTransaction(put.Transaction.GetServicePath()).GetServicePath())
}
