package controller

import (
	"AuthInGo/service"
	"fmt"
	"net/http"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(_userService service.UserService) *UserController {
	return &UserController{
		UserService: _userService,
	}
}

func (uc *UserController) Register(w http.ResponseWriter, r *http.Request) {
	uc.UserService.CreateUser()
	w.Write([]byte("register endpoint in controller"))
}

func (uc *UserController) GetUserByIdService(w http.ResponseWriter, r *http.Request) {
	fmt.Println("fetching user detaild in controller")

	uc.UserService.LoginUser()
	uc.UserService.GetUserByIdService()
	w.Write([]byte("fetching user endpont"))

}
