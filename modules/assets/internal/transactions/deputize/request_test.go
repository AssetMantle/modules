/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package deputize

import (
	"encoding/json"
	cryptoCodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	vestingTypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	baseHelpers "github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
)

func Test_Deputize_Request(t *testing.T) {
	var Codec = codec.NewLegacyAmino()
	schema.RegisterLegacyAminoCodec(Codec)
	sdkTypes.RegisterLegacyAminoCodec(Codec)
	cryptocryptoCodec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vestingTypes.RegisterLegacyAminoCodec(Codec)
	Codec.Seal()
	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{flags.FromID, flags.ToID, flags.ClassificationID, flags.MaintainedProperties, flags.AddMaintainer, flags.RemoveMaintainer, flags.MutateMaintainer})
	cliContext := client.Context{}.WithLegacyAmino(Codec)

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, Error := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, Error)

	maintainedProperty := "maintainedProperties:S|maintainedProperties"
	maintainedProperties, Error := base.ReadProperties(maintainedProperty)
	require.Equal(t, nil, Error)

	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test", Fees: sdkTypes.NewCoins()}
	testTransactionRequest := newTransactionRequest(testBaseReq, "fromID", "toID", "classificationID", maintainedProperty, false, false, false)

	require.Equal(t, transactionRequest{BaseReq: testBaseReq, FromID: "fromID", ToID: "toID", ClassificationID: "classificationID", MaintainedProperties: maintainedProperty, AddMaintainer: false, RemoveMaintainer: false, MutateMaintainer: false}, testTransactionRequest)
	require.Equal(t, nil, testTransactionRequest.Validate())

	requestFromCLI, Error := transactionRequest{}.FromCLI(cliCommand, cliContext)
	require.Equal(t, nil, Error)
	require.Equal(t, transactionRequest{BaseReq: rest.BaseReq{From: cliContext.GetFromAddress().String(), ChainID: cliContext.ChainID, Simulate: cliContext.Simulate}, FromID: "", ToID: "", ClassificationID: "", MaintainedProperties: "", AddMaintainer: false, RemoveMaintainer: false, MutateMaintainer: false}, requestFromCLI)

	jsonMessage, _ := json.Marshal(testTransactionRequest)
	transactionRequestUnmarshalled, Error := transactionRequest{}.FromJSON(jsonMessage)
	require.Equal(t, nil, Error)
	require.Equal(t, testTransactionRequest, transactionRequestUnmarshalled)

	randomUnmarshall, Error := transactionRequest{}.FromJSON([]byte{})
	require.Equal(t, nil, randomUnmarshall)
	require.NotNil(t, Error)

	require.Equal(t, testBaseReq, testTransactionRequest.GetBaseReq())

	msg, Error := testTransactionRequest.MakeMsg()
	require.Equal(t, newMessage(fromAccAddress, base.NewID("fromID"), base.NewID("toID"), base.NewID("classificationID"), maintainedProperties, false, false, false), msg)
	require.Nil(t, Error)

	msg2, Error := newTransactionRequest(rest.BaseReq{From: "randomString", ChainID: "test", Fees: sdkTypes.NewCoins()}, "fromID", "toID", "classificationID", maintainedProperty, false, false, false).MakeMsg()
	require.NotNil(t, Error)
	require.Nil(t, msg2)

	msg2, Error = newTransactionRequest(testBaseReq, "fromID", "toID", "classificationID", "randomString", false, false, false).MakeMsg()
	require.NotNil(t, Error)
	require.Nil(t, msg2)

	require.Equal(t, transactionRequest{}, requestPrototype())
	require.NotPanics(t, func() {
		requestPrototype().RegisterLegacyAminoCodec(codec.NewLegacyAmino())
	})
}
