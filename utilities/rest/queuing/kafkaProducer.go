/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package queuing

import (
	"github.com/Shopify/sarama"

	"github.com/cosmos/cosmos-sdk/codec"
)

// NewProducer is a producer to send messages to kafka
func NewProducer(kafkaNodes []string) sarama.SyncProducer {
	producer, err := sarama.NewSyncProducer(kafkaNodes, nil)
	if err != nil {
		panic(err)
	}

	return producer
}

// KafkaProducerDeliverMessage : delivers messages to kafka
func KafkaProducerDeliverMessage(kafkaMsg KafkaMsg, topic string, producer sarama.SyncProducer, Codec *codec.Codec) error {
	kafkaStoreBytes, err := Codec.MarshalJSON(kafkaMsg)
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
func SendToKafka(kafkaMsg KafkaMsg, kafkaState kafkaState, Codec *codec.Codec) []byte {
	Error := KafkaProducerDeliverMessage(kafkaMsg, "Topic", kafkaState.Producer, Codec)
	if Error != nil {
		jsonResponse, Error := Codec.MarshalJSON(struct {
			Response string `json:"response"`
		}{Response: "Something is up with kafka server, restart rest and kafka."})
		if Error != nil {
			panic(Error)
		}

		SetTicketIDtoDB(kafkaMsg.TicketID, kafkaState.KafkaDB, Codec, jsonResponse)
	} else {
		jsonResponse, err := Codec.MarshalJSON(struct {
			Error string `json:"error"`
		}{Error: "Request in process, wait and try after some time"})
		if err != nil {
			panic(err)
		}
		SetTicketIDtoDB(kafkaMsg.TicketID, kafkaState.KafkaDB, Codec, jsonResponse)
	}

	jsonResponse, Error := Codec.MarshalJSON(struct {
		TicketID TicketID `json:"ticketID"`
	}{TicketID: kafkaMsg.TicketID})
	if Error != nil {
		panic(Error)
	}

	return jsonResponse
}
