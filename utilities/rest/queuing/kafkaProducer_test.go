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

func TestKafkaProducerDeliverMessage(t *testing.T) {
	testProducer := []string{"testProducer"}
	var Codec = codec.New()
	require.Panics(t, func() {
		schema.RegisterLegacyAminoCodec(Codec)
		sdkTypes.RegisterCodec(Codec)
		codec.RegisterCrypto(Codec)
		codec.RegisterEvidences(Codec)
		vesting.RegisterCodec(Codec)
		testKafkaMessage := kafkaMsg{Msg: nil}
		producer, err := sarama.NewSyncProducer(testProducer, nil)
		if err != nil {
		}
		require.Equal(t, kafkaProducerDeliverMessage(testKafkaMessage, "Topic", producer, Codec), nil)
	})

}

func TestNewProducer(t *testing.T) {
	testProducer := []string{"testProducer"}
	producer, err := sarama.NewSyncProducer(testProducer, nil)

	if err != nil {
	}
	require.Panics(t, func() {
		require.Equal(t, newProducer(testProducer), producer)
	})
}
