package rest

import (
	"net/http"
	"os"
	"strconv"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-externals/kafka"

	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/RediSearch/redisearch-go/redisearch"
	kudakiredisearch "github.com/ilhammhdd/kudaki-externals/redisearch"

	"github.com/ilhammhdd/kudaki-gateway-service/adapters"
)

type RetrieveItems struct{}

func (ri *RetrieveItems) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := ri.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}

	adapter := adapters.RetrieveItems{Producer: kafka.NewProduction()}
	adapters.HandleEventDrivenUpstream(r, &adapter, ri).WriteResponse(&w)
}

func (ri *RetrieveItems) validate(r *http.Request) (errs *[]string, ok bool) {
	urlValidation := URLParamValidation{
		Rules: map[string]string{
			"offset": RegexNumber,
			"limit":  RegexNumber},
		Values: r.URL.Query()}

	return urlValidation.Validate()
}

func (ri *RetrieveItems) Process(r *http.Request, out proto.Message) (result interface{}) {
	offset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 32)
	errorkit.ErrorHandled(err)
	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)
	errorkit.ErrorHandled(err)

	client := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.Item.Name())
	client.CreateIndex(kudakiredisearch.Item.Schema())

	itemDocs, total, err := client.Search(redisearch.NewQuery(`*`).Limit(int(offset), int(limit)))
	errorkit.ErrorHandled(err)

	return adapters.RetrieveItemsProcessResult{
		Total:     total,
		ItemsDocs: &itemDocs}
}

type SearchItems struct{}

func (si *SearchItems) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if errs, valid := si.validate(r); !valid {
		resBody := adapters.ResponseBody{Errs: errs}
		adapters.NewResponse(http.StatusBadRequest, &resBody).WriteResponse(&w)
		return
	}
}

func (si *SearchItems) validate(r *http.Request) (errs *[]string, ok bool) {
	urlValidation := URLParamValidation{
		Rules: map[string]string{
			"keyword":   RegexNotEmpty,
			"quantity":  RegexNumber,
			"min_price": RegexNumber,
			"max_price": RegexNumber,
			"rating":    RegexNumber},
		Values: r.URL.Query()}

	return urlValidation.ValidateIfExists()
}

func (si *SearchItems) Process(r *http.Request, out proto.Message) (result interface{}) {

	return nil
}
