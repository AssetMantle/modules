/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package queuing

import (
	"github.com/Shopify/sarama"
	"github.com/cosmos/cosmos-sdk/codec"
)

// newConsumer : is a consumer which is needed to create child consumers to consume topics
func newConsumer(kafkaNodes []string) sarama.Consumer {
	config := sarama.NewConfig()

	consumer, Error := sarama.NewConsumer(kafkaNodes, config)
	if Error != nil {
		panic(Error)
	}

	return consumer
}

// partitionConsumers : is a child consumer
func partitionConsumers(consumer sarama.Consumer, topic string) sarama.PartitionConsumer {
	// partition and offset defined in configurations.go
	partitionConsumer, Error := consumer.ConsumePartition(topic, partition, offset)
	if Error != nil {
		panic(Error)
	}

	return partitionConsumer
}

// kafkaTopicConsumer : Takes a consumer and makes it consume a topic message at a time
func kafkaTopicConsumer(topic string, consumers map[string]sarama.PartitionConsumer, cdc *codec.Codec) kafkaMsg {
	partitionConsumer := consumers[topic]

	if len(partitionConsumer.Messages()) == 0 {
		return kafkaMsg{Msg: nil}
	}

	var consumedKafkaMsg kafkaMsg
	err := cdc.UnmarshalJSON((<-partitionConsumer.Messages()).Value, &consumedKafkaMsg)

	if err != nil {
		panic(err)
	}

	return consumedKafkaMsg
}
