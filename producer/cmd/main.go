package main

import (
	kafkaConfig "github.com/Sufiyan33/Kafka-go/config"
	"github.com/Sufiyan33/Kafka-go/producer"
)

func main() {
	topic := kafkaConfig.CONST_TOPIC
	producer.Produce(topic, 1000)
}
