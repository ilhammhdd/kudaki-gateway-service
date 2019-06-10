package adapters

import (
	"log"
	"net/http"

	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/google/uuid"
	"github.com/ilhammhdd/kudaki-entities/events"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-gateway-service/usecases"
)

type Signup struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (s *Signup) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := events.SignupRequested{
		Email:       r.MultipartForm.Value["email"][0],
		FullName:    r.MultipartForm.Value["full_name"][0],
		Password:    r.MultipartForm.Value["password"][0],
		PhoneNumber: r.MultipartForm.Value["phone_number"][0],
		Photo:       r.MultipartForm.Value["photo"][0],
		Role:        r.MultipartForm.Value["role"][0],
		Uid:         uuid.New().String(),
	}
	out, err := proto.Marshal(&outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (s *Signup) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.Signedup)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &inEvent.EventStatus.Errors
	}

	return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
}

func (s *Signup) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	inUnmarshal := usecases.KafkaMessageUnmarshal(
		func(key []byte, val []byte) (proto.Message, bool) {
			var inEvent events.Signedup

			if err := proto.Unmarshal(val, &inEvent); err == nil {
				if outKey == string(key) {
					return &inEvent, true
				}
			}
			return nil, false
		})

	return &usecases.EventDrivenUsecase{
		Consumer:    s.Consumer,
		InTopic:     events.UserTopic_SIGNED_UP.String(),
		OutTopic:    events.UserTopic_SIGN_UP_REQUESTED.String(),
		Producer:    s.Producer,
		InUnmarshal: &inUnmarshal}
}

type Login struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (l *Login) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.LoginRequested)
	outEvent.Email = r.MultipartForm.Value["email"][0]
	outEvent.Password = r.MultipartForm.Value["password"][0]
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (l *Login) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.Loggedin)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &inEvent.EventStatus.Errors
	} else {
		resBody.Data = map[string]string{"token": inEvent.User.Token}
	}

	return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
}

func (l *Login) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	inUnmarshal := usecases.KafkaMessageUnmarshal(func(key []byte, val []byte) (proto.Message, bool) {
		var inEvent events.Loggedin
		if err := proto.Unmarshal(val, &inEvent); err == nil {
			if outKey == string(key) {
				return &inEvent, true
			}
		}

		return nil, false
	})

	return &usecases.EventDrivenUsecase{
		Consumer:    l.Consumer,
		InTopic:     events.UserTopic_LOGGED_IN.String(),
		InUnmarshal: &inUnmarshal,
		OutTopic:    events.UserTopic_LOGIN_REQUESTED.String(),
		Producer:    l.Producer}
}

type VerifyUser struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (vu *VerifyUser) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.VerifyUserRequested)
	outEvent.Uid = uuid.New().String()
	outEvent.VerifyToken = r.URL.Query().Get("verify_token")

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (vu *VerifyUser) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.UserVerified)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &inEvent.EventStatus.Errors
	}
	return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
}

func (vu *VerifyUser) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	inUnmarshal := usecases.KafkaMessageUnmarshal(func(key []byte, val []byte) (proto.Message, bool) {
		var inEvent events.UserVerified
		if err := proto.Unmarshal(val, &inEvent); err == nil {
			if outKey == string(key) {
				return &inEvent, true
			}
		}
		return nil, false
	})

	return &usecases.EventDrivenUsecase{
		Consumer:    vu.Consumer,
		InTopic:     events.UserTopic_USER_VERIFIED.String(),
		InUnmarshal: &inUnmarshal,
		OutTopic:    events.UserTopic_VERIFY_USER_REQUESTED.String(),
		Producer:    vu.Producer,
	}
}

type ChangePassword struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (cp *ChangePassword) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.ChangePasswordRequested)
	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.NewPassword = r.MultipartForm.Value["new_password"][0]
	outEvent.OldPassword = r.MultipartForm.Value["old_password"][0]
	outEvent.Uid = uuid.New().String()

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}

func (cp *ChangePassword) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.PasswordChanged)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &inEvent.EventStatus.Errors
	}
	return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
}

func (cp *ChangePassword) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	inUnmarshal := usecases.KafkaMessageUnmarshal(func(key []byte, val []byte) (proto.Message, bool) {
		var inEvent events.PasswordChanged

		if err := proto.Unmarshal(val, &inEvent); err == nil {
			if outKey == string(key) {
				return &inEvent, true
			}
		}
		return nil, false
	})

	return &usecases.EventDrivenUsecase{
		Consumer:    cp.Consumer,
		InTopic:     events.UserTopic_PASSWORD_CHANGED.String(),
		InUnmarshal: &inUnmarshal,
		OutTopic:    events.UserTopic_CHANGE_PASSWORD_REQUESTED.String(),
		Producer:    cp.Producer}
}

type ResetPasswordSendEmail struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (rpse *ResetPasswordSendEmail) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.SendResetPasswordEmailRequested)
	outEvent.Email = r.MultipartForm.Value["email"][0]
	outEvent.Uid = uuid.New().String()

	log.Println("outEvent : ", outEvent)

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (rpse *ResetPasswordSendEmail) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.ResetPasswordEmailSent)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &inEvent.EventStatus.Errors
	}
	return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
}

func (rpse *ResetPasswordSendEmail) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	inUnmarshal := usecases.KafkaMessageUnmarshal(func(key []byte, val []byte) (proto.Message, bool) {
		var inEvent events.ResetPasswordEmailSent

		if proto.Unmarshal(val, &inEvent) == nil {
			if outKey == string(key) {
				return &inEvent, true
			}
		}

		return nil, false
	})

	edu := usecases.EventDrivenUsecase{
		Consumer:    rpse.Consumer,
		InTopic:     events.UserTopic_RESET_PASSWORD_EMAIL_SENT.String(),
		InUnmarshal: &inUnmarshal,
		OutTopic:    events.UserTopic_SEND_RESET_PASSWORD_EMAIL_REQUESTED.String(),
		Producer:    rpse.Producer}

	return &edu
}

type ResetPassword struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (rp *ResetPassword) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	inUnmarshal := usecases.KafkaMessageUnmarshal(func(key []byte, val []byte) (proto.Message, bool) {
		var inEvent events.PasswordReseted

		if proto.Unmarshal(val, &inEvent) == nil {
			if outKey == string(key) {
				return &inEvent, true
			}
		}
		return nil, false
	})

	return &usecases.EventDrivenUsecase{
		Consumer:    rp.Consumer,
		InTopic:     events.UserTopic_PASSWORD_RESETED.String(),
		InUnmarshal: &inUnmarshal,
		OutTopic:    events.UserTopic_RESET_PASSWORD_REQUESTED.String(),
		Producer:    rp.Producer}
}

func (rp *ResetPassword) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.PasswordReseted)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &inEvent.EventStatus.Errors
	}

	return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
}

func (rp *ResetPassword) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := &events.ResetPasswordRequested{
		NewPassword: r.MultipartForm.Value["new_password"][0],
		ResetToken:  r.URL.Query().Get("reset_token"),
		Uid:         uuid.New().String()}

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}
