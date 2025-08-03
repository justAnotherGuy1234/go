package app

import (
	config "AuthInGo/config/env"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Addr string // port
}

type Application struct {
	Config Config
}

func NewConifg() Config {

	port := config.GetString("PORT", ":8080")
	return Config{
		Addr: port,
	}
}

func NewApplication(cfg Config) *Application {
	return &Application{
		Config: cfg,
	}
}

func (app *Application) Run() error {
	server := http.Server{
		Addr:         app.Config.Addr,
		Handler:      nil, // setup and add chi router here .
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("server started at ", app.Config.Addr)

	return server.ListenAndServe()
}
