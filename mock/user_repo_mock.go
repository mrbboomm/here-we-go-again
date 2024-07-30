package mock

import (
	"go-nf/entities"

	"github.com/stretchr/testify/mock"
)

type countryRepoMock struct {
	mock.Mock
}

func NewCountryRepoMock() *countryRepoMock{
	return &countryRepoMock{}
}

func(u *countryRepoMock)FindAll()[]entities.CountryEntity{
	args := u.Called()
	return args.Get(0).([]entities.CountryEntity)
}