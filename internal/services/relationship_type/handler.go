package relationship_type

import (
	"family-tree/internal/models"
	"family-tree/internal/repository/postgres/relationship_type"
)

type service struct {
	repo relationship_type.RelationshipTypeRepository
}

type Service interface {
	Create(relationshipType *models.RelationshipType) error
	GetByType(relationType string) (*models.RelationshipType, error)
	List() ([]models.RelationshipType, error)
}

func NewService(repo relationship_type.RelationshipTypeRepository) Service {
	return &service{
		repo: repo,
	}
}

func (s service) Create(relationshipType *models.RelationshipType) error {
	rltType, err := s.repo.GetByType(relationshipType.Type)
	if rltType.Type == "" {

		err = s.repo.Create(relationshipType)
		if err != nil {
			return err
		}

		return nil
	}

	return err
}

func (s service) GetByType(relationType string) (*models.RelationshipType, error) {
	relationshipType, err := s.repo.GetByType(relationType)
	if err != nil {
		return nil, err
	}

	return relationshipType, nil
}

func (s service) List() ([]models.RelationshipType, error) {
	relationshipTypes, err := s.repo.List()

	return relationshipTypes, err
}
