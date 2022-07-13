package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Kafka struct {
	consumer *kafka.Consumer
}

func NewKafkaConsumer() (*Kafka, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":               "localhost:9092,host2:9092",
		"group.id":                        "foo",
		"go.application.rebalance.enable": true})
	if err != nil {
		return nil, err
	}
	return &Kafka{
		consumer: consumer,
	}, nil
}

func (k Kafka) StartConsume() {

	msg_count := 0
	for true {
		ev := k.consumer.Poll(0)
		switch e := ev.(type) {
		case kafka.AssignedPartitions:
			_ = k.consumer.Assign(e.Partitions)
		case kafka.RevokedPartitions:
			_ = k.consumer.Unassign()
		case *kafka.Message:
			msg_count += 1

			/*			if msg_count%MIN_COMMIT_COUNT == 0 {
							k.consumer.Commit()
						}
			*/
			fmt.Printf("%% Message on %s:\n%s\n",
				e.TopicPartition, string(e.Value))

		case kafka.PartitionEOF:
			fmt.Printf("%% Reached %v\n", e)
		case kafka.Error:
			k.StopConsume()
			break
		default:
			fmt.Printf("Ignored %v\n", e)
		}
	}
}

func (k Kafka) StopConsume() {
	_ = k.consumer.Close()
}
