/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package auxiliaries

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Auxiliaries_Prototype(t *testing.T) {
	require.Equal(t, base.NewAuxiliaries(), Prototype())
}
