package main

import (
	"AuthInGo/app"
)

func main() {

	cfg := app.NewConifg()

	app := app.NewApplication(cfg)
	app.Run()

}
