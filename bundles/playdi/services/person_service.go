package services

import (
	"dbstair/bundles/playdi"
	"dbstair/bundles/playdi/models"
)

type PersonService struct {
	config     *playdi.Config
	repository *models.PersonRepository
}

func (service *PersonService) FindAll() []*Person {
	if service.config.Enabled {
		return service.repository.FindAll()
	}

	return []*Person{}
}

func NewPersonService(config *Config, repository *PersonRepository) *PersonService {
	return &PersonService{config: config, repository: repository}
}
