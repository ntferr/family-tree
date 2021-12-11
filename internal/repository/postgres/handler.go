package postgres

import (
	"family-tree/internal/models"
	"family-tree/internal/settings"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func startDB() (*gorm.DB, error) {
	log.Println("Connecting  to Database")

	log.Println(settings.GetSettings().Host)

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v",
		settings.GetSettings().Host, settings.GetSettings().Database.User, settings.GetSettings().Database.Password,
		settings.GetSettings().Database.Name, settings.GetSettings().Database.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, models.ErrOnConnectDB
	}

	db.Debug()

	dbConfig, err := db.DB()
	if err != nil {
		return nil, err
	}

	dbConfig.SetConnMaxLifetime(time.Minute * 5)
	dbConfig.SetMaxOpenConns(15)
	dbConfig.SetMaxIdleConns(5)

	return db, nil
}

func GetDB() (*gorm.DB, error) {
	return startDB()
}
