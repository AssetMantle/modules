/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package queuing

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/utilities/random"
)

func Test_Kafka_Types(t *testing.T) {

	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()
	cliContext := context.NewCLIContext().WithCodec(Codec)

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test", Fees: sdkTypes.NewCoins()}

	testMessage := sdkTypes.NewTestMsg()

	ticketID := TicketID(random.GenerateUniqueIdentifier("name"))
	testKafkaMsg := NewKafkaMsgFromRest(testMessage, ticketID, testBaseReq, cliContext)
	kafkaCliCtx := kafkaCliCtx{
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
	require.Equal(t, kafkaMsg{Msg: testMessage, TicketID: ticketID, BaseRequest: testBaseReq, KafkaCliCtx: kafkaCliCtx}, testKafkaMsg)
	require.Equal(t, cliContext, cliCtxFromKafkaMsg(testKafkaMsg, cliContext))
	// require
}
