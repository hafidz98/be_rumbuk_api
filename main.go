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
	var basePath = "/rumbuk"

	// Students
	router.GET(basePath+"/students", studentController.FindAll)
	router.GET(basePath+"/students/:studentId", studentController.FetchById)
	router.POST(basePath+"/students", studentController.Create)
	router.PATCH(basePath+"/students/:studentId", studentController.Update)
	router.DELETE(basePath+"/students/:studentId", studentController.Delete)

	// Rooms
	router.GET("/rooms", nil)
	router.GET("/rooms", nil)
	router.POST("/rooms", nil)
	router.PATCH("/rooms", nil)
	router.DELETE("/rooms", nil)

	//
	/*
		router.GET("", nil)
		router.GET("", nil)
		router.POST("", nil)
		router.PATCH("", nil)
		router.DELETE("", nil)
	*/

	server := http.Server{
		Addr:    "localhost:8991",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
