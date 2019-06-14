package rest

import (
	"net/http"

	"github.com/ilhammhdd/kudaki-externals/kafka"

	"github.com/ilhammhdd/kudaki-gateway-service/adapters"
)

type AddCartItem struct{}

func (aci *AddCartItem) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := aci.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	adapter := adapters.AddCartItem{Producer: kafka.NewProduction(), Consumer: kafka.NewConsumption()}
	adapters.HandleEventDriven(r, &adapter).WriteResponse(&w)
}

func (aci *AddCartItem) validate(r *http.Request) (errs *[]string, ok bool) {
	restValidation := RestValidation{
		Rules: map[string]string{
			"storefront_uuid": RegexUUIDV4,
			"item_uuid":       RegexUUIDV4,
			"item_amount":     RegexNumber},
		request: r}

	return restValidation.Validate()
}
