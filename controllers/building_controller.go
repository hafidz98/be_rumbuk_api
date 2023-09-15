package controllers

import (
	"net/http"
	"strconv"

	"github.com/hafidz98/be_rumbuk_api/helper"
	"github.com/hafidz98/be_rumbuk_api/models/rest"
	"github.com/hafidz98/be_rumbuk_api/services"
	"github.com/julienschmidt/httprouter"
)

type BuildingController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FetchById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type BuildingControllerImpl struct {
	BuildingService services.BuildingService
}

func NewBuildingController(buildingService services.BuildingService) BuildingController {
	return &BuildingControllerImpl{
		BuildingService: buildingService,
	}
}

func (controller *BuildingControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	buildingCreateRequest := rest.BuildingCreateRequest{}
	helper.ReadFromRequestBody(request, &buildingCreateRequest)

	buildingResponse := controller.BuildingService.Create(request.Context(), buildingCreateRequest)
	webResponse := rest.WebResponse{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data:   buildingResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BuildingControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	buildingUpdateRequest := rest.BuildingUpdateRequest{}
	helper.ReadFromRequestBody(request, &buildingUpdateRequest)

	b, err := strconv.Atoi(params.ByName("buildingId"))
	helper.PanicIfError(err)

	buildingUpdateRequest.ID = b

	buildingResponse := controller.BuildingService.Update(request.Context(), buildingUpdateRequest)
	webResponse := rest.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   buildingResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BuildingControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	buildingId, err := strconv.Atoi(params.ByName("buildingId"))
	helper.PanicIfError(err)

	controller.BuildingService.Delete(request.Context(), buildingId)
	webResponse := rest.WebResponse{
		Code:   http.StatusNoContent,
		Status: http.StatusText(http.StatusOK),
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BuildingControllerImpl) FetchById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	buildingId, err := strconv.Atoi(params.ByName("buildingId"))
	helper.PanicIfError(err)

	buildingResponse := controller.BuildingService.FetchById(request.Context(), buildingId)
	webResponse := rest.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   buildingResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BuildingControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	buildingResponses := controller.BuildingService.FetchAll(request.Context())

	webResponse := rest.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   buildingResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
