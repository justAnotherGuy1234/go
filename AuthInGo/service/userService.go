package service

import (
	env "AuthInGo/config/env"
	db "AuthInGo/db/repo"
	"AuthInGo/dto"
	"AuthInGo/util"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	GetUserByIdService() error
	CreateUser(userData *dto.SignUpUserDto) error
	LoginUser() any
}

type UserServiceImpl struct {
	userRepo db.UserRepo
}

func NewUserService(_userRepo db.UserRepo) UserService {
	return &UserServiceImpl{
		userRepo: _userRepo,
	}
}

func (u *UserServiceImpl) CreateUser(userData *dto.SignUpUserDto) error {

	hashed, err := util.HashPassword(userData.Password)

	if err != nil {
		fmt.Println("error hashing password", err)
	}

	u.userRepo.Create(userData.Username, userData.Email, hashed)

	return nil
}

func (u *UserServiceImpl) GetUserByIdService() error {
	fmt.Println("fetching user id details")

	err := u.userRepo.GetUserById()

	if err != nil {
		fmt.Println("error getting user data ", err)
	}
	return nil
}

func (u *UserServiceImpl) LoginUser() any {
	// create login func in repo which expects email and password

	payload := jwt.MapClaims{
		"email": "test",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenString, err := token.SignedString([]byte(env.GetString("JWT_SECRET", "TOKEN")))

	if err != nil {
		fmt.Println("failed to create jwt", err)
	}

	fmt.Println("token", tokenString)
	return tokenString
}
