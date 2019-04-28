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
		resBody := adapters.ResponseBody{
			Errs:    errs,
			Success: valid,
		}
		adapters.NewResponse(http.StatusAccepted, &resBody).WriteResponse(&w)

		return
	}

	log.Printf("request validated, multipart form : type = %T, value = %v", r.MultipartForm.Value, r.MultipartForm.Value)

	adapters.AddStorefrontItem(r, kafka.NewProduction(), kafka.NewConsumption()).WriteResponse(&w)
}
