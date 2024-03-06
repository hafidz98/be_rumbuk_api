package controllers

import (
	"net/http"
	//"strconv"

	"github.com/hafidz98/be_rumbuk_api/helper"
	"github.com/hafidz98/be_rumbuk_api/models/rest"
	"github.com/hafidz98/be_rumbuk_api/services"
	"github.com/julienschmidt/httprouter"
)

type ReservationController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetReservationByStudentID(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	//GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type ReservationControllerImpl struct {
	ReserveService services.ReservationService
}

func NewReservationController(reserveService services.ReservationService) ReservationController {
	return &ReservationControllerImpl{
		ReserveService: reserveService,
	}
}

func (controller *ReservationControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	reservationCreateRequest := rest.ReserveCreateRequest{}
	helper.ReadFromRequestBody(request, &reservationCreateRequest)

	reserve := controller.ReserveService.CreateReservation(request.Context(), reservationCreateRequest)

	webResponse := rest.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   reserve,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ReservationControllerImpl) GetReservationByStudentID(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	studentId := params.ByName("studentId")

	reservationResponse := controller.ReserveService.SelectReservationByStudentID(request.Context(), studentId)

	webResponse := rest.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   reservationResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
