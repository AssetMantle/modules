/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package asset

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cryptoCodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	vestingTypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/common"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/key"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	baseHelpers "github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Asset_Request(t *testing.T) {

	var Codec = codec.NewLegacyAmino()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterLegacyAminoCodec(Codec)
	cryptoCodec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vestingTypes.RegisterLegacyAminoCodec(Codec)
	Codec.Seal()
	classificationID := base.NewID("classificationID")
	immutableProperties := base.NewProperties(base.NewProperty(base.NewID("ID1"), base.NewFact(base.NewStringData("ImmutableData"))))

	testAssetID := key.NewAssetID(classificationID, immutableProperties)
	testQueryRequest := newQueryRequest(testAssetID)
	require.Equal(t, nil, testQueryRequest.Validate())
	require.Equal(t, queryRequest{}, requestPrototype())

	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{flags.AssetID})
	cliContext := client.Context{}.WithLegacyAmino(Codec)
	require.Equal(t, newQueryRequest(base.NewID("")), queryRequest{}.FromCLI(cliCommand, cliContext))

	vars := make(map[string]string)
	vars["assets"] = "randomString"
	require.Equal(t, newQueryRequest(base.NewID("randomString")), queryRequest{}.FromMap(vars))

	encodedRequest, Error := testQueryRequest.Encode()
	encodedResult, _ := common.Codec.MarshalJSON(testQueryRequest)
	require.Equal(t, encodedResult, encodedRequest)
	require.Nil(t, Error)

	decodedRequest, Error := queryRequest{}.Decode(encodedRequest)
	require.Equal(t, testQueryRequest, decodedRequest)
	require.Equal(t, nil, Error)

	randomDecode, _ := queryRequest{}.Decode(base.NewID("").Bytes())
	require.Equal(t, nil, randomDecode)
	require.Equal(t, testQueryRequest, queryRequestFromInterface(testQueryRequest))
	require.Equal(t, queryRequest{}, queryRequestFromInterface(nil))
}
