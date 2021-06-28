package queuing

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"time"
)

func InitializeKafka(nodeList []string, cliContext context.CLIContext) {
	KafkaState = *NewKafkaState(nodeList)
	if KafkaState.IsEnabled {
		go func() {
			for {
				kafkaConsumerMessages(cliContext)
				time.Sleep(sleepRoutine)
			}
		}()
	}
}
