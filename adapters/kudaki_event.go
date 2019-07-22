package adapters

import (
	"encoding/json"
	"log"
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
	adDurationFrom, err := strconv.ParseInt(r.MultipartForm.Value["ad_duration_from"][0], 10, 64)
	errorkit.ErrorHandled(err)
	adDurationTo, err := strconv.ParseInt(r.MultipartForm.Value["ad_duration_to"][0], 10, 64)
	errorkit.ErrorHandled(err)

	outEvent.AdDurationFrom = adDurationFrom
	outEvent.AdDurationTo = adDurationTo
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
	inEvent := in.(*events.KudakiEventDokuInvoiceIssued)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	}

	var responseData struct {
		Amount           float32 `json:"AMOUNT"`
		PurchaseAmount   float32 `json:"PURCHASEAMOUNT"`
		TransIDMerchant  string  `json:"TRANSIDMERCHANT"`
		Words            string  `json:"WORDS"`
		RequestDateTime  int64   `json:"REQUESTDATETIME"`
		Currency         int32   `json:"CURRENCY"`
		PurchaseCurrency int32   `json:"PURCHASECURRENCY"`
		SessionID        string  `json:"SESSIONID"`
		Name             string  `json:"NAME"`
		Email            string  `json:"EMAIL"`
		Basket           string  `json:"BASKET"`
	}

	responseData.Amount = inEvent.DokuInvoice.Amount
	responseData.Basket = inEvent.DokuInvoice.Basket
	responseData.Currency = inEvent.DokuInvoice.Currency
	responseData.Email = inEvent.DokuInvoice.Email
	responseData.Name = inEvent.DokuInvoice.Name
	responseData.PurchaseAmount = inEvent.DokuInvoice.PurchaseAmount
	responseData.PurchaseCurrency = inEvent.DokuInvoice.PurchaseCurrency
	responseData.RequestDateTime = inEvent.DokuInvoice.RequestDateTime
	responseData.SessionID = inEvent.DokuInvoice.SessionId
	responseData.TransIDMerchant = inEvent.DokuInvoice.TransactionIdMerchant
	responseData.Words = inEvent.DokuInvoice.Words

	resBody.Data = responseData

	log.Println("response body data : ", resBody.Data)

	return NewResponse(http.StatusOK, &resBody)
}

func (ake *AddKudakiEvent) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       ake.Consumer,
		InEventChecker: ake,
		InTopic:        events.EventPaymentServiceEventTopic_EVENT_DOKU_INVOICE_ISSUED.String(),
		OutTopic:       events.EventServiceCommandTopic_ADD_KUDAKI_EVENT.String(),
		Producer:       ake.Producer}
}

func (ake *AddKudakiEvent) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.KudakiEventDokuInvoiceIssued
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

type RetrieveOrganizerInvoices struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (ake *RetrieveOrganizerInvoices) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.RetrieveOrganizerInvoices)

	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (ake *RetrieveOrganizerInvoices) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.OrganizerInvoicesRetrieved)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	} else {
		resBody.Data = json.RawMessage(inEvent.Result)
	}

	return NewResponse(http.StatusOK, &resBody)
}

func (ake *RetrieveOrganizerInvoices) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       ake.Consumer,
		InEventChecker: ake,
		InTopic:        events.EventServiceEventTopic_ORGANIZER_INVOICES_RETRIEVED.String(),
		OutTopic:       events.EventServiceCommandTopic_RETRIEVE_ORGANIZER_INVOICES.String(),
		Producer:       ake.Producer}
}

func (ake *RetrieveOrganizerInvoices) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.OrganizerInvoicesRetrieved
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}
