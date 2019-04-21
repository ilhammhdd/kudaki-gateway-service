package rest

import (
	"net/http"

	"github.com/ilhammhdd/kudaki-gateway-service/adapters"
	"github.com/ilhammhdd/kudaki-gateway-service/externals/kafka"
)

func CreateMountain(w http.ResponseWriter, r *http.Request) {
	restValidation := RestValidation{
		Rules: map[string]string{
			"name":        RegexNotEmpty,
			"height":      RegexNumber,
			"latitude":    RegexLatitude,
			"longitude":   RegexLongitude,
			"description": RegexNotEmpty,
		},
		request: r}

	if errs, ok := restValidation.Validate(); !ok {
		resBody := adapters.ResponseBody{
			Errs:    errs,
			Success: ok}

		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
	} else {
		adapters.CreateMountain(r, kafka.NewProduction(), kafka.NewConsumption()).WriteResponse(&w)
	}
}

func RetrieveMountains(w http.ResponseWriter, r *http.Request) {

	adapters.RetrieveMountains(kafka.NewProduction(), kafka.NewConsumption()).WriteResponse(&w)
}
