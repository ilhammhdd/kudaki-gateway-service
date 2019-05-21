package rest

import (
	"net/http"
)

type HandlerMiddleware struct {
	Handler http.Handler
}

type MethodRouting struct {
	PostHandler   http.Handler
	PutHandler    http.Handler
	DeleteHandler http.Handler
	GetHandler    http.Handler
	PatchHandler  http.Handler
}

func (mr MethodRouting) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		mr.PostHandler.ServeHTTP(w, r)
	case http.MethodPut:
		mr.PutHandler.ServeHTTP(w, r)
	case http.MethodDelete:
		mr.DeleteHandler.ServeHTTP(w, r)
	case http.MethodGet:
		mr.GetHandler.ServeHTTP(w, r)
	case http.MethodPatch:
		mr.PatchHandler.ServeHTTP(w, r)
	}
}
