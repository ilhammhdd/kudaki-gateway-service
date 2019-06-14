package adapters

import (
	"net/http"
	"strconv"

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
