package rest

import (
	"net/http"

	"github.com/ilhammhdd/kudaki-externals/kafka"

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

type Login struct{}

func (l *Login) validate(r *http.Request) (errs *[]string, ok bool) {
	r.ParseMultipartForm(32 << 20)

	restValidation := RestValidation{
		Rules: map[string]string{
			"email":    RegexEmail,
			"password": RegexPassword},
		request: r,
	}

	return restValidation.Validate()
}

func (l *Login) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, ok := l.validate(r); !ok {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	adapter := adapters.Login{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction(),
	}
	adapters.HandleEventDriven(r, &adapter).WriteResponse(&w)
}

type VerifyUser struct{}

func (vu *VerifyUser) validate(r *http.Request) (errs *[]string, ok bool) {
	paramValidation := URLParamValidation{
		Rules:  map[string]string{"verify_token": RegexJWT},
		Values: r.URL.Query(),
	}
	return paramValidation.Validate()
}

func (vu *VerifyUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, ok := vu.validate(r); !ok {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	adapter := &adapters.VerifyUser{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, adapter).WriteResponse(&w)
}

type ChangePassword struct{}

func (cp *ChangePassword) validate(r *http.Request) (errs *[]string, ok bool) {
	r.ParseMultipartForm(32 << 20)

	restValidation := RestValidation{
		Rules: map[string]string{
			"old_password": RegexPassword,
			"new_password": RegexPassword},
		request: r}

	return restValidation.Validate()
}

func (cp *ChangePassword) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, ok := cp.validate(r); !ok {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	adapter := adapters.ChangePassword{Consumer: kafka.NewConsumption(), Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, &adapter).WriteResponse(&w)
}
