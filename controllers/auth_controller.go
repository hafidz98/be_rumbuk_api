package controllers

import (
	"net/http"

	"github.com/hafidz98/be_rumbuk_api/helper"
	service "github.com/hafidz98/be_rumbuk_api/models/rest"
	"github.com/hafidz98/be_rumbuk_api/services"
	"github.com/julienschmidt/httprouter"
)

type AuthController interface {
	Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type AuthControllerImpl struct {
	AuthService services.AuthService
}

func NewAuthController(authService services.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

func (controller *AuthControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	authLoginRequest := service.AuthLoginRequest{}
	helper.ReadFromRequestBody(request, &authLoginRequest)

	tokenString := controller.AuthService.Login(request.Context(), authLoginRequest)
	webResponse := service.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   "Authorized",
	}

	writer.Header().Add("X-JWT-Token-Key", tokenString)
	helper.WriteToResponseBody(writer, webResponse)
}
