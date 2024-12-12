// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transactions

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/metas/transactions/reveal"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, Prototype().GetTransaction("reveal").GetServicePath(), baseHelpers.NewTransactions(
		reveal.Transaction,
	).GetTransaction("reveal").GetServicePath())
}
