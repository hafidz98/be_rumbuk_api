package middlewares

import (
	"net/http"

	"github.com/hafidz98/be_rumbuk_api/exception"
	"github.com/julienschmidt/httprouter"
)

func RequiredStudentOrAdmin(handler httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		role := request.Header.Get("X-User-Role")
		id := request.Header.Get("X-User-Id")
		studentId := params.ByName("studentId")

		if role == "Student" {
			if studentId != id {
				panic(exception.NewAuthorization(exception.AccessUnauthorized))
			}
			//handler(writer, request, params)
		} else if role != "Admin" && role != "Staff" {
			panic(exception.NewAuthorization(exception.AccessUnauthorized))
		}
		handler(writer, request, params)
	}
}
