package adapters

import (
	"net/http"

	"github.com/ilhammhdd/go-toolkit/safekit"

	"github.com/ilhammhdd/kudaki-gateway-service/usecases"

	"github.com/golang/protobuf/proto"
)

type Event struct {
	Message proto.Message
	Uid     string
}

type EventDrivenSourceProcessor interface {
	Process(r *http.Request) (result interface{})
}

type EventDrivenSourceHandler interface {
	ParseRequestToKafkaMessage(r *http.Request) (outKey string, outMsg []byte)
	ParseResultToResponse(result interface{}) *Response
	initUseCaseSourceHandler(outKey string) usecases.EventDrivenSourceHandler
}

func HandleEventDrivenSource(r *http.Request, edsh EventDrivenSourceHandler, edsp EventDrivenSourceProcessor) *Response {
	outKey, outMsg := edsh.ParseRequestToKafkaMessage(r)
	result := edsp.Process(r)
	safekit.Do(func() {
		usecaseHandler := edsh.initUseCaseSourceHandler(outKey)
		usecaseHandler.Handle(outKey, outMsg)
	})
	return edsh.ParseResultToResponse(result)
}

type EventDrivenHandler interface {
	ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte)
	ParseEventToResponse(in proto.Message) *Response
	initUsecaseHandler(outKey string) usecases.EventDrivenHandler
}

func HandleEventDriven(r *http.Request, edha EventDrivenHandler) *Response {
	outKey, outMsg := edha.ParseRequestToKafkaMessage(r)
	usecaseHandler := edha.initUsecaseHandler(outKey)
	inEvent := usecaseHandler.Handle(outKey, outMsg)
	return edha.ParseEventToResponse(inEvent)
}
