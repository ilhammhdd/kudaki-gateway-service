package usecases

import (
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-entities/events"
	"gopkg.in/Shopify/sarama.v1"
)

type Mountain struct {
	Esp EventSourceProducer
	Esc EventSourceConsumer
}

func (m Mountain) CreateMountain(key string, msg []byte) *events.MountainCreated {
	m.Esp.Set(events.Mountain_name[int32(events.Mountain_CREATE_MOUNTAIN_REQUESTED)])
	_, _, err := m.Esp.SyncProduce(key, msg)
	errorkit.ErrorHandled(err)

	m.Esc.Set(events.Mountain_name[int32(events.Mountain_MOUNTAIN_CREATED)], 0, sarama.OffsetNewest)
	partCons, sig, closeChan := m.Esc.Consume()
	defer close(closeChan)

	var mc events.MountainCreated

	for {
		select {
		case msg := <-partCons.Messages():
			if unmarshalErr := proto.Unmarshal(msg.Value, &mc); unmarshalErr == nil {
				if mc.Uid == key {
					return &mc
				}
			}
		case consErr := <-partCons.Errors():
			mc.EventStatus.Errors = []string{consErr.Error()}
			mc.EventStatus.HttpCode = http.StatusInternalServerError
			mc.Uid = key
			return &mc
		case <-sig:
			return nil
		}
	}
}

func (m Mountain) RetrieveMountains(key string, msg []byte) *events.MountainsRetrieved {

	m.Esp.Set(events.Mountain_name[int32(events.Mountain_RETRIEVE_MOUNTAINS_REQUESTED)])
	partition, offset, err := m.Esp.SyncProduce(key, msg)
	errorkit.ErrorHandled(err)

	log.Printf("produced : partition = %d, offset = %d, key = %s, message = %s", partition, offset, key, msg)

	m.Esc.Set(events.Mountain_name[int32(events.Mountain_MOUNTAINS_RETRIEVED)], 0, sarama.OffsetNewest)
	partCons, sig, closeChan := m.Esc.Consume()

	var mr events.MountainsRetrieved

	for {
		select {
		case msg := <-partCons.Messages():
			log.Printf("consumed : partition = %d, offset = %d, key = %s, message = %s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			if err := proto.Unmarshal(msg.Value, &mr); err == nil {
				if mr.Uid == key {
					log.Println("uid match!")
					partCons.Close()
					close(closeChan)
					return &mr
				}
			}
		case errs := <-partCons.Errors():
			errorkit.ErrorHandled(errs)
		case <-sig:
			return nil
		}
	}
}
