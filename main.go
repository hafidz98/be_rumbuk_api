package main

import (
	"github.com/hafidz98/be_rumbuk_api/app"
	"github.com/hafidz98/be_rumbuk_api/helper"
)

func main() {
	helper.Info.Println("RUMBUK API STARTING")

	app.NewDB()
	app.NewDB().Close()
}
