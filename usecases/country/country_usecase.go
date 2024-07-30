package usecases

import (
	"go-nf/domains"
	"go-nf/entities"
)

type countryUseCase struct {
	repo domains.CountryRepo
}

func NewCountryUseCase(repo domains.CountryRepo) domains.CountryUseCase {
	return &countryUseCase{repo}
}

func (c *countryUseCase) CreateCountry(data entities.CountryEntity) error {
	err := c.repo.Create(data)
	if err != nil {
		return err
	}
	return nil
}
