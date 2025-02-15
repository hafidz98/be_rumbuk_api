package middlewares

import (
	"net/http"

	"github.com/hafidz98/be_rumbuk_api/exception"
	"github.com/hafidz98/be_rumbuk_api/helper"
	"github.com/julienschmidt/httprouter"
)

func AuthMiddleware(handler httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		tokenString := request.Header.Get("X-JWT-Token-Key")
		if tokenString == "" {
			panic(exception.NewAuthorization(exception.InvalidOrMissingToken))
		}

		err := helper.ValidateToken(tokenString)
		if err != nil {
			panic(exception.NewAuthorization(exception.InvalidOrMissingToken))
		}

		userData, err := helper.ExtractUserDataFromToken(tokenString)

		helper.Info.Println(userData)
		helper.Info.Println(userData.Role)

		if err != nil {
			panic(exception.NewAuthorization(exception.InvalidCredentials))
		}

		if userData.Role == "Admin" || userData.Role == "Staff" {
			request.Header.Set("X-User-Role", userData.Role)
			handler(writer, request, params)
			return
		} else if userData.Role == "Student" {
			request.Header.Set("X-User-Role", userData.Role)
			request.Header.Set("X-User-Id", userData.UserID)
			handler(writer, request, params)
			return
		}

		panic(exception.NewAuthorization(exception.InvalidCredentials))
	}
}
