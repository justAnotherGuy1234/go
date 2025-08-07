package app

import (
	dbConfig "AuthInGo/config/db"
	config "AuthInGo/config/env"
	"AuthInGo/controller"
	db "AuthInGo/db/repo"
	"AuthInGo/router"
	"AuthInGo/service"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Addr string // port
}

type Application struct {
	Config Config
	Store  db.Storage
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
		Store:  *db.NewStore(),
	}
}

func (app *Application) Run() error {

	dbConf, err := dbConfig.SetupDb()

	if err != nil {
		fmt.Println("error connecting to db", err)
		return nil
	}

	ur := db.NewUserRepo(dbConf)
	us := service.NewUserService(ur)
	uc := controller.NewUserController(us)

	uRouter := router.NewUserRouter(uc)

	server := http.Server{
		Addr:         app.Config.Addr,
		Handler:      router.SetupRouter(uRouter),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("server started at ", app.Config.Addr)

	return server.ListenAndServe()
}
