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
		RoleID:    2,
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

func (s UserService) Login(req models.UserLoginRequest) (models.UserLoginResponse, error) {

	user, err := s.service.UserRepo.Login(req.Email)
	if err != nil {
		log.Println("Error finding user by email: ", err)
		return models.UserLoginResponse{}, errors.New("user not found")
	}

	isValidPassword, err := s.service.Generator.ComparePassword(user.Password, req.Password)
	if !isValidPassword || err != nil {
		log.Println("Error comparing password: ", err)
		return models.UserLoginResponse{}, errors.New("invalid password")
	}

	accessToken, err := s.service.Generator.GenerateJWT(user.ID, user.Email, user.RoleID)
	if err != nil {
		log.Println("Error generating JWT: ", err)
		return models.UserLoginResponse{}, errors.New("failed to generate access token")
	}

	refreshToken, err := s.service.Generator.GenerateRefreshToken(user.ID)
	if err != nil {
		log.Println("Error generating refresh token: ", err)
		return models.UserLoginResponse{}, errors.New("failed to generate refresh token")
	}

	permissions, err := s.service.UserRepo.FindUserPermissions(user.ID)
	if err != nil {
		log.Println("Error finding user permissions: ", err)
		return models.UserLoginResponse{}, errors.New("failed to find user permissions")
	}
	result := models.UserLoginResponse{
		UserID:       user.ID,
		RoleID:       user.RoleID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Permission:   permissions,
	}

	return result, nil
}

func (s UserService) RefreshToken(accessToken string) (models.UserLoginResponse, error) {

	userID, err := s.service.Generator.ValidateRefreshToken(accessToken)
	if err != nil {
		log.Println("Error validating access token: ", err)
		return models.UserLoginResponse{}, errors.New("invalid access token")
	}

	user, err := s.service.UserRepo.FindUserByID(userID)
	if err != nil {
		log.Println("Error finding user by ID: ", err)
		return models.UserLoginResponse{}, errors.New("user not found")
	}

	accessToken, err = s.service.Generator.GenerateJWT(user.ID, user.Email, user.RoleID)
	if err != nil {
		log.Println("Error generating JWT: ", err)
		return models.UserLoginResponse{}, errors.New("failed to generate access token")
	}
	refreshToken, err := s.service.Generator.GenerateRefreshToken(user.ID)
	if err != nil {
		log.Println("Error generating refresh token: ", err)
		return models.UserLoginResponse{}, errors.New("failed to generate refresh token")
	}

	result := models.UserLoginResponse{
		UserID:       user.ID,
		RoleID:       user.RoleID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return result, nil
}
