package relationship_type

import (
	"family-tree/internal/models"
	"family-tree/internal/repository/postgres"
)

type relationshipTypeRepository struct{}

type RelationshipTypeRepository interface {
	Create(obj *models.RelationshipType) error
	GetByType(relationshipType string) (*models.RelationshipType, error)
	List() ([]models.RelationshipType, error)
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

	relationshipType := models.RelationshipType{}

	err = db.Where(&models.RelationshipType{Type: relationType}).First(&relationshipType).Error

	return &relationshipType, err
}

func (rtr relationshipTypeRepository) List() ([]models.RelationshipType, error) {
	db, err := postgres.GetDB()
	if err != nil {
		return nil, err
	}

	relationshipTypes := []models.RelationshipType{}

	err = db.Find(&relationshipTypes).Error
	if err != nil {
		return nil, err
	}

	return relationshipTypes, nil
}
