package routes

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/hafidz98/be_rumbuk_api/controllers"
	"github.com/hafidz98/be_rumbuk_api/repositories"
	"github.com/hafidz98/be_rumbuk_api/services"
	group "github.com/mythrnr/httprouter-group"
)

func FloorRoute(db *sql.DB, validate *validator.Validate) *group.RouteGroup {
	floorRepository := repositories.NewFloorRepo()
	floorService := services.NewFloorService(floorRepository, db, validate)
	floorController := controllers.NewFloorController(floorService)

	floorEndpoint := group.New("/floor").POST(floorController.Create).GET(floorController.GetAll).Children(
		group.New("/:floorId").GET(floorController.GetByID),
	)

	return floorEndpoint
}