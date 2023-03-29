package main

import (
	"context"
	"flag"

	//"fmt"
	"net/http"
	"os"
	"os/signal"

	//"syscall"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/hafidz98/be_rumbuk_api/app"
	"github.com/hafidz98/be_rumbuk_api/routes"

	//controller "github.com/hafidz98/be_rumbuk_api/controllers"

	middleware "github.com/hafidz98/be_rumbuk_api/middlewares"
	"github.com/joho/godotenv"
	group "github.com/mythrnr/httprouter-group"

	//"github.com/hafidz98/be_rumbuk_api/exception"
	"github.com/hafidz98/be_rumbuk_api/helper"

	//repository "github.com/hafidz98/be_rumbuk_api/repositories"
	//service "github.com/hafidz98/be_rumbuk_api/services"
	"github.com/julienschmidt/httprouter"
)

func init() {
	err := godotenv.Load(".env")
	helper.PanicIfError(err)
}

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	//helper.Info.Println("RUMBUK API STARTING")

	basepath := "/v1/rumbuk"
	db := app.NewDB()
	validate := validator.New()
	router := httprouter.New()

	routes.StudentRoute(db, validate)

	mainRoute := group.New(basepath).Middleware(middleware.CommonMiddleware).Children(
		routes.AuthRoute(db, validate),
		routes.StudentRoute(db, validate),
	)

	// studentRepository := repository.NewStudentRepo()
	// studentService := service.NewStudentService(studentRepository, db, validate)
	// studentController := controller.NewStudentController(studentService)

	// g := group.New(basePath).Middleware(middleware.CommonMiddleware).Children(
	// 	group.New("/students").Middleware(middleware.AuthMiddleware).GET(studentController.FindAll).Children(
	// 		group.New("/:studentId").GET(studentController.FetchById),
	// 	),
	// )

	// Students
	// router.POST(basePath+"/auth/students", studentController.Login)
	// router.GET(basePath+"/students", middleware.AuthMiddleware(studentController.FindAll))
	// router.GET(basePath+"/students/:studentId", middleware.AuthMiddleware(studentController.FetchById))
	// router.POST(basePath+"/students", studentController.Create)
	// router.PATCH(basePath+"/students/:studentId", middleware.AuthMiddleware(studentController.Update))
	// router.DELETE(basePath+"/students/:studentId", studentController.Delete)

	// Rooms
	// router.GET("/rooms", nil)
	// router.GET("/rooms", nil)
	// router.POST("/rooms", nil)
	// router.PATCH("/rooms", nil)
	// router.DELETE("/rooms", nil)

	//
	/*
		router.GET("", nil)
		router.GET("", nil)
		router.POST("", nil)
		router.PATCH("", nil)
		router.DELETE("", nil)
	*/

	//router.PanicHandler = exception.ErrorHandler

	//fmt.Println(g.Routes().String())

	for _, r := range mainRoute.Routes() {
		router.Handle(r.Method(), r.Path(), r.Handler())
	}

	server := http.Server{
		Addr:    "localhost:8991",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			helper.Error.Println(err)
		}
	}()

	helper.Info.Printf("Listening and serving on port %v\n", server.Addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer func() {
		err := db.Close()
		helper.Error.Panicln(err)
		cancel()
	}()

	helper.Info.Println("Shutting down server...")
	if err := server.Shutdown(ctx); err != nil {
		helper.Error.Fatalf("Server forced to shutdown: %v\n", err)
	}
	helper.Info.Println("Server Exited Properly")
	os.Exit(0)
}
