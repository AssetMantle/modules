// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package verify

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/schema/errors/constants"
)

func Test_Verify_Response(t *testing.T) {
	testAuxiliaryResponse := newAuxiliaryResponse(nil)
	require.Equal(t, auxiliaryResponse{Success: true, Error: nil}, testAuxiliaryResponse)
	require.Equal(t, true, testAuxiliaryResponse.IsSuccessful())
	require.Equal(t, nil, testAuxiliaryResponse.GetError())

	testAuxiliaryResponse2 := newAuxiliaryResponse(constants.IncorrectFormat)
	require.Equal(t, auxiliaryResponse{Success: false, Error: constants.IncorrectFormat}, testAuxiliaryResponse2)
	require.Equal(t, false, testAuxiliaryResponse2.IsSuccessful())
	require.Equal(t, constants.IncorrectFormat, testAuxiliaryResponse2.GetError())
}
