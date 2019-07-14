package adapters

import (
	"net/http"
	"strconv"

	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-gateway-service/usecases"
	"github.com/ilhammhdd/kudaki-gateway-service/usecases/events"
)

type AddKudakiEvent struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (ake *AddKudakiEvent) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.AddKudakiEvent)

	durationFrom, err := strconv.ParseInt(r.MultipartForm.Value["duration_from"][0], 10, 64)
	errorkit.ErrorHandled(err)
	durationTo, err := strconv.ParseInt(r.MultipartForm.Value["duration_to"][0], 10, 64)
	errorkit.ErrorHandled(err)

	outEvent.Description = r.MultipartForm.Value["description"][0]
	outEvent.DurationFrom = durationFrom
	outEvent.DurationTo = durationTo
	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Name = r.MultipartForm.Value["name"][0]
	outEvent.Uid = uuid.New().String()
	outEvent.Venue = r.MultipartForm.Value["venue"][0]

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (ake *AddKudakiEvent) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.KudakiEventAdded)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	}

	return NewResponse(http.StatusOK, &resBody)
}

func (ake *AddKudakiEvent) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       ake.Consumer,
		InEventChecker: ake,
		InTopic:        events.EventServiceEventTopic_KUDAKI_EVENT_ADDED.String(),
		OutTopic:       events.EventServiceCommandTopic_ADD_KUDAKI_EVENT.String(),
		Producer:       ake.Producer}
}

func (ake *AddKudakiEvent) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.KudakiEventAdded
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

// -------------------------------------------------------------------------------------------

type DeleteKudakiEvent struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (ake *DeleteKudakiEvent) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.DeleteKudakiEvent)

	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.EventUuid = r.MultipartForm.Value["event_uuid"][0]
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (ake *DeleteKudakiEvent) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.KudakiEventDeleted)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	}

	return NewResponse(http.StatusOK, &resBody)
}

func (ake *DeleteKudakiEvent) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       ake.Consumer,
		InEventChecker: ake,
		InTopic:        events.EventServiceEventTopic_KUDAKI_EVENT_DELETED.String(),
		OutTopic:       events.EventServiceCommandTopic_DELETE_KUDAKI_EVENT.String(),
		Producer:       ake.Producer}
}

func (ake *DeleteKudakiEvent) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.KudakiEventDeleted
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

// -------------------------------------------------------------------------------------------

type AddPrice struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (ap *AddPrice) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.AddPrice)

	duration, err := strconv.ParseInt(r.MultipartForm.Value["duration"][0], 10, 64)
	errorkit.ErrorHandled(err)

	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Duration = duration
	outEvent.DurationUnit = r.MultipartForm.Value["durataion_unit"][0]
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (ap *AddPrice) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.PriceAdded)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	}

	return NewResponse(http.StatusOK, &resBody)
}

func (ap *AddPrice) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       ap.Consumer,
		InEventChecker: ap,
		InTopic:        events.EventServiceEventTopic_PRICE_ADDED.String(),
		OutTopic:       events.EventServiceCommandTopic_ADD_PRICE.String(),
		Producer:       ap.Producer}
}

func (ap *AddPrice) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.PriceAdded
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}
