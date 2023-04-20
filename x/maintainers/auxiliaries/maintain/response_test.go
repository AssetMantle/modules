// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package maintain

import (
	"testing"

	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	"github.com/stretchr/testify/require"
)

func Test_Maintain_Response(t *testing.T) {
	testAuxiliaryResponse := newAuxiliaryResponse(nil)
	require.Equal(t, auxiliaryResponse{Success: true, Error: nil}, testAuxiliaryResponse)
	require.Equal(t, true, testAuxiliaryResponse.IsSuccessful())
	require.Equal(t, nil, testAuxiliaryResponse.GetError())

	testAuxiliaryResponse2 := newAuxiliaryResponse(errorConstants.IncorrectFormat)
	require.Equal(t, auxiliaryResponse{Success: false, Error: errorConstants.IncorrectFormat}, testAuxiliaryResponse2)
	require.Equal(t, false, testAuxiliaryResponse2.IsSuccessful())
	require.Equal(t, errorConstants.IncorrectFormat, testAuxiliaryResponse2.GetError())
}