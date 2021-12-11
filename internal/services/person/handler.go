package person

import (
	"family-tree/internal/models"
	repository "family-tree/internal/repository/postgres/person"
)

type service struct {
	repo repository.PersonRepository
}

type Service interface {
	Create(person *models.Person) error
	GetBy(id string, name string) (*models.Person, error)
}

func NewService(repo repository.PersonRepository) Service {
	return &service{
		repo: repo,
	}
}

func (s service) Create(person *models.Person) error {
	return s.repo.Create(person)
}

func (s service) GetBy(id string, name string) (*models.Person, error) {
	if id != "" {

		return nil, nil
	}
	if name != "" {

		return nil, nil
	}

	return nil, models.ErrToGetParamPerson
}
