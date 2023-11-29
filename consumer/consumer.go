package consumer

import (
	"log"

	"github.com/IBM/sarama"
	kafkaConfig "github.com/Sufiyan33/Kafka-go/config"
)

func Consme(topic string) {
	config := sarama.NewConfig()

	consumer, err := sarama.NewConsumer([]string{kafkaConfig.CONST_HOST}, config)
	if err != nil {
		log.Fatal("NewConsumer err : ", err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatal("ConsumerPartition err: ", err)
	}
	defer partitionConsumer.Close()

	// displayting message
	for message := range partitionConsumer.Messages() {
		log.Printf("[Consumer] partitionid: %d; offset:%d, value: %s\n", message.Partition, message.Offset, string(message.Value))
	}
}
