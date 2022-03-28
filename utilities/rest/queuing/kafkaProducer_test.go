package queuing

import (
	"testing"

	"github.com/Shopify/sarama"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/schema"
)

func TestKafkaProducerDeliverMessage(t *testing.T) {
	testProducer := []string{"testProducer"}
	var Codec = codec.NewLegacyAmino()
	require.Panics(t, func() {
		schema.RegisterLegacyAminoCodec(Codec)
		sdkTypes.RegisterLegacyAminoCodec(Codec)
		cryptoCodec.RegisterCrypto(Codec)
		codec.RegisterEvidences(Codec)
		vesting.RegisterCodec(Codec)

		testKafkaMessage := kafkaMsg{Msg: nil}

		producer, err := sarama.NewSyncProducer(testProducer, nil)
		require.Nil(t, err)

		require.Equal(t, kafkaProducerDeliverMessage(testKafkaMessage, "Topic", producer, Codec), nil)
	})

}

func TestNewProducer(t *testing.T) {
	testProducer := []string{"testProducer"}

	producer, err := sarama.NewSyncProducer(testProducer, nil)

	require.Nilf(t, err, "should not happened. err %v", err)

	require.Panics(t, func() {
		require.Equal(t, newProducer(testProducer), producer)
	})
}
