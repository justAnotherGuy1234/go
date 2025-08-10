package router

import (
	"AuthInGo/controller"
	"AuthInGo/middleware"

	"github.com/go-chi/chi/v5"
)

type Router interface {
	Register(r chi.Router)
}

func SetupRouter(UserRouter Router) *chi.Mux {
	chiRouter := chi.NewRouter()
	chiRouter.Use(middleware.RequestLogger)
	chiRouter.Get("/ping", controller.PingController)

	UserRouter.Register(chiRouter)

	return chiRouter
}
