package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/hafidz98/be_rumbuk_api/helper"
	"github.com/hafidz98/be_rumbuk_api/models/rest"
	"github.com/hafidz98/be_rumbuk_api/services"
	"github.com/julienschmidt/httprouter"
)

type RoomController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FetchAllRooms(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FetchAllRoomsDetail(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type RoomControllerImpl struct {
	RoomService services.RoomService
}

func NewRoomController(roomService services.RoomService) RoomController {
	return &RoomControllerImpl{
		RoomService: roomService,
	}
}

func (ctrl *RoomControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roomCreateRequest := rest.RoomCreateRequest{}
	helper.ReadFromRequestBody(request, &roomCreateRequest)

	room := ctrl.RoomService.Create(request.Context(), roomCreateRequest)
	roomResponse := ctrl.RoomService.FetchByID(request.Context(), room.ID)

	webResponse := rest.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   roomResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (ctrl *RoomControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roomUpdateRequest := rest.RoomUpdateRequest{}
	helper.ReadFromRequestBody(request, &roomUpdateRequest)

	id, err := strconv.Atoi(params.ByName("roomId"))
	helper.PanicIfError(err)

	roomUpdateRequest.ID = id
	ctrl.RoomService.Update(request.Context(), roomUpdateRequest)

	roomResponse := ctrl.RoomService.FetchByID(request.Context(), id)
	webResponse := rest.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   roomResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (ctrl *RoomControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roomId := params.ByName("roomId")
	res, err := strconv.Atoi(roomId)
	helper.PanicIfError(err)

	ctrl.RoomService.Delete(request.Context(), res)
	webResponse := rest.WebResponse{
		Code:   http.StatusNoContent,
		Status: http.StatusText(http.StatusOK),
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (ctrl *RoomControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roomId := params.ByName("roomId")
	res, err := strconv.Atoi(roomId)
	helper.PanicIfError(err)

	roomResponse := ctrl.RoomService.FetchByID(request.Context(), res)
	webResponse := rest.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   roomResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (ctrl *RoomControllerImpl) FetchAllRooms(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	roomResponses := ctrl.RoomService.FetchAll(request.Context())

	webResponse := rest.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   roomResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (ctrl *RoomControllerImpl) FetchAllRoomsDetail(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	rQ := request.URL.Query()
	dt := time.Now()

	dateRsv := rQ.Get("date")
	if dateRsv == "" {
		dateRsv = dt.AddDate(0, 0, 3).Format("2006-01-02")
	}

	helper.Info.Printf("date: %v", dateRsv)
	helper.Info.Printf("date: %v", dt)

	roomResponses := ctrl.RoomService.FetchAllRooms(request.Context(), dateRsv)

	webResponse := rest.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   map[string]interface{}{"date": dateRsv, "buildings": roomResponses},
	}

	helper.WriteToResponseBody(writer, webResponse)
}
