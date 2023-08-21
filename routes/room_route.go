package routes

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/hafidz98/be_rumbuk_api/controllers"
	"github.com/hafidz98/be_rumbuk_api/repositories"
	"github.com/hafidz98/be_rumbuk_api/services"
	group "github.com/mythrnr/httprouter-group"
)

func RoomRoute(db *sql.DB, validate *validator.Validate) *group.RouteGroup {
	roomRepository := repositories.NewRoomRepo()
	roomService := services.NewRoomService(roomRepository, db, validate)
	roomCtrl := controllers.NewRoomController(roomService)

	apiRoomRoute := group.New("rooms").GET(roomCtrl.FetchAll)

	return apiRoomRoute
}
