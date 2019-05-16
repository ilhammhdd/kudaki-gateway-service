package usecases

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-entities/events"
	sarama "gopkg.in/Shopify/sarama.v1"
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
				if sia.Uid == key {
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

type StorefrontItemDeletion struct {
	Esc EventSourceConsumer
	Esp EventSourceProducer
}

func (s StorefrontItemDeletion) DeleteStorefrontItem(key string, msg []byte) *events.StorefrontItemDeleted {

	err := s.produce(key, msg)
	errorkit.ErrorHandled(err)

	storefrontItemDeleted, err := s.consume(key)
	errorkit.ErrorHandled(err)

	return storefrontItemDeleted
}

func (s StorefrontItemDeletion) produce(key string, msg []byte) error {
	s.Esp.Set(events.StoreTopic_name[int32(events.StoreTopic_DELETE_STOREFRONT_ITEM_REQUESTED)])
	start := time.Now()
	prodPart, prodOffset, err := s.Esp.SyncProduce(key, msg)
	duration := time.Since(start)

	log.Printf("produced DeleteStorefrontItemRequested : partition = %d, offset = %d, key = %s, duration = %f seconds", prodPart, prodOffset, key, duration.Seconds())
	return err
}

func (s StorefrontItemDeletion) consume(key string) (*events.StorefrontItemDeleted, error) {
	var storefrontItemDeleted events.StorefrontItemDeleted
	s.Esc.Set(events.StoreTopic_name[int32(events.StoreTopic_STOREFRONT_ITEM_DELETED)], 0, sarama.OffsetNewest)
	cons, sig, closeCons := s.Esc.Consume()
	defer close(closeCons)
	for {
		select {
		case msg := <-cons.Messages():
			if !errorkit.ErrorHandled(proto.Unmarshal(msg.Value, &storefrontItemDeleted)) {
				if key == string(msg.Key) {
					log.Printf("consumed StorefrontItemDeleted : partition = %d, offset = %d, key = %s", msg.Partition, msg.Offset, msg.Key)
					return &storefrontItemDeleted, nil
				}
			}
		case errs := <-cons.Errors():
			return nil, errs.Err
		case <-sig:
			return nil, errors.New("service terminated")
		}
	}
}

type StorefrontItemsRetrieval struct {
	Consumer EventSourceConsumer
	Producer EventSourceProducer
}

func (s StorefrontItemsRetrieval) Retrieve(key string, value []byte) *events.StorefrontItemsRetrieved {
	s.produce(key, value)
	return s.consume(key)
}

func (s StorefrontItemsRetrieval) produce(key string, value []byte) {
	s.Producer.Set(events.StoreTopic_name[int32(events.StoreTopic_RETRIEVE_STOREFRONT_ITEMS_REQUESTED)])
	start := time.Now()
	prodPart, prodOffset, err := s.Producer.SyncProduce(key, value)
	duration := time.Since(start)
	errorkit.ErrorHandled(err)

	log.Printf("produced RetrieveStorefrontItemRequested : partition = %d, offset = %d, key = %s, duration = %f seconds", prodPart, prodOffset, key, duration.Seconds())
}

func (s StorefrontItemsRetrieval) consume(key string) *events.StorefrontItemsRetrieved {

	s.Consumer.Set(events.StoreTopic_name[int32(events.StoreTopic_STOREFRONT_ITEMS_RETRIEVED)], 0, sarama.OffsetNewest)
	partCons, sig, closeChan := s.Consumer.Consume()
	defer close(closeChan)

	var sir events.StorefrontItemsRetrieved

	for {
		select {
		case msg := <-partCons.Messages():
			if unmarshallErr := proto.Unmarshal(msg.Value, &sir); unmarshallErr == nil {
				if string(msg.Key) == key {
					log.Printf("consumed StorefrontItemsRetrieved : partition = %d, offset = %d, key = %s", msg.Partition, msg.Offset, msg.Key)
					return &sir
				}
			}
		case errs := <-partCons.Errors():
			errorkit.ErrorHandled(errs.Err)
		case <-sig:
			return nil
		}
	}
}
