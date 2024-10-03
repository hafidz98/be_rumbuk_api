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
	CancelReservation(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
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

	reserve, msg := controller.ReserveService.CreateReservation(request.Context(), reservationCreateRequest)

	// type data struct {
	// 	Msg string `json:"message,omitempty"`
	// }

	//message := data{Msg: msg}

	webResponse := rest.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   map[string]interface{}{"reservation": reserve},
		Msg:    msg,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ReservationControllerImpl) GetReservationByStudentID(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	studentIdParam := params.ByName("studentId")
	studentIdBody := rest.StudentDataRequest{}
	helper.ReadFromRequestBody(request, &studentIdBody)

	var studentId string
	if studentIdParam == "" {
		studentId = studentIdBody.StudentID
	} else {
		studentId = studentIdParam
	}

	reservationResponse := controller.ReserveService.SelectReservationByStudentID(request.Context(), studentId)

	webResponse := rest.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   reservationResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ReservationControllerImpl) CancelReservation(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	reservationUpdateRequest := rest.ReserveUpdateRequest{}
	helper.ReadFromRequestBody(request, &reservationUpdateRequest)

	controller.ReserveService.CancelReservation(request.Context(), reservationUpdateRequest.ReservationID)

	webResponse := rest.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   map[string]interface{}{"message": "Ok"},
	}

	helper.WriteToResponseBody(writer, webResponse)
}
