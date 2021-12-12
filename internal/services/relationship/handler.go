package relationship

import (
	"family-tree/internal/models"
	"family-tree/internal/repository/postgres/person"
	"family-tree/internal/repository/postgres/relationship"
	"family-tree/internal/repository/postgres/relationship_type"

	"github.com/google/uuid"
)

type service struct {
	personRepo           person.PersonRepository
	relationshipRepo     relationship.RelationshipRepository
	relationshipTypeRepo relationship_type.RelationshipTypeRepository
}

type Service interface {
	Create(relationship *models.Relationship) error
	GetByID(id string) (*models.Relationship, error)
	List() ([]models.Relationship, error)
}

func NewService(personRepo person.PersonRepository,
	relationshipRepo relationship.RelationshipRepository,
	relationshipTypeRepo relationship_type.RelationshipTypeRepository,
) Service {
	return &service{
		personRepo:           personRepo,
		relationshipRepo:     relationshipRepo,
		relationshipTypeRepo: relationshipTypeRepo,
	}
}

func (s service) Create(relationship *models.Relationship) error {
	relationship.ID = uuid.New().String()
	return s.relationshipRepo.Create(relationship)
}

func (s service) GetByID(id string) (*models.Relationship, error) {
	return s.relationshipRepo.GetByID(id)
}

func (s service) List() ([]models.Relationship, error) {
	return s.relationshipRepo.List()
}
