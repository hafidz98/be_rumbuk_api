package controllers

import (
	"net/http"
	"strconv"

	"github.com/hafidz98/be_rumbuk_api/helper"
	"github.com/hafidz98/be_rumbuk_api/models/rest"
	"github.com/hafidz98/be_rumbuk_api/services"
	"github.com/julienschmidt/httprouter"
)

type FloorController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetByID(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type FloorControllerImpl struct {
	FloorService services.FloorService
}

func NewFloorController(floorService services.FloorService) FloorController {
	return &FloorControllerImpl{
		FloorService: floorService,
	}
}

func (controller *FloorControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	floorCreateRequest := rest.FloorCreateRequest{}
	helper.ReadFromRequestBody(request, &floorCreateRequest)

	floor := controller.FloorService.Create(request.Context(), floorCreateRequest)
	floorResponse := controller.FloorService.GetById(request.Context(), floor.ID)

	webResponse := rest.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   floorResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *FloorControllerImpl) GetByID(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
	floorId:= params.ByName("floorId")
	id, err := strconv.Atoi(floorId)
	helper.PanicIfError(err)

	floorResponse := controller.FloorService.GetById(request.Context(), id)

	webResponse := rest.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   floorResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *FloorControllerImpl) GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
	floorResponse := controller.FloorService.GetAll(request.Context())

	webResponse := rest.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   floorResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
