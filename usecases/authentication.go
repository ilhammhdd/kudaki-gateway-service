package usecases

import (
	"context"
	"log"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	entities "github.com/ilhammhdd/kudaki-entities"
	"github.com/ilhammhdd/kudaki-entities/events"
	sarama "gopkg.in/Shopify/sarama.v1"
)

type User struct {
	Esp EventSourceProducer
	Esc EventSourceConsumer
}

func (u User) Signup(ctx context.Context, in *events.SignupRequested) (*events.UserVerificationEmailSent, error) {
	// u.Esp.Set(entities.Topics_name[int32(entities.Topics_USER_command)], int32(commands.Partition_SIGN_UP), sarama.OffsetNewest)
	u.Esp.Set(entities.Topics_name[int32(entities.Topics_SIGN_UP_REQUESTED)], 0, sarama.OffsetNewest)
	start := time.Now()
	partition, offset, err := u.Esp.SyncProduce(in.Uid, in)
	errorkit.ErrorHandled(err)

	duration := time.Since(start)
	log.Println(duration.Seconds(), " seconds passed after producing")

	log.Println("Signup command produced at :", partition, offset)

	u.Esc.Set(entities.Topics_name[int32(entities.Topics_USER_VERIFICATION_EMAIL_SENT)], 0, sarama.OffsetNewest)
	partCons, sig, closeChan := u.Esc.Consume()

	var resultedEvent events.UserVerificationEmailSent
	var eventErr error

ConsLoop:
	for {
		select {
		case msg := <-partCons.Messages():
			err = proto.Unmarshal(msg.Value, &resultedEvent)
			if !errorkit.ErrorHandled(err) {
				if resultedEvent.User.Uuid == in.Profile.User.Uuid {
					break ConsLoop
				}
			}
		case err := <-partCons.Errors():
			eventErr = err.Err
			break ConsLoop
		case <-sig:
			break ConsLoop
		}
	}

	close(closeChan)

	return &resultedEvent, eventErr
}

func (u User) VerifyUser(ctx context.Context, in *events.VerifyUserRequested) (*events.Signedup, error) {

	u.Esp.Set(entities.Topics_name[int32(entities.Topics_VERIFY_USER_REQUESTED)], 0, sarama.OffsetNewest)
	start := time.Now()
	_, _, err := u.Esp.SyncProduce(in.Uid, in)
	errorkit.ErrorHandled(err)
	duration := time.Since(start)
	log.Println(duration.Seconds(), " seconds passed after producing")

	u.Esc.Set(entities.Topics_name[int32(entities.Topics_SIGNED_UP)], 0, sarama.OffsetNewest)
	partCons, sig, closeChan := u.Esc.Consume()

	var resultedEvent events.Signedup
	var eventErr error

ConsLoop:
	for {
		select {
		case msg := <-partCons.Messages():
			err = proto.Unmarshal(msg.Value, &resultedEvent)
			if !errorkit.ErrorHandled(err) {
				if resultedEvent.Uid == in.Uid {
					break ConsLoop
				}
			}
		case err := <-partCons.Errors():
			eventErr = err.Err
			break ConsLoop
		case <-sig:
			break ConsLoop
		}
	}

	close(closeChan)

	return &resultedEvent, eventErr
}

func (u User) Login(ctx context.Context, in *events.LoginRequested) (*events.Loggedin, error) {

	u.Esp.Set(events.User_name[int32(events.User_LOGIN_REQUESTED)], 0, sarama.OffsetNewest)
	_, _, err := u.Esp.SyncProduce(in.Uid, in)
	errorkit.ErrorHandled(err)

	u.Esc.Set(events.User_name[int32(events.User_LOGGED_IN)], 0, sarama.OffsetNewest)
	partCons, sig, closeChan := u.Esc.Consume()

	var loggedin events.Loggedin
	var eventErr error

ConsLoop:
	for {
		select {
		case msg := <-partCons.Messages():
			if !errorkit.ErrorHandled(proto.Unmarshal(msg.Value, &loggedin)) {
				if loggedin.Uid == in.Uid {
					log.Println("Logged in event : ", loggedin)
					break ConsLoop
				}
			}
		case err := <-partCons.Errors():
			errorkit.ErrorHandled(err)
			eventErr = err.Err
			break ConsLoop
		case <-sig:
			break ConsLoop
		}
	}

	close(closeChan)

	return &loggedin, eventErr
}
