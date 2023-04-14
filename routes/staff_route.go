package routes

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/hafidz98/be_rumbuk_api/controllers"

	"github.com/hafidz98/be_rumbuk_api/middlewares"
	"github.com/hafidz98/be_rumbuk_api/repositories"
	"github.com/hafidz98/be_rumbuk_api/services"
	group "github.com/mythrnr/httprouter-group"
)

func StaffRoute(db *sql.DB, validate *validator.Validate) *group.RouteGroup {
	staffRepository := repositories.NewStaffRepo()
	staffService := services.NewStaffService(staffRepository, db, validate)
	staffController := controllers.NewStaffController(staffService)

	apiStudentRoute := group.New("/staff").Middleware(middlewares.AuthMiddleware).Children(
		group.New("").GET(staffController.FetchAllFilter).POST(staffController.Create).Middleware(middlewares.RequiredAdmin),
		group.New("/:staffId").GET(staffController.FetchById).PATCH(staffController.Update).DELETE(staffController.Delete).Middleware(middlewares.RequiredAdmin),
	)

	return apiStudentRoute
}
