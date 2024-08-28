package entity

type Category struct {
	ID   uint    `json:"id" gorm:"primaryKey;type:int;autoIncrement"`
	Name string `json:"name" gorm:"type:varchar(255)"`

	Products []Product `json:"-" gorm:"foreignKey:category_id;references:id"`
}
