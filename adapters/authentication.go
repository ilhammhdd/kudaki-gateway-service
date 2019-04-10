package adapters

import (
	"context"
	"net/http"

	"github.com/ilhammhdd/go-toolkit/errorkit"
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

	user := usecases.User{Esp: esp, Esc: esc}
	uves, err := user.Signup(r.Context(), &su)
	errorkit.ErrorHandled(err)

	var resBody ResponseBody

	if uves.EventStatus.HttpCode != http.StatusOK {
		resBody.Success = false
		resBody.Errs = &uves.EventStatus.Errors
	} else {
		resBody.Success = true
	}

	return NewResponse(int(uves.EventStatus.HttpCode), &resBody)
}

func VerifyUser(jwt string, ctx context.Context, esp usecases.EventSourceProducer, esc usecases.EventSourceConsumer) *Response {

	vu := events.VerifyUserRequested{
		Uid:           uuid.New().String(),
		VerifyUserJwt: jwt}

	userUsecase := usecases.User{Esc: esc, Esp: esp}

	sdu, err := userUsecase.VerifyUser(ctx, &vu)
	errorkit.ErrorHandled(err)

	var resBody ResponseBody

	if sdu.EventStatus.HttpCode != http.StatusOK {
		resBody.Success = false
		resBody.Errs = &sdu.EventStatus.Errors
	} else {
		resBody.Success = true
	}

	return NewResponse(int(sdu.EventStatus.HttpCode), &resBody)
}

func Login(r *http.Request, esp usecases.EventSourceProducer, esc usecases.EventSourceConsumer) *Response {

	lr := events.LoginRequested{
		Uid: uuid.New().String(),
		User: &user.User{
			Email:    r.MultipartForm.Value["email"][0],
			Password: r.MultipartForm.Value["password"][0]}}

	userUsecase := usecases.User{Esp: esp, Esc: esc}

	loggedin, err := userUsecase.Login(r.Context(), &lr)
	errorkit.ErrorHandled(err)

	var resBody ResponseBody

	if loggedin.EventStatus.HttpCode != http.StatusOK {
		resBody.Success = false
		resBody.Errs = &loggedin.EventStatus.Errors
	} else {
		resBody.Success = true
		resBody.Data = &map[string]interface{}{
			"token": loggedin.User.Token}
	}

	return NewResponse(int(loggedin.EventStatus.HttpCode), &resBody)
}

func ResetPassword(r *http.Request, esp usecases.EventSourceProducer, esc usecases.EventSourceConsumer) *Response {

	return nil
}
