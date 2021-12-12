package relationship

import (
	"errors"
	"family-tree/internal/models"
	"family-tree/internal/repository/postgres"
)

type relationshipRepository struct{}

type RelationshipRepository interface {
	Create(relationship *models.Relationship) error
	GetByID(id string) (*models.Relationship, error)
	List() ([]models.Relationship, error)
	GetByPersonID(id string) ([]models.Relationship, error)
}

func NewRelationshipRepository() RelationshipRepository {
	return &relationshipRepository{}
}

func (rr relationshipRepository) Create(relationship *models.Relationship) error {
	db, err := postgres.GetDB()
	if err != nil {
		return err
	}

	return db.Create(relationship).Error
}

func (rr relationshipRepository) GetByID(id string) (*models.Relationship, error) {
	db, err := postgres.GetDB()
	if err != nil {
		return nil, err
	}

	relationship := models.Relationship{}

	err = db.
		Where(&models.Relationship{ID: id}).
		First(&relationship).Error

	return &relationship, err
}

func (rr relationshipRepository) List() ([]models.Relationship, error) {
	db, err := postgres.GetDB()
	if err != nil {
		return nil, err
	}

	relationships := []models.Relationship{}

	err = db.
		Find(&relationships).Error

	return relationships, err
}

func (rr relationshipRepository) GetByPersonID(id string) ([]models.Relationship, error) {
	db, err := postgres.GetDB()
	if err != nil {
		return nil, err
	}

	relationships := []models.Relationship{}
	relationships_b := []models.Relationship{}

	err = db.
		Where(&models.Relationship{PersonAID: id}).
		Find(&relationships).Error

	err = db.
		Where(&models.Relationship{PersonBID: id}).
		Find(&relationships_b).Error

	if relationships == nil || relationships_b == nil {
		return nil, errors.New("No data found")
	}

	relationships = append(relationships, relationships_b...)

	return relationships, nil
}
