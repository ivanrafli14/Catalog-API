package postgres

import (
	"log"

	"github.com/ivanrafli14/CatalogAPI/entity"
	"gorm.io/gorm"
)

func generateCategory(db *gorm.DB) error {
	var roles []*entity.Category

	roles = append(roles,
		&entity.Category{
			ID:   1,
			Name: "Facial Wash",
		},
		&entity.Category{
			ID:   2,
			Name: "Moisturizer",
		},&entity.Category{
			ID:   3,
			Name: "Sunscreen",
		},&entity.Category{
			ID:   4,
			Name: "Serum",
		},&entity.Category{
			ID:   5,
			Name: "Toner",
		})

	if err := db.CreateInBatches(roles, 5).Error; err != nil {
		return err
	}
	return nil
}

func SeedDB(db *gorm.DB) {
	var len int64

	if err := db.Model(&entity.Category{}).Count(&len).Error; err != nil {
		log.Fatalf("Error while counting categories: %v", err)
	}

	if len == 0 {
		if err := generateCategory(db); err != nil {
			log.Fatalf("Error while generating book: %v", err)

		}
	}
}