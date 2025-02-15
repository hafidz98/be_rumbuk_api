package exception

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/hafidz98/be_rumbuk_api/helper"
	"github.com/hafidz98/be_rumbuk_api/models/rest"
)

const (
	AccessUnauthorized    = "Access unauthorized"
	AccessForbidden       = "Access forbidden"
	InvalidOrMissingToken = "Invalid or missing access token"
	InvalidCredentials    = "Invalid credentials"
	DuplicateEmail        = "An account with that email already exists"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {

	if authError(writer, request, err) {
		return
	}

	if accessForbiddenError(writer, request, err) {
		return
	}

	if notFoundError(writer, request, err) {
		return
	}

	if validationError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func authError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(AuthError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")

		webResponse := rest.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: http.StatusText(http.StatusUnauthorized),
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func accessForbiddenError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(AccessForbiddenError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")

		webResponse := rest.WebResponse{
			Code:   http.StatusForbidden,
			Status: http.StatusText(http.StatusForbidden),
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := rest.WebResponse{
			Code:   http.StatusNotFound,
			Status: http.StatusText(http.StatusNotFound),
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func validationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := rest.WebResponse{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := rest.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: http.StatusText(http.StatusInternalServerError),
		Data:   err,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
