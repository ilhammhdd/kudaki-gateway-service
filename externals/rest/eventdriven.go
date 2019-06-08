package rest

import (
	"net/http"

	"github.com/golang/protobuf/proto"
)

type EventDrivenHandler interface {
	http.Handler
	validate(out proto.Message) (errs *[]string, ok bool)
}
