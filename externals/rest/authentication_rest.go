package rest

import (
	"net/http"

	"github.com/ilhammhdd/kudaki-gateway-service/externals/kafka"

	"github.com/ilhammhdd/kudaki-gateway-service/adapters"
)

type Signup struct{}

func (s *Signup) validate(r *http.Request) (errs *[]string, ok bool) {
	r.ParseMultipartForm(32 << 20)

	v := RestValidation{
		Rules: map[string]string{
			"email":        RegexEmail,
			"password":     RegexPassword,
			"role":         RegexNotEmpty,
			"full_name":    RegexNotEmpty,
			"phone_number": RegexNotEmpty,
			"photo":        RegexURL},
		request: r,
	}

	return v.Validate()
}

func (s *Signup) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, ok := s.validate(r); !ok {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	adapter := &adapters.Signup{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, adapter).WriteResponse(&w)
}
