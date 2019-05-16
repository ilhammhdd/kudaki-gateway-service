package adapters

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ilhammhdd/kudaki-entities/events"

	"github.com/golang/protobuf/proto"

	"github.com/google/uuid"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-entities/store"
	"github.com/ilhammhdd/kudaki-entities/user"

	"github.com/ilhammhdd/go-toolkit/jwtkit"

	"github.com/ilhammhdd/kudaki-gateway-service/usecases"
)

func ParseUserFromJWT(jwt *jwtkit.JWT) *user.User {
	usr := jwt.Payload.Claims["user"].(map[string]interface{})
	return &user.User{
		AccountType: user.AccountType(user.AccountType_value[usr["account_type"].(string)]),
		Email:       usr["email"].(string),
		PhoneNumber: usr["phone_number"].(string),
		Role:        user.Role(user.Role_value[usr["role"].(string)]),
		Uuid:        usr["uuid"].(string),
	}
}

func AddStorefrontItem(r *http.Request, esp usecases.EventSourceProducer, esc usecases.EventSourceConsumer) *Response {

	jwt, err := jwtkit.GetJWT(jwtkit.JWTString(r.Header.Get("Kudaki-Token")))
	errorkit.ErrorHandled(err)

	amount, _ := strconv.ParseInt(r.MultipartForm.Value["amount"][0], 10, 32)
	price, _ := strconv.ParseInt(r.MultipartForm.Value["price"][0], 10, 32)
	parsedUser := ParseUserFromJWT(jwt)

	asir := events.AddStorefrontItemRequested{
		Item: &store.Item{
			Amount:      int32(amount),
			Description: r.MultipartForm.Value["description"][0],
			Name:        r.MultipartForm.Value["name"][0],
			Photo:       r.MultipartForm.Value["photo"][0],
			Price:       int32(price),
			Storefront: &store.Storefront{
				UserUuid: parsedUser.Uuid,
			},
			Unit: r.MultipartForm.Value["unit"][0],
			Uuid: uuid.New().String(),
		},
		Uid: uuid.New().String(),
	}

	asirBytes, err := proto.Marshal(&asir)
	errorkit.ErrorHandled(err)

	store := usecases.Store{
		Esc: esc,
		Esp: esp,
	}

	sia := store.AddStorefrontItem(asir.Uid, asirBytes)
	log.Printf("StorefrontItemAdded : type = %T, value = %v", sia, sia)

	var resBody ResponseBody
	if sia.EventStatus.HttpCode != int32(http.StatusOK) {
		resBody.Errs = &sia.EventStatus.Errors
	}

	return NewResponse(int(sia.EventStatus.HttpCode), &resBody)
}

type StorefrontItemDeletion struct {
	Request  *http.Request
	Producer usecases.EventSourceProducer
	Consumer usecases.EventSourceConsumer
}

func (s StorefrontItemDeletion) DeleteStorefrontItem() *Response {

	dsir := s.parseRequestToEvent()

	marshalledDsir, err := proto.Marshal(dsir)
	errorkit.ErrorHandled(err)

	storefrontItemDeletion := usecases.StorefrontItemDeletion{
		Esc: s.Consumer,
		Esp: s.Producer,
	}

	storefrontItemDeletedEvent := storefrontItemDeletion.DeleteStorefrontItem(dsir.Uid, marshalledDsir)

	return s.parseEventToResponse(storefrontItemDeletedEvent)
}

func (s StorefrontItemDeletion) parseRequestToEvent() *events.DeleteStorefrontItemRequested {
	jwt, err := jwtkit.GetJWT(jwtkit.JWTString(s.Request.Header.Get("Kudaki-Token")))
	errorkit.ErrorHandled(err)
	parsedUser := ParseUserFromJWT(jwt)

	dsir := events.DeleteStorefrontItemRequested{
		Item: &store.Item{
			Uuid: s.Request.URL.Query().Get("item_uuid"),
			Storefront: &store.Storefront{
				UserUuid: parsedUser.Uuid,
			},
		},
		Uid: uuid.New().String(),
	}

	return &dsir
}

func (s StorefrontItemDeletion) parseEventToResponse(event *events.StorefrontItemDeleted) *Response {

	var resBody ResponseBody

	if event.EventStatus.HttpCode != int32(http.StatusOK) {
		resBody.Errs = &event.EventStatus.Errors
	}

	return NewResponse(http.StatusOK, &resBody)
}

type StorefrontItemsRetrieval struct {
	Request  *http.Request
	Producer usecases.EventSourceProducer
	Consumer usecases.EventSourceConsumer
}

func (s StorefrontItemsRetrieval) Retrieve() *Response {

	retrieveStorefrontItemRequested := s.parseRequestToEvent()

	rsirMarshalled, err := proto.Marshal(retrieveStorefrontItemRequested)
	errorkit.ErrorHandled(err)

	sir := usecases.StorefrontItemsRetrieval{
		Consumer: s.Consumer,
		Producer: s.Producer,
	}

	sird := sir.Retrieve(retrieveStorefrontItemRequested.Uid, rsirMarshalled)
	return s.parseEventToResponse(sird)
}

func (s StorefrontItemsRetrieval) parseEventToResponse(in *events.StorefrontItemsRetrieved) *Response {
	type responseData struct {
		First int32         `json:"first"`
		Last  int32         `json:"last,omitempty"`
		Limit int32         `json:"limit,omitempty"`
		Items []*store.Item `json:"items,omitempty"`
	}

	var resBody ResponseBody

	if int32(in.EventStatus.HttpCode) != http.StatusOK {
		resBody.Errs = &in.EventStatus.Errors

		return NewResponse(int(in.EventStatus.HttpCode), &resBody)
	}

	log.Printf("first value from event : %v", in.First)

	resData := responseData{
		First: in.First,
		Items: in.Items.Items,
		Limit: in.Limit,
		Last:  in.Last,
	}

	resBody.Data = resData

	return NewResponse(http.StatusOK, &resBody)
}

func (s StorefrontItemsRetrieval) parseRequestToEvent() *events.RetrieveStorefrontItemsRequested {
	jwt, err := jwtkit.GetJWT(jwtkit.JWTString(s.Request.Header.Get("Kudaki-Token")))
	errorkit.ErrorHandled(err)

	usr := ParseUserFromJWT(jwt)

	from, err := strconv.ParseInt(s.Request.URL.Query().Get("from"), 10, 32)
	errorkit.ErrorHandled(err)
	limit, err := strconv.ParseInt(s.Request.URL.Query().Get("limit"), 10, 32)
	errorkit.ErrorHandled(err)

	rsir := events.RetrieveStorefrontItemsRequested{
		Uid:   uuid.New().String(),
		User:  &user.User{Uuid: usr.Uuid},
		From:  int32(from),
		Limit: int32(limit),
	}

	return &rsir
}
