/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package queuing

import (
	"testing"

	cryptoCodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	vestingTypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/stretchr/testify/require"
)

func Test_Kafka(t *testing.T) {

	var Codec = codec.NewLegacyAmino()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterLegacyAminoCodec(Codec)
	cryptoCodec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vestingTypes.RegisterLegacyAminoCodec(Codec)

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, Error := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, Error)
	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test"}
	ticketID := TicketIDGenerator("ticket")
	kafkaPorts := []string{"localhost:9092"}
	testKafkaState := NewKafkaState(kafkaPorts)
	message := bankTypes.NewMsgSend(fromAccAddress, fromAccAddress, sdkTypes.NewCoins(sdkTypes.NewCoin("stake", sdkTypes.NewInt(123))))
	cliContext := client.Context{}.WithLegacyAmino(Codec)

	testKafkaMsg := NewKafkaMsgFromRest(message, ticketID, testBaseReq, cliContext)
	SendToKafka(testKafkaMsg, testKafkaState, Codec)

	kafkaMsg := KafkaTopicConsumer("Topic", testKafkaState.Consumers, Codec)
	require.Equal(t, testKafkaMsg.TicketID, kafkaMsg.TicketID)
	require.Equal(t, testKafkaMsg.BaseRequest, kafkaMsg.BaseRequest)

}
