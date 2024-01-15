package controllers

import (
	"net/http"

	//"github.com/hafidz98/be_rumbuk_api/exception"
	"github.com/hafidz98/be_rumbuk_api/helper"
	service_model "github.com/hafidz98/be_rumbuk_api/models/rest"
	service "github.com/hafidz98/be_rumbuk_api/services"
	"github.com/julienschmidt/httprouter"
)

type StudentController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FetchById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type StudentControllerImpl struct {
	StudentService service.StudentService
}

func NewStudentController(studentService service.StudentService) StudentController {
	return &StudentControllerImpl{
		StudentService: studentService,
	}
}

func (controller *StudentControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	studentCreateRequest := service_model.StudentCreateRequest{}
	helper.ReadFromRequestBody(request, &studentCreateRequest)

	studentResponse := controller.StudentService.Create(request.Context(), studentCreateRequest)
	webResponse := service_model.WebResponse{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data:   studentResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *StudentControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	studentUpdateRequest := service_model.StudentUpdateRequest{}
	helper.ReadFromRequestBody(request, &studentUpdateRequest)

	studentUpdateRequest.StudentID = params.ByName("studentId")

	studentResponse := controller.StudentService.Update(request.Context(), studentUpdateRequest)
	webResponse := service_model.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   studentResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *StudentControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	studentId := params.ByName("studentId")

	controller.StudentService.Delete(request.Context(), studentId)
	webResponse := service_model.WebResponse{
		Code:   http.StatusNoContent,
		Status: http.StatusText(http.StatusOK),
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *StudentControllerImpl) FetchById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	studentId := params.ByName("studentId")

	studentResponse := controller.StudentService.FetchById(request.Context(), studentId)
	webResponse := service_model.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   studentResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *StudentControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	studentResponses := controller.StudentService.FindAll(request.Context())

	webResponse := service_model.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   studentResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
