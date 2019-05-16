package adapters

import (
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"

	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/go-toolkit/jwtkit"
	"github.com/ilhammhdd/kudaki-entities/events"

	"github.com/ilhammhdd/kudaki-gateway-service/usecases"

	"github.com/google/uuid"
	"github.com/ilhammhdd/kudaki-entities/user"
)

func Signup(r *http.Request, esp usecases.EventSourceProducer, esc usecases.EventSourceConsumer) *Response {
	usr := user.User{
		Uuid:     uuid.New().String(),
		Email:    r.MultipartForm.Value["email"][0],
		Password: r.MultipartForm.Value["password"][0],
		Role:     user.Role(user.Role_value[r.MultipartForm.Value["role"][0]]),
	}

	profile := user.Profile{
		Uuid:       uuid.New().String(),
		User:       &usr,
		FullName:   r.MultipartForm.Value["full_name"][0],
		Photo:      r.MultipartForm.Value["photo"][0],
		Reputation: 0,
	}

	su := events.SignupRequested{
		Uid:     uuid.New().String(),
		Profile: &profile,
	}
	suBytes, err := proto.Marshal(&su)
	errorkit.ErrorHandled(err)

	user := usecases.User{Esp: esp, Esc: esc}
	sdu, err := user.Signup(su.Uid, suBytes)
	errorkit.ErrorHandled(err)

	var resBody ResponseBody

	if sdu.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &sdu.EventStatus.Errors
	}

	return NewResponse(int(sdu.EventStatus.HttpCode), &resBody)
}

func VerifyUser(jwt string, esp usecases.EventSourceProducer, esc usecases.EventSourceConsumer) *Response {

	vu := events.VerifyUserRequested{
		Uid:           uuid.New().String(),
		VerifyUserJwt: jwt}
	vuBytes, err := proto.Marshal(&vu)
	errorkit.ErrorHandled(err)

	userUsecase := usecases.User{Esc: esc, Esp: esp}

	sdu, err := userUsecase.VerifyUser(vu.Uid, vuBytes)
	errorkit.ErrorHandled(err)

	var resBody ResponseBody

	if sdu.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &sdu.EventStatus.Errors
	}

	return NewResponse(int(sdu.EventStatus.HttpCode), &resBody)
}

func Login(r *http.Request, esp usecases.EventSourceProducer, esc usecases.EventSourceConsumer) *Response {

	lr := events.LoginRequested{
		Uid: uuid.New().String(),
		User: &user.User{
			Email:    r.MultipartForm.Value["email"][0],
			Password: r.MultipartForm.Value["password"][0]}}
	lrBytes, err := proto.Marshal(&lr)
	errorkit.ErrorHandled(err)

	userUsecase := usecases.User{Esp: esp, Esc: esc}

	loggedin, err := userUsecase.Login(lr.Uid, lrBytes)
	errorkit.ErrorHandled(err)

	var resBody ResponseBody

	if loggedin.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &loggedin.EventStatus.Errors
	} else {
		resBody.Data = DataMap{"token": string(loggedin.User.Token)}
	}

	return NewResponse(int(loggedin.EventStatus.HttpCode), &resBody)
}

func ResetPassword(r *http.Request, esp usecases.EventSourceProducer, esc usecases.EventSourceConsumer) *Response {

	jwt, err := jwtkit.GetJWT(jwtkit.JWTString(r.Header.Get("Kudaki-Token")))
	errorkit.ErrorHandled(err)

	jwtMap := jwt.Payload.Claims["user"].(map[string]interface{})
	log.Printf("jwt map : type = %T, value = %v", jwtMap, jwtMap)

	rpr := events.ResetPasswordRequested{
		NewPassword: r.MultipartForm.Value["new_password"][0],
		OldPassword: r.MultipartForm.Value["old_password"][0],
		Uid:         uuid.New().String(),
		Profile: &user.Profile{
			User: &user.User{
				Uuid:  "",
				Token: r.Header.Get("Kudaki-Token")}}}
	rprBytes, err := proto.Marshal(&rpr)
	errorkit.ErrorHandled(err)

	usr := usecases.User{
		Esc: esc,
		Esp: esp}

	pr, err := usr.ResetPassword(rpr.Uid, rprBytes)
	errorkit.ErrorHandled(err)

	var resBody ResponseBody

	if pr.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &pr.EventStatus.Errors
	}

	return NewResponse(int(pr.EventStatus.HttpCode), &resBody)
}
