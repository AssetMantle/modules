/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package auxiliaries

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

func Test_Auxiliary_Prototype(t *testing.T) {
	require.Equal(t, base.NewAuxiliaries(), Prototype())
}
