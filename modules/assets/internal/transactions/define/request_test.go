// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	"encoding/json"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/constants/flags"
	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/base"
)

func Test_Define_Request(t *testing.T) {
	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()
	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{flags.FromID, flags.ImmutableMetaProperties, flags.ImmutableProperties, flags.MutableMetaProperties, flags.MutableProperties})
	cliContext := context.NewCLIContext().WithCodec(Codec)

	immutableMetaPropertiesString := "defaultImmutableMeta1:S|defaultImmutableMeta1"
	immutablePropertiesString := "defaultMutableMeta1:S|defaultMutableMeta1"
	mutableMetaPropertiesString := "defaultMutableMeta1:S|defaultMutableMeta1"
	mutablePropertiesString := "defaultMutable1:S|defaultMutable1"

	immutableMetaProperties, err := base.ReadMetaProperties(immutableMetaPropertiesString)
	require.Equal(t, nil, err)
	immutableProperties, err := base.ReadProperties(immutablePropertiesString)
	require.Equal(t, nil, err)
	mutableMetaProperties, err := base.ReadMetaProperties(mutableMetaPropertiesString)
	require.Equal(t, nil, err)
	mutableProperties, err := base.ReadProperties(mutablePropertiesString)
	require.Equal(t, nil, err)

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test", Fees: sdkTypes.NewCoins()}
	testTransactionRequest := newTransactionRequest(testBaseReq, "fromID", immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString)

	require.Equal(t, transactionRequest{BaseReq: testBaseReq, FromID: "fromID", ImmutableMetaProperties: immutableMetaPropertiesString, ImmutableProperties: immutablePropertiesString, MutableMetaProperties: mutableMetaPropertiesString, MutableProperties: mutablePropertiesString}, testTransactionRequest)
	require.Equal(t, nil, testTransactionRequest.Validate())

	requestFromCLI, err := transactionRequest{}.FromCLI(cliCommand, cliContext)
	require.Equal(t, nil, err)
	require.Equal(t, transactionRequest{BaseReq: rest.BaseReq{From: cliContext.GetFromAddress().String(), ChainID: cliContext.ChainID, Simulate: cliContext.Simulate}, FromID: "", ImmutableMetaProperties: "", ImmutableProperties: "", MutableMetaProperties: "", MutableProperties: ""}, requestFromCLI)

	jsonMessage, _ := json.Marshal(testTransactionRequest)
	transactionRequestUnmarshalled, err := transactionRequest{}.FromJSON(jsonMessage)
	require.Equal(t, nil, err)
	require.Equal(t, testTransactionRequest, transactionRequestUnmarshalled)

	randomUnmarshall, err := transactionRequest{}.FromJSON([]byte{})
	require.Equal(t, nil, randomUnmarshall)
	require.NotNil(t, err)

	require.Equal(t, testBaseReq, testTransactionRequest.GetBaseReq())

	msg, err := testTransactionRequest.MakeMsg()
	require.Equal(t, newMessage(fromAccAddress, baseIDs.NewID("fromID"), immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties), msg)
	require.Nil(t, err)

	msg, err = newTransactionRequest(rest.BaseReq{From: "randomFromAddress", ChainID: "test"}, "fromID", immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString).MakeMsg()
	require.Equal(t, nil, msg)
	require.NotNil(t, err)

	msg, err = newTransactionRequest(testBaseReq, "fromID", "randomString", immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString).MakeMsg()
	require.Equal(t, nil, msg)
	require.NotNil(t, err)

	msg, err = newTransactionRequest(testBaseReq, "fromID", immutableMetaPropertiesString, "randomString", mutableMetaPropertiesString, mutablePropertiesString).MakeMsg()
	require.Equal(t, nil, msg)
	require.NotNil(t, err)

	msg, err = newTransactionRequest(testBaseReq, "fromID", immutableMetaPropertiesString, immutablePropertiesString, "randomString", mutablePropertiesString).MakeMsg()
	require.Equal(t, nil, msg)
	require.NotNil(t, err)

	msg, err = newTransactionRequest(testBaseReq, "fromID", immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, "randomString").MakeMsg()
	require.Equal(t, nil, msg)
	require.NotNil(t, err)

	require.Equal(t, transactionRequest{}, requestPrototype())
	require.NotPanics(t, func() {
		requestPrototype().RegisterCodec(codec.New())
	})

}
