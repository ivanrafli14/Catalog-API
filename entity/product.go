package entity

import "mime/multipart"

type Product struct {
	ID         string `json:"id" gorm:"primaryKey;type:varchar(50)"`
	Name       string `json:"name" gorm:"type:varchar(255)"`
	ImageUrl   string `json:"image_url" gorm:"type:varchar(255)"`
	Stock      int    `json:"stock" gorm:"type:int"`
	Price      int    `json:"price" gorm:"type:int"`
	CategoryID uint   `json:"category_id" gomr:"type:int"`
    Category   Category `json:"category" gorm:"foreignKey:CategoryID"`
}

type ProductRequest struct {
	Name       string `json:"name" binding:"required"`
	ImageUrl   string `json:"image_url" binding:"required"`
	Stock      int    `json:"stock" binding:"required"`
	Price      int    `json:"price" binding:"required"`
	CategoryID uint   `json:"category_id" binding:"required"`
}

type ProductPhotoRequest struct {
	File     *multipart.FileHeader `form:"file" binding:"required"`
	Type string                `form:"type" binding:"required"`

}
