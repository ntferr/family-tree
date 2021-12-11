package person

import (
	"family-tree/internal/models"
	repository "family-tree/internal/repository/postgres/person"

	"github.com/google/uuid"
)

type service struct {
	repo repository.PersonRepository
}

type Service interface {
	Create(person *models.Person) error
	GetBy(id string, name string) (*models.Person, error)
	List() ([]models.Person, error)
	Update(person *models.Person) error
}

func NewService(repo repository.PersonRepository) Service {
	return &service{
		repo: repo,
	}
}

func (s service) Create(person *models.Person) error {
	person.ID = uuid.New().String()

	return s.repo.Create(person)
}

func (s service) GetBy(id string, name string) (*models.Person, error) {
	if id != "" {
		return s.repo.GetByID(id)
	}

	if name != "" {
		return s.repo.GetByName(id)
	}

	return nil, models.ErrToGetParamPerson
}

func (s service) List() ([]models.Person, error) {
	return s.repo.List()
}

func (s service) Update(person *models.Person) error {
	_, err := s.repo.GetByID(person.ID)
	if err != nil {
		return err
	}

	return s.repo.Update(person)
}

func (s service) Delete(id string) error {
	return s.repo.Delete(id)
}
