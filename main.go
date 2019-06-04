package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/RediSearch/redisearch-go/redisearch"

	"github.com/ilhammhdd/kudaki-entities/kudakiredisearch"
	"github.com/ilhammhdd/kudaki-entities/user"

	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/go-toolkit/safekit"
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
	// testing again, how to trigger it? publish or push with tag? COME ON!
	// still unsure, COME ON!!
	wp := safekit.NewWorkerPool()

	wp.Work <- restListener

	wp.PoolWG.Wait()
}

func restListener() {
	http.HandleFunc("/testing", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("tested"))
	}))

	// user
	http.Handle("/signup", rest.MethodValidator(http.MethodPost, http.HandlerFunc(rest.Signup)))
	http.Handle("/user/verify", rest.MethodValidator(http.MethodGet, http.HandlerFunc(rest.VerifyUser)))
	http.Handle("/login", rest.MethodValidator(http.MethodPost, http.HandlerFunc(rest.Login)))
	http.Handle("/test/authenticate/jwt", rest.MethodValidator(http.MethodGet, rest.Authenticate(http.HandlerFunc(rest.TestAuthenticateJWT))))
	http.Handle("/user/change-password", rest.MethodValidator(http.MethodPut, rest.Authenticate(http.HandlerFunc(rest.ChangePassword))))
	http.Handle("/team/add", rest.MethodValidator(http.MethodPost, rest.Authorize(user.Role_ADMIN, http.HandlerFunc(rest.AddTeam))))
	http.Handle("/test/authorize/user", rest.Authorize(user.Role_USER, http.HandlerFunc(rest.TestAuthorizeUser)))
	http.Handle("/user/reset-password/send-email", rest.MethodValidator(http.MethodPut, http.HandlerFunc(rest.SendResetPasswordEmail)))
	http.Handle("/user/reset-password", rest.MethodRouting{
		GetHandler: http.HandlerFunc(rest.ResetPasswordPage),
		PutHandler: http.HandlerFunc(rest.ResetPassword),
	})

	// mountain
	http.Handle("/mountain/create", rest.MethodValidator(http.MethodPost, rest.Authorize(user.Role_KUDAKI_TEAM, http.HandlerFunc(rest.CreateMountain))))
	http.Handle("/mountain/retrieve", rest.MethodValidator(http.MethodGet, rest.Authenticate(http.HandlerFunc(rest.RetrieveMountains))))

	// store
	http.Handle("/storefront/items", rest.MethodValidator(http.MethodGet, rest.Authenticate(http.HandlerFunc(rest.RetrieveStorefrontItems))))
	http.Handle("/storefront/item", rest.MethodRouting{
		PostHandler:   rest.Authenticate(http.HandlerFunc(rest.AddFrontstoreItem)),
		DeleteHandler: rest.Authenticate(http.HandlerFunc(rest.DeleteStorefrontItem)),
		PutHandler:    rest.Authenticate(http.HandlerFunc(rest.UpdateStorefrontItem)),
	})
	http.Handle("/store/items", rest.MethodValidator(http.MethodGet, rest.Authenticate(http.HandlerFunc(rest.RetrieveItems))))
	http.Handle("/store/item", rest.MethodValidator(http.MethodGet, rest.Authenticate(http.HandlerFunc(rest.RetrieveItem))))
	http.Handle("/store/search-items", rest.MethodValidator(http.MethodGet, rest.Authenticate(http.HandlerFunc(rest.SearchItems))))

	// rental
	http.Handle("/rental/checkout", rest.MethodValidator(http.MethodPost, rest.Authenticate(new(rest.Checkout))))
	http.Handle("/rental/cart/item", rest.MethodRouting{
		PostHandler: rest.Authenticate(new(rest.AddCartItem)),
	})

	// mock
	http.Handle("/redisearch/index/reset-all", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rsClient := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.Cart.Name())
		rsClient.Drop()

		rsClient = redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.Item.Name())
		rsClient.Drop()

		rsClient = redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.CartItem.Name())
		rsClient.Drop()

		w.WriteHeader(http.StatusOK)
	}))
	http.Handle("/redisearch/search", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(100000)

		type responseData struct {
			Total int                   `json:"total"`
			Docs  []redisearch.Document `json:"docs"`
		}

		rsClient := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.CartItem.Name())
		rawQuery := fmt.Sprintf(`@cart_item_uuid:"%s"`, kudakiredisearch.RedisearchText(r.MultipartForm.Value["cart_item_uuid"][0]).Sanitize())
		docs, total, err := rsClient.Search(redisearch.NewQuery(rawQuery))
		errorkit.ErrorHandled(err)

		resCartItem := responseData{
			Docs:  docs,
			Total: total,
		}
		resCartItemByte, err := json.Marshal(resCartItem)
		errorkit.ErrorHandled(err)
		log.Printf("resCartItemByte : %v", string(resCartItemByte))

		rsClient = redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.Cart.Name())
		rawQuery = fmt.Sprintf(`@cart_uuid:"%s"`, kudakiredisearch.RedisearchText(r.MultipartForm.Value["cart_uuid"][0]).Sanitize())
		cartDocs, cartTotal, err := rsClient.Search(redisearch.NewQuery(rawQuery))
		errorkit.ErrorHandled(err)

		resCart := responseData{
			Docs:  cartDocs,
			Total: cartTotal,
		}
		resCartByte, err := json.Marshal(resCart)
		errorkit.ErrorHandled(err)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(resCartByte)
	}))

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", os.Getenv("REST_PORT")),
		/* ReadTimeout:  time.Second * 3,
		WriteTimeout: time.Second * 7 */}

	defer server.Close()

	errorkit.ErrorHandled(server.ListenAndServe())
}
