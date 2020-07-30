package rest

import (
	"github.com/persistenceOne/persistenceSDK/kafka"
	"time"

	cTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/cosmos/cosmos-sdk/client/context"

)

// KafkaConsumerMsgs : msgs to consume 5 second delay
func KafkaConsumerMsgs(cliCtx context.CLIContext, kafkaState kafka.KafkaState) {

	quit := make(chan bool)

	//modes := []string{}
	//passwords := []string{}
	cliCtxs := []context.CLIContext{}
	baseRequests := []rest.BaseReq{}
	ticketIDs := []kafka.Ticket{}
	msgs := []cTypes.Msg{}

	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				kafkaMsg := kafka.KafkaTopicConsumer("Topic", kafkaState.Consumers, cliCtx.Codec)
				if kafkaMsg.Msg != nil {
					cliCtxs = append(cliCtxs, kafka.CliCtxFromKafkaMsg(kafkaMsg, cliCtx))
					baseRequests = append(baseRequests, kafkaMsg.BaseRequest)
					ticketIDs = append(ticketIDs, kafkaMsg.TicketID)
					msgs = append(msgs, kafkaMsg.Msg)
				}
			}
		}
	}()

	time.Sleep(kafka.SleepTimer)
	quit <- true
	if len(msgs) == 0 {
		return
	}

	output, err := SignAndBroadcastMultiples(baseRequests, cliCtxs, msgs)
	if err != nil {
		jsonError, e := cliCtx.Codec.MarshalJSON(struct {
			Error string `json:"error"`
		}{Error: err.Error()})
		if e != nil {
			panic(err)
		}
		for _, ticketID := range ticketIDs {
			kafka.AddResponseToDB(ticketID, jsonError, kafkaState.KafkaDB, cliCtx.Codec)
		}
		return
	}

	for _, ticketID := range ticketIDs {
		kafka.AddResponseToDB(ticketID, output, kafkaState.KafkaDB, cliCtx.Codec)
	}
}
