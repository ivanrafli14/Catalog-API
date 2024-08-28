package repository

import (
	"errors"

	"github.com/ivanrafli14/CatalogAPI/entity"
	"gorm.io/gorm"
)

var (
	ErrRepoNotFound = errors.New("repository: data not found")
) 

type IUserRepository interface {
	CreateUser(user *entity.User) (*entity.User, error)
	FindByEmail(email string) (*entity.User,error)
	FindByID(id string) (*entity.User,error)

}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) CreateUser(user *entity.User) (*entity.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (*entity.User,error){
	var user *entity.User
	err := r.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRepoNotFound
		}

		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindByID(id string) (*entity.User,error){
	var user *entity.User
	err := r.db.Where("id = ?", id).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRepoNotFound
		}

		return nil, err
	}
	return user, nil
}