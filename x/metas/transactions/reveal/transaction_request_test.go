// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package reveal

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/client"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/constants"
)

func Test_Reveal_Request(t *testing.T) {
	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.Data})

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	data := "S|newData"
	commonTransactionRequest := helpers.PrototypeCommonTransactionRequest()
	testTransactionRequest := newTransactionRequest(commonTransactionRequest, data)

	require.Equal(t, transactionRequest{CommonTransactionRequest: commonTransactionRequest, Data: data}, testTransactionRequest)
	require.Equal(t, nil, testTransactionRequest.Validate())

	requestFromCLI, err := transactionRequest{}.FromCLI(cliCommand, client.Context{}.WithCodec(baseHelpers.CodecPrototype()))
	require.Equal(t, nil, err)
	require.Equal(t, transactionRequest{CommonTransactionRequest: helpers.PrototypeCommonTransactionRequest(), Data: ""}, requestFromCLI)

	jsonMessage, _ := json.Marshal(testTransactionRequest)
	transactionRequestUnmarshalled, err := transactionRequest{}.FromJSON(jsonMessage)
	require.Equal(t, nil, err)
	require.Equal(t, testTransactionRequest, transactionRequestUnmarshalled)

	randomUnmarshall, err := transactionRequest{}.FromJSON([]byte{})
	require.Equal(t, nil, randomUnmarshall)
	require.NotNil(t, err)

	require.Equal(t, commonTransactionRequest, testTransactionRequest.GetCommonTransactionRequest())

	require.Nil(t, err)

	msg2, err := newTransactionRequest(helpers.PrototypeCommonTransactionRequest(), data).MakeMsg()
	require.NotNil(t, err)
	require.Nil(t, msg2)

	msg2, err = newTransactionRequest(helpers.PrototypeCommonTransactionRequest().SetFrom(fromAddress), "randomString").MakeMsg()
	require.NotNil(t, err)
	require.Nil(t, msg2)

	require.Equal(t, transactionRequest{}, requestPrototype())
	require.NotPanics(t, func() {
		requestPrototype().RegisterLegacyAminoCodec(codec.NewLegacyAmino())
	})
}
