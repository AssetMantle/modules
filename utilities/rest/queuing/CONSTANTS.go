/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package queuing

import (
	"os"
	"time"
)

// SleepTimer : the time the kafka messages are to be taken in
const SleepTimer = time.Duration(1000000000)

// SleepRoutine : the time the kafka messages are to be taken in
const SleepRoutine = time.Duration(2500000000)

// DefaultCLIHome : is the home path
var DefaultCLIHome = os.ExpandEnv("$HOME/.kafka")

const partition = int32(0)
const offset = int64(0)

// Topics : is list of topics
var Topics = []string{
	"Topic",
}
