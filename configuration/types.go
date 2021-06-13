package configuration

import "github.com/spf13/viper"

type Config struct {
	Kafka KafkaConfig
}

func NewConfig() Config {
	return Config{
		Kafka: NewKafkaConfig(),
	}
}

type KafkaConfig struct {
	KafkaBool bool
}

func NewKafkaConfig() KafkaConfig {
	return KafkaConfig{
		KafkaBool: viper.GetBool("kafka"),
	}
}
