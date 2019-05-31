package usecases

import (
	"log"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
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

func Consume(unmarshalProto proto.Message, topic string, key string, consumer EventSourceConsumer) proto.Message {
	consumer.Set(topic, 0, sarama.OffsetNewest)
	partCons, sig, closeChan := consumer.Consume()
	defer close(closeChan)

	for {
		select {
		case msg := <-partCons.Messages():
			log.Printf("consumed : topic = %s, partition = %d, offset = %d, key = %s", msg.Topic, msg.Partition, msg.Offset, string(msg.Key))
			if unmarshallErr := proto.Unmarshal(msg.Value, unmarshalProto); unmarshallErr == nil {
				if string(msg.Key) == (key) {
					return unmarshalProto
				} else {
					log.Printf("out and in key not matched, out = %s, in = %s", key, string(msg.Key))
				}
			} else {
				log.Printf("unmarshal error : %v", unmarshallErr)
			}
		case errs := <-partCons.Errors():
			errorkit.ErrorHandled(errs.Err)
		case <-sig:
			return nil
		}
	}
}
