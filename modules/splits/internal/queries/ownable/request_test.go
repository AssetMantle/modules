// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package ownable

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/modules/splits/internal/common"
	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

func Test_Split_Request(t *testing.T) {
	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()

	testSplitID := baseIDs.NewStringID("OwnableID")
	testQueryRequest := newQueryRequest(testSplitID)
	require.Equal(t, nil, testQueryRequest.Validate())
	require.Equal(t, queryRequest{}, requestPrototype())

	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.SplitID})
	cliContext := context.NewCLIContext().WithCodec(Codec)
	require.Panics(t, func() {
		require.Equal(t, newQueryRequest(baseIDs.NewStringID("")), queryRequest{}.FromCLI(cliCommand, cliContext))
	})

	vars := make(map[string]string)
	vars["ownables"] = "randomString"
	require.Equal(t, newQueryRequest(baseIDs.NewStringID("randomString")), queryRequest{}.FromMap(vars))

	encodedRequest, err := testQueryRequest.Encode()
	encodedResult, _ := common.Codec.MarshalJSON(testQueryRequest)
	require.Equal(t, encodedResult, encodedRequest)
	require.Nil(t, err)

	decodedRequest, err := queryRequest{}.Decode(encodedRequest)
	require.Equal(t, testQueryRequest, decodedRequest)
	require.Equal(t, nil, err)

	randomDecode, _ := queryRequest{}.Decode(baseIDs.NewStringID("").Bytes())
	require.Equal(t, nil, randomDecode)
	require.Equal(t, testQueryRequest, queryRequestFromInterface(testQueryRequest))
	require.Equal(t, queryRequest{}, queryRequestFromInterface(nil))
}
