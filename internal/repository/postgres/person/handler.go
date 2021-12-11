package person

import (
	"family-tree/internal/models"
	"family-tree/internal/repository/postgres"
)

type personRepository struct{}

type PersonRepository interface {
	Create(person *models.Person) error
	GetByID(id string) (*models.Person, error)
	GetByName(name string) (*models.Person, error)
	List() ([]models.Person, error)
	Update(person *models.Person) error
	Delete(id string) error
}

func NewPersonRepository() PersonRepository {
	return &personRepository{}
}

func (pr personRepository) Create(person *models.Person) error {
	db, err := postgres.GetDB()
	if err != nil {
		return err
	}

	return db.Create(person).Error
}

func (pr personRepository) GetByID(id string) (*models.Person, error) {
	db, err := postgres.GetDB()
	if err != nil {
		return nil, err
	}

	person := models.Person{}

	err = db.
		Where(&models.Person{ID: id}).
		First(&person).Error

	return &person, err
}

func (pr personRepository) GetByName(name string) (*models.Person, error) {
	db, err := postgres.GetDB()
	if err != nil {
		return nil, err
	}

	person := models.Person{}

	err = db.
		Where(&models.Person{Name: name}).
		First(&person).Error

	return &person, err
}

func (pr personRepository) List() ([]models.Person, error) {
	db, err := postgres.GetDB()
	if err != nil {
		return nil, err
	}

	people := []models.Person{}

	err = db.
		Find(&people).Error

	return people, err
}

func (pr personRepository) Update(person *models.Person) error {
	db, err := postgres.GetDB()
	if err != nil {
		return err
	}

	return db.
		Where(&models.Person{ID: person.ID}).
		UpdateColumns(&models.Person{Name: person.Name}).Error

}

func (pr personRepository) Delete(id string) error {
	db, err := postgres.GetDB()
	if err != nil {
		return err
	}

	return db.
		Delete(&models.Person{}, id).Error
}
