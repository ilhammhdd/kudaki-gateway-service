package usecases

import (
	"os"

	sarama "gopkg.in/Shopify/sarama.v1"
)

type EventSourceKafka struct {
	Esp EventSourceProducer
	Esc EventSourceConsumer
}

type EventSourceProducer interface {
	Set(topic string)
	Get() (topic string)
	SyncProduce(key string, value []byte) (producedPartition int32, producedOffset int64, err error)
}

type EventSourceConsumer interface {
	Set(topic string, partition int32, offset int64)
	Get() (topic string, partition int32, offset int64)
	Consume() (partCons sarama.PartitionConsumer, signals chan os.Signal, close chan bool)
}

type EventDrivenConsumerGroup interface {
	Set(groupID string, topics []string, offset int64)
	Messages() chan *sarama.ConsumerMessage
	Errors() chan error
	Close()
}
