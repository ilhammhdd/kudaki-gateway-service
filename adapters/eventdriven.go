package adapters

import (
	"net/http"

	"github.com/ilhammhdd/kudaki-gateway-service/usecases"

	"github.com/golang/protobuf/proto"
)

type EventDrivenHandler interface {
	ParseRequestToEvent(*http.Request) proto.Message
	ParseEventToKafkaMessage(proto.Message) (key string, message []byte)
	ParseEventToResponse(proto.Message) *Response
}

type EventDrivenAdapterProp struct {
	UsecaseHandler usecases.EventDrivenHandler
	Request        *http.Request
	UsecaseProp    *usecases.EventDrivenUsecaseProp
}

func HandleEventDriven(prop EventDrivenAdapterProp, edha EventDrivenHandler) *Response {
	out := edha.ParseRequestToEvent(prop.Request)
	outKey, outMsg := edha.ParseEventToKafkaMessage(out)
	prop.UsecaseProp.ProducerKey = outKey
	prop.UsecaseProp.ProducerMsg = outMsg
	in := usecases.HandleEventDriven(prop.UsecaseProp, prop.UsecaseHandler)
	return edha.ParseEventToResponse(in)
}
