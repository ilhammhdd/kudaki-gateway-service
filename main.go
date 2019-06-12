package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/RediSearch/redisearch-go/redisearch"

	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/go-toolkit/safekit"
	kudakiredisearch "github.com/ilhammhdd/kudaki-externals/redisearch"
	"github.com/ilhammhdd/kudaki-gateway-service/externals/rest"
)

func init() {
	if len(os.Args) > 1 {
		for _, val := range os.Args[1:] {
			flag := strings.Split(val, " ")
			os.Setenv(flag[1], flag[2])
		}
	}

	log.Println(os.Getenv("KAFKA_VERSION"))
}

func main() {
	wp := safekit.NewWorkerPool()

	wp.Work <- restListener

	wp.PoolWG.Wait()
}

func restListener() {
	http.Handle("/signup", rest.MethodValidator(http.MethodPost, new(rest.Signup)))
	http.Handle("/user/verify", rest.MethodValidator(http.MethodGet, new(rest.VerifyUser)))
	http.Handle("/login", rest.MethodValidator(http.MethodPost, new(rest.Login)))
	http.Handle("/user/password/change", rest.MethodValidator(http.MethodPut, rest.Authenticate(new(rest.ChangePassword))))
	http.Handle("/user/password/reset", rest.MethodRouting{
		PostHandler: new(rest.ResetPasswordSendEmail),
		PutHandler:  new(rest.ResetPassword),
	})
	http.Handle("/store/storefront/item", rest.MethodRouting{
		PostHandler:   rest.Authenticate(new(rest.AddStorefrontItem)),
		PutHandler:    rest.Authenticate(new(rest.UpdateStorefrontItem)),
		DeleteHandler: rest.Authenticate(new(rest.DeleteStorefrontItem)),
	})
	http.Handle("/store/storefront/items", rest.Authenticate(new(rest.GetAllStorefrontItems)))

	http.Handle("/mock/index/drop/all", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		client := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.Item.Name())
		client.Drop()
		client = redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.Storefront.Name())
		client.Drop()
		client = redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.Profile.Name())
		client.Drop()
		client = redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.User.Name())
		client.Drop()

		w.WriteHeader(http.StatusOK)
	}))
	http.Handle("/mock/storefront/items", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		client := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.Item.Name())
		rawQuery := fmt.Sprintf(`@storefront_uuid:"%s"`, kudakiredisearch.RedisearchText("6647cd7a-25b2-41ee-9ac7-f15959139274").Sanitize())
		docs, total, err := client.Search(redisearch.NewQuery(rawQuery))
		errorkit.ErrorHandled(err)
		log.Printf("searched %d storefront items", total)
		log.Printf("item docs : %v", docs)

		type responseData struct {
			ItemDocs []redisearch.Document `json:"item_docs"`
		}

		var resData responseData
		resData.ItemDocs = docs
		resDataMarshalled, err := json.Marshal(resData)
		errorkit.ErrorHandled(err)

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resDataMarshalled)
	}))

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", os.Getenv("REST_PORT"))}

	defer server.Close()

	errorkit.ErrorHandled(server.ListenAndServe())
}
