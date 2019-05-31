package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/RediSearch/redisearch-go/redisearch"
	kudaki_entities "github.com/ilhammhdd/kudaki-entities"

	"github.com/ilhammhdd/kudaki-entities/user"

	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/go-toolkit/safekit"
	_ "github.com/ilhammhdd/kudaki-entities/rental"
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
	http.Handle("/rental/mock-index-cart", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rsClient := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudaki_entities.ClientName_CARTS.String())
		rsClient.Drop()
	}))

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", os.Getenv("REST_PORT")),
		/* ReadTimeout:  time.Second * 3,
		WriteTimeout: time.Second * 7 */}

	defer server.Close()

	errorkit.ErrorHandled(server.ListenAndServe())
}
