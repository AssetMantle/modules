package queuing

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Kafka_Types(t *testing.T) {

	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()
	//cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{flags.IdentityID, flags.To})
	cliContext := context.NewCLIContext().WithCodec(Codec)

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	//fromAccAddress, Error := sdkTypes.AccAddressFromBech32(fromAddress)
	//require.Nil(t, Error)
	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test", Fees: sdkTypes.NewCoins()}

	//testAssetID := base.NewID("assetID")
	//testFromID := base.NewID("fromID")

	testMessage := sdkTypes.NewTestMsg()

	//testMsg:=newMessage(burn)
	ticketID := TicketIDGenerator("name")
	testKafkaMsg := NewKafkaMsgFromRest(testMessage, ticketID, testBaseReq, cliContext)
	kafkaCli := KafkaCliCtx{
		OutputFormat:  cliContext.OutputFormat,
		ChainID:       cliContext.ChainID,
		Height:        cliContext.Height,
		HomeDir:       cliContext.HomeDir,
		NodeURI:       cliContext.NodeURI,
		From:          cliContext.From,
		TrustNode:     cliContext.TrustNode,
		UseLedger:     cliContext.UseLedger,
		BroadcastMode: cliContext.BroadcastMode,
		Simulate:      cliContext.Simulate,
		GenerateOnly:  cliContext.GenerateOnly,
		FromAddress:   cliContext.FromAddress,
		FromName:      cliContext.FromName,
		Indent:        cliContext.Indent,
		SkipConfirm:   cliContext.SkipConfirm,
	}
	require.Equal(t, KafkaMsg{Msg: testMessage, TicketID: ticketID, BaseRequest: testBaseReq, KafkaCli: kafkaCli}, testKafkaMsg)
	require.Equal(t, cliContext, CliCtxFromKafkaMsg(testKafkaMsg, cliContext))
	//require
}
