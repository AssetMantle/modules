/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package classification

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/common"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	baseHelpers "github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Classification_Request(t *testing.T) {

	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()

	testClassificationID := base.NewID("ClassificationID")
	testQueryRequest := newQueryRequest(testClassificationID)
	require.Equal(t, nil, testQueryRequest.Validate())
	require.Equal(t, queryRequest{}, requestPrototype())

	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{flags.ClassificationID})
	cliContext := context.NewCLIContext().WithCodec(Codec)
	require.Equal(t, newQueryRequest(base.NewID("")), queryRequest{}.FromCLI(cliCommand, cliContext))

	vars := make(map[string]string)
	vars["classifications"] = "randomString"
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
