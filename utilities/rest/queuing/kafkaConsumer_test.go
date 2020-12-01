package queuing

import (
	"github.com/Shopify/sarama"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Kafka_Consumer(t *testing.T) {

	kafkaPorts := []string{"9092"}
	newComsumer := NewConsumer(kafkaPorts)
	config := sarama.NewConfig()
	consumer, _ := sarama.NewConsumer(kafkaPorts, config)
	require.Equal(t, consumer, newComsumer)
	//require.Equal(t,nil,err)

}
