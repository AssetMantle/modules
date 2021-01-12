/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package queuing

import (
	"github.com/Shopify/sarama"
	"github.com/cosmos/cosmos-sdk/codec"
)

// NewConsumer : is a consumer which is needed to create child consumers to consume topics
func NewConsumer(kafkaPorts []string) sarama.Consumer {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer(kafkaPorts, config)
	if err != nil {
		panic(err)
	}
	return consumer
}

// PartitionConsumers : is a child consumer
func PartitionConsumers(consumer sarama.Consumer, topic string) sarama.PartitionConsumer {
	// partition and offset defined in CONSTANTS.go
	partitionConsumer, err := consumer.ConsumePartition(topic, partition, offset)
	if err != nil {
		panic(err)
	}
	return partitionConsumer
}

//KafkaTopicConsumer : Takes a consumer and makes it consume a topic message at a time
func KafkaTopicConsumer(topic string, consumers map[string]sarama.PartitionConsumer, cdc *codec.Codec) KafkaMsg {

	partitionConsumer := consumers[topic]

	if len(partitionConsumer.Messages()) == 0 {
		var kafkaStore = KafkaMsg{Msg: nil}
		return kafkaStore
	}
	kafkaMsg := <-partitionConsumer.Messages()
	var kafkaStore KafkaMsg
	err := cdc.UnmarshalJSON(kafkaMsg.Value, &kafkaStore)
	if err != nil {
		panic(err)
	}
	return kafkaStore
}
