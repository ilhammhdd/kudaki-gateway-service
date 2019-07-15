package rest

import (
	"log"
	"net/http"

	"github.com/ilhammhdd/kudaki-gateway-service/adapters"
	"github.com/ilhammhdd/kudaki-gateway-service/externals/kafka"
)

type AddKudakiEvent struct{}

func (ake *AddKudakiEvent) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := ake.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	edha := &adapters.AddKudakiEvent{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, edha).WriteResponse(&w)
}

func (ake *AddKudakiEvent) validate(r *http.Request) (errs *[]string, ok bool) {
	r.ParseMultipartForm(32 << 20)

	restValidation := RestValidation{
		Rules: map[string]string{
			"venue":            RegexNotEmpty,
			"description":      RegexNotEmpty,
			"duration_from":    RegexNumber,
			"duration_to":      RegexNumber,
			"name":             RegexNotEmpty,
			"ad_duration_from": RegexNotEmpty,
			"ad_duration_to":   RegexNotEmpty},
		request: r}
	return restValidation.Validate()
}

// -------------------------------------------------------------------------------------------

type DeleteKudakiEvent struct{}

func (dke *DeleteKudakiEvent) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := dke.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	edha := &adapters.DeleteKudakiEvent{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, edha).WriteResponse(&w)
}

func (dke *DeleteKudakiEvent) validate(r *http.Request) (errs *[]string, ok bool) {
	r.ParseMultipartForm(32 << 20)

	restValidation := RestValidation{
		Rules: map[string]string{
			"event_uuid": RegexUUIDV4},
		request: r}
	return restValidation.Validate()
}

// -------------------------------------------------------------------------------------------

type RedirectDoku struct{}

func (ap *RedirectDoku) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := ap.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	log.Println("request header : ", r.Header)
	log.Println("request body : ", r.Body)
}

func (ap *RedirectDoku) validate(r *http.Request) (errs *[]string, ok bool) {
	return nil, true
}

// -------------------------------------------------------------------------------------------

type NotifyDoku struct{}

func (ap *NotifyDoku) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := ap.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	log.Println("request header : ", r.Header)
	log.Println("request body : ", r.Body)
}

func (ap *NotifyDoku) validate(r *http.Request) (errs *[]string, ok bool) {
	return nil, true
}

// -------------------------------------------------------------------------------------------

type IdentifyDoku struct{}

func (ap *IdentifyDoku) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := ap.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	log.Println("request header : ", r.Header)
	log.Println("request body : ", r.Body)
}

func (ap *IdentifyDoku) validate(r *http.Request) (errs *[]string, ok bool) {
	return nil, true
}
