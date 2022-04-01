// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queuing

import (
	"github.com/Shopify/sarama"

	"github.com/cosmos/cosmos-sdk/codec"
)

// newProducer is a producer to send messages to kafka
func newProducer(kafkaNodes []string) sarama.SyncProducer {
	producer, err := sarama.NewSyncProducer(kafkaNodes, nil)
	if err != nil {
		panic(err)
	}

	return producer
}

// kafkaProducerDeliverMessage : delivers messages to kafka
func kafkaProducerDeliverMessage(kafkaMsg kafkaMsg, topic string, producer sarama.SyncProducer, codec *codec.Codec) error {
	kafkaStoreBytes, err := codec.MarshalJSON(kafkaMsg)
	if err != nil {
		panic(err)
	}

	sendMsg := sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(kafkaStoreBytes),
	}

	_, _, err = producer.SendMessage(&sendMsg)
	if err != nil {
		return err
	}

	return nil
}

// SendToKafka : handles sending message to kafka
func SendToKafka(kafkaMsg kafkaMsg, codec *codec.Codec) []byte {
	var jsonResponse []byte

	err := kafkaProducerDeliverMessage(kafkaMsg, "Topic", KafkaState.Producer, codec)
	if err != nil {
		jsonResponse, err = codec.MarshalJSON(struct {
			Response string `json:"response"`
		}{Response: "Something is up with kafka server, restart rest and kafka."})
		if err != nil {
			panic(err)
		}

		setTicketIDtoDB(kafkaMsg.TicketID, KafkaState.KafkaDB, codec, jsonResponse)
	} else {
		jsonResponse, err = codec.MarshalJSON(struct {
			Error string `json:"error"`
		}{Error: "Request in process, wait and try after some time"})
		if err != nil {
			panic(err)
		}

		setTicketIDtoDB(kafkaMsg.TicketID, KafkaState.KafkaDB, codec, jsonResponse)
	}

	jsonResponse, err = codec.MarshalJSON(struct {
		TicketID TicketID `json:"TicketID"`
	}{TicketID: kafkaMsg.TicketID})
	if err != nil {
		panic(err)
	}

	return jsonResponse
}
