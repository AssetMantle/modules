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
	tconsumers := []string{"testconsumer"}
	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)

	require.Panics(t, func() {
		testKafkaState := NewKafkaState(tconsumers)
		partitionConsumer := testKafkaState.Consumers["Topic"]
		var kafkaStore KafkaMsg
		if len(partitionConsumer.Messages()) == 0 {
			kafkaStore = KafkaMsg{Msg: nil}
		}
		kafkaMsg := <-partitionConsumer.Messages()
		err := Codec.UnmarshalJSON(kafkaMsg.Value, &kafkaStore)

		if err != nil {
			panic(err)
		}
		require.Equal(t, KafkaTopicConsumer("Topic", testKafkaState.Consumers, Codec), kafkaStore)
	})
}

func TestNewConsumer(t *testing.T) {
	consumers := []string{"testconsumer"}
	config := sarama.NewConfig()
	consumer, Error := sarama.NewConsumer(consumers, config)

	if Error != nil {
	}
	require.Panics(t, func() {
		require.Equal(t, NewConsumer(consumers), consumer)
	})
}
