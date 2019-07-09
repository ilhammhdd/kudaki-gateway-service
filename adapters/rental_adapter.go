package adapters

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"

	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/ilhammhdd/kudaki-gateway-service/usecases/events"

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
	durationFrom, err := strconv.ParseInt(r.MultipartForm.Value["duration_from"][0], 10, 64)
	errorkit.ErrorHandled(err)
	durationTo, err := strconv.ParseInt(r.MultipartForm.Value["duration_to"][0], 10, 64)
	errorkit.ErrorHandled(err)

	outEvent := new(events.AddCartItem)
	outEvent.DurationFrom = durationFrom
	outEvent.DurationTo = durationTo
	outEvent.ItemAmount = int32(itemAmount)
	outEvent.ItemUuid = r.MultipartForm.Value["item_uuid"][0]
	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Uid = uuid.New().String()

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}

func (aci *AddCartItem) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.CartItemAdded)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &inEvent.EventStatus.Errors
	}

	return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
}

func (aci *AddCartItem) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.CartItemAdded
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

func (aci *AddCartItem) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       aci.Consumer,
		InTopic:        events.RentalServiceEventTopic_CART_ITEM_ADDED.String(),
		InEventChecker: aci,
		OutTopic:       events.RentalServiceCommandTopic_ADD_CART_ITEM.String(),
		Producer:       aci.Producer}
}

type RetrieveCartItems struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (rci *RetrieveCartItems) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)
	errorkit.ErrorHandled(err)
	offset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 32)
	errorkit.ErrorHandled(err)

	outEvent := new(events.RetrieveCartItems)
	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Limit = int32(limit)
	outEvent.Offset = int32(offset)
	outEvent.Uid = uuid.New().String()

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}

func (rci *RetrieveCartItems) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.CartItemsRetrieved)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &inEvent.EventStatus.Errors
	} else {
		resBody.Data = inEvent.Result
	}

	return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
}

func (rci *RetrieveCartItems) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.CartItemsRetrieved
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

func (rci *RetrieveCartItems) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       rci.Consumer,
		InEventChecker: rci,
		InTopic:        events.RentalServiceEventTopic_CART_ITEMS_RETRIEVED.String(),
		OutTopic:/* events.RentalTopic_RETRIEVE_CART_ITEMS_REQUESTED.String() */ events.RentalServiceCommandTopic_RETRIEVE_CART_ITEMS.String(),
		Producer: rci.Producer}
}

type DeleteCartItem struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (dci *DeleteCartItem) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := /* new(events.DeleteCartItemRequested) */ new(events.DeleteCartItem)
	outEvent.CartItemUuid = r.URL.Query().Get("cart_item_uuid")
	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Uid = uuid.New().String()

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}

func (dci *DeleteCartItem) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.CartItemDeleted)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &inEvent.EventStatus.Errors
	}

	return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
}

func (dci *DeleteCartItem) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.CartItemDeleted
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
		return nil, false
	}
	return nil, false
}

func (dci *DeleteCartItem) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer: dci.Consumer,
		InTopic:/* events.RentalTopic_CART_ITEM_DELETED.String() */ events.RentalServiceEventTopic_CART_ITEM_DELETED.String(),
		InEventChecker: dci,
		OutTopic:/* events.RentalTopic_DELETE_CART_ITEM_REQUESTED.String() */ events.RentalServiceCommandTopic_DELETE_CART_ITEM.String(),
		Producer: dci.Producer}
}

type UpdateCartItem struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (uci *UpdateCartItem) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	totalItem, err := strconv.ParseInt(r.URL.Query().Get("total_item"), 10, 32)
	errorkit.ErrorHandled(err)

	outEvent := &events.UpdateCartItem{
		CartItemUuid: r.URL.Query().Get("cart_item_uuid"),
		KudakiToken:  r.Header.Get("Kudaki-Token"),
		TotalItem:    int32(totalItem),
		Uid:          uuid.New().String()}

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
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

func (uci *UpdateCartItem) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.CartItemUpdated
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

func (uci *UpdateCartItem) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer: uci.Consumer,
		InTopic:/* events.RentalTopic_CART_ITEM_UPDATED.String() */ events.RentalServiceEventTopic_CART_ITEM_UPDATED.String(),
		InEventChecker: uci,
		OutTopic:/* events.RentalTopic_UPDATE_CART_ITEM_REQUESTED.String() */ events.RentalServiceCommandTopic_UPDATE_CART_ITEM.String(),
		Producer: uci.Producer}
}

type TenantConfirmReturnment struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (tcr *TenantConfirmReturnment) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.TenantConfirmReturnment)

	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.OrderUuid = r.MultipartForm.Value["order_uuid"][0]
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (tcr *TenantConfirmReturnment) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.TenantConfirmedReturnment)
	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	}

	return NewResponse(http.StatusOK, &resBody)
}

func (tcr *TenantConfirmReturnment) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       tcr.Consumer,
		InEventChecker: tcr,
		InTopic:        events.RentalServiceEventTopic_TENANT_CONFIRMED_RETURNMENT.String(),
		OutTopic:       events.RentalServiceCommandTopic_TENANT_CONFIRM_RETURNMENT.String(),
		Producer:       tcr.Producer}
}

func (tcr *TenantConfirmReturnment) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.TenantConfirmedReturnment
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

type OwnerConfirmReturnment struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (ocr *OwnerConfirmReturnment) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.OwnerConfirmReturnment)

	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.OrderUuid = r.MultipartForm.Value["order_uuid"][0]
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (ocr *OwnerConfirmReturnment) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.OwnerConfirmedReturnment)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	}

	return NewResponse(http.StatusOK, &resBody)
}

func (ocr *OwnerConfirmReturnment) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       ocr.Consumer,
		InEventChecker: ocr,
		InTopic:        events.RentalServiceEventTopic_OWNER_CONFIRMED_RETURNMENT.String(),
		OutTopic:       events.RentalServiceCommandTopic_OWNER_CONFIRM_RETURNMENT.String(),
		Producer:       ocr.Producer}
}

func (ocr *OwnerConfirmReturnment) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.OwnerConfirmedReturnment
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}
