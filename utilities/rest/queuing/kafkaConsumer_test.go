package queuing

import (
	"github.com/Shopify/sarama"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKafkaTopicConsumer(t *testing.T) {
	testConsumers := []string{"testConsumers"}
	var Codec = codec.NewLegacyAmino()
	schema.RegisterLegacyAminoCodec(Codec)
	sdkTypes.RegisterLegacyAminoCodec(Codec)
	cryptoCodec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)

	require.Panics(t, func() {
		testKafkaState := NewKafkaState(testConsumers)
		partitionConsumer := testKafkaState.Consumers["Topic"]
		var kafkaStore kafkaMsg
		if len(partitionConsumer.Messages()) == 0 {
			kafkaStore = kafkaMsg{Msg: nil}
		}
		kafkaMsg := <-partitionConsumer.Messages()
		err := Codec.UnmarshalJSON(kafkaMsg.Value, &kafkaStore)

		if err != nil {
			panic(err)
		}
		require.Equal(t, kafkaTopicConsumer("Topic", testKafkaState.Consumers, Codec), kafkaStore)
	})
}

func TestNewConsumer(t *testing.T) {
	consumers := []string{"testConsumers"}
	config := sarama.NewConfig()
	consumer, Error := sarama.NewConsumer(consumers, config)

	if Error != nil {
	}
	require.Panics(t, func() {
		require.Equal(t, newConsumer(consumers), consumer)
	})
}
