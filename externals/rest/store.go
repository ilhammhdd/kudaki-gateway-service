package rest

import (
	"log"
	"net/http"

	"github.com/ilhammhdd/kudaki-gateway-service/externals/kafka"

	"github.com/ilhammhdd/kudaki-gateway-service/adapters"
)

func AddFrontstoreItem(w http.ResponseWriter, r *http.Request) {
	restValidation := RestValidation{
		Rules: map[string]string{
			"name":        RegexNotEmpty,
			"amount":      RegexNumber,
			"unit":        RegexNotEmpty,
			"price":       RegexNumber,
			"description": RegexNotEmpty,
			"photo":       RegexNotEmpty,
		},
		request: r,
	}

	if errs, valid := restValidation.Validate(); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusAccepted, &resBody).WriteResponse(&w)

		return
	}

	log.Printf("request validated, multipart form : type = %T, value = %v", r.MultipartForm.Value, r.MultipartForm.Value)

	adapters.AddStorefrontItem(r, kafka.NewProduction(), kafka.NewConsumption()).WriteResponse(&w)
}

func DeleteStorefrontItem(w http.ResponseWriter, r *http.Request) {
	upv := URLParamValidation{
		Rules:  map[string]string{"item_uuid": RegexUUIDV4},
		Values: r.URL.Query(),
	}

	if errs, valid := upv.Validate(); !valid {
		resBody := adapters.ResponseBody{Errs: errs}

		adapters.NewResponse(http.StatusBadGateway, &resBody).WriteResponse(&w)
		return
	}

	storefrontItemDeletionAdapter := adapters.StorefrontItemDeletion{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction(),
		Request:  r,
	}

	storefrontItemDeletionAdapter.DeleteStorefrontItem().WriteResponse(&w)
}

func RetrieveStorefrontItems(w http.ResponseWriter, r *http.Request) {

	urlParamValid := URLParamValidation{
		Rules: map[string]string{
			"from":  RegexNumber,
			"limit": RegexNumber,
		},
		Values: r.URL.Query(),
	}

	if errs, valid := urlParamValid.Validate(); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)

		return
	}

	sir := adapters.StorefrontItemsRetrieval{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction(),
		Request:  r,
	}
	sir.Retrieve().WriteResponse(&w)
}

func UpdateStorefrontItem(w http.ResponseWriter, r *http.Request) {
	restValidation := RestValidation{
		Rules: map[string]string{
			"uuid":        RegexUUIDV4,
			"name":        RegexNotEmpty,
			"amount":      RegexNumber,
			"unit":        RegexNotEmpty,
			"price":       RegexNumber,
			"description": RegexNotEmpty,
			"photo":       RegexNotEmpty,
		},
		request: r,
	}

	var resBody adapters.ResponseBody

	if errs, valid := restValidation.Validate(); !valid {
		resBody.Errs = errs
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)

		return
	}

	siu := adapters.StorefrontItemUpdate{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction(),
		Request:  r,
	}
	siu.Update().WriteResponse(&w)
}

func RetrieveItems(w http.ResponseWriter, r *http.Request) {
	urlParamValid := URLParamValidation{
		Rules: map[string]string{
			"from":  RegexNumber,
			"limit": RegexNumber,
		},
		Values: r.URL.Query()}

	if errs, valid := urlParamValid.Validate(); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)

		return
	}

	ir := adapters.ItemsRetrieval{
		Consumer:  kafka.NewConsumption(),
		Producer:  kafka.NewProduction(),
		URLParams: r.URL.Query(),
	}

	ir.Retrieve().WriteResponse(&w)
}

func RetrieveItem(w http.ResponseWriter, r *http.Request) {
	urlValidation := URLParamValidation{
		Rules:  map[string]string{"item_uuid": RegexUUIDV4},
		Values: r.URL.Query(),
	}

	if errs, valid := urlValidation.Validate(); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)

		return
	}

	irAdapter := adapters.ItemRetrieval{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction(),
		Request:  r,
	}
	irAdapter.Retrieve().WriteResponse(&w)
}

func SearchItems(w http.ResponseWriter, r *http.Request) {

}
