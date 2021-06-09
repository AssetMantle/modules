package configuration

import (
	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
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
		KafkaBool: flags.KafkaBool,
	}
}
