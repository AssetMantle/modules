package queuing

import (
	"time"

	"github.com/cosmos/cosmos-sdk/client/context"
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
