package usecases_test

import (
	"go-nf/entities"
	"go-nf/mock"
	usecases "go-nf/usecases/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T){
	userRepo := mock.NewUserRepoMock()
	userRepo.On("FindAll").Return([]entities.UserEntity{{Id: "111", Name: "name111"},{Id: "222", Name: "name222"}})

	userUseCase := usecases.NewUserUseCase(userRepo)
	
	results := userUseCase.GetUsers()

	expected := []entities.UserEntity{{Id: "111", Name: "name111"},{Id: "222", Name: "name222"}}
	
	assert.Equal(t, expected, results)
}