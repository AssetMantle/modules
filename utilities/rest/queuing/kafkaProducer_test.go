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
	tproducer := []string{"testproducer"}
	var Codec = codec.New()
	require.Panics(t, func() {
		schema.RegisterCodec(Codec)
		sdkTypes.RegisterCodec(Codec)
		codec.RegisterCrypto(Codec)
		codec.RegisterEvidences(Codec)
		vesting.RegisterCodec(Codec)
		tkmsg := KafkaMsg{Msg: nil}
		producer, err := sarama.NewSyncProducer(tproducer, nil)
		if err != nil {
		}
		require.Equal(t, KafkaProducerDeliverMessage(tkmsg, "Topic",producer,Codec),nil)
	})

}

func TestNewProducer(t *testing.T) {
	tproducer := []string{"testproducer"}
	producer, err := sarama.NewSyncProducer(tproducer, nil)

	if err != nil {
	}
	require.Panics(t, func() {
		require.Equal(t, NewProducer(tproducer),producer)
	})
}


