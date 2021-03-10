/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package queuing

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	vestingTypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/stretchr/testify/require"
)

func Test_Kafka_Types(t *testing.T) {

	var Codec = codec.NewLegacyAmino()
	schema.RegisterCodec(Codec)
	codec.RegisterEvidences(Codec)
	vestingTypes.RegisterLegacyAminoCodec(Codec)
	std.RegisterLegacyAminoCodec(Codec)

	cliContext := client.Context{}.WithLegacyAmino(Codec)

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test", Fees: sdkTypes.NewCoins()}

	testMessage := testdata.NewTestMsg()

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
		UseLedger:     cliContext.UseLedger,
		BroadcastMode: cliContext.BroadcastMode,
		Simulate:      cliContext.Simulate,
		GenerateOnly:  cliContext.GenerateOnly,
		FromAddress:   cliContext.FromAddress,
		FromName:      cliContext.FromName,
		SkipConfirm:   cliContext.SkipConfirm,
	}
	require.Equal(t, KafkaMsg{Msg: testMessage, TicketID: ticketID, BaseRequest: testBaseReq, KafkaCli: kafkaCli}, testKafkaMsg)
	require.Equal(t, cliContext, CliCtxFromKafkaMsg(testKafkaMsg, cliContext))
	//require
}
