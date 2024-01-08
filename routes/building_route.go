package routes

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/hafidz98/be_rumbuk_api/controllers"

	// "github.com/hafidz98/be_rumbuk_api/middlewares"
	"github.com/hafidz98/be_rumbuk_api/repositories"
	"github.com/hafidz98/be_rumbuk_api/services"
	group "github.com/mythrnr/httprouter-group"
)

func BuildingRoute(db *sql.DB, validate *validator.Validate) *group.RouteGroup {
	buildingRepository := repositories.NewBuildingRepo()
	buildingService := services.NewBuildingService(buildingRepository, db, validate)
	buildingController := controllers.NewBuildingController(buildingService)

	apiStudentRoute := group.New("/building").Children(
		group.New("").GET(buildingController.FindAll).POST(buildingController.Create),
		group.New("/:buildingId").GET(buildingController.FetchById).POST(buildingController.Update),
	)

	return apiStudentRoute
}
