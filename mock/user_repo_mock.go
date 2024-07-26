package mock

import (
	"go-nf/entities"

	"github.com/stretchr/testify/mock"
)

type userRepoMock struct {
	mock.Mock
}

func NewUserRepoMock() *userRepoMock{
	return &userRepoMock{}
}

func(u *userRepoMock)FindAll()[]entities.UserEntity{
	args := u.Called()
	return args.Get(0).([]entities.UserEntity)
}