// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

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
