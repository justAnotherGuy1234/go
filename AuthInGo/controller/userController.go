package controller

import (
	"AuthInGo/dto"
	"AuthInGo/service"
	"AuthInGo/util"
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

	var payload dto.SignUpUserDto

	if jsonErr := util.ReadJson(r, &payload); jsonErr != nil {
		w.Write([]byte("failed to read json"))
		return
	}

	if validationErr := util.Validator.Struct(payload); validationErr != nil {
		w.Write([]byte("incorrect input failed"))
		return
	}

	res := uc.UserService.CreateUser(&payload)

	response := map[string]any{
		"msg":  "sign up done",
		"data": res,
	}

	util.JsonReponse(w, http.StatusOK, response)
}

func (uc *UserController) GetUserByIdService(w http.ResponseWriter, r *http.Request) {
	fmt.Println("fetching user detaild in controller")

	jwtToken := uc.UserService.LoginUser()

	response := map[string]any{
		"msg":  "login sucessfull",
		"data": jwtToken,
	}

	util.JsonReponse(w, http.StatusOK, response)
}
