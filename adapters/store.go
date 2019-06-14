package adapters

import (
	"net/http"
	"strconv"

	"github.com/RediSearch/redisearch-go/redisearch"

	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/ilhammhdd/kudaki-entities/events"

	"github.com/ilhammhdd/kudaki-gateway-service/usecases"
)

type RetrieveItemsProcessResult struct {
	Total     int
	ItemsDocs *[]redisearch.Document `json:"items"`
}

type RetrieveItems struct {
	Producer usecases.EventDrivenProducer
}

func (ri *RetrieveItems) ParseRequestToKafkaMessage(r *http.Request) (outKey string, outMsg []byte) {
	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)
	errorkit.ErrorHandled(err)
	offset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 32)
	errorkit.ErrorHandled(err)

	outEvent := new(events.RetrieveItemsRequested)
	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Limit = int32(limit)
	outEvent.Offset = int32(offset)
	outEvent.Uuid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uuid, out
}

func (ri *RetrieveItems) ParseResultToResponse(result interface{}) *Response {
	assertedResult := result.(RetrieveItemsProcessResult)

	resBody := ResponseBody{Data: assertedResult}

	return NewResponse(http.StatusOK, &resBody)
}

func (ri *RetrieveItems) initUseCaseUpstreamHandler(outKey string) usecases.EventDrivenUpstreamHandler {
	return &usecases.EventDrivenUpstreamUsecase{
		OutTopic: events.StoreTopic_RETRIEVE_ITEMS_REQUESTED.String(),
		Producer: ri.Producer}
}
