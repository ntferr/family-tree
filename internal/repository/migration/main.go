package main

import (
	"log"

	"family-tree/internal/models"
	"family-tree/internal/repository/postgres"
)

func main() {
	log.Println("Runing migration")

	db, err := postgres.GetDB()
	if err != nil {
		log.Fatal(err)
		return
	}

	if err = db.AutoMigrate(
		&models.Person{},
		&models.Relationship{},
		&models.RelationshipType{},
	); err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Migration successful")
}
