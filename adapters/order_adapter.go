package adapters

import (
	"net/http"
	"strconv"

	"github.com/ilhammhdd/kudaki-gateway-service/entities/aggregates/order"

	"github.com/google/uuid"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-gateway-service/usecases/events"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-gateway-service/usecases"
)

type RetrieveOwnerOrderHistories struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (rooh *RetrieveOwnerOrderHistories) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.RetrieveOwnersOrderHistories)

	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)
	errorkit.ErrorHandled(err)
	offset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 32)
	errorkit.ErrorHandled(err)

	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Limit = int32(limit)
	outEvent.Offset = int32(offset)
	outEvent.OrderStatus = order.OrderStatus(order.OrderStatus_value[r.URL.Query().Get("order_status")])
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (rooh *RetrieveOwnerOrderHistories) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.OwnersOrderHistoriesRetrieved)
	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	}

	return NewResponse(http.StatusOK, &resBody)
}

func (rooh *RetrieveOwnerOrderHistories) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       rooh.Consumer,
		InEventChecker: rooh,
		InTopic:        events.OrderServiceEventTopic_OWNERS_ORDER_HISTORIES_RETRIEVED.String(),
		OutTopic:       events.OrderServiceCommandTopic_RETRIEVE_OWNERS_ORDER_HISTORIES.String(),
		Producer:       rooh.Producer}
}

func (rooh *RetrieveOwnerOrderHistories) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.OwnersOrderHistoriesRetrieved
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

type RetrieveTenantOrderHistories struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (rtoh *RetrieveTenantOrderHistories) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.RetrieveTenantsOrderHistories)

	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)
	errorkit.ErrorHandled(err)
	offset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 32)
	errorkit.ErrorHandled(err)

	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Limit = int32(limit)
	outEvent.Offset = int32(offset)
	outEvent.OrderStatus = order.OrderStatus(order.OrderStatus_value[r.URL.Query().Get("order_status")])
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (rtoh *RetrieveTenantOrderHistories) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.TenantOrderHistoriesRetrieved)
	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	} else {
		resBody.Data = inEvent.Result
	}

	return NewResponse(http.StatusOK, &resBody)
}

func (rtoh *RetrieveTenantOrderHistories) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       rtoh.Consumer,
		InEventChecker: rtoh,
		InTopic:        events.OrderServiceEventTopic_TENANTS_ORDER_HISTORIES_RETRIEVED.String(),
		OutTopic:       events.OrderServiceCommandTopic_RETRIEVE_TENANTS_ORDER_HISTORIES.String(),
		Producer:       rtoh.Producer}
}

func (rtoh *RetrieveTenantOrderHistories) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.TenantOrderHistoriesRetrieved
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

type TenantReviewOwner struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (tro *TenantReviewOwner) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.TenantReviewOwner)

	rating, err := strconv.ParseFloat(r.MultipartForm.Value["rating"][0], 32)
	errorkit.ErrorHandled(err)

	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.OrderUuid = r.MultipartForm.Value["order_uuid"][0]
	outEvent.Rating = float32(rating)
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (tro *TenantReviewOwner) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.TenantReviewedOwner)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	}

	return NewResponse(http.StatusOK, &resBody)
}

func (tro *TenantReviewOwner) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       tro.Consumer,
		InEventChecker: tro,
		InTopic:        events.OrderServiceEventTopic_TENANT_REVIEWED_OWNER.String(),
		OutTopic:       events.OrderServiceCommandTopic_TENANT_REVIEW_OWNER.String(),
		Producer:       tro.Producer}
}

func (tro *TenantReviewOwner) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.TenantReviewedOwner
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

type TenantReviewItems struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (tri *TenantReviewItems) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.TenantReviewItems)

	rating, err := strconv.ParseFloat(r.MultipartForm.Value["rating"][0], 32)
	errorkit.ErrorHandled(err)

	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Rating = float32(rating)
	outEvent.Review = r.MultipartForm.Value["review"][0]
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (tri *TenantReviewItems) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.TenantReviewedItems)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	}

	return NewResponse(http.StatusOK, &resBody)
}

func (tri *TenantReviewItems) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       tri.Consumer,
		InEventChecker: tri,
		InTopic:        events.OrderServiceEventTopic_TENANT_REVIEWED_ITEMS.String(),
		OutTopic:       events.OrderServiceCommandTopic_TENANT_REVIEW_ITEMS.String(),
		Producer:       tri.Producer}
}

func (tri *TenantReviewItems) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.TenantReviewedItems
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

type ApproveOrder struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (ao *ApproveOrder) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.ApproveOrder)

	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.OrderUuid = r.MultipartForm.Value["order_uuid"][0]
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (ao *ApproveOrder) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.OrderApproved)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	}

	return NewResponse(http.StatusOK, &resBody)
}

func (ao *ApproveOrder) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       ao.Consumer,
		InEventChecker: ao,
		InTopic:        events.OrderServiceEventTopic_ORDER_APPROVED.String(),
		OutTopic:       events.OrderServiceCommandTopic_APPROVE_ORDER.String(),
		Producer:       ao.Producer}
}

func (ao *ApproveOrder) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.OrderApproved
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

type DisapproveOrder struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (do *DisapproveOrder) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.DisapproveOrder)

	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.OrderUuid = r.MultipartForm.Value["order_uuid"][0]
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (do *DisapproveOrder) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.OrderDisapproved)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	}

	return NewResponse(http.StatusOK, &resBody)
}

func (do *DisapproveOrder) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       do.Consumer,
		InEventChecker: do,
		InTopic:        events.OrderServiceEventTopic_ORDER_DISAPPROVED.String(),
		OutTopic:       events.OrderServiceCommandTopic_DISAPPROVE_ORDER.String(),
		Producer:       do.Producer}
}

func (do *DisapproveOrder) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.OrderDisapproved
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}
