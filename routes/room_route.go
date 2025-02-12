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
	roomTimeslotRepo := repositories.NewRoomTimeslotRepo()
	roomService := services.NewRoomService(roomRepository, db, validate)
	roomTimeslotService := services.NewRoomTimeslotService(roomTimeslotRepo, db, validate)
	roomCtrl := controllers.NewRoomController(roomService)
	roomTimeslotCtrl := controllers.NewRoomTimeslotController(roomTimeslotService)

	apiRoomRoute := group.New("room").Children(
		group.New("").GET(roomCtrl.FetchAllRooms).POST(roomCtrl.Create).Children(
			group.New("/:roomId").GET(roomCtrl.FindById).PATCH(roomCtrl.Update),
			group.New("/:roomId").DELETE(roomCtrl.Delete).Children(
				group.New("/change_status").PATCH(roomCtrl.UpdateRoomStatus),
			),
			group.New("/add_timeslot").POST(roomTimeslotCtrl.AddRoomTimeslot),
		),
	)

	return apiRoomRoute
}
