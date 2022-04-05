// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

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
	KafkaCliCtx kafkaCliCtx  `json:"kafkaCliCtx"`
}

// NewKafkaMsgFromRest : makes a msg to send to kafka queue
func NewKafkaMsgFromRest(msg sdk.Msg, ticketID TicketID, baseRequest rest.BaseReq, cliCtx context.CLIContext) kafkaMsg {
	kafkaCtx := kafkaCliCtx{
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

	// TODO return pointer
	return kafkaMsg{
		Msg:         msg,
		TicketID:    ticketID,
		BaseRequest: baseRequest,
		KafkaCliCtx: kafkaCtx,
	}
}

// cliCtxFromKafkaMsg : sets the transaction and cli contexts again to consume
func cliCtxFromKafkaMsg(kafkaMsg kafkaMsg, cliContext context.CLIContext) context.CLIContext {
	cliContext.OutputFormat = kafkaMsg.KafkaCliCtx.OutputFormat
	cliContext.ChainID = kafkaMsg.KafkaCliCtx.ChainID
	cliContext.Height = kafkaMsg.KafkaCliCtx.Height
	cliContext.HomeDir = kafkaMsg.KafkaCliCtx.HomeDir
	cliContext.NodeURI = kafkaMsg.KafkaCliCtx.NodeURI
	cliContext.From = kafkaMsg.KafkaCliCtx.From
	cliContext.TrustNode = kafkaMsg.KafkaCliCtx.TrustNode
	cliContext.UseLedger = kafkaMsg.KafkaCliCtx.UseLedger
	cliContext.BroadcastMode = kafkaMsg.KafkaCliCtx.BroadcastMode
	cliContext.Simulate = kafkaMsg.KafkaCliCtx.Simulate
	cliContext.GenerateOnly = kafkaMsg.KafkaCliCtx.GenerateOnly
	cliContext.FromAddress = kafkaMsg.KafkaCliCtx.FromAddress
	cliContext.FromName = kafkaMsg.KafkaCliCtx.FromName
	cliContext.Indent = kafkaMsg.KafkaCliCtx.Indent
	cliContext.SkipConfirm = kafkaMsg.KafkaCliCtx.SkipConfirm

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
func NewKafkaState(nodeList []string) *kafkaState {
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
