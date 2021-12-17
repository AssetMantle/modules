/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package deputize

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

func Test_Deputize_Request(t *testing.T) {
	var Codec = codec.New()

	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()

	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{flags.FromID, flags.ToID, flags.ClassificationID, flags.MaintainedProperties, flags.AddMaintainer, flags.RemoveMaintainer, flags.MutateMaintainer})
	cliContext := context.NewCLIContext().WithCodec(Codec)

	const fromAddress = "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"

	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	const maintainedProperty = "maintainedProperties:S|maintainedProperties"

	var maintainedProperties types.Properties
	maintainedProperties, err = base.ReadProperties(maintainedProperty)
	require.Equal(t, nil, err)

	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test", Fees: sdkTypes.NewCoins()}
	testTransactionRequest := newTransactionRequest(testBaseReq, "fromID", "toID", "classificationID", maintainedProperty, false, false, false)

	require.Equal(t, transactionRequest{BaseReq: testBaseReq, FromID: "fromID", ToID: "toID", ClassificationID: "classificationID", MaintainedProperties: maintainedProperty, AddMaintainer: false, RemoveMaintainer: false, MutateMaintainer: false}, testTransactionRequest)
	require.Equal(t, nil, testTransactionRequest.Validate())

	var requestFromCLI helpers.TransactionRequest
	requestFromCLI, err = transactionRequest{}.FromCLI(cliCommand, cliContext)
	require.Equal(t, nil, err)
	require.Equal(t, transactionRequest{BaseReq: rest.BaseReq{From: cliContext.GetFromAddress().String(), ChainID: cliContext.ChainID, Simulate: cliContext.Simulate}, FromID: "", ToID: "", ClassificationID: "", MaintainedProperties: "", AddMaintainer: false, RemoveMaintainer: false, MutateMaintainer: false}, requestFromCLI)

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
	require.Equal(t, newMessage(fromAccAddress, base.NewID("fromID"), base.NewID("toID"), base.NewID("classificationID"), maintainedProperties, false, false, false), msg)
	require.Nil(t, err)

	var msg2 sdkTypes.Msg
	msg2, err = newTransactionRequest(rest.BaseReq{From: "randomString", ChainID: "test", Fees: sdkTypes.NewCoins()}, "fromID", "toID", "classificationID", maintainedProperty, false, false, false).MakeMsg()
	require.NotNil(t, err)
	require.Nil(t, msg2)

	msg2, err = newTransactionRequest(testBaseReq, "fromID", "toID", "classificationID", "randomString", false, false, false).MakeMsg()
	require.NotNil(t, err)
	require.Nil(t, msg2)

	require.Equal(t, transactionRequest{}, requestPrototype())
	require.NotPanics(t, func() {
		requestPrototype().RegisterCodec(codec.New())
	})
}
