package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"time"

	"github.com/google/uuid"
	"github.com/ilhammhdd/kudaki-entities/events"

	"github.com/ilhammhdd/kudaki-entities/rpc"

	"google.golang.org/grpc"

	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-gateway-service/adapters"
)

const (
	RegexEmail              = `^[A-Za-z0-9](([_\.\-]?[a-zA-Z0-9]+)*)@([A-Za-z0-9]+)(([\.\-]?[a-zA-Z0-9]+)*)\.([A-Za-z]{2,})$`
	RegexEmailErrMessage    = "not a valid email address"
	RegexPassword           = `^[\w_!@#$%*]{6,30}$`
	RegexPasswordErrMessage = "not a valid password, allowed alphanumeric with _!@#$%* symbols, minimal 6 and maximal 30 in length"
	RegexNotEmpty           = `.*`
	RegexNotEmptyErrMessage = "can't be empty"
	RegexURL                = `^((((h)(t)|(f))(t)(p)((s)?))\://)?(www.|[a-zA-Z0-9].)[a-zA-Z0-9\-\.]+\.[a-zA-Z]{2,6}(\:[0-9]{1,5})*(/($|[a-zA-Z0-9\.\,\;\?\'\\\+&amp;%\$#\=~_\-]+))*$`
	RegexURLErrMessage      = "not a valid url"
	RegexJWT                = `^[A-Za-z0-9-_=]+\.[A-Za-z0-9-_=]+\.?[A-Za-z0-9-_.+/=]*$`
	RegexJWTErrMessage      = "not a valid jwt"
)

type RestValidator interface {
	Validate() (errs *[]string, ok bool)
}

type RestValidation struct {
	Rules   map[string]string
	request *http.Request
}

func (rv RestValidation) Validate() (*[]string, bool) {
	rv.request.ParseMultipartForm(50000000)
	valid := true
	var errs []string

	for param, rule := range rv.Rules {
		if rv.request.MultipartForm == nil {
			valid = false
			errs = append(errs, "multipart form required")
		} else if val, ok := rv.request.MultipartForm.Value[param]; !ok {
			log.Println("validating, ", param, "missing")
			valid = false
			errs = append(errs, fmt.Sprintf("%s not exists", param))
		} else {
			regexOk, regexErr := regexp.MatchString(rule, val[0])
			errorkit.ErrorHandled(regexErr)
			if !regexOk {
				log.Println("regex not ok, rule : ", rule, "value : ", val[0])
				valid = false
				switch rule {
				case RegexEmail:
					errs = append(errs, fmt.Sprintf("%s %s", param, RegexEmailErrMessage))
				case RegexPassword:
					errs = append(errs, fmt.Sprintf("%s %s", param, RegexPassword))
				case RegexNotEmpty:
					errs = append(errs, fmt.Sprintf("%s %s", param, RegexNotEmptyErrMessage))
				case RegexURL:
					errs = append(errs, fmt.Sprintf("%s %s", param, RegexURLErrMessage))
				case RegexJWT:
					errs = append(errs, fmt.Sprintf("%s %s", param, RegexJWTErrMessage))
				}
			}
		}
	}

	return &errs, valid
}

func MethodValidator(m string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != m {
			resBody := adapters.ResponseBody{Success: false}
			adapters.NewResponse(http.StatusMethodNotAllowed, &resBody).WriteResponse(&w)

			return
		}

		h.ServeHTTP(w, r)
	})
}

func HeaderParamValidator(rules map[string]string, h http.Header) (*[]string, bool) {

	valid := true
	var errs []string

	for param, rule := range rules {
		if vals, ok := h[param]; !ok {
			log.Println("Http Header, ", param, "missing")
			valid = false
			errs = append(errs, fmt.Sprintf("%s not exists", param))
		} else {
			for _, val := range vals {
				regexOk, regexErr := regexp.MatchString(rule, val)
				errorkit.ErrorHandled(regexErr)
				if !regexOk {
					log.Println("regex not ok, rule : ", rule, "value : ", val[0])
					valid = false
					switch rule {
					case RegexEmail:
						errs = append(errs, fmt.Sprintf("%s %s", param, RegexEmailErrMessage))
					case RegexPassword:
						errs = append(errs, fmt.Sprintf("%s %s", param, RegexPassword))
					case RegexNotEmpty:
						errs = append(errs, fmt.Sprintf("%s %s", param, RegexNotEmptyErrMessage))
					case RegexURL:
						errs = append(errs, fmt.Sprintf("%s %s", param, RegexURLErrMessage))
					case RegexJWT:
						errs = append(errs, fmt.Sprintf("%s %s", param, RegexJWTErrMessage))
					}
				}
			}
		}
	}

	return &errs, valid
}

func URLParamValidator(rules map[string]string, val url.Values) (*[]string, bool) {

	valid := true
	var errs []string

	for param, rule := range rules {
		vals, ok := val[param]
		if !ok && len(val) > 0 {
			log.Println("URLEncodingating, ", param, "missing")
			valid = false
			errs = append(errs, fmt.Sprintf("%s not exists", param))
		} else {
			for _, val := range vals {
				regexOk, regexErr := regexp.MatchString(rule, val)
				errorkit.ErrorHandled(regexErr)
				if !regexOk {
					log.Println("regex not ok, rule : ", rule, "value : ", val[0])
					valid = false
					switch rule {
					case RegexEmail:
						errs = append(errs, fmt.Sprintf("%s %s", param, RegexEmailErrMessage))
					case RegexPassword:
						errs = append(errs, fmt.Sprintf("%s %s", param, RegexPassword))
					case RegexNotEmpty:
						errs = append(errs, fmt.Sprintf("%s %s", param, RegexNotEmptyErrMessage))
					case RegexURL:
						errs = append(errs, fmt.Sprintf("%s %s", param, RegexURLErrMessage))
					case RegexJWT:
						errs = append(errs, fmt.Sprintf("%s %s", param, RegexJWTErrMessage))
					}
				}
			}
		}
	}

	return &errs, valid
}

func Authenticate(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rules := map[string]string{
			"Kudaki-Token": RegexJWT}

		errs, valid := HeaderParamValidator(rules, r.Header)
		if !valid {
			resBody := adapters.ResponseBody{Success: false, Errs: errs}
			adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)

			return
		}

		uar := events.UserAuthenticationRequested{
			Uid: uuid.New().String(),
			Jwt: r.Header.Get("Kudaki-Token")}

		conn, err := grpc.Dial(os.Getenv("USER_SERVICE_GRPC_ADDRESS"), grpc.WithInsecure())
		errorkit.ErrorHandled(err)

		defer conn.Close()

		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		client := rpc.NewUserClient(conn)

		log.Println("calling UserAuthentication grpc, token : ", uar.Jwt)
		ua, err := client.UserAuthentication(ctx, &uar)
		if err != nil {
			resBody := adapters.ResponseBody{Success: false, Errs: &[]string{err.Error()}}
			adapters.NewResponse(http.StatusUnauthorized, &resBody).WriteResponse(&w)

			return
		}

		if ua.EventStatus.HttpCode != http.StatusOK {
			resBody := adapters.ResponseBody{Success: false, Errs: &ua.EventStatus.Errors}
			adapters.NewResponse(int(ua.EventStatus.HttpCode), &resBody).WriteResponse(&w)

			return
		}

		h.ServeHTTP(w, r)
	})
}

func TestAuthenticate(w http.ResponseWriter, r *http.Request) {

	resBody := adapters.ResponseBody{Success: true}
	adapters.NewResponse(http.StatusOK, &resBody).WriteResponse(&w)
}
