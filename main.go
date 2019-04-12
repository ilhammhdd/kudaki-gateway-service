package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/go-toolkit/safekit"
	"github.com/ilhammhdd/kudaki-gateway-service/externals/rest"
)

func init() {
	if len(os.Args) == 6 {
		os.Setenv("ADDRESS", os.Args[1])
		os.Setenv("GRPC_PORT", os.Args[2])
		os.Setenv("REST_PORT", os.Args[3])
		os.Setenv("KAFKA_BROKERS", os.Args[4])
		os.Setenv("USER_SERVICE_GRPC_ADDRESS", os.Args[5])
	}
}

func main() {
	// testing
	wp := safekit.NewWorkerPool()

	wp.Work <- restListener

	wp.PoolWG.Wait()
}

func restListener() {
	http.HandleFunc("/testing", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("tested"))
	}))
	http.Handle("/signup", rest.MethodValidator(http.MethodPost, http.HandlerFunc(rest.Signup)))
	http.Handle("/user/verify", rest.MethodValidator(http.MethodGet, http.HandlerFunc(rest.VerifyUser)))
	http.Handle("/login", rest.MethodValidator(http.MethodPost, http.HandlerFunc(rest.Login)))
	http.Handle("/test/authenticate/jwt", rest.MethodValidator(http.MethodGet, rest.Authenticate(http.HandlerFunc(rest.TestAuthenticateJWT))))
	http.Handle("/user/reset/password", rest.MethodValidator(http.MethodPost, rest.Authenticate(http.HandlerFunc(rest.ResetPassword))))

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("REST_PORT")),
		ReadTimeout:  time.Second * 3,
		WriteTimeout: time.Second * 7}

	defer server.Close()

	errorkit.ErrorHandled(server.ListenAndServe())
}
