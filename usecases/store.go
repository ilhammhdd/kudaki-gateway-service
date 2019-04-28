package usecases

import (
	"log"
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"

	"gopkg.in/Shopify/sarama.v1"

	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-entities/events"
)

type Store struct {
	Esp EventSourceProducer
	Esc EventSourceConsumer
}

func (s Store) AddStorefrontItem(key string, msg []byte) *events.StorefrontItemAdded {

	s.Esp.Set(events.StoreTopic_name[int32(events.StoreTopic_ADD_STOREFRONT_ITEM_REQUESTED)])
	start := time.Now()
	prodPart, prodOffset, err := s.Esp.SyncProduce(key, msg)
	duration := time.Since(start)
	errorkit.ErrorHandled(err)

	log.Printf("produced AddStorefrontItemRequested : partition = %d, offset = %d, key = %s, duration = %f seconds", prodPart, prodOffset, key, duration.Seconds())

	s.Esc.Set(events.StoreTopic_name[int32(events.StoreTopic_STOREFRONT_ITEM_ADDED)], 0, sarama.OffsetNewest)
	partCons, sig, closeChan := s.Esc.Consume()

	var sia events.StorefrontItemAdded

ConsLoop:
	for {
		select {
		case msg := <-partCons.Messages():
			if !errorkit.ErrorHandled(proto.Unmarshal(msg.Value, &sia)) {
				if sia.Uuid == key {
					log.Printf("consumed StorefrontItemAdded : partition = %d, offset = %d, key = %s", msg.Partition, msg.Offset, msg.Key)
					break ConsLoop
				}
			}
		case consErr := <-partCons.Errors():
			errorkit.ErrorHandled(consErr.Err)
			sia.EventStatus = new(events.Status)
			sia.EventStatus.Errors = []string{consErr.Error()}
			sia.EventStatus.HttpCode = http.StatusInternalServerError
			break ConsLoop
		case <-sig:
			break ConsLoop
		}
	}

	close(closeChan)

	return &sia
}
