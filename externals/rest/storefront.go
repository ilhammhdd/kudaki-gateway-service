package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/ilhammhdd/kudaki-entities/store"

	"github.com/RediSearch/redisearch-go/redisearch"

	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/go-toolkit/jwtkit"

	"github.com/ilhammhdd/kudaki-entities/user"

	"github.com/ilhammhdd/kudaki-externals/kafka"
	kudakiredisearch "github.com/ilhammhdd/kudaki-externals/redisearch"
	"github.com/ilhammhdd/kudaki-gateway-service/adapters"
)

type AddStorefrontItem struct{}

func (asi *AddStorefrontItem) validate(r *http.Request) (errs *[]string, ok bool) {
	r.ParseMultipartForm(32 << 20)

	restValidation := RestValidation{
		Rules: map[string]string{
			"name":        RegexNotEmpty,
			"amount":      RegexNumber,
			"unit":        RegexNotEmpty,
			"price":       RegexNumber,
			"description": RegexNotEmpty,
			"photo":       RegexNotEmpty},
		request: r}

	return restValidation.Validate()
}

func (asi *AddStorefrontItem) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := asi.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	adapter := adapters.AddStorefrontItem{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}

	adapters.HandleEventDriven(r, &adapter).WriteResponse(&w)
}

type UpdateStorefrontItem struct{}

func (usi *UpdateStorefrontItem) validate(r *http.Request) (errs *[]string, ok bool) {
	restValidation := RestValidation{
		Rules: map[string]string{
			"uuid":        RegexUUIDV4,
			"name":        RegexNotEmpty,
			"amount":      RegexNumber,
			"unit":        RegexNotEmpty,
			"price":       RegexNumber,
			"description": RegexNotEmpty,
			"photo":       RegexURL},
		request: r}

	return restValidation.Validate()
}

func (usi *UpdateStorefrontItem) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := usi.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	adapter := adapters.UpdateStorefrontItem{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}

	adapters.HandleEventDriven(r, &adapter).WriteResponse(&w)
}

type DeleteStorefrontItem struct{}

func (dsi *DeleteStorefrontItem) validate(r *http.Request) (errs *[]string, ok bool) {
	urlValidation := URLParamValidation{
		Rules: map[string]string{
			"item_uuid": RegexUUIDV4},
		Values: r.URL.Query()}

	return urlValidation.Validate()
}

func (dsi *DeleteStorefrontItem) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := dsi.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	adapter := adapters.DeleteStorefrontItem{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}

	adapters.HandleEventDriven(r, &adapter).WriteResponse(&w)
}

type GetAllStorefrontItems struct{}

func (gasi *GetAllStorefrontItems) validate(r *http.Request) (errs *[]string, ok bool) {
	urlValidation := URLParamValidation{
		Rules: map[string]string{
			"offset": RegexNumber,
			"limit":  RegexNumber},
		Values: r.URL.Query()}

	return urlValidation.Validate()
}

func (gasi *GetAllStorefrontItems) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := gasi.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	adapter := &adapters.GetAllStorefrontItems{Producer: kafka.NewProduction()}
	adapters.HandleEventDrivenSource(r, adapter, gasi).WriteResponse(&w)
}

func (gasi *GetAllStorefrontItems) Process(r *http.Request) (result interface{}) {
	offset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 64)
	errorkit.ErrorHandled(err)
	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)
	errorkit.ErrorHandled(err)

	usr := gasi.getUserFromKudakiToken(r.Header.Get("Kudaki-Token"))
	storefront := gasi.retrieveStorefront(usr.Uuid)
	var storefrontItemDocs []redisearch.Document
	if storefront != nil {
		storefrontItemDocs = gasi.retrieveStorefrontItems(int(offset), int(limit), storefront.Uuid)
	}

	return &adapters.GetAllStorefrontItemsProcessResult{
		Storefront:         storefront,
		StorefrontItemDocs: storefrontItemDocs}
}

func (gasi *GetAllStorefrontItems) getUserFromKudakiToken(kudakiToken string) *user.User {
	jwt, err := jwtkit.GetJWT(jwtkit.JWTString(kudakiToken))
	errorkit.ErrorHandled(err)

	userClaim := jwt.Payload.Claims["user"].(map[string]interface{})
	usr := &user.User{
		AccountType: user.AccountType(user.AccountType_value[userClaim["account_type"].(string)]),
		Email:       userClaim["email"].(string),
		PhoneNumber: userClaim["phone_number"].(string),
		Role:        user.Role(user.Role_value[userClaim["role"].(string)]),
		Uuid:        userClaim["uuid"].(string),
	}

	return usr
}

func (gasi *GetAllStorefrontItems) retrieveStorefront(userUUID string) *store.Storefront {
	client := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.Storefront.Name())
	client.CreateIndex(kudakiredisearch.Storefront.Schema())
	rawQuery := fmt.Sprintf(`@user_uuid:"%s"`, kudakiredisearch.RedisearchText(userUUID).Sanitize())
	docs, total, err := client.Search(redisearch.NewQuery(rawQuery))
	errorkit.ErrorHandled(err)

	if total != 0 {
		totalItem, err := strconv.ParseInt(docs[0].Properties["storefront_total_item"].(string), 10, 32)
		errorkit.ErrorHandled(err)
		rating, err := strconv.ParseFloat(docs[0].Properties["storefront_rating"].(string), 10)
		errorkit.ErrorHandled(err)

		var storefront store.Storefront
		unsanitizedStorefrontUUID := kudakiredisearch.RedisearchText(docs[0].Properties["storefront_uuid"].(string)).UnSanitize()
		storefront.Uuid = unsanitizedStorefrontUUID
		storefront.TotalItem = int32(totalItem)
		storefront.Rating = float32(rating)
		return &storefront
	}
	return nil
}

func (gasi *GetAllStorefrontItems) retrieveStorefrontItems(offset int, num int, storefrontUUID string) []redisearch.Document {
	client := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.Item.Name())
	client.CreateIndex(kudakiredisearch.Item.Schema())
	rawQuery := fmt.Sprintf(`@storefront_uuid:"%s"`, kudakiredisearch.RedisearchText(storefrontUUID).Sanitize())
	itemDocs, total, err := client.Search(redisearch.NewQuery(rawQuery).Limit(offset, num))
	errorkit.ErrorHandled(err)

	if total != 0 {
		return itemDocs
	}
	return nil
}
