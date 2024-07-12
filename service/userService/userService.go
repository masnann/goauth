package userservice

import (
	"errors"
	"go-auth/helpers"
	"go-auth/models"
	"go-auth/service"
	"log"
)

type UserService struct {
	service service.Service
}

func NewUserService(service service.Service) UserService {
	return UserService{
		service: service,
	}
}

func (s UserService) Register(req models.UserRegisterRequest) (int64, error) {

	hash, err := s.service.Generator.GenerateHash(req.Password)
	if err != nil {
		log.Println("Error generating hash: ", err)
		return 0, errors.New("failed to generate hash")
	}

	newData := models.UserModels{
		Username:  req.Username,
		Email:     req.Email,
		Password:  hash,
		Role:      "user",
		Status:    "",
		CreatedAt: helpers.TimeStampNow(),
		UpdatedAt: "",
	}

	result, err := s.service.UserRepo.Register(newData)
	if err != nil {
		log.Println("Error registering user: ", err)
		return 0, errors.New("failed to register user")
	}
	return result, nil
}

func (s UserService) FindUserByID(req models.RequestID) (models.UserModels, error) {
	result, err := s.service.UserRepo.FindUserByID(req.ID)
	if err != nil {
		log.Println("Error finding user by ID: ", err)
		return result, errors.New("user not found")
	}
	return result, nil
}
