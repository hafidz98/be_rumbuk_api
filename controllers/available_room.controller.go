package controllers

import (
	"net/http"
	"time"

	"github.com/hafidz98/be_rumbuk_api/helper"
	"github.com/hafidz98/be_rumbuk_api/models/rest"
	"github.com/hafidz98/be_rumbuk_api/services"
	"github.com/julienschmidt/httprouter"
)

type AvailableRoomController interface {
	FetchAllAvailableRoom(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type AvailableRoomControllerImpl struct {
	AvailableRoomService services.AvailableRoomService
}

func NewAvailableRoomController(roomService services.AvailableRoomService) AvailableRoomController {
	return &AvailableRoomControllerImpl{
		AvailableRoomService: roomService,
	}
}

func (controller *AvailableRoomControllerImpl) FetchAllAvailableRoom(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	rQ := request.URL.Query()
	dt := time.Now()

	dateRsv := rQ.Get("date")
	if dateRsv == "" {
		dateRsv = dt.AddDate(0, 0, 3).Format("2006-01-02")
	}

	helper.Info.Printf("date: %v", dateRsv)
	helper.Info.Printf("date: %v", dt)

	availableRoomResponses := controller.AvailableRoomService.GetAllAvailableRoom(request.Context(), dateRsv)

	helper.Info.Printf("Room Data: %v", availableRoomResponses)
	
	webResponse := rest.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   map[string]interface{}{"date": dateRsv, "buildings": availableRoomResponses},
	}

	helper.WriteToResponseBody(writer, webResponse)
}
