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

func StudentRoute(db *sql.DB, validate *validator.Validate) *group.RouteGroup {
	studentRepository := repositories.NewStudentRepo()
	studentService := services.NewStudentService(studentRepository, db, validate)
	studentController := controllers.NewStudentController(studentService)

	apiStudentRoute := group.New("/students").Middleware(middlewares.AuthMiddleware).Children(
		group.New("").GET(studentController.FindAll).POST(studentController.Create).Middleware(middlewares.RequiredAdmin),
		group.New("/:studnetId").DELETE(studentController.Delete).Middleware(middlewares.RequiredAdmin),
		group.New("/:studentId").GET(studentController.FetchById).PATCH(studentController.Update).Middleware(middlewares.RequiredStudentOrAdmin).Children(
			ReservationRoute(db,validate),
		),
	)

	return apiStudentRoute
}
