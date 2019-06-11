package rest

import (
	"net/http"

	"github.com/ilhammhdd/kudaki-externals/kafka"
	"github.com/ilhammhdd/kudaki-gateway-service/adapters"
)

type AddStorefrontItem struct{}

func (asi *AddStorefrontItem) validate(r *http.Request) (errs *[]string, ok bool) {
	r.ParseMultipartForm(32 << 20)

	restValidation := RestValidation{
		Rules: map[string]string{
			"name":        RegexNotEmpty,
			"amount":      RegexNumber,
			"unit":        RegexNotEmpty,
			"price":       RegexNumber,
			"description": RegexNotEmpty,
			"photo":       RegexNotEmpty},
		request: r}

	return restValidation.Validate()
}

func (asi *AddStorefrontItem) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := asi.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	adapter := adapters.AddStorefrontItem{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}

	adapters.HandleEventDriven(r, &adapter).WriteResponse(&w)
}

type UpdateStorefrontItem struct{}

func (usi *UpdateStorefrontItem) validate(r *http.Request) (errs *[]string, ok bool) {
	restValidation := RestValidation{
		Rules: map[string]string{
			"uuid":        RegexUUIDV4,
			"name":        RegexNotEmpty,
			"amount":      RegexNumber,
			"unit":        RegexNotEmpty,
			"price":       RegexNumber,
			"description": RegexNotEmpty,
			"photo":       RegexURL},
		request: r}

	return restValidation.Validate()
}

func (usi *UpdateStorefrontItem) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := usi.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	adapter := adapters.UpdateStorefrontItem{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}

	adapters.HandleEventDriven(r, &adapter).WriteResponse(&w)
}

type DeleteStorefrontItem struct{}

func (dsi *DeleteStorefrontItem) validate(r *http.Request) (errs *[]string, ok bool) {
	urlValidation := URLParamValidation{
		Rules: map[string]string{
			"item_uuid": RegexUUIDV4},
		Values: r.URL.Query()}

	return urlValidation.Validate()
}

func (dsi *DeleteStorefrontItem) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := dsi.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	adapter := adapters.DeleteStorefrontItem{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}

	adapters.HandleEventDriven(r, &adapter).WriteResponse(&w)
}

type GetAllStorefrontItems struct{}

func (gasi *GetAllStorefrontItems) validate(r *http.Request) (errs *[]string, ok bool) {
	urlValidation := URLParamValidation{
		Rules: map[string]string{
			"from":  RegexNumber,
			"limit": RegexNumber},
		Values: r.URL.Query()}

	return urlValidation.Validate()
}

func (gasi *GetAllStorefrontItems) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := gasi.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	adapter := new(adapters.GetAllStorefrontItems)
	adapters.HandleEventDrivenSource(r, adapter, gasi)
}

func (gasi *GetAllStorefrontItems) Process(r *http.Request) (result interface{}) {

	return nil
}
