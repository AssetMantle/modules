package kafka

import (
	"github.com/Shopify/sarama"
)

// KafkaAdmin : is admin to create topics
func KafkaAdmin(kafkaPorts []string) sarama.ClusterAdmin {

	config := sarama.NewConfig()
	config.Version = sarama.V0_11_0_0 // hardcoded
	admin, err := sarama.NewClusterAdmin(kafkaPorts, config)
	if err != nil {
		panic(err)
	}
	return admin
}

// TopicsInit : is needed to initialise topics
func TopicsInit(admin sarama.ClusterAdmin, topic string) {

	err := admin.CreateTopic(topic, &topicDetail, true)
	if err != nil {
		panic(err)
	}
}
