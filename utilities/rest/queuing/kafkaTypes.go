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

// Ticket : is a type that implements string
type Ticket string

// KafkaMsg : is a store that can be stored in kafka queues
type KafkaMsg struct {
	Msg         sdk.Msg      `json:"msg"`
	TicketID    Ticket       `json:"ticketID"`
	BaseRequest rest.BaseReq `json:"base_req"`
	KafkaCli    KafkaCliCtx  `json:"kafkaCliCtx"`
}

// NewKafkaMsgFromRest : makes a msg to send to kafka queue
func NewKafkaMsgFromRest(msg sdk.Msg, ticketID Ticket, baseRequest rest.BaseReq, cliCtx context.CLIContext) KafkaMsg {
	kafkaCli := KafkaCliCtx{
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

	return KafkaMsg{
		Msg:         msg,
		TicketID:    ticketID,
		BaseRequest: baseRequest,
		KafkaCli:    kafkaCli,
	}
}

// CliCtxFromKafkaMsg : sets the transaction and cli contexts again to consume
func CliCtxFromKafkaMsg(kafkaMsg KafkaMsg, cliContext context.CLIContext) context.CLIContext {
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

// KafkaCliCtx : client tx without codec
type KafkaCliCtx struct {
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

// TicketIDResponse : is a json structure to send TicketID to user
type TicketIDResponse struct {
	TicketID Ticket `json:"ticketID" valid:"required~ticketID is mandatory,length(20)~ticketID length should be 20" `
}

// KafkaState : is a struct showing the state of kafka
type KafkaState struct {
	KafkaDB   *dbm.GoLevelDB
	Admin     sarama.ClusterAdmin
	Consumer  sarama.Consumer
	Consumers map[string]sarama.PartitionConsumer
	Producer  sarama.SyncProducer
	Topics    []string
}

// NewKafkaState : returns a kafka state
func NewKafkaState(kafkaPorts []string) KafkaState {
	kafkaDB, _ := dbm.NewGoLevelDB("KafkaDB", DefaultCLIHome)
	admin := KafkaAdmin(kafkaPorts)
	producer := NewProducer(kafkaPorts)
	consumer := NewConsumer(kafkaPorts)

	var consumers = make(map[string]sarama.PartitionConsumer)

	for _, topic := range Topics {
		partitionConsumer := PartitionConsumers(consumer, topic)
		consumers[topic] = partitionConsumer
	}

	return KafkaState{
		KafkaDB:   kafkaDB,
		Admin:     admin,
		Consumer:  consumer,
		Consumers: consumers,
		Producer:  producer,
		Topics:    Topics,
	}
}
