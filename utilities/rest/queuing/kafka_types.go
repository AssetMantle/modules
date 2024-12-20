// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queuing

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/IBM/sarama"
	dbm "github.com/cometbft/cometbft-db"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TicketID : is a type that implements string
type TicketID string

// kafkaMsg : is a store that can be stored in kafka queues
type kafkaMsg struct {
	sdk.Msg                          `json:"msg"`
	TicketID                         `json:"TicketID"`
	helpers.CommonTransactionRequest `json:"commonTransactionRequest"`
	KafkaCliCtx                      kafkaCliCtx
}

// newKafkaMsgFromRest : makes a msg to send to kafka queue
func newKafkaMsgFromRest(msg sdk.Msg, ticketID TicketID, commonTransactionRequest helpers.CommonTransactionRequest, context client.Context) kafkaMsg {
	kafkaCtx := kafkaCliCtx{
		OutputFormat:  context.OutputFormat,
		ChainID:       context.ChainID,
		Height:        context.Height,
		HomeDir:       context.HomeDir,
		NodeURI:       context.NodeURI,
		From:          context.From,
		UseLedger:     context.UseLedger,
		BroadcastMode: context.BroadcastMode,
		Simulate:      context.Simulate,
		GenerateOnly:  context.GenerateOnly,
		FromAddress:   context.FromAddress,
		FromName:      context.FromName,
		SkipConfirm:   context.SkipConfirm,
	}

	// TODO return pointer
	return kafkaMsg{
		Msg:                      msg,
		TicketID:                 ticketID,
		CommonTransactionRequest: commonTransactionRequest,
		KafkaCliCtx:              kafkaCtx,
	}
}

// cliCtxFromKafkaMsg : sets the transaction and cli contexts again to consume
func cliCtxFromKafkaMsg(kafkaMsg kafkaMsg, context client.Context) client.Context {
	context.OutputFormat = kafkaMsg.KafkaCliCtx.OutputFormat
	context.ChainID = kafkaMsg.KafkaCliCtx.ChainID
	context.Height = kafkaMsg.KafkaCliCtx.Height
	context.HomeDir = kafkaMsg.KafkaCliCtx.HomeDir
	context.NodeURI = kafkaMsg.KafkaCliCtx.NodeURI
	context.From = kafkaMsg.KafkaCliCtx.From
	context.UseLedger = kafkaMsg.KafkaCliCtx.UseLedger
	context.BroadcastMode = kafkaMsg.KafkaCliCtx.BroadcastMode
	context.Simulate = kafkaMsg.KafkaCliCtx.Simulate
	context.GenerateOnly = kafkaMsg.KafkaCliCtx.GenerateOnly
	context.FromAddress = kafkaMsg.KafkaCliCtx.FromAddress
	context.FromName = kafkaMsg.KafkaCliCtx.FromName
	context.SkipConfirm = kafkaMsg.KafkaCliCtx.SkipConfirm

	return context
}

// kafkaCliCtx : client tx without codec
type kafkaCliCtx struct {
	FromAddress   sdk.AccAddress
	OutputFormat  string
	ChainID       string
	HomeDir       string
	NodeURI       string
	From          string
	BroadcastMode string
	FromName      string
	Height        int64
	UseLedger     bool
	Simulate      bool
	GenerateOnly  bool
	Offline       bool
	Indent        bool
	SkipConfirm   bool
}

// kafkaState : is a struct showing the state of kafka
type kafkaState struct {
	KafkaDB   *dbm.GoLevelDB
	Admin     sarama.ClusterAdmin
	Consumer  sarama.Consumer
	Consumers map[string]sarama.PartitionConsumer
	Producer  sarama.SyncProducer
	Topics    []string
	IsEnabled bool
}

// newKafkaState : returns a kafka state
func newKafkaState(nodeList []string) *kafkaState {
	kafkaDB, _ := dbm.NewGoLevelDB("KafkaDB", defaultCLIHome)
	admin := kafkaAdmin(nodeList)
	producer := newProducer(nodeList)
	consumer := newConsumer(nodeList)

	var consumers = make(map[string]sarama.PartitionConsumer)

	for _, topic := range topics {
		partitionConsumer := partitionConsumers(consumer, topic)
		consumers[topic] = partitionConsumer
	}

	return &kafkaState{
		KafkaDB:   kafkaDB,
		Admin:     admin,
		Consumer:  consumer,
		Consumers: consumers,
		Producer:  producer,
		Topics:    topics,
		IsEnabled: true,
	}
}
