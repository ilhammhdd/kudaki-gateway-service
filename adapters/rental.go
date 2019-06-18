package adapters

import (
	"net/http"
	"strconv"

	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/google/uuid"

	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/ilhammhdd/kudaki-entities/events"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-gateway-service/usecases"
)

type AddCartItem struct {
	Producer usecases.EventDrivenProducer
	Consumer usecases.EventDrivenConsumer
}

func (aci *AddCartItem) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	itemAmount, err := strconv.ParseInt(r.MultipartForm.Value["item_amount"][0], 10, 32)
	errorkit.ErrorHandled(err)

	outEvent := new(events.AddCartItemRequested)
	if len(r.MultipartForm.Value["cart_uuid"]) != 0 {
		outEvent.CartUuid = r.MultipartForm.Value["cart_uuid"][0]
	}
	outEvent.ItemAmount = int32(itemAmount)
	outEvent.ItemUuid = r.MultipartForm.Value["item_uuid"][0]
	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.StorefrontUuid = r.MultipartForm.Value["storefront_uuid"][0]
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (aci *AddCartItem) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.CartItemAdded)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &inEvent.EventStatus.Errors
	}

	return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
}

func (aci *AddCartItem) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	inUnmarshal := usecases.KafkaMessageUnmarshal(func(key []byte, val []byte) (proto.Message, bool) {
		var inEvent events.CartItemAdded
		if proto.Unmarshal(val, &inEvent) == nil {
			if outKey == string(key) {
				return &inEvent, true
			}
		}
		return nil, false
	})

	return &usecases.EventDrivenUsecase{
		Consumer:    aci.Consumer,
		InTopic:     events.RentalTopic_CART_ITEM_ADDED.String(),
		InUnmarshal: &inUnmarshal,
		OutTopic:    events.RentalTopic_ADD_CART_ITEM_REQUESTED.String(),
		Producer:    aci.Producer}
}

type RetrieveCartItems struct {
	Producer usecases.EventDrivenProducer
}

func (rci *RetrieveCartItems) ParseRequestToKafkaMessage(r *http.Request) (outKey string, outMsg []byte) {
	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)
	errorkit.ErrorHandled(err)
	offset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 32)
	errorkit.ErrorHandled(err)

	outEvent := new(events.RetreiveCartItemsRequested)
	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Limit = int32(limit)
	outEvent.Offset = int32(offset)
	outEvent.Uuid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uuid, out
}

type RetrieveCartItemsResult struct {
	CartDoc       *redisearch.Document   `json:"cart"`
	CartItemsDocs *[]redisearch.Document `json:"cart_item"`
}

func (rci *RetrieveCartItems) ParseResultToResponse(result interface{}) *Response {
	resBody := ResponseBody{Data: result}

	return NewResponse(http.StatusOK, &resBody)
}

func (rci *RetrieveCartItems) initUseCaseUpstreamHandler(outKey string) usecases.EventDrivenUpstreamHandler {
	return &usecases.EventDrivenUpstreamUsecase{
		OutTopic: events.RentalTopic_RETRIEVE_CART_ITEMS_REQUESTED.String(),
		Producer: rci.Producer}
}

type DeleteCartItem struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (dci *DeleteCartItem) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.DeleteCartItemRequested)
	outEvent.CartItemUuid = r.URL.Query().Get("cart_item_uuid")
	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (dci *DeleteCartItem) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.CartItemDeleted)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &inEvent.EventStatus.Errors
	}

	return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
}

func (dci *DeleteCartItem) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	inUnmarshal := usecases.KafkaMessageUnmarshal(func(key []byte, val []byte) (proto.Message, bool) {
		var inEvent events.CartItemDeleted
		if proto.Unmarshal(val, &inEvent) == nil {
			if outKey == string(key) {
				return &inEvent, true
			}
			return nil, false
		}
		return nil, false
	})

	return &usecases.EventDrivenUsecase{
		Consumer:    dci.Consumer,
		InTopic:     events.RentalTopic_CART_ITEM_DELETED.String(),
		InUnmarshal: &inUnmarshal,
		OutTopic:    events.RentalTopic_DELETE_CART_ITEM_REQUESTED.String(),
		Producer:    dci.Producer}
}

type UpdateCartItem struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (uci *UpdateCartItem) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	totalItem, err := strconv.ParseInt(r.URL.Query().Get("total_item"), 10, 32)
	errorkit.ErrorHandled(err)

	outEvent := &events.UpdateCartItemRequested{
		CartItemUuid: r.URL.Query().Get("cart_item_uuid"),
		KudakiToken:  r.Header.Get("Kudaki-Token"),
		TotalItem:    int32(totalItem),
		Uid:          uuid.New().String()}

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (uci *UpdateCartItem) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.CartItemUpdated)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	}

	return NewResponse(http.StatusOK, &resBody)
}

func (uci *UpdateCartItem) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	inUnmarshal := usecases.KafkaMessageUnmarshal(func(key []byte, val []byte) (proto.Message, bool) {
		var inEvent events.CartItemUpdated
		if proto.Unmarshal(val, &inEvent) == nil {
			if outKey == string(key) {
				return &inEvent, true
			}
		}
		return nil, false
	})

	return &usecases.EventDrivenUsecase{
		Consumer:    uci.Consumer,
		InTopic:     events.RentalTopic_CART_ITEM_UPDATED.String(),
		InUnmarshal: &inUnmarshal,
		OutTopic:    events.RentalTopic_UPDATE_CART_ITEM_REQUESTED.String(),
		Producer:    uci.Producer}
}
