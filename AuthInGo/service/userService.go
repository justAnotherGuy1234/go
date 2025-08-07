package service

import (
	db "AuthInGo/db/repo"
	"fmt"
)

type UserService interface {
	GetUserByIdService() error
	CreateUser() error
}

type UserServiceImpl struct {
	userRepo db.UserRepo
}

func NewUserService(_userRepo db.UserRepo) UserService {
	return &UserServiceImpl{
		userRepo: _userRepo,
	}
}

func (u *UserServiceImpl) CreateUser() error {
	fmt.Println("create user in service")
	u.userRepo.Create()
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
