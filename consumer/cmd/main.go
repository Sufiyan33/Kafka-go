package main

import (
	kafkaConfig "github.com/Sufiyan33/Kafka-go/config"
	"github.com/Sufiyan33/Kafka-go/consumer"
)

func main() {
	topic := kafkaConfig.CONST_TOPIC
	consumer.Consme(topic)
}
