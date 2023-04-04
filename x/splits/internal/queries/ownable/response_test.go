// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package ownable

import (
	"testing"

	errorConstants "github.com/AssetMantle/schema/x/errors/constants"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/x/splits/internal/common"
)

func Test_Split_Response(t *testing.T) {
	split := sdkTypes.SmallestDec()

	testQueryResponse := newQueryResponse(split, nil)
	testQueryResponseWithError := newQueryResponse(split, errorConstants.IncorrectFormat)

	require.Equal(t, true, testQueryResponse.IsSuccessful())
	require.Equal(t, false, testQueryResponseWithError.IsSuccessful())
	require.Equal(t, nil, testQueryResponse.GetError())
	require.Equal(t, errorConstants.IncorrectFormat, testQueryResponseWithError.GetError())

	encodedResponse, _ := testQueryResponse.Encode()
	bytes, _ := common.LegacyAmino.MarshalJSON(testQueryResponse)
	require.Equal(t, bytes, encodedResponse)

	decodedResponse, _ := (&QueryResponse{}).Decode(bytes)
	require.Equal(t, testQueryResponse, decodedResponse)

	decodedResponse2, _ := (&QueryResponse{}).Decode([]byte{})
	require.Equal(t, nil, decodedResponse2)

	require.Equal(t, &QueryResponse{}, responsePrototype())
}
