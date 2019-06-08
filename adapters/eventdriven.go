package adapters

import (
	"net/http"

	"github.com/ilhammhdd/kudaki-gateway-service/usecases"

	"github.com/golang/protobuf/proto"
)

type EventDrivenHandler interface {
	ParseRequestToKafkaMessage(*http.Request) (key string, message []byte)
	ParseEventToResponse(proto.Message) *Response
	initUsecaseHandler(outKey string) usecases.EventDrivenHandler
}

func HandleEventDriven(r *http.Request, edha EventDrivenHandler) *Response {
	outKey, outMsg := edha.ParseRequestToKafkaMessage(r)
	usecaseHandler := edha.initUsecaseHandler(outKey)
	inEvent := usecaseHandler.Handle(outKey, outMsg)
	return edha.ParseEventToResponse(inEvent)
}
