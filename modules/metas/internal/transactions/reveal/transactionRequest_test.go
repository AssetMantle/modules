// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package reveal

import (
	"encoding/json"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/std"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/data/utilities"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
)

func Test_Reveal_Request(t *testing.T) {
	var legacyAmino = codec.NewLegacyAmino()
	schema.RegisterLegacyAminoCodec(legacyAmino)
	std.RegisterLegacyAminoCodec(legacyAmino)
	legacyAmino.Seal()

	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.Data})

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	data := "S|newData"
	newData, err := utilities.ReadData(data)
	require.Equal(t, nil, err)

	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test", Fees: sdkTypes.NewCoins()}
	testTransactionRequest := newTransactionRequest(testBaseReq, data)

	require.Equal(t, transactionRequest{BaseReq: testBaseReq, Data: data}, testTransactionRequest)
	require.Equal(t, nil, testTransactionRequest.Validate())

	requestFromCLI, err := transactionRequest{}.FromCLI(cliCommand, baseHelpers.TestClientContext)
	require.Equal(t, nil, err)
	require.Equal(t, transactionRequest{BaseReq: rest.BaseReq{From: baseHelpers.TestClientContext.GetFromAddress().String(), ChainID: baseHelpers.TestClientContext.ChainID, Simulate: baseHelpers.TestClientContext.Simulate}, Data: ""}, requestFromCLI)

	jsonMessage, _ := json.Marshal(testTransactionRequest)
	transactionRequestUnmarshalled, err := transactionRequest{}.FromJSON(jsonMessage)
	require.Equal(t, nil, err)
	require.Equal(t, testTransactionRequest, transactionRequestUnmarshalled)

	randomUnmarshall, err := transactionRequest{}.FromJSON([]byte{})
	require.Equal(t, nil, randomUnmarshall)
	require.NotNil(t, err)

	require.Equal(t, testBaseReq, testTransactionRequest.GetBaseReq())

	msg, err := testTransactionRequest.MakeMsg()
	require.Equal(t, newMessage(fromAccAddress, newData), msg)
	require.Nil(t, err)

	msg2, err := newTransactionRequest(rest.BaseReq{From: "randomFromAddress", ChainID: "test"}, data).MakeMsg()
	require.NotNil(t, err)
	require.Nil(t, msg2)

	msg2, err = newTransactionRequest(rest.BaseReq{From: fromAddress, ChainID: "test"}, "randomString").MakeMsg()
	require.NotNil(t, err)
	require.Nil(t, msg2)

	require.Equal(t, transactionRequest{}, requestPrototype())
	require.NotPanics(t, func() {
		requestPrototype().RegisterLegacyAminoCodec(codec.NewLegacyAmino())
	})
}
