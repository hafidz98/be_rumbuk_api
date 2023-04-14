package controllers

import (
	"net/http"
	"strconv"

	"github.com/hafidz98/be_rumbuk_api/helper"
	"github.com/hafidz98/be_rumbuk_api/models/domain"
	"github.com/hafidz98/be_rumbuk_api/models/service"
	"github.com/hafidz98/be_rumbuk_api/services"
	"github.com/julienschmidt/httprouter"
)

type StaffController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FetchById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FetchAllFilter(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type StaffControllerImpl struct {
	StaffService services.StaffService
}

func NewStaffController(staffService services.StaffService) StaffController {
	return &StaffControllerImpl{
		StaffService: staffService,
	}
}

func (controller *StaffControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	staffCreateRequest := service.StaffCreateRequest{}
	helper.ReadFromRequestBody(request, &staffCreateRequest)

	staff := controller.StaffService.Create(request.Context(), staffCreateRequest)
	staffResponse := controller.StaffService.FetchById(request.Context(), staff.StaffID)
	webResponse := service.WebResponse{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data:   []interface{}{staffResponse},
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *StaffControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	staffUpdateRequest := service.StaffUpdateRequest{}
	helper.ReadFromRequestBody(request, &staffUpdateRequest)

	staffUpdateRequest.StaffID = params.ByName("staffId")
	controller.StaffService.Update(request.Context(), staffUpdateRequest)

	staffResponse := controller.StaffService.FetchById(request.Context(), staffUpdateRequest.StaffID)
	webResponse := service.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   []interface{}{staffResponse},
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *StaffControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	controller.StaffService.Delete(request.Context(), params.ByName("staffId"))
	webResponse := service.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   []interface{}{},
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *StaffControllerImpl) FetchById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	staffResponse := controller.StaffService.FetchById(request.Context(), params.ByName("staffId"))
	webResponse := service.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   []interface{}{staffResponse},
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *StaffControllerImpl) FetchAllFilter(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	r := request.URL.Query()
	limit, _ := strconv.Atoi(r.Get("limit"))
	helper.Info.Printf("ctrl limit: %v", limit)

	offset, _ := strconv.Atoi(r.Get("offset"))
	helper.Info.Printf("ctrl offset: %v", offset)

	staffFilter := domain.FilterParams{
		Limit:  uint64(limit),
		Offset: uint64(offset),
	}

	staffResponse := controller.StaffService.FetchAllFilter(request.Context(), &staffFilter)
	res := service.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   staffResponse,
	}
	helper.WriteToResponseBody(writer, res)

}
