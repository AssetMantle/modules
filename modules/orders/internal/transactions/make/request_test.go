/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package make

import (
	"encoding/json"
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
	"testing"
)

func Test_Define_Request(t *testing.T) {

	var Codec = codec.NewLegacyAmino()
	schema.RegisterLegacyAminoCodec(Codec)
	sdkTypes.RegisterLegacyAminoCodec(Codec)
	cryptoCodec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()
	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{flags.FromID, flags.ClassificationID, flags.MakerOwnableSplit, flags.MakerOwnableID, flags.TakerOwnableID, flags.ExpiresIn, flags.TakerOwnableSplit, flags.ImmutableMetaProperties, flags.ImmutableProperties, flags.MutableMetaProperties, flags.MutableProperties})
	cliContext := context.NewCLIContext().WithCodec(Codec)

	immutableMetaPropertiesString := "defaultImmutableMeta1:S|defaultImmutableMeta1"
	immutablePropertiesString := "defaultMutableMeta1:S|defaultMutableMeta1"
	mutableMetaPropertiesString := "defaultMutableMeta1:S|defaultMutableMeta1"
	mutablePropertiesString := "defaultMutable1:S|defaultMutable1"

	immutableMetaProperties, Error := base.ReadMetaProperties(immutableMetaPropertiesString)
	require.Equal(t, nil, Error)
	immutableProperties, Error := base.ReadProperties(immutablePropertiesString)
	require.Equal(t, nil, Error)
	mutableMetaProperties, Error := base.ReadMetaProperties(mutableMetaPropertiesString)
	require.Equal(t, nil, Error)
	mutableProperties, Error := base.ReadProperties(mutablePropertiesString)
	require.Equal(t, nil, Error)

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, Error := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, Error)

	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test", Fees: sdkTypes.NewCoins()}
	testTransactionRequest := newTransactionRequest(testBaseReq, "fromID", "classificationID", "makerOwnableID", "takerOwnableID", 123, "2", sdkTypes.OneDec().String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString)

	require.Equal(t, transactionRequest{BaseReq: testBaseReq, FromID: "fromID", ClassificationID: "classificationID", MakerOwnableID: "makerOwnableID", TakerOwnableID: "takerOwnableID", ExpiresIn: 123, MakerOwnableSplit: "2", TakerOwnableSplit: sdkTypes.OneDec().String(), ImmutableMetaProperties: immutableMetaPropertiesString, ImmutableProperties: immutablePropertiesString, MutableMetaProperties: mutableMetaPropertiesString, MutableProperties: mutablePropertiesString}, testTransactionRequest)
	require.Equal(t, nil, testTransactionRequest.Validate())

	requestFromCLI, Error := transactionRequest{}.FromCLI(cliCommand, cliContext)
	require.Equal(t, nil, Error)
	require.Equal(t, transactionRequest{BaseReq: rest.BaseReq{From: cliContext.GetFromAddress().String(), ChainID: cliContext.ChainID, Simulate: cliContext.Simulate}, FromID: "", ImmutableMetaProperties: "", ImmutableProperties: "", MutableMetaProperties: "", MutableProperties: ""}, requestFromCLI)

	jsonMessage, _ := json.Marshal(testTransactionRequest)
	transactionRequestUnmarshalled, error3 := transactionRequest{}.FromJSON(jsonMessage)
	require.Equal(t, nil, error3)
	require.Equal(t, testTransactionRequest, transactionRequestUnmarshalled)

	randomUnmarshall, Error := transactionRequest{}.FromJSON([]byte{})
	require.Equal(t, nil, randomUnmarshall)
	require.NotNil(t, Error)

	require.Equal(t, testBaseReq, testTransactionRequest.GetBaseReq())

	msg, Error := testTransactionRequest.MakeMsg()
	require.Equal(t, newMessage(fromAccAddress, base.NewID("fromID"), base.NewID("classificationID"), base.NewID("makerOwnableID"), base.NewID("takerOwnableID"), base.NewHeight(123), sdkTypes.NewDec(2), sdkTypes.OneDec(), immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties), msg)
	require.Nil(t, Error)

	msg, Error = newTransactionRequest(rest.BaseReq{From: "randomFromAddress", ChainID: "test"}, "fromID", "classificationID", "makerOwnableID", "takerOwnableID", 123, "2", sdkTypes.OneDec().String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString).MakeMsg()
	require.Equal(t, nil, msg)
	require.NotNil(t, Error)

	msg, Error = newTransactionRequest(rest.BaseReq{From: fromAddress, ChainID: "test"}, "fromID", "classificationID", "makerOwnableID", "takerOwnableID", 123, "randomInput", sdkTypes.OneDec().String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString).MakeMsg()
	require.Equal(t, nil, msg)
	require.NotNil(t, Error)

	msg, Error = newTransactionRequest(rest.BaseReq{From: "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c", ChainID: "test", Fees: sdkTypes.NewCoins()}, "fromID", "classificationID", "makerOwnableID", "takerOwnableID", 123, "2", sdkTypes.OneDec().String(), "randomString", immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString).MakeMsg()
	require.Equal(t, nil, msg)
	require.NotNil(t, Error)

	msg, Error = newTransactionRequest(rest.BaseReq{From: "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c", ChainID: "test", Fees: sdkTypes.NewCoins()}, "fromID", "classificationID", "makerOwnableID", "takerOwnableID", 123, "2", sdkTypes.OneDec().String(), immutableMetaPropertiesString, "randomString", mutableMetaPropertiesString, mutablePropertiesString).MakeMsg()
	require.Equal(t, nil, msg)
	require.NotNil(t, Error)

	msg, Error = newTransactionRequest(rest.BaseReq{From: "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c", ChainID: "test", Fees: sdkTypes.NewCoins()}, "fromID", "classificationID", "makerOwnableID", "takerOwnableID", 123, "2", sdkTypes.OneDec().String(), immutableMetaPropertiesString, immutablePropertiesString, "randomString", mutablePropertiesString).MakeMsg()
	require.Equal(t, nil, msg)
	require.NotNil(t, Error)

	msg, Error = newTransactionRequest(rest.BaseReq{From: "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c", ChainID: "test", Fees: sdkTypes.NewCoins()}, "fromID", "classificationID", "makerOwnableID", "takerOwnableID", 123, "2", sdkTypes.OneDec().String(), immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, "randomString").MakeMsg()
	require.Equal(t, nil, msg)
	require.NotNil(t, Error)

	msg, Error = newTransactionRequest(rest.BaseReq{From: "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c", ChainID: "test", Fees: sdkTypes.NewCoins()}, "fromID", "classificationID", "makerOwnableID", "takerOwnableID", 123, "2", "test", immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, "randomString").MakeMsg()
	require.Equal(t, nil, msg)
	require.NotNil(t, Error)

	require.Equal(t, transactionRequest{}, requestPrototype())
	require.NotPanics(t, func() {
		requestPrototype().RegisterLegacyAminoCodec(codec.New())
	})
}
