package relationship

import (
	"family-tree/internal/models"
	"family-tree/internal/repository/postgres/person"
	"family-tree/internal/repository/postgres/relationship"
	"family-tree/internal/repository/postgres/relationship_type"
	"family-tree/internal/utils"
	"log"

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
	GetRelationByName(name string) (*models.Data, error)
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

func (s service) GetRelationByName(name string) (*models.Data, error) {
	person, err := s.personRepo.GetByName(name)
	if err != nil {
		return nil, err
	}

	relationships, err := s.relationshipRepo.GetByPersonID(person.ID)
	if err != nil {
		return nil, err
	}

	peopleID := []string{}
	mapType := make(map[string]int)

	for _, relationship := range relationships {
		peopleID = append(peopleID, relationship.PersonAID, relationship.PersonBID)
		log.Println(peopleID)
		mapType[relationship.PersonAID] = relationship.TypeID
		mapType[relationship.PersonBID] = relationship.TypeID
	}

	filteredPeople := utils.RemoveDuplicateValues(peopleID, person.ID)

	data := models.Data{}

	data.Name = person.Name

	for _, filteredPerson := range filteredPeople {
		member := models.Member{}

		person, err := s.personRepo.GetByID(filteredPerson)
		if err != nil {
			return nil, err
		}

		member.Name = person.Name

		relationshipType, err := s.relationshipTypeRepo.GetByID(mapType[filteredPerson])
		if err != nil {
			return nil, err
		}

		member.Relationship = relationshipType.Type
		data.Member = append(data.Member, member)
	}

	return &data, nil
}
