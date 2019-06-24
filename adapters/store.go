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
	Total     int                    `json:"total"`
	ItemsDocs *[]redisearch.Document `json:"items"`
}

type RetrieveItems struct {
	Producer usecases.EventDrivenProducer
}

func (ri *RetrieveItems) ParseRequestToKafkaMessage(r *http.Request) (out proto.Message, outKey string, outMsg []byte) {
	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)
	errorkit.ErrorHandled(err)
	offset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 32)
	errorkit.ErrorHandled(err)

	outEvent := new(events.RetrieveItemsRequested)
	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Limit = int32(limit)
	outEvent.Offset = int32(offset)
	outEvent.Uid = uuid.New().String()

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent, outEvent.Uid, outByte
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

type SearchItems struct {
	Producer usecases.EventDrivenProducer
}

func (si *SearchItems) ParseRequestToKafkaMessage(r *http.Request) (out proto.Message, outKey string, outMsg []byte) {
	outEvent := new(events.SearchItemsRequested)

	if len(r.MultipartForm.Value["amount"]) > 0 {
		amount, err := strconv.ParseInt(r.MultipartForm.Value["amount"][0], 10, 32)
		errorkit.ErrorHandled(err)
		outEvent.Amount = int32(amount)
	}
	if len(r.MultipartForm.Value["keyword"]) > 0 {
		outEvent.Keyword = r.MultipartForm.Value["keyword"][0]
	}
	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	if len(r.MultipartForm.Value["limit"]) > 0 {
		limit, err := strconv.ParseInt(r.MultipartForm.Value["limit"][0], 10, 32)
		errorkit.ErrorHandled(err)
		outEvent.Limit = int32(limit)
	}
	if len(r.MultipartForm.Value["offset"]) > 0 {
		offset, err := strconv.ParseInt(r.MultipartForm.Value["offset"][0], 10, 32)
		errorkit.ErrorHandled(err)
		outEvent.Offset = int32(offset)
	}
	if len(r.MultipartForm.Value["min_price"]) > 0 {
		priceFrom, err := strconv.ParseInt(r.MultipartForm.Value["min_price"][0], 10, 32)
		errorkit.ErrorHandled(err)
		outEvent.PriceFrom = int32(priceFrom)
	}
	if len(r.MultipartForm.Value["max_price"]) > 0 {
		priceTo, err := strconv.ParseInt(r.MultipartForm.Value["max_price"][0], 10, 32)
		errorkit.ErrorHandled(err)
		outEvent.PriceTo = int32(priceTo)
	}
	if len(r.MultipartForm.Value["min_rating"]) > 0 {
		ratingFrom, err := strconv.ParseFloat(r.MultipartForm.Value["min_rating"][0], 32)
		errorkit.ErrorHandled(err)
		outEvent.RatingFrom = float32(ratingFrom)
	}
	if len(r.MultipartForm.Value["max_rating"]) > 0 {
		ratingTo, err := strconv.ParseFloat(r.MultipartForm.Value["max_rating"][0], 32)
		errorkit.ErrorHandled(err)
		outEvent.RatingTo = float32(ratingTo)
	}
	outEvent.Uuid = uuid.New().String()

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent, outEvent.Uuid, outByte
}

type SearchItemsProcessResult struct {
	Total     int                    `json:"total"`
	ItemsDocs *[]redisearch.Document `json:"items"`
}

func (si *SearchItems) ParseResultToResponse(result interface{}) *Response {
	assertedResult := result.(SearchItemsProcessResult)
	resBody := ResponseBody{Data: assertedResult}
	return NewResponse(http.StatusOK, &resBody)
}

func (si *SearchItems) initUseCaseUpstreamHandler(outKey string) usecases.EventDrivenUpstreamHandler {
	return &usecases.EventDrivenUpstreamUsecase{
		OutTopic: events.StoreTopic_SEARCH_ITEMS_REQUESTED.String(),
		Producer: si.Producer}
}
