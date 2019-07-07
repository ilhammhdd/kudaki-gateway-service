package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"

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
	/*
		mountain aggregate
	*/
	http.Handle("/recommendation/item", rest.MethodRouting{
		DeleteHandler: rest.Authenticate(new(rest.DeleteRecommendedGearItem)),
		PostHandler:   rest.Authenticate(new(rest.AddRecommendedGearItem)),
	})
	http.Handle("/recommendation", rest.MethodRouting{
		PostHandler:   rest.Authenticate(new(rest.AddRecommendedGear)),
		DeleteHandler: rest.Authenticate(new(rest.DeleteRecommendedGear)),
	})
	http.Handle("/recommendations", rest.MethodValidator(http.MethodGet, rest.Authenticate(new(rest.RetrieveRecommendedGears))))
	http.Handle("/recommendation/items", rest.MethodValidator(http.MethodGet, rest.Authenticate(new(rest.RetrieveRecommendedGearItems))))
	http.Handle("/recommendation/upvote", rest.MethodValidator(http.MethodPost, rest.Authenticate(new(rest.UpVoteRecommendedGear))))
	http.Handle("/recommendation/downvote", rest.MethodValidator(http.MethodPost, rest.Authenticate(new(rest.DownVoteRecommendedGear))))
	/*
		order aggregate
	*/
	http.Handle("/order/owner", rest.MethodValidator(http.MethodGet, rest.Authenticate(new(rest.RetrieveOwnerOrderHistories))))
	http.Handle("/order/tenant", rest.MethodValidator(http.MethodGet, rest.Authenticate(new(rest.RetrieveTenantOrderHistories))))
	http.Handle("/order/tenant/review/owner", rest.MethodValidator(http.MethodPost, rest.Authenticate(new(rest.TenantReviewOwner))))
	http.Handle("/order/tenant/review/items", rest.MethodValidator(http.MethodPost, rest.Authenticate(new(rest.TenantReviewItems))))
	http.Handle("/order/approve", rest.MethodValidator(http.MethodPost, rest.Authenticate(new(rest.ApproveOrder))))
	http.Handle("/order/disapprove", rest.MethodValidator(http.MethodPost, rest.Authenticate(new(rest.DisapproveOrder))))
	/*
		rental aggregate
	*/
	http.Handle("/rental/cart/item", rest.MethodRouting{
		PostHandler:   rest.Authenticate(new(rest.AddCartItem)),
		DeleteHandler: rest.Authenticate(new(rest.DeleteCartItem)),
		PatchHandler:  rest.Authenticate(new(rest.UpdateCartItem)),
	})
	http.Handle("/rental/cart/items", rest.MethodValidator(http.MethodGet, rest.Authenticate(new(rest.RetrieveCartItems))))
	http.Handle("/rental/confirm-returnment/tenant", rest.MethodValidator(http.MethodPost, rest.Authenticate(new(rest.TenantConfirmReturnment))))
	http.Handle("/rental/confirm-returnment/owner", rest.MethodValidator(http.MethodPost, rest.Authenticate(new(rest.OwnerConfirmReturnment))))
	/*
		store aggregate
	*/
	http.Handle("/storefront/item", rest.MethodRouting{
		PostHandler:   rest.Authenticate(new(rest.AddStorefrontItem)),
		DeleteHandler: rest.Authenticate(new(rest.DeleteStorefrontItem)),
		PutHandler:    rest.Authenticate(new(rest.UpdateStorefrontItem)),
	})
	http.Handle("/storefront/items", rest.MethodValidator(http.MethodGet, rest.Authenticate(new(rest.RetrieveStorefrontItems))))
	http.Handle("/items", rest.MethodValidator(http.MethodGet, rest.Authenticate(new(rest.RetrieveItems))))
	http.Handle("/item/search", rest.MethodValidator(http.MethodGet, rest.Authenticate(new(rest.SearchItems))))
	http.Handle("/item-review/review", rest.MethodValidator(http.MethodPost, rest.Authenticate(new(rest.ReviewItem))))
	http.Handle("/item-review/reviews", rest.MethodValidator(http.MethodGet, rest.Authenticate(new(rest.RetrieveItemReviews))))
	http.Handle("/item-review/review/comment", rest.MethodValidator(http.MethodPost, rest.Authenticate(new(rest.CommentItemReview))))
	http.Handle("/item-review/review/comments", rest.MethodValidator(http.MethodGet, rest.Authenticate(new(rest.RetrieveItemReviewComments))))
	/*
		user aggregate
	*/
	http.Handle("/login", rest.MethodValidator(http.MethodPost, new(rest.Login)))
	http.Handle("/user/password/reset", rest.MethodRouting{
		GetHandler:   nil,
		PatchHandler: new(rest.ResetPassword),
		PostHandler:  new(rest.ResetPasswordSendEmail),
	})
	http.Handle("/signup", rest.MethodValidator(http.MethodPost, new(rest.Signup)))
	http.Handle("/user-info/address", rest.MethodRouting{
		PostHandler: rest.Authenticate(new(rest.AddAddress)),
		PutHandler:  rest.Authenticate(new(rest.UpdateAddress)),
	})
	http.Handle("/user-info/profile", rest.MethodValidator(http.MethodPatch, rest.Authenticate(new(rest.UpdateProfile))))
	http.Handle("/user-info/addresses", rest.MethodValidator(http.MethodGet, rest.Authenticate(new(rest.RetrieveAddresses))))
	http.Handle("/user/password/change", rest.MethodValidator(http.MethodPatch, rest.Authenticate(new(rest.ChangePassword))))
	http.Handle("/user/verify", rest.MethodValidator(http.MethodGet, new(rest.VerifyUser)))

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", os.Getenv("REST_PORT"))}

	defer server.Close()

	errorkit.ErrorHandled(server.ListenAndServe())
}
