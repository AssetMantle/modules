// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queuing

import (
	"testing"

	"github.com/Shopify/sarama"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/schema"
)

func TestKafkaProducerDeliverMessage(t *testing.T) {
	testProducer := []string{"testProducer"}
	var Codec = codec.New()
	require.Panics(t, func() {
		schema.RegisterCodec(Codec)
		sdkTypes.RegisterCodec(Codec)
		codec.RegisterCrypto(Codec)
		codec.RegisterEvidences(Codec)
		vesting.RegisterCodec(Codec)

		testKafkaMessage := kafkaMsg{Msg: nil}

		producer, _ := sarama.NewSyncProducer(testProducer, nil)
		// TODO: Add test cases.
		// require.Nil(t, err)

		require.Equal(t, kafkaProducerDeliverMessage(testKafkaMessage, "Topic", producer, Codec), nil)
	})

}

func TestNewProducer(t *testing.T) {
	testProducer := []string{"testProducer"}

	producer, _ := sarama.NewSyncProducer(testProducer, nil)

	// TODO: Add test cases.
	// require.Nilf(t, err, "should not happened. err %v", err)

	require.Panics(t, func() {
		require.Equal(t, newProducer(testProducer), producer)
	})
}
