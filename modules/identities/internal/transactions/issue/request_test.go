/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package issue

import (
	"encoding/json"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	baseHelpers "github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func Test_Mint_Request(t *testing.T) {
	var Codec = codec.New()

	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()

	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{flags.FromID, flags.To, flags.ClassificationID, flags.ImmutableMetaProperties, flags.ImmutableProperties, flags.MutableMetaProperties, flags.MutableProperties})
	cliContext := context.NewCLIContext().WithCodec(Codec)

	immutableMetaPropertiesString := "defaultImmutableMeta1:S|defaultImmutableMeta1"
	immutablePropertiesString := "defaultMutableMeta1:S|defaultMutableMeta1"
	mutableMetaPropertiesString := "defaultMutableMeta1:S|defaultMutableMeta1"
	mutablePropertiesString := "defaultMutable1:S|defaultMutable1"

	immutableMetaProperties, err := base.ReadMetaProperties(immutableMetaPropertiesString)
	require.Equal(t, nil, err)

	var immutableProperties types.Properties
	immutableProperties, err = base.ReadProperties(immutablePropertiesString)
	require.Equal(t, nil, err)

	var mutableMetaProperties types.MetaProperties
	mutableMetaProperties, err = base.ReadMetaProperties(mutableMetaPropertiesString)
	require.Equal(t, nil, err)

	var mutableProperties types.Properties
	mutableProperties, err = base.ReadProperties(mutablePropertiesString)
	require.Equal(t, nil, err)

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	var fromAccAddress sdkTypes.AccAddress
	fromAccAddress, err = sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	toAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	var toAccAddress sdkTypes.AccAddress
	toAccAddress, err = sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test", Fees: sdkTypes.NewCoins()}
	testTransactionRequest := newTransactionRequest(testBaseReq, toAddress, "fromID", "classificationID", immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString)
	require.Equal(t, transactionRequest{BaseReq: testBaseReq, To: toAddress, FromID: "fromID", ClassificationID: "classificationID", ImmutableMetaProperties: immutableMetaPropertiesString, ImmutableProperties: immutablePropertiesString, MutableMetaProperties: mutableMetaPropertiesString, MutableProperties: mutablePropertiesString}, testTransactionRequest)
	require.Equal(t, nil, testTransactionRequest.Validate())

	var requestFromCLI helpers.TransactionRequest
	requestFromCLI, err = transactionRequest{}.FromCLI(cliCommand, cliContext)
	require.Equal(t, nil, err)
	require.Equal(t, transactionRequest{BaseReq: rest.BaseReq{From: cliContext.GetFromAddress().String(), ChainID: cliContext.ChainID, Simulate: cliContext.Simulate}, To: "", FromID: "", ClassificationID: "", ImmutableMetaProperties: "", ImmutableProperties: "", MutableMetaProperties: "", MutableProperties: ""}, requestFromCLI)

	var jsonMessage []byte
	jsonMessage, err = json.Marshal(testTransactionRequest)
	require.Equal(t, nil, err)

	var transactionRequestUnmarshalled helpers.TransactionRequest
	transactionRequestUnmarshalled, err = transactionRequest{}.FromJSON(jsonMessage)
	require.Equal(t, nil, err)
	require.Equal(t, testTransactionRequest, transactionRequestUnmarshalled)

	var randomUnmarshall helpers.TransactionRequest
	randomUnmarshall, err = transactionRequest{}.FromJSON([]byte{})
	require.Equal(t, nil, randomUnmarshall)
	require.NotNil(t, err)

	require.Equal(t, testBaseReq, testTransactionRequest.GetBaseReq())

	var msg sdkTypes.Msg
	msg, err = testTransactionRequest.MakeMsg()
	require.Equal(t, newMessage(fromAccAddress, toAccAddress, base.NewID("fromID"), base.NewID("classificationID"), immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties), msg)
	require.Nil(t, err)

	msg, err = newTransactionRequest(rest.BaseReq{From: "randomFromAddress", ChainID: "test"}, toAddress, "fromID", "classificationID", immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString).MakeMsg()
	require.Equal(t, nil, msg)
	require.NotNil(t, err)

	msg, err = newTransactionRequest(rest.BaseReq{From: fromAddress, ChainID: "test"}, "randomString", "fromID", "classificationID", immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString).MakeMsg()
	require.Equal(t, nil, msg)
	require.NotNil(t, err)

	msg, err = newTransactionRequest(testBaseReq, toAddress, "fromID", "classificationID", "randomString", immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString).MakeMsg()
	require.Equal(t, nil, msg)
	require.NotNil(t, err)

	msg, err = newTransactionRequest(testBaseReq, toAddress, "fromID", "classificationID", immutableMetaPropertiesString, "randomString", mutableMetaPropertiesString, mutablePropertiesString).MakeMsg()
	require.Equal(t, nil, msg)
	require.NotNil(t, err)

	msg, err = newTransactionRequest(testBaseReq, toAddress, "fromID", "classificationID", immutableMetaPropertiesString, immutablePropertiesString, "randomString", mutablePropertiesString).MakeMsg()
	require.Equal(t, nil, msg)
	require.NotNil(t, err)

	msg, err = newTransactionRequest(testBaseReq, toAddress, "fromID", "classificationID", immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, "randomString").MakeMsg()
	require.Equal(t, nil, msg)
	require.NotNil(t, err)

	require.Equal(t, transactionRequest{}, requestPrototype())
	require.NotPanics(t, func() {
		requestPrototype().RegisterCodec(codec.New())
	})
}
