package flags

import "github.com/persistenceOne/persistenceSDK/schema/helpers/base"

// Note: Arranged alphabetically
var (
	Queuing    = base.NewCLIFlag("queuing", false, "Enable kafka queuing and squashing of transactions")
	KafkaNodes = base.NewCLIFlag("kafkaNodes", "localhost:9092", "Space separated addresses in quotes of the kafka listening node: example: --kafkaPort \"addr1 addr2\" ")
)
