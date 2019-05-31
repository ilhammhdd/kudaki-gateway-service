package rest

import (
	"fmt"
	"net/http"
	"os"

	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/golang/protobuf/proto"

	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-gateway-service/externals/kudakiredisearch"

	"github.com/ilhammhdd/kudaki-gateway-service/externals/kafka"

	kudaki_entities "github.com/ilhammhdd/kudaki-entities"
	"github.com/ilhammhdd/kudaki-entities/events"
	"github.com/ilhammhdd/kudaki-gateway-service/usecases"

	"github.com/ilhammhdd/kudaki-gateway-service/adapters"
)

type EventDrivenHandler interface {
	http.Handler
	Validate(out proto.Message) (errs *[]string, ok bool)
}

type Checkout struct{}

func (c Checkout) Validate(out proto.Message) (errs *[]string, ok bool) {
	outEvent := out.(*events.CheckoutRequested)
	rsClient := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudaki_entities.ClientName_CARTS.String())
	unsanitaryUUID := adapters.RedisearchText(outEvent.Cart.Uuid)
	rawQuery := fmt.Sprintf(`@uuid:"%s"`, unsanitaryUUID.Sanitize())

	query := redisearch.NewQuery(rawQuery)
	_, total, err := rsClient.Search(query)
	errorkit.ErrorHandled(err)

	var validateErrs []string
	if total == 0 {
		errMsg := "cart with the given uuid not found"
		validateErrs = append(validateErrs, errMsg)
		return &validateErrs, false
	} else if total > 1 {
		errMsg := "duplicate cart uuid"
		validateErrs = append(validateErrs, errMsg)
		return &validateErrs, false
	}
	return nil, true
}

func (c Checkout) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	usecaseProp := usecases.EventDrivenUsecaseProp{
		ConsumerTopic: events.RentalTopic_name[int32(events.RentalTopic_CHECKEDOUT)],
		ProducerTopic: events.RentalTopic_name[int32(events.RentalTopic_CHECKOUT_REQUESTED)],
	}
	usecase := usecases.Checkout{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction(),
	}

	adapterProp := adapters.EventDrivenAdapterProp{
		Request:        r,
		UsecaseHandler: usecase,
		UsecaseProp:    &usecaseProp,
	}
	adapter := adapters.Checkout{
		CartsSchema: kudakiredisearch.CartItemsSchema.Schema(),
	}

	out := adapter.ParseRequestToEvent(r)
	if errs, ok := c.Validate(out); !ok {
		var resBody adapters.ResponseBody
		resBody.Errs = errs

		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	adapters.HandleEventDriven(adapterProp, adapter).WriteResponse(&w)
}

type AddCartItem struct{}

func (aci AddCartItem) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, errs, valid := aci.ValidateRequest(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	aci.Handle(&w, r)
}

func (aci AddCartItem) ValidateRequest(r *http.Request) (existsedParams *[]string, errs *[]string, valid bool) {
	restValidation := RestValidation{
		Rules: map[string]string{
			"cart_uuid":       RegexUUIDV4,
			"storefront_uuid": RegexUUIDV4,
			"item_uuid":       RegexUUIDV4,
			"item_amount":     RegexNumber,
		},
		request: r,
	}

	existsedParams, errs, valid = restValidation.ValidateIfExists()
	return
}

func (aci AddCartItem) Handle(w *http.ResponseWriter, r *http.Request) {

	usecaseProp := usecases.EventDrivenUsecaseProp{
		ConsumerTopic: events.RentalTopic_CART_ITEM_ADDED.String(),
		ProducerTopic: events.RentalTopic_ADD_CART_ITEM_REQUESTED.String(),
	}
	usecase := &usecases.AddCartItem{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction(),
	}

	adapterProp := adapters.EventDrivenAdapterProp{
		Request:        r,
		UsecaseHandler: usecase,
		UsecaseProp:    &usecaseProp,
	}
	adapter := &adapters.AddCartItem{
		CartItemSchema: kudakiredisearch.CartItemsSchema.Schema(),
	}
	adapters.HandleEventDriven(adapterProp, adapter).WriteResponse(w)
}
