package usecases

import (
	"log"

	"github.com/golang/protobuf/proto"
)

type EventDrivenHandler interface {
	Consume(topic string, key string) (inEvent proto.Message)
	Produce(topic string, key string, value []byte) (partition int32, offset int64, durationSeconds float64)
}

type EventDrivenUsecaseProp struct {
	ConsumerTopic string
	ProducerTopic string
	ProducerKey   string
	ProducerMsg   []byte
}

func HandleEventDriven(edup *EventDrivenUsecaseProp, edhu EventDrivenHandler) (inEvent proto.Message) {
	partition, offset, duration := edhu.Produce(edup.ProducerTopic, edup.ProducerKey, edup.ProducerMsg)
	log.Printf("produced %s : partition = %d, offset = %d, duration = %f seconds", edup.ProducerTopic, partition, offset, duration)
	return edhu.Consume(edup.ConsumerTopic, edup.ProducerKey)
}
