/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package queuing

import (
	"github.com/Shopify/sarama"
)

// kafkaAdmin : is admin to create topics
func kafkaAdmin(kafkaNodes []string) sarama.ClusterAdmin {
	config := sarama.NewConfig()
	config.Version = sarama.V0_11_0_0 // hardcoded

	admin, Error := sarama.NewClusterAdmin(kafkaNodes, config)
	if Error != nil {
		panic(Error)
	}

	return admin
}
