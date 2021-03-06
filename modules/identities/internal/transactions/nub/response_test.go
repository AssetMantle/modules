/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package nub

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/constants/errors"
)

func Test_Nub_Response(t *testing.T) {
	testTransactionResponse := newTransactionResponse(errors.IncorrectFormat)
	testTransactionResponse2 := newTransactionResponse(nil)

	require.Equal(t, transactionResponse{Success: false, Error: errors.IncorrectFormat}, testTransactionResponse)
	require.Equal(t, false, testTransactionResponse.IsSuccessful())
	require.Equal(t, true, testTransactionResponse2.IsSuccessful())

	require.Equal(t, errors.IncorrectFormat, testTransactionResponse.GetError())
	require.Equal(t, nil, testTransactionResponse2.GetError())
}
