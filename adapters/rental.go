package adapters

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/RediSearch/redisearch-go/redisearch"

	"github.com/ilhammhdd/kudaki-entities/rental"

	"github.com/google/uuid"

	"github.com/ilhammhdd/kudaki-entities/events"
	"github.com/ilhammhdd/kudaki-entities/store"

	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/go-toolkit/jwtkit"

	"github.com/golang/protobuf/proto"
	kudaki_entities "github.com/ilhammhdd/kudaki-entities"
)

type checkoutRequest struct {
	CartUUID string `json:"cart_uuid"`
	// Items    []item `json:"items"`
	Items []*store.Item `json:"items"`
}

func ValidateCheckout(cartSchema *redisearch.Schema, cartUUID string) {
	rsClient := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudaki_entities.ClientName_CARTS.String())
	unsanitaryUUID := RedisearchUnsanitary("1fee9956-b3d5-4693-88b8-c31ee03bc5a9")
	rawQuery := fmt.Sprintf(`@uuid:"%s"`, unsanitaryUUID.Sanitize())
	// rawQuery := "@uuid:\"f600d707\""
	query := redisearch.NewQuery(rawQuery).Highlight([]string{"uuid"}, "<b>", "</b>")
	explanation, err := rsClient.Explain(query)
	errorkit.ErrorHandled(err)
	log.Println("explanation : ", explanation)
	docs, total, err := rsClient.Search(query)
	errorkit.ErrorHandled(err)
	log.Printf("search result : total = %d, docs = %v", total, docs)
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

	// check cart by uuid
	ValidateCheckout(c.CartsSchema, cr.CartUUID)
	// check cart by uuid

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
