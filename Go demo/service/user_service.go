package service

import (
	domain "demo/domain/dto"
	"fmt"
	"time"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) Login(user *domain.UserDto) (*domain.UserDto, error) {
	time.Sleep(time.Second)
	fmt.Println(user)
	return &domain.UserDto{Name: "lala", Password: "tete"}, nil
}
