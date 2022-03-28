package queuing

import (
	"time"

	"github.com/cosmos/cosmos-sdk/client"
)

func InitializeKafka(nodeList []string, cliContext client.Context) {
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
