/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

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
