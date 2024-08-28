package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/ivanrafli14/CatalogAPI/entity"
	"github.com/ivanrafli14/CatalogAPI/internal/repository"
	"github.com/ivanrafli14/CatalogAPI/pkg/bcrypt"
	"github.com/ivanrafli14/CatalogAPI/pkg/database/redis"
	"github.com/ivanrafli14/CatalogAPI/pkg/jwt"
)

var (
	ErrDuplicateEmail = errors.New("email already registered")
	ErrInvalidEmail = errors.New("email not registered")
)

type IUserService interface {
	Login(*entity.UserRequest) (*entity.LoginResponse, error)
	Register(*entity.UserRequest) (*entity.User, error)
	FindByID(id string) (*entity.User, error)
}

type UserService struct {
	ur repository.IUserRepository
	bcrypt bcrypt.Interface
	jwtAuth jwt.Interface
	redis redis.Interface
}

func NewUserService(ur repository.IUserRepository, bcrypt bcrypt.Interface, jwtAuth jwt.Interface, redis redis.Interface) IUserService {
	return &UserService{ur, bcrypt, jwtAuth, redis}
}

func(s * UserService) Register(userReq *entity.UserRequest) (*entity.User, error){
	userRes, err := s.ur.FindByEmail(userReq.Email)
	if userRes != nil {
		return nil, ErrDuplicateEmail
	}

	hashedPasssword, err := s.bcrypt.GenerateFromPassword(userReq.Password)

	if err != nil {
		return nil, err
	}

	userParse := &entity.User{
		ID: uuid.New().String(),
		Email: userReq.Email,
		Password: hashedPasssword,
		Role: "merchant",
	}

	user, err := s.ur.CreateUser(userParse)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) Login(userReq *entity.UserRequest) (*entity.LoginResponse,error){

	userRegistered, err := s.ur.FindByEmail(userReq.Email)
	response := &entity.LoginResponse{}
	if err != nil {
		if errors.Is(err, repository.ErrRepoNotFound) {
			return response, ErrInvalidEmail
			
		}
		return response, err
	}

	
	if err := s.bcrypt.CompareAndHashPasswrord(userRegistered.Password, userReq.Password); err != nil {
		return response, err
	}

	accessToken, err := s.jwtAuth.CreateJWTToken(userRegistered.ID)

	if err != nil {
		return response, err
	}

	if err := s.redis.SetData(accessToken, "true", 10 * time.Minute); err != nil {
		return response, err
	}

	response.Role = userRegistered.Role
	response.Token = accessToken

	return response,nil

}

func (s *UserService) FindByID(id string) (*entity.User, error){
	user, err := s.ur.FindByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}