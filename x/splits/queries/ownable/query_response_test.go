// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package ownable

import (
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_Split_Response(t *testing.T) {

	testQueryResponse := newQueryResponse(sdkTypes.OneInt())

	encodedResponse, _ := testQueryResponse.Encode()
	bytes, _ := baseHelpers.CodecPrototype().GetLegacyAmino().MarshalJSON(testQueryResponse)
	require.Equal(t, bytes, encodedResponse)

	decodedResponse, _ := (&QueryResponse{}).Decode(bytes)
	require.Equal(t, testQueryResponse, decodedResponse)

	decodedResponse2, _ := (&QueryResponse{}).Decode([]byte{})
	require.Equal(t, nil, decodedResponse2)

	require.Equal(t, &QueryResponse{}, responsePrototype())
}
