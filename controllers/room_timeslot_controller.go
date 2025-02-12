package controllers

import (
	"net/http"

	"github.com/hafidz98/be_rumbuk_api/helper"
	"github.com/hafidz98/be_rumbuk_api/models/rest"
	"github.com/hafidz98/be_rumbuk_api/services"
	"github.com/julienschmidt/httprouter"
)

type RoomTimeslotController interface {
	AddRoomTimeslot(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type RoomTimeslotControllerImpl struct {
	RoomTimeslotService services.RoomTimeslotService
}

func NewRoomTimeslotController(roomTimeslotService services.RoomTimeslotService) RoomTimeslotController {
	return &RoomTimeslotControllerImpl{
		RoomTimeslotService: roomTimeslotService,
	}
}

func (controller *RoomTimeslotControllerImpl) AddRoomTimeslot(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roomTimeslotRequest := rest.RoomTimeslotRequest{}
	helper.ReadFromRequestBody(request, &roomTimeslotRequest)

	roomTimeslot, err := controller.RoomTimeslotService.AddRoomTimeslot(request.Context(), roomTimeslotRequest)
	webResponse := rest.WebResponse{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data:   []interface{}{roomTimeslot, err},
	}

	helper.WriteToResponseBody(writer, webResponse)
}