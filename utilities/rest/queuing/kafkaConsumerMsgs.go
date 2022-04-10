// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queuing

import (
	"time"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// kafkaConsumerMessages : messages to consume 5-second delay
func kafkaConsumerMessages(cliCtx context.CLIContext) {
	quit := make(chan bool)

	var kafkaMsgList []kafkaMsg

	go func() {
		var msg kafkaMsg

		for {
			select {
			case <-quit:
				return
			default:
				msg = kafkaTopicConsumer("Topic", KafkaState.Consumers, cliCtx.Codec)
				if msg.Msg != nil {
					kafkaMsgList = append(kafkaMsgList, msg)
				}
			}
		}
	}()

	time.Sleep(sleepTimer)
	quit <- true

	if len(kafkaMsgList) == 0 {
		return
	}

	output, err := signAndBroadcastMultiple(kafkaMsgList, cliCtx)
	if err != nil {
		jsonError, e := cliCtx.Codec.MarshalJSON(struct {
			Error string `json:"error"`
		}{Error: err.Error()})
		if e != nil {
			panic(e)
		}

		for _, kafkaMsg := range kafkaMsgList {
			addResponseToDB(kafkaMsg.TicketID, jsonError, KafkaState.KafkaDB, cliCtx.Codec)
		}

		return
	}

	for _, kafkaMsg := range kafkaMsgList {
		addResponseToDB(kafkaMsg.TicketID, output, KafkaState.KafkaDB, cliCtx.Codec)
	}
}
