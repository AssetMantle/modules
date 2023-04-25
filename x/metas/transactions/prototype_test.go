// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transactions

import (
	"testing"

	"github.com/AssetMantle/modules/x/metas/transactions/reveal"

	"github.com/stretchr/testify/require"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, Prototype().Get("reveal").GetName(), baseHelpers.NewTransactions(
		reveal.Transaction,
	).Get("reveal").GetName())
}
