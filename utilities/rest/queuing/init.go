package queuing

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"strings"
	"time"
)

func InitializeKafka(cliContext context.CLIContext) {
	KafkaState = *NewKafkaState(strings.Split(strings.Trim(flags.KafkaNodes.ReadCLIValue().(string), "\" "), " "))
	if KafkaState.IsEnabled {
		go func() {
			for {
				kafkaConsumerMessages(cliContext)
				time.Sleep(sleepRoutine)
			}
		}()
	}
}
