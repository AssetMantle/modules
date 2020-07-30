package kafka

import (
	"os"
	"time"

	"github.com/Shopify/sarama"
)

// SleepTimer : the time the kafka msgs are to be taken in
var SleepTimer = time.Duration(1000000000)

// SleepRoutine : the time the kafka msgs are to be taken in
var SleepRoutine = time.Duration(2500000000)

// These are the config parameters for running kafka admins and producers and consumers. Declared very minimal
var replicaAssignment = map[int32][]int32{}
var configEntries = map[string]*string{}

// DefaultCLIHome : is the home path
var DefaultCLIHome = os.ExpandEnv("$HOME/.kafka")
var partition = int32(0)
var offset = int64(0)

// topicDetail : configs
var topicDetail = sarama.TopicDetail{
	NumPartitions:     1,
	ReplicationFactor: 1,
	ReplicaAssignment: replicaAssignment,
	ConfigEntries:     configEntries,
}

// Topics : is list of topics
var Topics = []string{
	"Topic",
}

/*
BUEO - buyer execute order
DEZO - define zone
DEAC - define acl
DEOR - define organization
ISAS - issue asset
ISFI - issue fiat
RDAS - redeem asset
RDFI - redeem fiat
RLAS - release asset
SEEO - seller execute order
SEAS - send asset
SEFI - send fiat
SECO - send coin > tx.go
CHBB - change buyer bid
CHSB - change seller bid
COBB - confirm buyer bid
COSB - confirm seller bid
SUBF -submit buyer feedback
SUSF - submit seller feedback
*/
