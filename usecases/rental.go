package usecases

import (
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-entities/events"
)

type Checkout struct {
	Producer EventSourceProducer
	Consumer EventSourceConsumer
}

func (c Checkout) Consume(topic string, key string) (inEvent proto.Message) {
	return Consume(&events.Checkedout{}, topic, key, c.Consumer).(*events.Checkedout)
}

func (c Checkout) Produce(topic string, key string, value []byte) (partition int32, offset int64, durationSeconds float64) {
	c.Producer.Set(topic)
	start := time.Now()
	partition, offset, err := c.Producer.SyncProduce(key, value)
	durationSeconds = time.Since(start).Seconds()
	errorkit.ErrorHandled(err)

	return
}
