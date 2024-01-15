package controllers

import (
	"net/http"
	"strconv"

	"github.com/hafidz98/be_rumbuk_api/helper"
	"github.com/hafidz98/be_rumbuk_api/models/rest"
	"github.com/hafidz98/be_rumbuk_api/services"
	"github.com/julienschmidt/httprouter"
)

type TimeslotController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	//Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	//Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type TimeslotControllerImpl struct {
	TimeslotService services.TimeslotService
}

func NewTimeslotController(timeslotService services.TimeslotService) TimeslotController {
	return &TimeslotControllerImpl{
		TimeslotService: timeslotService,
	}
}

func (controller *TimeslotControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	timeslotCreateRequest := rest.TimeSlotCreateRequest{}
	helper.ReadFromRequestBody(request, &timeslotCreateRequest)

	timeslot := controller.TimeslotService.Create(request.Context(), timeslotCreateRequest)
	timeslotResponse := controller.TimeslotService.GetById(request.Context(), timeslot.ID)

	webResponse := rest.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   timeslotResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

//func (controller *TimeslotControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
//func (controller *TimeslotControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)

// Mendapatkan Timeslot berdasarkan referensi id timeslot
func (controller *TimeslotControllerImpl) GetById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	timeslotId := params.ByName("timeslotId")
	id, err := strconv.Atoi(timeslotId)
	helper.PanicIfError(err)

	timeslotResponse := controller.TimeslotService.GetById(request.Context(), id)

	webResponse := rest.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   timeslotResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TimeslotControllerImpl) GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	timeslotResponse := controller.TimeslotService.GetAll(request.Context())

	webResponse := rest.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   timeslotResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
