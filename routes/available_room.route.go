package routes

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/hafidz98/be_rumbuk_api/controllers"
	"github.com/hafidz98/be_rumbuk_api/repositories"
	"github.com/hafidz98/be_rumbuk_api/services"
	group "github.com/mythrnr/httprouter-group"
)

func AvailableRoomRoute(db *sql.DB, validate *validator.Validate) *group.RouteGroup {
	availableRoomRepository := repositories.NewAvailableRoomRepo()
	availableRoomService := services.NewAvailableRoomService(availableRoomRepository, db, validate)
	availableRoomController := controllers.NewAvailableRoomController(availableRoomService)

	apiAvailableRoomRoute := group.New("available-room").Children(
		group.New("").GET(availableRoomController.FetchAllAvailableRoom),
	)

	return apiAvailableRoomRoute
}
