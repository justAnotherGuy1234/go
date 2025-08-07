package router

import (
	"AuthInGo/controller"

	"github.com/go-chi/chi/v5"
)

type UserRouter struct {
	UserController *controller.UserController
}

func NewUserRouter(_userController *controller.UserController) Router {
	return &UserRouter{
		UserController: _userController,
	}
}

func (uc *UserRouter) Register(r chi.Router) {
	r.Post("/signup", uc.UserController.Register)
	r.Get("/profile", uc.UserController.GetUserByIdService)
}
