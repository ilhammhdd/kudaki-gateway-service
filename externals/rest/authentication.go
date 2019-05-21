package rest

import (
	"fmt"
	"net/http"

	"github.com/ilhammhdd/kudaki-gateway-service/externals/kafka"

	"github.com/ilhammhdd/kudaki-gateway-service/adapters"
)

func AddTeam(w http.ResponseWriter, r *http.Request) {
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
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)

		return
	}

	adapters.Signup(r, kafka.NewProduction(), kafka.NewConsumption()).WriteResponse(&w)
}

func Signup(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)

	fmt.Println("multipart form : ", r.MultipartForm)

	v := RestValidation{
		Rules: map[string]string{
			"email":     RegexEmail,
			"password":  RegexPassword,
			"role":      RegexRole,
			"full_name": RegexNotEmpty,
			"photo":     RegexURL},
		request: r,
	}

	if errs, ok := v.Validate(); !ok {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)

		return
	}

	adapters.Signup(r, kafka.NewProduction(), kafka.NewConsumption()).WriteResponse(&w)
}

func VerifyUser(w http.ResponseWriter, r *http.Request) {
	param := "verify_token"

	upv := URLParamValidation{
		Rules:  map[string]string{param: RegexJWT},
		Values: r.URL.Query(),
	}

	if errs, ok := upv.Validate(); !ok {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)

		return
	}

	adapters.VerifyUser(r.URL.Query()[param][0], kafka.NewProduction(), kafka.NewConsumption())
}

func Login(w http.ResponseWriter, r *http.Request) {
	restValidation := RestValidation{
		Rules: map[string]string{
			"email":    RegexEmail,
			"password": RegexPassword},
		request: r}

	if errs, valid := restValidation.Validate(); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)

		return
	}

	adapters.Login(r, kafka.NewProduction(), kafka.NewConsumption()).WriteResponse(&w)
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	restValidation := RestValidation{
		Rules: map[string]string{
			"old_password": RegexPassword,
			"new_password": RegexPassword},
		request: r}

	if errs, ok := restValidation.Validate(); !ok {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)

		return
	}

	adapters.ChangePassword(r, kafka.NewProduction(), kafka.NewConsumption()).WriteResponse(&w)
}

func SendResetPasswordEmail(w http.ResponseWriter, r *http.Request) {
	restValidation := RestValidation{
		Rules:   map[string]string{"email": RegexEmail},
		request: r,
	}

	if errs, ok := restValidation.Validate(); !ok {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)

		return
	}

	rped := adapters.ResetPasswordEmailDelivery{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction(),
		Request:  r,
	}

	rped.SendEmail().WriteResponse(&w)
}

func ResetPasswordPage(w http.ResponseWriter, r *http.Request) {

}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	urlValidation := URLParamValidation{
		Rules:  map[string]string{"reset_token": RegexJWT},
		Values: r.URL.Query(),
	}

	if errs, valid := urlValidation.Validate(); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)

		return
	}

	restValidation := RestValidation{
		Rules:   map[string]string{"new_password": RegexPassword},
		request: r,
	}

	if aerrs, valid := restValidation.Validate(); !valid {
		resBody := adapters.ResponseBody{Errs: aerrs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)

		return
	}

	rp := adapters.ResetPassword{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction(),
		Request:  r,
	}
	rp.ResetPassword().WriteResponse(&w)
}
