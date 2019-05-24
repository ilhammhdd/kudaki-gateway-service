package adapters

import (
	"log"
	"net/http"
	"net/url"
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

	resData := responseData{Limit: in.Limit}

	if in.Items != nil {
		resData.First = in.First
		resData.Items = in.Items.Items
		resData.Last = in.Last
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

type StorefrontItemUpdate struct {
	Request  *http.Request
	Consumer usecases.EventSourceConsumer
	Producer usecases.EventSourceProducer
}

func (s StorefrontItemUpdate) Update() *Response {

	usir := s.parseRequestToEvent()
	usirBytes, err := proto.Marshal(usir)
	errorkit.ErrorHandled(err)

	siu := usecases.StorefrontItemUpdate{
		Consumer: s.Consumer,
		Key:      &usir.Uid,
		Message:  &usirBytes,
		Producer: s.Producer,
	}
	siud := siu.Update()
	return s.parseEventToResponse(siud)
}

func (s StorefrontItemUpdate) parseRequestToEvent() *events.UpdateStorefrontItemRequested {

	var usir events.UpdateStorefrontItemRequested

	usir.Item = s.parseItem()
	usir.Uid = uuid.New().String()

	jwt, err := jwtkit.GetJWT(jwtkit.JWTString(s.Request.Header.Get("Kudaki-Token")))
	errorkit.ErrorHandled(err)
	usir.User = ParseUserFromJWT(jwt)

	log.Printf("UpdateStorefrontItemRequested event : type = %T, value = %v", usir, usir)

	return &usir
}

func (s StorefrontItemUpdate) parseEventToResponse(in *events.StorefrontItemUpdated) *Response {

	var resBody ResponseBody

	if in.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &in.EventStatus.Errors

		return NewResponse(int(in.EventStatus.HttpCode), &resBody)
	}

	return NewResponse(int(in.EventStatus.HttpCode), &resBody)
}

func (s StorefrontItemUpdate) parseItem() *store.Item {

	var item store.Item

	amount, err := strconv.ParseInt(s.Request.MultipartForm.Value["amount"][0], 10, 32)
	errorkit.ErrorHandled(err)
	item.Amount = int32(amount)
	item.Description = s.Request.MultipartForm.Value["description"][0]
	item.Name = s.Request.MultipartForm.Value["name"][0]
	item.Photo = s.Request.MultipartForm.Value["photo"][0]
	price, err := strconv.ParseInt(s.Request.MultipartForm.Value["price"][0], 10, 32)
	errorkit.ErrorHandled(err)
	item.Price = int32(price)
	item.Unit = s.Request.MultipartForm.Value["unit"][0]
	item.Uuid = s.Request.MultipartForm.Value["uuid"][0]

	return &item
}

type ItemsRetrieval struct {
	Consumer  usecases.EventSourceConsumer
	Producer  usecases.EventSourceProducer
	URLParams url.Values
}

func (ir ItemsRetrieval) Retrieve() *Response {

	key, msg := ir.parseToKafkaMessage()
	irUsecase := usecases.ItemsRetrieval{
		Consumer: ir.Consumer,
		Producer: ir.Producer,
		Key:      key,
		Message:  &msg,
	}
	return ir.parseToResponse(irUsecase.Retrieve())
}

func (ir ItemsRetrieval) parseToKafkaMessage() (string, []byte) {

	var in events.RetrieveItemsRequested

	from, err := strconv.ParseInt(ir.URLParams.Get("from"), 10, 32)
	errorkit.ErrorHandled(err)
	limit, err := strconv.ParseInt(ir.URLParams.Get("limit"), 10, 32)
	errorkit.ErrorHandled(err)

	in.From = int32(from)
	in.Limit = int32(limit)
	in.Uid = uuid.New().String()

	inByte, err := proto.Marshal(&in)
	errorkit.ErrorHandled(err)

	return in.Uid, inByte
}

func (ir ItemsRetrieval) parseToResponse(in *events.ItemsRetrieved) *Response {

	type responseData struct {
		First int32         `json:"first,omitempty"`
		Last  int32         `json:"last,omitempty"`
		Limit int32         `json:"limit,omitempty"`
		Items []*store.Item `json:"items,omitempty"`
	}

	var resBody ResponseBody

	if int32(in.EventStatus.HttpCode) != http.StatusOK {
		resBody.Errs = &in.EventStatus.Errors

		return NewResponse(int(in.EventStatus.HttpCode), &resBody)
	}

	resData := responseData{Limit: in.Limit}

	if in.Items != nil {
		resData.First = in.First
		resData.Items = in.Items.Items
		resData.Last = in.Last
	}

	resBody.Data = resData

	return NewResponse(http.StatusOK, &resBody)
}

type ItemRetrieval struct {
	Consumer usecases.EventSourceConsumer
	Producer usecases.EventSourceProducer
	Request  *http.Request
}

func (ir ItemRetrieval) Retrieve() *Response {

	key, msg := ir.parseToKafkaMessage()

	irUsecase := usecases.ItemRetrieval{
		Consumer: ir.Consumer,
		Producer: ir.Producer,
		OutKey:   key,
		OutMsg:   msg,
	}
	return ir.parseToResponse(irUsecase.Retrieve())
}

func (ir ItemRetrieval) parseToKafkaMessage() (string, *[]byte) {

	var in events.RetrieveItemRequested
	in.ItemUuid = ir.Request.URL.Query().Get("item_uuid")
	in.Uid = uuid.New().String()

	userJWT, err := jwtkit.GetJWT(jwtkit.JWTString(ir.Request.Header.Get("Kudaki-Token")))
	errorkit.ErrorHandled(err)
	in.User = ParseUserFromJWT(userJWT)

	inByte, err := proto.Marshal(&in)
	errorkit.ErrorHandled(err)

	return in.Uid, &inByte
}

func (ir ItemRetrieval) parseToResponse(in *events.ItemRetrieved) *Response {

	var resBody ResponseBody

	if int32(in.EventStatus.HttpCode) != http.StatusOK {
		resBody.Errs = &in.EventStatus.Errors

		return NewResponse(int(in.EventStatus.HttpCode), &resBody)
	}

	resBody.Data = in.Item
	return NewResponse(int(in.EventStatus.HttpCode), &resBody)
}

type ItemSearch struct {
	Consumer usecases.EventSourceConsumer
	Producer usecases.EventSourceProducer
	Request  *http.Request
}

func (is ItemSearch) Search() *Response {

	key, msg := is.parseToKafkaMessage()

	isUsecase := usecases.ItemSearch{
		Consumer: is.Consumer,
		OutKey:   key,
		OutMsg:   msg,
		Producer: is.Producer,
	}

	return is.parseToResponse(isUsecase.Search())
}

func (is ItemSearch) parseToKafkaMessage() (string, *[]byte) {

	var sir events.SearchItemsRequested
	sir.Keyword = is.Request.URL.Query().Get("keyword")
	sir.Uid = uuid.New().String()
	from, err := strconv.ParseUint(is.Request.URL.Query().Get("from"), 10, 64)
	errorkit.ErrorHandled(err)
	limit, err := strconv.ParseInt(is.Request.URL.Query().Get("limit"), 10, 32)
	errorkit.ErrorHandled(err)
	sir.From = from
	sir.Limit = int32(limit)

	userJWT, err := jwtkit.GetJWT(jwtkit.JWTString(is.Request.Header.Get("Kudaki-Token")))
	errorkit.ErrorHandled(err)
	sir.User = ParseUserFromJWT(userJWT)
	sirByte, err := proto.Marshal(&sir)
	errorkit.ErrorHandled(err)

	return sir.Uid, &sirByte
}

func (is ItemSearch) parseToResponse(in *events.ItemsSearched) *Response {

	type ResponseData struct {
		First uint64        `json:"first,omitempty"`
		Last  uint64        `json:"last,omitempty"`
		Limit int32         `json:"limit,omitempty"`
		Items []*store.Item `json:"items,omitempty"`
	}

	var resBody ResponseBody

	if in.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &in.EventStatus.Errors

		return NewResponse(int(in.EventStatus.HttpCode), &resBody)
	}

	var resData ResponseData

	if in.Items != nil {
		resData.First = in.First
		resData.Items = in.Items.Items
		resData.Limit = in.Limit
		resData.Last = in.Last
	}

	resBody.Data = resData

	return NewResponse(int(in.EventStatus.HttpCode), &resBody)
}
