package adapters

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/golang/protobuf/proto"

	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/google/uuid"
	"github.com/ilhammhdd/kudaki-entities/events"
	"github.com/ilhammhdd/kudaki-entities/mountain"

	"github.com/ilhammhdd/kudaki-gateway-service/usecases"
)

func CreateMountain(r *http.Request, esp usecases.EventSourceProducer, esc usecases.EventSourceConsumer) *Response {

	height, err := strconv.ParseInt(r.MultipartForm.Value["height"][0], 10, 32)
	errorkit.ErrorHandled(err)
	latitude, err := strconv.ParseFloat(r.MultipartForm.Value["latitude"][0], 64)
	errorkit.ErrorHandled(err)
	longitude, err := strconv.ParseFloat(r.MultipartForm.Value["longitude"][0], 64)
	errorkit.ErrorHandled(err)

	cmr := &events.CreateMountainRequested{
		Mountain: &mountain.Mountain{
			Description: r.MultipartForm.Value["description"][0],
			Difficulty:  0,
			Height:      int32(height),
			Latitude:    latitude,
			Longitude:   longitude,
			Name:        r.MultipartForm.Value["name"][0],
			Uuid:        uuid.New().String(),
		},
		Uid: uuid.New().String()}
	cmrBytes, err := proto.Marshal(cmr)
	errorkit.ErrorHandled(err)

	mountain := usecases.Mountain{Esp: esp, Esc: esc}
	mc := mountain.CreateMountain(cmr.Uid, cmrBytes)

	resBody := ResponseBody{}

	if mc.EventStatus.HttpCode != int32(http.StatusOK) {
		resBody.Errs = &mc.EventStatus.Errors
	}

	return NewResponse(int(mc.EventStatus.HttpCode), &resBody)
}

func RetrieveMountains(esp usecases.EventSourceProducer, esc usecases.EventSourceConsumer) *Response {

	rmr := events.RetrieveMountainsRequested{Uid: uuid.New().String()}
	rmrBytes, err := proto.Marshal(&rmr)
	errorkit.ErrorHandled(err)

	mountain := usecases.Mountain{
		Esc: esc,
		Esp: esp,
	}

	mr := mountain.RetrieveMountains(rmr.Uid, rmrBytes)
	errorkit.ErrorHandled(err)

	// forgot the event status
	mrJSON, err := json.Marshal(mr)
	errorkit.ErrorHandled(err)

	resBody := ResponseBody{
		Data: mrJSON,
	}
	return NewResponse(http.StatusOK, &resBody)
}
