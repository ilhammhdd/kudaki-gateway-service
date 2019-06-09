package usecases

import (
	"log"
	"time"

	"gopkg.in/Shopify/sarama.v1"

	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/golang/protobuf/proto"
)

type KafkaMessageUnmarshal func(key []byte, val []byte) (proto.Message, bool)

type EventDrivenHandler interface {
	Handle(outKey string, outMsg []byte) (inEvent proto.Message)
	produce(outKey string, outMsg []byte)
	consume() (inEvent proto.Message)
}

type EventDrivenUsecase struct {
	OutTopic    string
	InTopic     string
	Producer    EventDrivenProducer
	Consumer    EventDrivenConsumer
	InUnmarshal *KafkaMessageUnmarshal
}

func (edu *EventDrivenUsecase) Handle(outKey string, outMsg []byte) (inEvent proto.Message) {
	edu.produce(outKey, outMsg)
	return edu.consume()
}

func (edu *EventDrivenUsecase) produce(outKey string, outMsg []byte) {
	edu.Producer.Set(edu.OutTopic)
	start := time.Now()
	part, offset, err := edu.Producer.SyncProduce(outKey, outMsg)
	errorkit.ErrorHandled(err)
	duration := time.Since(start)
	log.Printf("produced %s : partition = %d, offset = %d, duration = %f seconds", edu.OutTopic, part, offset, duration.Seconds())
}

func (edu *EventDrivenUsecase) consume() (inEvent proto.Message) {
	edu.Consumer.Set(edu.InTopic, 0, sarama.OffsetNewest)
	partCons, sig, closeChan := edu.Consumer.Consume()
	defer close(closeChan)

	ok := false
	for {
		select {
		case msg := <-partCons.Messages():
			if inEvent, ok = (*edu.InUnmarshal)(msg.Key, msg.Value); ok {
				return
			}
		case errs := <-partCons.Errors():
			log.Printf("error while consuming %s : %s", edu.InTopic, errs.Err.Error())
		case <-sig:
			return nil
		}
	}
}
