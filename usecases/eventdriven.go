package usecases

import (
	"log"
	"time"

	"gopkg.in/Shopify/sarama.v1"

	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/golang/protobuf/proto"
)

// type KafkaMessageUnmarshal func(key []byte, val []byte) (proto.Message, bool)
type InEventChecker interface {
	CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool)
}

type EventDrivenHandler interface {
	Handle(outKey string, outMsg []byte) (inEvent proto.Message)
	produce(outKey string, outMsg []byte)
	consume(outKey string) (inEvent proto.Message)
}

type EventDrivenUsecase struct {
	OutTopic string
	InTopic  string
	Producer EventDrivenProducer
	Consumer EventDrivenConsumer
	// InUnmarshal *KafkaMessageUnmarshal
	InEventChecker InEventChecker
}

func (edu *EventDrivenUsecase) Handle(outKey string, outMsg []byte) (inEvent proto.Message) {
	edu.produce(outKey, outMsg)
	return edu.consume(outKey)
}

func (edu *EventDrivenUsecase) produce(outKey string, outMsg []byte) {
	edu.Producer.Set(edu.OutTopic)
	start := time.Now()
	part, offset, err := edu.Producer.SyncProduce(outKey, outMsg)
	errorkit.ErrorHandled(err)
	duration := time.Since(start)
	log.Printf("produced %s : partition = %d, offset = %d, key = %s, duration = %f seconds", edu.OutTopic, part, offset, outKey, duration.Seconds())
}

func (edu *EventDrivenUsecase) consume(outKey string) proto.Message {
	edu.Consumer.Set(edu.InTopic, 0, sarama.OffsetNewest)
	partCons, sig := edu.Consumer.Consume()

	defer partCons.AsyncClose()

	for {
		select {
		case msg := <-partCons.Messages():
			if inEvent, ok := edu.InEventChecker.CheckInEvent(outKey, msg.Key, msg.Value); ok {
				log.Printf("consumed %s : partition = %d, offset = %d, key = %s", edu.InTopic, msg.Partition, msg.Offset, string(msg.Key))
				return inEvent
			}
		case errs := <-partCons.Errors():
			log.Printf("error while consuming %s : %s", edu.InTopic, errs.Err.Error())
		case <-sig:
			return nil
		}
	}
}
