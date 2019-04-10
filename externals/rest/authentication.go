package rest

import (
	"fmt"
	"net/http"

	"github.com/ilhammhdd/kudaki-gateway-service/externals/kafka"

	"github.com/ilhammhdd/kudaki-gateway-service/adapters"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)

	fmt.Println("multipart form : ", r.MultipartForm)

	v := RestValidation{
		Rules: map[string]string{
			"email":     RegexEmail,
			"password":  RegexPassword,
			"role":      RegexNotEmpty,
			"full_name": RegexNotEmpty,
			"photo":     RegexURL},
		request: r,
	}

	if errs, ok := v.Validate(); !ok {
		resBody := adapters.ResponseBody{Success: ok, Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)

		return
	}

	adapters.Signup(r, kafka.NewProduction(), kafka.NewConsumption()).WriteResponse(&w)
}

func VerifyUser(w http.ResponseWriter, r *http.Request) {
	param := "verify_token"
	rules := map[string]string{param: RegexJWT}
	urlVals := r.URL.Query()

	if errs, ok := URLParamValidator(rules, urlVals); !ok {
		resBody := adapters.ResponseBody{Success: ok, Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)

		return
	}

	adapters.VerifyUser(urlVals[param][0], r.Context(), kafka.NewProduction(), kafka.NewConsumption())
}

func Login(w http.ResponseWriter, r *http.Request) {
	restValidation := RestValidation{
		Rules: map[string]string{
			"email":    RegexEmail,
			"password": RegexPassword},
		request: r}

	if errs, valid := restValidation.Validate(); !valid {
		resBody := adapters.ResponseBody{
			Errs:    errs,
			Success: valid}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)

		return
	}

	adapters.Login(r, kafka.NewProduction(), kafka.NewConsumption()).WriteResponse(&w)
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	restValidation := RestValidation{
		Rules: map[string]string{
			"new_password": RegexPassword},
		request: r}

	if errs, ok := restValidation.Validate(); !ok {
		resBody := adapters.ResponseBody{
			Errs:    errs,
			Success: ok}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)

		return
	}

	adapters.ResetPassword(r, kafka.NewProduction(), kafka.NewConsumption())
}
