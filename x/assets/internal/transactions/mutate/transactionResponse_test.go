// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mutate

import (
	"testing"

	"github.com/stretchr/testify/require"

	errorConstants "github.com/AssetMantle/schema/x/errors/constants"
)

func Test_Mutate_Response(t *testing.T) {
	testTransactionResponse := newTransactionResponse(errorConstants.IncorrectFormat)
	testTransactionResponse2 := newTransactionResponse(nil)

	require.Equal(t, transactionResponse{Success: false, Error: errorConstants.IncorrectFormat}, testTransactionResponse)
	require.Equal(t, false, testTransactionResponse.IsSuccessful())
	require.Equal(t, true, testTransactionResponse2.IsSuccessful())

	require.Equal(t, errorConstants.IncorrectFormat, testTransactionResponse.GetError())
	require.Equal(t, nil, testTransactionResponse2.GetError())
}
