// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transactions

import (
	"testing"

	baseHelpers "github.com/AssetMantle/schema/x/helpers/base"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/modules/metas/internal/transactions/reveal"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, Prototype().Get("reveal").GetName(), baseHelpers.NewTransactions(
		reveal.Transaction,
	).Get("reveal").GetName())
}
