package adapters

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/RediSearch/redisearch-go/redisearch"

	"github.com/ilhammhdd/kudaki-entities/rental"

	"github.com/google/uuid"

	"github.com/ilhammhdd/kudaki-entities/events"
	"github.com/ilhammhdd/kudaki-entities/store"

	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/go-toolkit/jwtkit"

	"github.com/golang/protobuf/proto"
)

type checkoutRequest struct {
	CartUUID string        `json:"cart_uuid"`
	Items    []*store.Item `json:"items"`
}

type Checkout struct {
	CartsSchema *redisearch.Schema
}

func (c Checkout) ParseRequestToEvent(r *http.Request) proto.Message {
	bodyDecoder := json.NewDecoder(r.Body)

	var cr checkoutRequest
	err := bodyDecoder.Decode(&cr)
	errorkit.ErrorHandled(err)

	jwt, err := jwtkit.GetJWT(jwtkit.JWTString(r.Header.Get("Kudaki-Token")))
	errorkit.ErrorHandled(err)
	var out events.CheckoutRequested
	out.Cart = &rental.Cart{
		Items: cr.Items,
		User:  ParseUserFromJWT(jwt),
		Uuid:  cr.CartUUID,
	}
	out.Uid = uuid.New().String()

	return &out
}

func (c Checkout) ParseEventToKafkaMessage(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.CheckoutRequested)

	outByte, err := proto.Marshal(out)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}

func (c Checkout) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.Checkedout)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != int32(http.StatusOK) {
		resBody.Errs = &inEvent.EventStatus.Errors

		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	}

	resBody.Data = inEvent.Cart

	return NewResponse(http.StatusOK, &resBody)
}

type AddCartItem struct {
	CartItemSchema *redisearch.Schema
}

func (aci *AddCartItem) ParseRequestToEvent(r *http.Request) proto.Message {
	var out events.AddCartItemRequested
	out.Uid = uuid.New().String()
	out.CartUuid = r.MultipartForm.Value["cart_uuid"][0]
	out.StorefrontUuid = r.MultipartForm.Value["storefront_uuid"][0]
	out.ItemUuid = r.MultipartForm.Value["item_uuid"][0]
	itemAmount64, err := strconv.ParseUint(r.MultipartForm.Value["item_amount"][0], 10, 32)
	errorkit.ErrorHandled(err)
	out.ItemAmount = uint32(itemAmount64)

	return &out
}

func (aci *AddCartItem) ParseEventToKafkaMessage(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.AddCartItemRequested)
	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}

func (aci *AddCartItem) ParseEventToResponse(in proto.Message) *Response {
	var resBody ResponseBody

	inEvent := in.(*events.CartItemAdded)
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &inEvent.EventStatus.Errors
	}

	return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
}
