package postgres

import (
	"log"

	"github.com/ivanrafli14/CatalogAPI/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB){
	if err:= db.AutoMigrate(
		&entity.User{},
		&entity.Category{},
		&entity.Product{},
	); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}
}