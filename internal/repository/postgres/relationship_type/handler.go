package relationship_type

import (
	"family-tree/internal/models"
	"family-tree/internal/repository/postgres"
)

type relationshipTypeRepository struct{}

type RelationshipTypeRepository interface {
	Create(obj *models.RelationshipType) error
	GetByType(relationshipType string) (*models.RelationshipType, error)
}

func NewRelationshipTypeRepository() RelationshipTypeRepository {
	return &relationshipTypeRepository{}
}

func (rtr relationshipTypeRepository) Create(obj *models.RelationshipType) error {
	db, err := postgres.GetDB()
	if err != nil {
		return err
	}

	return db.Create(obj).Error
}

func (rtr relationshipTypeRepository) GetByType(relationType string) (*models.RelationshipType, error) {
	db, err := postgres.GetDB()
	if err != nil {
		return nil, err
	}

	var relationshipType models.RelationshipType

	err = db.Where(&models.RelationshipType{Type: relationType}).First(&relationshipType).Error

	return &relationshipType, err
}
