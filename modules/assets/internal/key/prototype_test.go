// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, Prototype(), assetIDFromInterface(base.NewID("")))
}
