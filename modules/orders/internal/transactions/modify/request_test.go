/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package modify

import (
	"encoding/json"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	baseHelpers "github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
)

func Test_Define_Request(t *testing.T) {

	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()
	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{flags.FromID, flags.OrderID, flags.MakerOwnableSplit, flags.ExpiresIn, flags.TakerOwnableSplit, flags.MutableMetaProperties, flags.MutableProperties})
	cliContext := context.NewCLIContext().WithCodec(Codec)

	mutableMetaPropertiesString := "defaultMutableMeta1:S|defaultMutableMeta1"
	mutablePropertiesString := "defaultMutable1:S|defaultMutable1"

	mutableMetaProperties, Error := base.ReadMetaProperties(mutableMetaPropertiesString)
	require.Equal(t, nil, Error)
	mutableProperties, Error := base.ReadProperties(mutablePropertiesString)
	require.Equal(t, nil, Error)

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, Error := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, Error)

	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test", Fees: sdkTypes.NewCoins()}
	testTransactionRequest := newTransactionRequest(testBaseReq, "fromID", "orderID", sdkTypes.OneDec().String(), "1.0", 123, mutableMetaPropertiesString, mutablePropertiesString)

	require.Equal(t, transactionRequest{BaseReq: testBaseReq, FromID: "fromID", OrderID: "orderID", ExpiresIn: 123, MakerOwnableSplit: "1.0", TakerOwnableSplit: sdkTypes.OneDec().String(), MutableMetaProperties: mutableMetaPropertiesString, MutableProperties: mutablePropertiesString}, testTransactionRequest)
	require.Equal(t, nil, testTransactionRequest.Validate())

	requestFromCLI, Error := transactionRequest{}.FromCLI(cliCommand, cliContext)
	require.Equal(t, nil, Error)
	require.Equal(t, transactionRequest{BaseReq: rest.BaseReq{From: cliContext.GetFromAddress().String(), ChainID: cliContext.ChainID, Simulate: cliContext.Simulate}, FromID: "", MutableMetaProperties: "", MutableProperties: ""}, requestFromCLI)

	jsonMessage, _ := json.Marshal(testTransactionRequest)
	transactionRequestUnmarshalled, error3 := transactionRequest{}.FromJSON(jsonMessage)
	require.Equal(t, nil, error3)
	require.Equal(t, testTransactionRequest, transactionRequestUnmarshalled)

	randomUnmarshall, Error := transactionRequest{}.FromJSON([]byte{})
	require.Equal(t, nil, randomUnmarshall)
	require.NotNil(t, Error)

	require.Equal(t, testBaseReq, testTransactionRequest.GetBaseReq())

	msg, Error := testTransactionRequest.MakeMsg()
	require.Equal(t, newMessage(fromAccAddress, base.NewID("fromID"), base.NewID("orderID"), sdkTypes.OneDec(), sdkTypes.NewDec(1), base.NewHeight(123), mutableMetaProperties, mutableProperties), msg)
	require.Nil(t, Error)

	msg, Error = newTransactionRequest(rest.BaseReq{From: "fromAddress", ChainID: "test", Fees: sdkTypes.NewCoins()}, "fromID", "orderID", sdkTypes.OneDec().String(), "aa", 123, mutableMetaPropertiesString, mutablePropertiesString).MakeMsg()
	require.Equal(t, nil, msg)
	require.NotNil(t, Error)

	msg, Error = newTransactionRequest(testBaseReq, "fromID", "orderID", sdkTypes.OneDec().String(), "aa", 123, mutableMetaPropertiesString, mutablePropertiesString).MakeMsg()
	require.Equal(t, nil, msg)
	require.NotNil(t, Error)

	msg, Error = newTransactionRequest(testBaseReq, "fromID", "orderID", "aa", "2.0", 123, mutableMetaPropertiesString, mutablePropertiesString).MakeMsg()
	require.Equal(t, nil, msg)
	require.NotNil(t, Error)

	msg, Error = newTransactionRequest(rest.BaseReq{From: "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c", ChainID: "test", Fees: sdkTypes.NewCoins()}, "fromID", "orderID", sdkTypes.OneDec().String(), sdkTypes.OneDec().String(), 123, "randomString", mutableMetaPropertiesString).MakeMsg()
	require.Equal(t, nil, msg)
	require.NotNil(t, Error)

	msg, Error = newTransactionRequest(rest.BaseReq{From: "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c", ChainID: "test", Fees: sdkTypes.NewCoins()}, "fromID", "orderID", sdkTypes.OneDec().String(), sdkTypes.OneDec().String(), 123, mutablePropertiesString, "randomString").MakeMsg()
	require.Equal(t, nil, msg)
	require.NotNil(t, Error)

	require.Equal(t, transactionRequest{}, requestPrototype())
	require.NotPanics(t, func() {
		requestPrototype().RegisterCodec(codec.New())
	})
}
