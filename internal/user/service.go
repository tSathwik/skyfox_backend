package user

import (

)

type UserService interface{
	CreateUser (user *User) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService{
	return &userService{repo: repo}
}

func (u *userService) CreateUser(user *User) error {
	return u.repo.CreateUser(user)
}