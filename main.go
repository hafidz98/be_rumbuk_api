package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"text/template"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/hafidz98/be_rumbuk_api/app"

	//"github.com/hafidz98/be_rumbuk_api/exception"
	"github.com/hafidz98/be_rumbuk_api/helper"
	middleware "github.com/hafidz98/be_rumbuk_api/middlewares"
	"github.com/hafidz98/be_rumbuk_api/routes"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	group "github.com/mythrnr/httprouter-group"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		helper.Warning.Println(err)
		return
	}
}

func StartNonTLSServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		helper.Info.Println("Redirecting to :443")
		http.Redirect(w, r, "https://purwacode.my.id", http.StatusTemporaryRedirect)
	})

	http.ListenAndServe(":80", nil)
}

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	basepath := os.Getenv("API_BASE_PATH")
	// address := os.Getenv("APP_ADDRESS") + ":" + os.Getenv("APP_PORT")
	address := ":" + os.Getenv("APP_PORT")

	db := app.NewDB()
	validate := validator.New()
	router := httprouter.New()
	//router.PanicHandler = exception.ErrorHandler

	//go StartNonTLSServer()

	mainRoute := group.New(basepath).Middleware(middleware.CommonMiddleware).Children(
		routes.AuthRoute(db, validate),
		routes.StudentRoute(db, validate),
		routes.StaffRoute(db, validate),
		routes.RoomRoute(db, validate),
		routes.BuildingRoute(db, validate),
		routes.TimeslotRoute(db, validate),
		routes.FloorRoute(db, validate),
		routes.AvailableRoomRoute(db, validate),
	)

	router.HandlerFunc(http.MethodGet, "/routes", func(w http.ResponseWriter, r *http.Request) {
		type ListRoute struct {
			Method string
			Path string
		}

		routeList := []ListRoute{}
		for _, r := range mainRoute.Routes() {
			routeList = append(routeList, ListRoute{Method: r.Method(), Path: r.Path()})
		}
		
		tmpl := template.Must(template.ParseFiles("fe/routes.view.html"))
		if err := tmpl.Execute(w, routeList); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	router.HandlerFunc(http.MethodGet, "/routes/list", func(w http.ResponseWriter, r *http.Request) {

		type ListRoute struct {
			Method string `json:"method"`
			Path string `json:"path"`
		}

		routeList := []ListRoute{}

		for _, r := range mainRoute.Routes() {
			routeList = append(routeList, ListRoute{Method: r.Method(), Path: r.Path()})
		}
		helper.WriteToResponseBody(w, routeList)
	})

	//helper.Info.Print("\n", mainRoute.Routes().String())

	for _, r := range mainRoute.Routes() {
		router.Handle(r.Method(), r.Path(), r.Handler())
	}

	server := http.Server{
		Addr:    address,
		Handler: router,
	}

	// go func() {
	// 	if err := server.ListenAndServeTLS("cert/certificate.crt", "cert/private.key"); err != nil && err != http.ErrServerClosed {
	// 		helper.Error.Println(err)
	// 	}
	// }()

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// helper.Info.Println(address)
			helper.Error.Println(err)
		}
	}()

	helper.Info.Printf("Listening and serving on %v\n", server.Addr)

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
	helper.Info.Println("Server terminated properly")
	os.Exit(0)
}

// TODO: Make test case for api endpoint
// TODO: Staff api
// TODO: Multi username login
