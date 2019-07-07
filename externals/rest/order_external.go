package rest

import (
	"net/http"

	"github.com/ilhammhdd/kudaki-gateway-service/adapters"
	"github.com/ilhammhdd/kudaki-gateway-service/externals/kafka"
)

type RetrieveOwnerOrderHistories struct{}

func (rooh *RetrieveOwnerOrderHistories) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := rooh.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	edha := &adapters.RetrieveOwnerOrderHistories{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, edha).WriteResponse(&w)
}

func (rooh *RetrieveOwnerOrderHistories) validate(r *http.Request) (errs *[]string, ok bool) {
	paramValidation := URLParamValidation{
		Rules: map[string]string{
			"order_status": RegexOrderStatus,
			"offset":       RegexNumber,
			"limit":        RegexNumber},
		Values: r.URL.Query()}

	return paramValidation.Validate()
}

type RetrieveTenantOrderHistories struct{}

func (rtoh *RetrieveTenantOrderHistories) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := rtoh.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	edha := &adapters.RetrieveTenantOrderHistories{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, edha).WriteResponse(&w)
}

func (rtoh *RetrieveTenantOrderHistories) validate(r *http.Request) (errs *[]string, ok bool) {
	paramValidation := URLParamValidation{
		Rules: map[string]string{
			"order_status": RegexOrderStatus,
			"offset":       RegexNumber,
			"limit":        RegexNumber},
		Values: r.URL.Query()}

	return paramValidation.Validate()
}

type TenantReviewOwner struct{}

func (tro *TenantReviewOwner) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := tro.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	edha := &adapters.TenantReviewOwner{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, edha).WriteResponse(&w)
}

func (tro *TenantReviewOwner) validate(r *http.Request) (errs *[]string, ok bool) {
	r.ParseMultipartForm(32 << 20)

	restValidation := RestValidation{
		Rules: map[string]string{
			"rating":     RegexNotEmpty,
			"order_uuid": RegexUUIDV4},
		request: r}

	return restValidation.Validate()
}

type TenantReviewItems struct{}

func (tri *TenantReviewItems) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := tri.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	edha := &adapters.TenantReviewOwner{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, edha).WriteResponse(&w)
}

func (tri *TenantReviewItems) validate(r *http.Request) (errs *[]string, ok bool) {
	r.ParseMultipartForm(32 << 20)

	restValidation := RestValidation{
		Rules: map[string]string{
			"rating":     RegexNotEmpty,
			"review":     RegexNotEmpty,
			"order_uuid": RegexUUIDV4},
		request: r}

	return restValidation.Validate()
}

type ApproveOrder struct{}

func (ao *ApproveOrder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := ao.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	edha := &adapters.ApproveOrder{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, edha).WriteResponse(&w)
}

func (ao *ApproveOrder) validate(r *http.Request) (errs *[]string, ok bool) {
	r.ParseMultipartForm(32 << 20)

	restValidation := RestValidation{
		Rules: map[string]string{
			"order_uuid": RegexUUIDV4},
		request: r}

	return restValidation.Validate()
}

type DisapproveOrder struct{}

func (do *DisapproveOrder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := do.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	edha := &adapters.DisapproveOrder{
		Consumer: kafka.NewConsumption(),
		Producer: kafka.NewProduction()}
	adapters.HandleEventDriven(r, edha).WriteResponse(&w)
}

func (do *DisapproveOrder) validate(r *http.Request) (errs *[]string, ok bool) {
	r.ParseMultipartForm(32 << 20)

	restValidation := RestValidation{
		Rules: map[string]string{
			"order_uuid": RegexUUIDV4},
		request: r}

	return restValidation.Validate()
}