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

	apiStudentRoute := group.New("/students").GET(studentController.FindAll).POST(studentController.Create).Middleware(middlewares.AuthMiddleware).Children(
		group.New("/:studentId").GET(studentController.FetchById).PATCH(studentController.Update).DELETE(studentController.Delete),
	)

	return apiStudentRoute
}
