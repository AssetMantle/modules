/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package sign

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_SignTx_Response(t *testing.T) {
	testFee := authTypes.NewStdFee(12, sdkTypes.NewCoins())

	testStdTx := authTypes.NewStdTx([]sdkTypes.Msg{}, testFee, []authTypes.StdSignature{}, "")
	require.Equal(t, response{Success: true, Error: nil, StdTx: testStdTx}, newResponse(testStdTx, nil))
	testResponse := newResponse(testStdTx, errors.IncorrectFormat)
	require.Equal(t, false, testResponse.IsSuccessful())
	require.Equal(t, errors.IncorrectFormat, testResponse.GetError())
}
