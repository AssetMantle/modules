/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package queuing

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/persistenceOne/persistenceSDK/schema"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
	"github.com/stretchr/testify/require"
	"testing"
)

type testMessage struct {
	Name string `json:"name"`
}

var _ sdkTypes.Msg = testMessage{}

func (message testMessage) Route() string { return "testModule" }
func (message testMessage) Type() string  { return "" }
func (message testMessage) ValidateBasic() error {
	return nil
}
func (message testMessage) GetSignBytes() []byte {
	return []byte{}
}
func (message testMessage) GetSigners() []sdkTypes.AccAddress {
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, _ := sdkTypes.AccAddressFromBech32(fromAddress)
	return []sdkTypes.AccAddress{fromAccAddress}
}
func (testMessage) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterXPRTConcrete(codec, "testModule", testMessage{})
}

func Test_Kafka(t *testing.T) {

	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, Error := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, Error)
	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test"}
	ticketID := TicketIDGenerator("ticket")
	kafkaPorts := []string{"localhost:9092"}
	require.Panics(t, func() {
		testKafkaState := NewKafkaState(kafkaPorts)
		bank.RegisterCodec(Codec)
		message := bank.NewMsgSend(fromAccAddress, fromAccAddress, sdkTypes.NewCoins(sdkTypes.NewCoin("stake", sdkTypes.NewInt(123))))
		cliContext := context.NewCLIContext().WithCodec(Codec)

		testKafkaMsg := NewKafkaMsgFromRest(message, ticketID, testBaseReq, cliContext)
		SendToKafka(testKafkaMsg, testKafkaState, Codec)

		kafkaMsg := KafkaTopicConsumer("Topic", testKafkaState.Consumers, Codec)
		require.Equal(t, testKafkaMsg.TicketID, kafkaMsg.TicketID)
		require.Equal(t, testKafkaMsg.BaseRequest, kafkaMsg.BaseRequest)
	})

}
