package controller

import (
	"net/http"

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

}

func (controller *StudentControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}

func (controller *StudentControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}

func (controller *StudentControllerImpl) FetchById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}

func (controller *StudentControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}
