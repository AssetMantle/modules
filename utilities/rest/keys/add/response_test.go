// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package add

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/schema/errors/constants"
)

func Test_Add_Response(t *testing.T) {
	testKeyOutput := keys.NewKeyOutput("name", "keyType", "address", "pubkey")
	testResponse := newResponse(testKeyOutput, nil)
	require.Equal(t, response{Success: true, Error: nil, KeyOutput: testKeyOutput}, testResponse)
	require.Equal(t, true, testResponse.IsSuccessful())
	require.Equal(t, nil, testResponse.GetError())
	testResponse2 := newResponse(testKeyOutput, constants.IncorrectFormat)
	require.Equal(t, false, testResponse2.IsSuccessful())
	require.Equal(t, constants.IncorrectFormat, testResponse2.GetError())
}
