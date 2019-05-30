package rest

import (
	"net/http"

	"github.com/ilhammhdd/kudaki-gateway-service/externals/kudakiredisearch"

	"github.com/ilhammhdd/kudaki-gateway-service/externals/kafka"

	"github.com/ilhammhdd/kudaki-entities/events"
	"github.com/ilhammhdd/kudaki-gateway-service/usecases"

	"github.com/ilhammhdd/kudaki-gateway-service/adapters"
)

func Checkout(w http.ResponseWriter, r *http.Request) {
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
	adapters.HandleEventDriven(adapterProp, adapter).WriteResponse(&w)
}
