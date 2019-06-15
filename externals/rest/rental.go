package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-externals/kafka"
	kudakiredisearch "github.com/ilhammhdd/kudaki-externals/redisearch"

	"github.com/ilhammhdd/kudaki-gateway-service/adapters"
)

type AddCartItem struct{}

func (aci *AddCartItem) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := aci.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	adapter := adapters.AddCartItem{Producer: kafka.NewProduction(), Consumer: kafka.NewConsumption()}
	adapters.HandleEventDriven(r, &adapter).WriteResponse(&w)
}

func (aci *AddCartItem) validate(r *http.Request) (errs *[]string, ok bool) {
	restValidation := RestValidation{
		Rules: map[string]string{
			"storefront_uuid": RegexUUIDV4,
			"item_uuid":       RegexUUIDV4,
			"item_amount":     RegexNumber},
		request: r}

	return restValidation.Validate()
}

type RetrieveCartItems struct{}

func (rci *RetrieveCartItems) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := rci.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	adapter := &adapters.RetrieveCartItems{Producer: kafka.NewProduction()}
	adapters.HandleEventDrivenUpstream(r, adapter, rci).WriteResponse(&w)
}

func (rci *RetrieveCartItems) validate(r *http.Request) (errs *[]string, ok bool) {
	urlValidation := URLParamValidation{
		Rules: map[string]string{
			"offset": RegexNumber,
			"limit":  RegexNumber},
		Values: r.URL.Query()}

	return urlValidation.Validate()
}

func (rci *RetrieveCartItems) Process(r *http.Request) (result interface{}) {
	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)
	errorkit.ErrorHandled(err)
	offset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 32)
	errorkit.ErrorHandled(err)

	usr := adapters.GetUserFromKudakiToken(r.Header.Get("Kudaki-Token"))
	cartDoc := rci.retrieveCart(usr.Uuid)
	cartItemsDocs := new([]redisearch.Document)
	if cartDoc != nil {
		cartItemsDocs = rci.retrieveCartItems(cartDoc.Properties["cart_uuid"].(string), int(offset), int(limit))
	}

	return adapters.RetrieveCartItemsResult{
		CartDoc:       cartDoc,
		CartItemsDocs: cartItemsDocs}
}

func (rci *RetrieveCartItems) retrieveCart(userUUID string) *redisearch.Document {
	client := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.Cart.Name())
	client.CreateIndex(kudakiredisearch.Cart.Schema())

	rawQuery := fmt.Sprintf(`@user_uuid:"%s"`, kudakiredisearch.RedisearchText(userUUID).Sanitize())
	docs, total, err := client.Search(redisearch.NewQuery(rawQuery))
	errorkit.ErrorHandled(err)

	if total != 0 {
		return &docs[0]
	}

	return nil
}

func (rci *RetrieveCartItems) retrieveCartItems(sanitizedCartUUID string, offset int, limit int) *[]redisearch.Document {
	client := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.CartItem.Name())
	client.CreateIndex(kudakiredisearch.CartItem.Schema())

	rawQuery := fmt.Sprintf(`@cart_uuid:"%s"`, sanitizedCartUUID)
	docs, total, err := client.Search(redisearch.NewQuery(rawQuery).Limit(offset, limit))
	errorkit.ErrorHandled(err)

	if total != 0 {
		return &docs
	}

	return nil
}