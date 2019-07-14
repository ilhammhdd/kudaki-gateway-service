package rest

import (
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
			"venue":         RegexNotEmpty,
			"description":   RegexNotEmpty,
			"duration_from": RegexNumber,
			"duration_to":   RegexNumber},
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

type AddPrice struct{}

func (ap *AddPrice) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := ap.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	edha := &adapters.AddPrice{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, edha).WriteResponse(&w)
}

func (ap *AddPrice) validate(r *http.Request) (errs *[]string, ok bool) {
	r.ParseMultipartForm(32 << 20)

	restValidation := RestValidation{
		Rules: map[string]string{
			"duration":      RegexNumber,
			"duration_unit": RegexNotEmpty},
		request: r}
	return restValidation.Validate()
}
