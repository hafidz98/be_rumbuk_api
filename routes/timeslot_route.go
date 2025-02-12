package routes

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/hafidz98/be_rumbuk_api/controllers"
	"github.com/hafidz98/be_rumbuk_api/repositories"
	"github.com/hafidz98/be_rumbuk_api/services"
	group "github.com/mythrnr/httprouter-group"
)

func TimeslotRoute(db *sql.DB, validate *validator.Validate) *group.RouteGroup {
	timeslotRepository := repositories.NewTimeSlotRepo()
	timeslotService := services.NewTimeslotService(timeslotRepository, db, validate)
	timeslotController := controllers.NewTimeslotController(timeslotService)

	timeslotEndpoint := group.New("/timeslot").POST(timeslotController.Create).GET(timeslotController.GetAll).Children(
		group.New("/:timeslotId").GET(timeslotController.GetById),
	)

	return timeslotEndpoint
}
