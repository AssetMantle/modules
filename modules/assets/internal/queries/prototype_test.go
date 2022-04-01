// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queries

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/queries/asset"
)

func Test_Prototype(t *testing.T) {
	prototype := Prototype()
	require.Equal(t, asset.Query.GetName(), prototype.Get("assets").GetName())
}
