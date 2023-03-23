package controller

import (
	"encoding/json"
	"net/http"

	"github.com/hafidz98/be_rumbuk_api/helper"
	service_model "github.com/hafidz98/be_rumbuk_api/model/service"
	"github.com/hafidz98/be_rumbuk_api/service"
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

func (controller *StudentControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	studentCreateRequest := service_model.StudentCreateRequest{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&studentCreateRequest)
	helper.PanicIfError(err)

	studentResponse := controller.StudentService.Create(request.Context(), studentCreateRequest)
	webResponse := service_model.WebResponse{
		Code: http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data: studentResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *StudentControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	studentUpdateRequest := service_model.StudentUpdateRequest{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&studentUpdateRequest)
	helper.PanicIfError(err)

	studentUpdateRequest.StudentID = params.ByName("studentId")

	studentResponse := controller.StudentService.Update(request.Context(), studentUpdateRequest)
	webResponse := service_model.WebResponse{
		Code: http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data: studentResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *StudentControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	studentId := params.ByName("studentId")

	controller.StudentService.Delete(request.Context(), studentId)
	webResponse := service_model.WebResponse{
		Code: http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *StudentControllerImpl) FetchById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	studentId := params.ByName("studentId")

	studentResponse := controller.StudentService.FetchById(request.Context(), studentId)
	webResponse := service_model.WebResponse{
		Code: http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data: studentResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *StudentControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	studentResponses := controller.StudentService.FindAll(request.Context())
	webResponse := service_model.WebResponse{
		Code: http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data: studentResponses,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
