package routes

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/hafidz98/be_rumbuk_api/controllers"
	"github.com/hafidz98/be_rumbuk_api/services"
	group "github.com/mythrnr/httprouter-group"
)

func AuthRoute(db *sql.DB, validate *validator.Validate) *group.RouteGroup {
	authService :=
		services.NewAuthService(db, validate)
	authController := controllers.NewAuthController(authService)

	apiAuthRoute := group.New("/auth").POST(authController.Login)
	return apiAuthRoute
}
