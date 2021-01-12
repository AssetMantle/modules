/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package add

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Add_Request(t *testing.T) {
	require.Equal(t, nil, request{Name: "name", Mnemonic: "mnemonic"}.Validate())
}
