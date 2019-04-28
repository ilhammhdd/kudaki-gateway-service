package adapters

import (
	"log"
	"net/http"
	"strconv"

	"github.com/golang/protobuf/proto"

	"github.com/google/uuid"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-entities/events"
	"github.com/ilhammhdd/kudaki-entities/store"
	"github.com/ilhammhdd/kudaki-entities/user"

	"github.com/ilhammhdd/go-toolkit/jwtkit"

	"github.com/ilhammhdd/kudaki-gateway-service/usecases"
)

func AddStorefrontItem(r *http.Request, esp usecases.EventSourceProducer, esc usecases.EventSourceConsumer) *Response {

	jwt, err := jwtkit.GetJWT(jwtkit.JWTString(r.Header.Get("Kudaki-Token")))
	errorkit.ErrorHandled(err)

	amount, _ := strconv.ParseInt(r.MultipartForm.Value["amount"][0], 10, 32)
	price, _ := strconv.ParseInt(r.MultipartForm.Value["price"][0], 10, 32)
	usr := jwt.Payload.Claims["user"].(map[string]interface{})
	parsedUser := &user.User{
		AccountType: user.AccountType(user.AccountType_value[usr["account_type"].(string)]),
		Email:       usr["email"].(string),
		PhoneNumber: usr["phone_number"].(string),
		Role:        user.Role(user.Role_value[usr["role"].(string)]),
		Uuid:        usr["uuid"].(string),
	}

	asir := events.AddStorefrontItemRequested{
		Item: &store.Item{
			Amount:      int32(amount),
			Description: r.MultipartForm.Value["description"][0],
			Name:        r.MultipartForm.Value["name"][0],
			Photo:       r.MultipartForm.Value["photo"][0],
			Price:       int32(price),
			Storefront: &store.Storefront{
				UserUuid: parsedUser.Uuid,
			},
			Unit: r.MultipartForm.Value["unit"][0],
			Uuid: uuid.New().String(),
		},
		Uuid: uuid.New().String(),
	}

	asirBytes, err := proto.Marshal(&asir)
	errorkit.ErrorHandled(err)

	store := usecases.Store{
		Esc: esc,
		Esp: esp,
	}

	sia := store.AddStorefrontItem(asir.Uuid, asirBytes)
	log.Printf("StorefrontItemAdded : type = %T, value = %v", sia, sia)

	var resBody ResponseBody
	if sia.EventStatus.HttpCode != int32(http.StatusOK) {
		resBody.Errs = &sia.EventStatus.Errors
		resBody.Success = false
	} else {
		resBody.Success = true
	}

	return NewResponse(int(sia.EventStatus.HttpCode), &resBody)
}
