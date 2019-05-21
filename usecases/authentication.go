package usecases

import (
	"log"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-entities/events"
	sarama "gopkg.in/Shopify/sarama.v1"
)

type User struct {
	Esp EventSourceProducer
	Esc EventSourceConsumer
}

func (u User) Signup(key string, msg []byte) (*events.Signedup, error) {
	u.Esp.Set(events.UserTopic_name[int32(events.UserTopic_SIGN_UP_REQUESTED)])
	start := time.Now()
	partition, offset, err := u.Esp.SyncProduce(key, msg)
	errorkit.ErrorHandled(err)
	duration := time.Since(start)

	log.Printf("produced SignupRequested : partition = %d, offset = %d, duration = %f seconds, key = %s", partition, offset, duration.Seconds(), key)

	u.Esc.Set(events.UserTopic_name[int32(events.UserTopic_SIGNED_UP)], 0, sarama.OffsetNewest)
	partCons, sig, closeChan := u.Esc.Consume()

	var sdu events.Signedup
	var eventErr error

ConsLoop:
	for {
		select {
		case msg := <-partCons.Messages():
			err = proto.Unmarshal(msg.Value, &sdu)
			if !errorkit.ErrorHandled(err) {
				if sdu.Uid == key {
					log.Printf("consumed Signedup : partition = %d, offset = %d, key = %s, event UID matched.", msg.Partition, msg.Offset, msg.Key)
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

	return &sdu, eventErr
}

func (u User) VerifyUser(key string, msg []byte) (*events.Signedup, error) {

	u.Esp.Set(events.UserTopic_name[int32(events.UserTopic_VERIFY_USER_REQUESTED)])
	start := time.Now()
	partition, offset, err := u.Esp.SyncProduce(key, msg)
	errorkit.ErrorHandled(err)
	duration := time.Since(start)
	log.Printf("produced VerifyUserRequested : partition = %d, offset = %d, duration = %f seconds, key = %s", partition, offset, duration.Seconds(), key)

	u.Esc.Set(events.UserTopic_name[int32(events.UserTopic_SIGNED_UP)], 0, sarama.OffsetNewest)
	partCons, sig, closeChan := u.Esc.Consume()

	var resultedEvent events.Signedup
	var eventErr error

ConsLoop:
	for {
		select {
		case msg := <-partCons.Messages():
			err = proto.Unmarshal(msg.Value, &resultedEvent)
			if !errorkit.ErrorHandled(err) {
				if resultedEvent.Uid == key {
					log.Printf("consumed Signedup : partition = %d, offset = %d, key = %s, event UID matched.", msg.Partition, msg.Offset, msg.Key)
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

func (u User) Login(key string, msg []byte) (*events.Loggedin, error) {

	u.Esp.Set(events.UserTopic_name[int32(events.UserTopic_LOGIN_REQUESTED)])
	start := time.Now()
	partition, offset, err := u.Esp.SyncProduce(key, msg)
	duration := time.Since(start)
	errorkit.ErrorHandled(err)
	log.Printf("produced LoginRequested : partition = %d, offset = %d, duration = %f seconds, key = %s", partition, offset, duration.Seconds(), key)

	u.Esc.Set(events.UserTopic_name[int32(events.UserTopic_LOGGED_IN)], 0, sarama.OffsetNewest)
	partCons, sig, closeChan := u.Esc.Consume()

	var loggedin events.Loggedin
	var eventErr error

ConsLoop:
	for {
		select {
		case msg := <-partCons.Messages():
			if !errorkit.ErrorHandled(proto.Unmarshal(msg.Value, &loggedin)) {
				if loggedin.Uid == key {
					log.Printf("consumed Loggedin : partition = %d, offset = %d, key = %s, event UID matched.", msg.Partition, msg.Offset, msg.Key)
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

type ChangeUserPassword struct {
	Consumer EventSourceConsumer
	Producer EventSourceProducer
}

func (rup ChangeUserPassword) ChangePassword(key string, msg []byte) (*events.PasswordChanged, error) {

	rup.Producer.Set(events.UserTopic_name[int32(events.UserTopic_CHANGE_PASSWORD_REQUESTED)])
	start := time.Now()
	partition, offset, err := rup.Producer.SyncProduce(key, msg)
	duration := time.Since(start)
	errorkit.ErrorHandled(err)
	log.Printf("produced ChangePasswordRequested : partition = %d, offset = %d, duration = %f seconds, key = %s", partition, offset, duration.Seconds(), key)

	rup.Consumer.Set(events.UserTopic_name[int32(events.UserTopic_PASSWORD_CHANGED)], 0, sarama.OffsetNewest)
	partCons, sig, closeChan := rup.Consumer.Consume()

	var pr events.PasswordChanged
	var returnErr error

ConsLoop:
	for {
		select {
		case msg := <-partCons.Messages():
			if !errorkit.ErrorHandled(proto.Unmarshal(msg.Value, &pr)) {
				if pr.Uid == key {
					log.Printf("consumed PasswordChanged : partition = %d, offset = %d, key = %s", msg.Partition, msg.Offset, msg.Key)
					break ConsLoop
				}
			}
		case consErr := <-partCons.Errors():
			returnErr = consErr.Err
			break ConsLoop
		case <-sig:
			break ConsLoop
		}
	}

	close(closeChan)

	return &pr, returnErr
}

type ResetPasswordEmailDelivery struct {
	Consumer EventSourceConsumer
	Producer EventSourceProducer
}

func (rped ResetPasswordEmailDelivery) SendEmail(key string, msg *[]byte) *events.ResetPasswordEmailSent {

	rped.produce(key, msg)
	return rped.consume(key)
}

func (rped ResetPasswordEmailDelivery) produce(key string, msg *[]byte) {
	rped.Producer.Set(events.UserTopic_name[int32(events.UserTopic_SEND_RESET_PASSWORD_EMAIL_REQUESTED)])
	start := time.Now()
	partition, offset, err := rped.Producer.SyncProduce(key, *msg)
	duration := time.Since(start)
	errorkit.ErrorHandled(err)

	log.Printf("produced SendResetPasswordEmailRequested : partition = %d, offset = %d, duration = %f seconds, key = %s", partition, offset, duration.Seconds(), key)
}

func (rped ResetPasswordEmailDelivery) consume(key string) *events.ResetPasswordEmailSent {
	rped.Consumer.Set(events.UserTopic_name[int32(events.UserTopic_RESET_PASSWORD_EMAIL_SENT)], 0, sarama.OffsetNewest)
	partCons, sig, closeChan := rped.Consumer.Consume()
	defer close(closeChan)

	var rpes events.ResetPasswordEmailSent

	for {
		select {
		case msg := <-partCons.Messages():
			if unmarshallErr := proto.Unmarshal(msg.Value, &rpes); unmarshallErr == nil {
				if string(msg.Key) == key {
					return &rpes
				}
			}
		case errs := <-partCons.Errors():
			errorkit.ErrorHandled(errs.Err)
		case <-sig:
			return nil
		}
	}
}

type ResetPassword struct {
	Consumer EventSourceConsumer
	Producer EventSourceProducer
}

func (rp ResetPassword) Reset(key string, msg *[]byte) *events.PasswordReseted {
	rp.produce(key, msg)

	return rp.consume(key)
}

func (rp ResetPassword) produce(key string, msg *[]byte) {
	rp.Producer.Set(events.UserTopic_name[int32(events.UserTopic_RESET_PASSWORD_REQUESTED)])
	start := time.Now()
	partition, offset, err := rp.Producer.SyncProduce(key, *msg)
	duration := time.Since(start)
	errorkit.ErrorHandled(err)

	log.Printf("produced ResetPasswordRequested : partition = %d, offset = %d, duration = %f seconds, key = %s", partition, offset, duration.Seconds(), key)
}

func (rp ResetPassword) consume(key string) *events.PasswordReseted {
	rp.Consumer.Set(events.UserTopic_name[int32(events.UserTopic_PASSWORD_RESETED)], 0, sarama.OffsetNewest)
	partCons, sig, closeChan := rp.Consumer.Consume()
	defer close(closeChan)

	var rpes events.PasswordReseted

	for {
		select {
		case msg := <-partCons.Messages():
			if unmarshallErr := proto.Unmarshal(msg.Value, &rpes); unmarshallErr == nil {
				if string(msg.Key) == key {
					return &rpes
				}
			}
		case errs := <-partCons.Errors():
			errorkit.ErrorHandled(errs.Err)
		case <-sig:
			return nil
		}
	}
}
