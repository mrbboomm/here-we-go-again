package deliveries

import (
	"go-nf/domains"

	"github.com/gofiber/fiber/v2"
)

type userHandlers struct {
	userUseCase domains.UserUseCase
}

func NewUserHandler(userUseCase domains.UserUseCase) *userHandlers{
	return &userHandlers{userUseCase}
}


func(u *userHandlers)GetAllUsers(c *fiber.Ctx) error{

users := u.userUseCase.GetUsers()
return c.Status(fiber.StatusOK).JSON(users)
}