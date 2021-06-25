/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package queuing

import (
	"time"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

// kafkaConsumerMessages : messages to consume 5 second delay
func kafkaConsumerMessages(cliCtx context.CLIContext) {
	quit := make(chan bool)

	var cliContextList []context.CLIContext

	var baseRequestList []rest.BaseReq

	var ticketIDList []TicketID

	var msgList []sdkTypes.Msg

	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				kafkaMsg := kafkaTopicConsumer("Topic", KafkaState.Consumers, cliCtx.Codec)
				if kafkaMsg.Msg != nil {
					cliContextList = append(cliContextList, cliCtxFromKafkaMsg(kafkaMsg, cliCtx))
					baseRequestList = append(baseRequestList, kafkaMsg.BaseRequest)
					ticketIDList = append(ticketIDList, kafkaMsg.TicketID)
					msgList = append(msgList, kafkaMsg.Msg)
				}
			}
		}
	}()

	time.Sleep(sleepTimer)
	quit <- true

	if len(msgList) == 0 {
		return
	}

	output, err := signAndBroadcastMultiple(baseRequestList, cliContextList, msgList)
	if err != nil {
		jsonError, e := cliCtx.Codec.MarshalJSON(struct {
			Error string `json:"error"`
		}{Error: err.Error()})
		if e != nil {
			panic(err)
		}

		for _, ticketID := range ticketIDList {
			addResponseToDB(ticketID, jsonError, KafkaState.KafkaDB, cliCtx.Codec)
		}

		return
	}

	for _, ticketID := range ticketIDList {
		addResponseToDB(ticketID, output, KafkaState.KafkaDB, cliCtx.Codec)
	}
}
