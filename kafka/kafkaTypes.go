package kafka

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
		Offline:       cliCtx.Offline,
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

// CliCtxFromKafkaMsg : sets the txctx and clictx again to consume
func CliCtxFromKafkaMsg(msg KafkaMsg, cliCtx context.CLIContext) context.CLIContext {

	//mvh := msg.KafkaCli.VerifierHome

	cliCtx.OutputFormat = msg.KafkaCli.OutputFormat
	cliCtx.ChainID = msg.KafkaCli.ChainID
	cliCtx.Height = msg.KafkaCli.Height
	cliCtx.HomeDir = msg.KafkaCli.HomeDir
	cliCtx.NodeURI = msg.KafkaCli.NodeURI
	cliCtx.From = msg.KafkaCli.From
	cliCtx.TrustNode = msg.KafkaCli.TrustNode
	cliCtx.UseLedger = msg.KafkaCli.UseLedger
	cliCtx.BroadcastMode = msg.KafkaCli.BroadcastMode
	cliCtx.Simulate = msg.KafkaCli.Simulate
	cliCtx.GenerateOnly = msg.KafkaCli.GenerateOnly
	cliCtx.FromAddress = msg.KafkaCli.FromAddress
	cliCtx.FromName = msg.KafkaCli.FromName
	cliCtx.Offline = msg.KafkaCli.Offline
	cliCtx.Indent = msg.KafkaCli.Indent
	cliCtx.SkipConfirm = msg.KafkaCli.SkipConfirm

	return cliCtx
}

// KafkaCliCtx : client tx without codec
type KafkaCliCtx struct {
	OutputFormat  string
	ChainID       string
	Height        int64
	HomeDir       string
	NodeURI       string
	From          string
	TrustNode     bool
	UseLedger     bool
	BroadcastMode string
	Simulate      bool
	GenerateOnly  bool
	FromAddress   sdk.AccAddress
	FromName      string
	Offline       bool
	Indent        bool
	SkipConfirm   bool
}

// TicketIDResponse : is a json structure to send TicketID to user
type TicketIDResponse struct {
	TicketID Ticket `json:"TicketID" valid:"required~TicketID is mandatory,length(20)~RelayerAddress length should be 20" `
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
