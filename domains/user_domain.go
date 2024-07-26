package domains

import "go-nf/entities"

type UserUseCase interface {
	GetUsers() []entities.UserEntity
}


type UserRepo interface {
	FindAll() []entities.UserEntity
}