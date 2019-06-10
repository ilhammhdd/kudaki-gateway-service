package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

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

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", os.Getenv("REST_PORT"))}

	defer server.Close()

	errorkit.ErrorHandled(server.ListenAndServe())
}
