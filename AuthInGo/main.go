package main

import (
	"AuthInGo/app"
	config "AuthInGo/config/env"
)

func main() {

	config.Load() // load env varibles

	cfg := app.NewConifg()

	app := app.NewApplication(cfg)
	app.Run()

}
