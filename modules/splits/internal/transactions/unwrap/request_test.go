package unwrap

import (
	"encoding/json"
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
	"testing"
)

func Test_Unwrap_Request(t *testing.T) {
	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()
	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{flags.FromID, flags.OwnableID, flags.Split})
	clicontext := context.NewCLIContext().WithCodec(Codec)

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, Error := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, Error)

	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test", Fees: sdkTypes.NewCoins()}
	testTransactionRequest := newTransactionRequest(testBaseReq, "fromID", "ownableID", "2")

	require.Equal(t, transactionRequest{BaseReq: testBaseReq, FromID: "fromID", OwnableID: "ownableID", Split: "2"}, testTransactionRequest)
	require.Equal(t, nil, testTransactionRequest.Validate())

	requestFromCLI, error := transactionRequest{}.FromCLI(cliCommand, clicontext)
	require.Equal(t, nil, error)
	require.Equal(t, transactionRequest{BaseReq: rest.BaseReq{From: clicontext.GetFromAddress().String(), ChainID: clicontext.ChainID, Simulate: clicontext.Simulate}, FromID: "", OwnableID: "", Split: ""}, requestFromCLI)

	jsonMessage, _ := json.Marshal(testTransactionRequest)
	transactionRequestUnmarshalled, error3 := transactionRequest{}.FromJSON(jsonMessage)
	require.Equal(t, nil, error3)
	require.Equal(t, testTransactionRequest, transactionRequestUnmarshalled)

	randomUnmarshall, error := transactionRequest{}.FromJSON([]byte{})
	require.Equal(t, nil, randomUnmarshall)
	require.NotNil(t, error)

	require.Equal(t, testBaseReq, testTransactionRequest.GetBaseReq())

	msg, error := testTransactionRequest.MakeMsg()
	require.Equal(t, newMessage(fromAccAddress, base.NewID("fromID"), base.NewID("ownableID"), sdkTypes.NewDec(2)), msg)
	require.Nil(t, error)

	msg2, error2 := newTransactionRequest(rest.BaseReq{From: "randomFromAddrs", ChainID: "test"}, "fromID", "ownableID", "2").MakeMsg()
	require.NotNil(t, error2)
	require.Nil(t, msg2)

	msg2, error2 = newTransactionRequest(testBaseReq, "fromID", "ownableID", "randomString").MakeMsg()
	require.NotNil(t, error2)
	require.Nil(t, msg2)

	require.Equal(t, transactionRequest{}, requestPrototype())
}
