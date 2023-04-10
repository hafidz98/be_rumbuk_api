package middlewares

import (
	"net/http"

	"github.com/hafidz98/be_rumbuk_api/exception"
	"github.com/julienschmidt/httprouter"
)

func RequiredAdmin(handler httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		if role := request.Header.Get("X-User-Role"); role != "Admin" && role != "Staff" {
			panic(exception.NewAuthorization(exception.AccessUnauthorized))
		}
		handler(writer, request, params)
	}
}
