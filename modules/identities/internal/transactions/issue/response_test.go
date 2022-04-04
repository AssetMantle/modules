// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package issue

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/constants/errors"
)

func Test_Issue_Response(t *testing.T) {
	testTransactionResponse := newTransactionResponse(errors.IncorrectFormat)
	testTransactionResponse2 := newTransactionResponse(nil)

	require.Equal(t, transactionResponse{Success: false, Error: errors.IncorrectFormat}, testTransactionResponse)
	require.Equal(t, false, testTransactionResponse.IsSuccessful())
	require.Equal(t, true, testTransactionResponse2.IsSuccessful())

	require.Equal(t, errors.IncorrectFormat, testTransactionResponse.GetError())
	require.Equal(t, nil, testTransactionResponse2.GetError())
}
