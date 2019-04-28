package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ilhammhdd/kudaki-entities/user"

	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/go-toolkit/safekit"
	"github.com/ilhammhdd/kudaki-gateway-service/externals/rest"
)

func init() {
	if len(os.Args) == 7 {
		os.Setenv("ADDRESS", os.Args[1])
		os.Setenv("GRPC_PORT", os.Args[2])
		os.Setenv("REST_PORT", os.Args[3])
		os.Setenv("KAFKA_BROKERS", os.Args[4])
		os.Setenv("USER_SERVICE_GRPC_ADDRESS", os.Args[5])
		os.Setenv("KAFKA_VERSION", os.Args[6])
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
	http.Handle("/user/reset/password", rest.MethodValidator(http.MethodPut, rest.Authenticate(http.HandlerFunc(rest.ResetPassword))))
	http.Handle("/team/add", rest.MethodValidator(http.MethodPost, rest.Authorize(user.Role_ADMIN, http.HandlerFunc(rest.AddTeam))))
	http.Handle("/test/authorize/user", rest.Authorize(user.Role_USER, http.HandlerFunc(rest.TestAuthorizeUser)))

	// mountain
	http.Handle("/mountain/create", rest.MethodValidator(http.MethodPost, rest.Authorize(user.Role_KUDAKI_TEAM, http.HandlerFunc(rest.CreateMountain))))
	http.Handle("/mountain/retrieve", rest.MethodValidator(http.MethodGet, rest.Authenticate(http.HandlerFunc(rest.RetrieveMountains))))

	// rental
	http.Handle("/storefront/item/add", rest.MethodValidator(http.MethodPost, rest.Authenticate(http.HandlerFunc(rest.AddFrontstoreItem))))

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", os.Getenv("REST_PORT")),
		/* ReadTimeout:  time.Second * 3,
		WriteTimeout: time.Second * 7 */}

	defer server.Close()

	errorkit.ErrorHandled(server.ListenAndServe())
}
