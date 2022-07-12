// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

func TestPrototype(t *testing.T) {
	assetID, err := assetIDFromInterface(baseIDs.NewStringID(""))
	require.Equal(t, Prototype(), assetID)
	require.Equal(t, nil, err)
}
