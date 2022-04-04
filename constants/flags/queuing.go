// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package flags

import "github.com/AssetMantle/modules/schema/helpers/base"

// Note: Arranged alphabetically
var (
	Queuing    = base.NewCLIFlag("queuing", false, "Enable kafka queuing and squashing of transactions")
	KafkaNodes = base.NewCLIFlag("kafkaNodes", "localhost:9092", "Space separated addresses in quotes of the kafka listening node: example: --kafkaPort \"addr1 addr2\" ")
)
