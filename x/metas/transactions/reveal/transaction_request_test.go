// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package reveal

import (
	"encoding/json"
	"github.com/AssetMantle/modules/utilities/rest"
	"github.com/cosmos/cosmos-sdk/client"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/constants"
)

func Test_Reveal_Request(t *testing.T) {
	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.Data})

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	data := "S|newData"
	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test", Fees: sdkTypes.NewCoins()}
	testTransactionRequest := newTransactionRequest(testBaseReq, data)

	require.Equal(t, transactionRequest{BaseReq: testBaseReq, Data: data}, testTransactionRequest)
	require.Equal(t, nil, testTransactionRequest.Validate())

	requestFromCLI, err := transactionRequest{}.FromCLI(cliCommand, client.Context{}.WithCodec(baseHelpers.CodecPrototype()))
	require.Equal(t, nil, err)
	require.Equal(t, transactionRequest{BaseReq: rest.BaseReq{From: client.Context{}.WithCodec(baseHelpers.CodecPrototype()).GetFromAddress().String(), ChainID: client.Context{}.WithCodec(baseHelpers.CodecPrototype()).ChainID, Simulate: client.Context{}.WithCodec(baseHelpers.CodecPrototype()).Simulate}, Data: ""}, requestFromCLI)

	jsonMessage, _ := json.Marshal(testTransactionRequest)
	transactionRequestUnmarshalled, err := transactionRequest{}.FromJSON(jsonMessage)
	require.Equal(t, nil, err)
	require.Equal(t, testTransactionRequest, transactionRequestUnmarshalled)

	randomUnmarshall, err := transactionRequest{}.FromJSON([]byte{})
	require.Equal(t, nil, randomUnmarshall)
	require.NotNil(t, err)

	require.Equal(t, testBaseReq, testTransactionRequest.GetBaseReq())

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
