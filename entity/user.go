package entity

type User struct {
	ID       string `json:"id" gorm:"primaryKey;type:varchar(50)"`
	Email    string `json:"email" gorm:"type:varchar(255)"`
	Password string `json:"password" gorm:"type:varchar(255)"`
	Role     string `json:"role" gorm:"type:varchar(255)"`
}

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginResponse struct {
	Token string `json:"access_token"`
	Role  string `json:"role"`
}
