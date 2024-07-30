package domains

import "go-nf/entities"

type CountryUseCase interface {
	CreateCountry(data entities.CountryEntity) error
}

type CountryRepo interface {
	Create(data entities.CountryEntity) error
}