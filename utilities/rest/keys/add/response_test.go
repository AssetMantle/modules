/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package add

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Add_Response(t *testing.T) {

	testKeyOutput := keys.NewKeyOutput("name", "keyType", "address", "pubkey")
	testResponse := newResponse(testKeyOutput, nil)
	require.Equal(t, response{Success: true, Error: nil, KeyOutput: testKeyOutput}, testResponse)
	require.Equal(t, true, testResponse.IsSuccessful())
	require.Equal(t, nil, testResponse.GetError())
	testResponse2 := newResponse(testKeyOutput, errors.IncorrectFormat)
	require.Equal(t, false, testResponse2.IsSuccessful())
	require.Equal(t, errors.IncorrectFormat, testResponse2.GetError())
}
