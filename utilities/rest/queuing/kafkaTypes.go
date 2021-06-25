/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package queuing

import (
	"github.com/Shopify/sarama"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	dbm "github.com/tendermint/tm-db"
)

// TicketID : is a type that implements string
type TicketID string

// kafkaMsg : is a store that can be stored in kafka queues
type kafkaMsg struct {
	Msg         sdk.Msg      `json:"msg"`
	TicketID    TicketID     `json:"TicketID"`
	BaseRequest rest.BaseReq `json:"base_req"`
	KafkaCli    kafkaCliCtx  `json:"kafkaCliCtx"`
}

// NewKafkaMsgFromRest : makes a msg to send to kafka queue
func NewKafkaMsgFromRest(msg sdk.Msg, ticketID TicketID, baseRequest rest.BaseReq, cliCtx context.CLIContext) kafkaMsg {
	kafkaCli := kafkaCliCtx{
		OutputFormat:  cliCtx.OutputFormat,
		ChainID:       cliCtx.ChainID,
		Height:        cliCtx.Height,
		HomeDir:       cliCtx.HomeDir,
		NodeURI:       cliCtx.NodeURI,
		From:          cliCtx.From,
		TrustNode:     cliCtx.TrustNode,
		UseLedger:     cliCtx.UseLedger,
		BroadcastMode: cliCtx.BroadcastMode,
		Simulate:      cliCtx.Simulate,
		GenerateOnly:  cliCtx.GenerateOnly,
		FromAddress:   cliCtx.FromAddress,
		FromName:      cliCtx.FromName,
		Indent:        cliCtx.Indent,
		SkipConfirm:   cliCtx.SkipConfirm,
	}

	return kafkaMsg{
		Msg:         msg,
		TicketID:    ticketID,
		BaseRequest: baseRequest,
		KafkaCli:    kafkaCli,
	}
}

// cliCtxFromKafkaMsg : sets the transaction and cli contexts again to consume
func cliCtxFromKafkaMsg(kafkaMsg kafkaMsg, cliContext context.CLIContext) context.CLIContext {
	cliContext.OutputFormat = kafkaMsg.KafkaCli.OutputFormat
	cliContext.ChainID = kafkaMsg.KafkaCli.ChainID
	cliContext.Height = kafkaMsg.KafkaCli.Height
	cliContext.HomeDir = kafkaMsg.KafkaCli.HomeDir
	cliContext.NodeURI = kafkaMsg.KafkaCli.NodeURI
	cliContext.From = kafkaMsg.KafkaCli.From
	cliContext.TrustNode = kafkaMsg.KafkaCli.TrustNode
	cliContext.UseLedger = kafkaMsg.KafkaCli.UseLedger
	cliContext.BroadcastMode = kafkaMsg.KafkaCli.BroadcastMode
	cliContext.Simulate = kafkaMsg.KafkaCli.Simulate
	cliContext.GenerateOnly = kafkaMsg.KafkaCli.GenerateOnly
	cliContext.FromAddress = kafkaMsg.KafkaCli.FromAddress
	cliContext.FromName = kafkaMsg.KafkaCli.FromName
	cliContext.Indent = kafkaMsg.KafkaCli.Indent
	cliContext.SkipConfirm = kafkaMsg.KafkaCli.SkipConfirm

	return cliContext
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
	TrustNode     bool
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

// NewKafkaState : returns a kafka state
func NewKafkaState(kafkaPorts []string) *kafkaState {
	kafkaDB, _ := dbm.NewGoLevelDB("KafkaDB", defaultCLIHome)
	admin := kafkaAdmin(kafkaPorts)
	producer := newProducer(kafkaPorts)
	consumer := newConsumer(kafkaPorts)

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
