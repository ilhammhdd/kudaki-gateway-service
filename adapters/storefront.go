package adapters

import (
	"net/http"
	"strconv"

	"github.com/RediSearch/redisearch-go/redisearch"

	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/google/uuid"
	"github.com/ilhammhdd/kudaki-entities/events"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-gateway-service/usecases"
)

type AddStorefrontItem struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (asi *AddStorefrontItem) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.AddStorefrontItemRequested)

	amount, err := strconv.ParseInt(r.MultipartForm.Value["amount"][0], 10, 32)
	errorkit.ErrorHandled(err)
	price, err := strconv.ParseInt(r.MultipartForm.Value["price"][0], 10, 32)
	errorkit.ErrorHandled(err)

	outEvent.Amount = int32(amount)
	outEvent.Description = r.MultipartForm.Value["description"][0]
	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Name = r.MultipartForm.Value["name"][0]
	outEvent.Photo = r.MultipartForm.Value["photo"][0]
	outEvent.Price = int32(price)
	outEvent.Uid = uuid.New().String()
	outEvent.Unit = r.MultipartForm.Value["unit"][0]

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (asi *AddStorefrontItem) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.StorefrontItemAdded)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &inEvent.EventStatus.Errors
	}

	return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
}

func (asi *AddStorefrontItem) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	inUnmarshal := usecases.KafkaMessageUnmarshal(
		func(key []byte, val []byte) (proto.Message, bool) {
			var inEvent events.StorefrontItemAdded

			if err := proto.Unmarshal(val, &inEvent); err == nil {
				if outKey == string(key) {
					return &inEvent, true
				}
			}
			return nil, false
		})

	return &usecases.EventDrivenUsecase{
		Consumer:    asi.Consumer,
		InTopic:     events.StoreTopic_STOREFRONT_ITEM_ADDED.String(),
		OutTopic:    events.StoreTopic_ADD_STOREFRONT_ITEM_REQUESTED.String(),
		Producer:    asi.Producer,
		InUnmarshal: &inUnmarshal}
}

type UpdateStorefrontItem struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (usi *UpdateStorefrontItem) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.UpdateStorefrontItemRequested)

	amount, err := strconv.ParseInt(r.MultipartForm.Value["amount"][0], 10, 32)
	errorkit.ErrorHandled(err)
	price, err := strconv.ParseInt(r.MultipartForm.Value["price"][0], 10, 32)
	errorkit.ErrorHandled(err)

	outEvent.Amount = int32(amount)
	outEvent.Description = r.MultipartForm.Value["description"][0]
	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Name = r.MultipartForm.Value["name"][0]
	outEvent.Photo = r.MultipartForm.Value["photo"][0]
	outEvent.Price = int32(price)
	outEvent.Uid = uuid.New().String()
	outEvent.Unit = r.MultipartForm.Value["unit"][0]
	outEvent.Uuid = r.MultipartForm.Value["uuid"][0]

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (usi *UpdateStorefrontItem) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.StorefrontItemUpdated)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &inEvent.EventStatus.Errors
	}

	return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
}

func (usi *UpdateStorefrontItem) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	inUnmarshal := usecases.KafkaMessageUnmarshal(
		func(key []byte, val []byte) (proto.Message, bool) {
			var inEvent events.StorefrontItemUpdated

			if err := proto.Unmarshal(val, &inEvent); err == nil {
				if outKey == string(key) {
					return &inEvent, true
				}
			}
			return nil, false
		})

	return &usecases.EventDrivenUsecase{
		Consumer:    usi.Consumer,
		InTopic:     events.StoreTopic_STOREFRONT_ITEM_UPDATED.String(),
		OutTopic:    events.StoreTopic_UPDATE_STOREFRONT_ITEM_REQUESTED.String(),
		Producer:    usi.Producer,
		InUnmarshal: &inUnmarshal}
}

type DeleteStorefrontItem struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (dsi *DeleteStorefrontItem) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.DeleteStorefrontItemRequested)

	outEvent.ItemUuid = r.URL.Query().Get("item_uuid")
	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (dsi *DeleteStorefrontItem) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.StorefrontItemDeleted)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &inEvent.EventStatus.Errors
	}

	return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
}

func (dsi *DeleteStorefrontItem) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	inUnmarshal := usecases.KafkaMessageUnmarshal(
		func(key []byte, val []byte) (proto.Message, bool) {
			var inEvent events.StorefrontItemDeleted

			if err := proto.Unmarshal(val, &inEvent); err == nil {
				if outKey == string(key) {
					return &inEvent, true
				}
			}
			return nil, false
		})

	return &usecases.EventDrivenUsecase{
		Consumer:    dsi.Consumer,
		InTopic:     events.StoreTopic_STOREFRONT_ITEM_DELETED.String(),
		OutTopic:    events.StoreTopic_DELETE_STOREFRONT_ITEM_REQUESTED.String(),
		Producer:    dsi.Producer,
		InUnmarshal: &inUnmarshal}
}

type GetAllUsersStorefrontItemsProcessResult struct {
	Storefront         *redisearch.Document
	StorefrontItemDocs *[]redisearch.Document
}

type GetAllUsersStorefrontItems struct {
	Producer usecases.EventDrivenProducer
}

func (gasi *GetAllUsersStorefrontItems) ParseRequestToKafkaMessage(r *http.Request) (out proto.Message, outKey string, outMsg []byte) {
	outEvent := new(events.RetrieveUsersStorefrontItemsRequested)

	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)
	errorkit.ErrorHandled(err)
	offset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 32)
	errorkit.ErrorHandled(err)

	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Limit = int32(limit)
	outEvent.Offset = int32(offset)
	outEvent.Uuid = uuid.New().String()

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent, outEvent.Uuid, outByte
}

func (gasi *GetAllUsersStorefrontItems) ParseResultToResponse(result interface{}) *Response {
	assertedResult := result.(*GetAllUsersStorefrontItemsProcessResult)

	type responseData struct {
		Storefront map[string]interface{}   `json:"storefront"`
		Items      []map[string]interface{} `json:"items"`
	}

	var storefrontData map[string]interface{}
	var itemsData []map[string]interface{}

	storefrontData = assertedResult.Storefront.Properties
	for i := 0; i < len(*assertedResult.StorefrontItemDocs); i++ {
		itemsData = append(itemsData, (*assertedResult.StorefrontItemDocs)[0].Properties)
	}

	resData := responseData{
		Items:      itemsData,
		Storefront: storefrontData}

	// resData := responseData{
	// 	Items:      assertedResult.StorefrontItemDocs,
	// 	Storefront: assertedResult.Storefront}

	resBody := ResponseBody{Data: resData}
	return NewResponse(http.StatusOK, &resBody)
}

func (gasi *GetAllUsersStorefrontItems) initUseCaseUpstreamHandler(outKey string) usecases.EventDrivenUpstreamHandler {
	return &usecases.EventDrivenUpstreamUsecase{
		OutTopic: events.StoreTopic_RETRIEVE_STOREFRONT_ITEMS_REQUESTED.String(),
		Producer: gasi.Producer}
}
