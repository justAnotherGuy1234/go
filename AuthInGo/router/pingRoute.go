package router

import (
	"AuthInGo/controller"

	"github.com/go-chi/chi/v5"
)

type Router interface {
	Register(r chi.Router)
}

func SetupRouter(UserRouter Router) *chi.Mux {
	chiRouter := chi.NewRouter()

	chiRouter.Get("/ping", controller.PingController)

	UserRouter.Register(chiRouter)

	return chiRouter
}
