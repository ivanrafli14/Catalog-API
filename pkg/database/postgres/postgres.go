package postgres

import (
	"log"

	"github.com/ivanrafli14/CatalogAPI/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.LoadDataSourceName()), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	return db
}	