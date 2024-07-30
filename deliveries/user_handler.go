package deliveries

import (
	"go-nf/domains"
	"go-nf/entities"

	"github.com/gofiber/fiber/v2"
)

type countryHandlers struct {
	countryUseCase domains.CountryUseCase
}

func NewCountryHandler(countryUseCase domains.CountryUseCase) *countryHandlers {
	return &countryHandlers{countryUseCase}
}

func (c *countryHandlers) CreateCountry(f *fiber.Ctx) error {
	var country entities.CountryEntity
	if err := f.BodyParser(&country); err != nil {
		return f.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if err := c.countryUseCase.CreateCountry(country); err != nil {
		return f.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return f.Status(fiber.StatusOK).JSON(true)
}
