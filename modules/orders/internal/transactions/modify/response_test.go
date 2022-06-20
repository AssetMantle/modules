// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package modify

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/schema/helpers/constants"
)

func Test_Make_Response(t *testing.T) {

	testTransactionResponse := newTransactionResponse(constants.IncorrectFormat)
	testTransactionResponse2 := newTransactionResponse(nil)

	require.Equal(t, transactionResponse{Success: false, Error: constants.IncorrectFormat}, testTransactionResponse)
	require.Equal(t, false, testTransactionResponse.IsSuccessful())
	require.Equal(t, true, testTransactionResponse2.IsSuccessful())

	require.Equal(t, constants.IncorrectFormat, testTransactionResponse.GetError())
	require.Equal(t, nil, testTransactionResponse2.GetError())
}
