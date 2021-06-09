package configuration

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

type Config struct {
	Kafka KafkaConfig
}

func NewConfig() Config {
	return Config{
		Kafka: NewKafkaConfig(),
	}
}

type KafkaConfig struct {
	KafkaBool helpers.CLIFlag

}

func NewKafkaConfig() KafkaConfig {
	return KafkaConfig{
		KafkaBool: base.NewCLIFlag("kafka", false, "kafka"),
	}
}
