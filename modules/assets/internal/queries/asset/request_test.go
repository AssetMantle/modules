// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package asset

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/modules/assets/internal/common"
	"github.com/AssetMantle/modules/modules/assets/internal/key"
	"github.com/AssetMantle/modules/schema"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
)

func Test_Asset_Request(t *testing.T) {
	var Codec = codec.New()

	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()

	classificationID := baseIDs.NewID("classificationID")
	immutableProperties := base.NewPropertyList(baseProperties.NewProperty(baseIDs.NewID("ID1"), baseData.NewStringData("ImmutableData")))

	testAssetID := key.NewAssetID(classificationID, immutableProperties)
	testQueryRequest := newQueryRequest(testAssetID)
	require.Equal(t, nil, testQueryRequest.Validate())
	require.Equal(t, queryRequest{}, requestPrototype())

	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.AssetID})
	cliContext := context.NewCLIContext().WithCodec(Codec)
	require.Equal(t, newQueryRequest(baseIDs.NewID("")), queryRequest{}.FromCLI(cliCommand, cliContext))

	vars := make(map[string]string)
	vars["assets"] = "randomString"
	require.Equal(t, newQueryRequest(baseIDs.NewID("randomString")), queryRequest{}.FromMap(vars))

	encodedRequest, err := testQueryRequest.Encode()
	require.Nil(t, err)

	var encodedResult []byte
	encodedResult, err = common.Codec.MarshalJSON(testQueryRequest)
	require.Nil(t, err)
	require.Equal(t, encodedResult, encodedRequest)

	var decodedRequest helpers.QueryRequest
	decodedRequest, err = queryRequest{}.Decode(encodedRequest)
	require.Equal(t, nil, err)
	require.Equal(t, testQueryRequest, decodedRequest)

	var randomDecode helpers.QueryRequest
	// we expect to get an error here, so ignore it
	randomDecode, _ = queryRequest{}.Decode(baseIDs.NewID("").Bytes())
	require.Equal(t, nil, randomDecode)
	require.Equal(t, testQueryRequest, queryRequestFromInterface(testQueryRequest))
	require.Equal(t, queryRequest{}, queryRequestFromInterface(nil))
}
