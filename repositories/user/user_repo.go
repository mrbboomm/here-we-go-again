package repositories

import (
	"go-nf/domains"
	"go-nf/entities"
)

// FIXME: will use db later once gom package is being set
type userRepo struct {
	data []entities.UserEntity
}

func NewUserRepo(data []entities.UserEntity) domains.UserRepo{
	return &userRepo{data}
}

func(u *userRepo)FindAll()[]entities.UserEntity{
return u.data
}