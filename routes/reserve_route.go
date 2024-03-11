package routes

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/hafidz98/be_rumbuk_api/controllers"
	"github.com/hafidz98/be_rumbuk_api/repositories"
	"github.com/hafidz98/be_rumbuk_api/services"
	group "github.com/mythrnr/httprouter-group"
)

func ReservationRoute(db *sql.DB, validate *validator.Validate) *group.RouteGroup {
	reserveRepository := repositories.NewReserveRoomRepo()
	reserveService := services.NewReservationService(reserveRepository, db, validate)
	reserveController := controllers.NewReservationController(reserveService)

	reservationEndpoint := group.New("/reservation").POST(reserveController.Create).GET(reserveController.GetReservationByStudentID).Children(
		group.New("/cancel").POST(reserveController.CancelReservation),
	)

	return reservationEndpoint
}