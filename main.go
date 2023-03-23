package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/hafidz98/be_rumbuk_api/app"
	"github.com/hafidz98/be_rumbuk_api/controller"
	"github.com/hafidz98/be_rumbuk_api/helper"
	"github.com/hafidz98/be_rumbuk_api/repository"
	"github.com/hafidz98/be_rumbuk_api/service"
	"github.com/julienschmidt/httprouter"
)

func main() {
	//helper.Info.Println("RUMBUK API STARTING")

	db := app.NewDB()
	validate := validator.New()
	router := httprouter.New()

	studentRepository := repository.NewStudentRepo()
	studentService := service.NewStudentService(studentRepository, db, validate)
	studentController := controller.NewStudentController(studentService)

	router.GET("/students", studentController.FindAll)
	router.GET("/students/:studentId", studentController.FetchById)
	router.POST("/students", studentController.Create)

	server := http.Server{
		Addr:    "localhost:2023",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
