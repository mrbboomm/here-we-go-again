package usecases

import (
	"go-nf/domains"
	"go-nf/entities"
)

type userUseCase struct {
	repo domains.UserRepo
}

func NewUserUseCase(repo domains.UserRepo) domains.UserUseCase{
	return &userUseCase{repo}
}

func(u *userUseCase)GetUsers()[]entities.UserEntity{
users :=	u.repo.FindAll()

return users
}

