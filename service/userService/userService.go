package userservice

import (
	"errors"
	"go-auth/helpers"
	"go-auth/models"
	"go-auth/service"
	"log"
	"time"
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
		Status:    "",
		CreatedAt: helpers.TimeStampNow(),
		UpdatedAt: "",
	}

	result, err := s.service.UserRepo.Register(newData)
	if err != nil {
		log.Println("Error registering user: ", err)
		return 0, errors.New("failed to register user")
	}

	// Assign Role
	newRole := models.AssignRoleToUserRequest{
		UserID: result,
		RoleID: 3,
	}

	err = s.service.PermissionRepo.AssignRoleToUserRequest(newRole)
	if err != nil {
		return 0, err
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

	isValidPassword, err := s.service.Generator.CompareHash(user.Password, req.Password)
	if !isValidPassword || err != nil {
		log.Println("Error comparing password: ", err)
		return models.UserLoginResponse{}, errors.New("invalid password")
	}

	role, err := s.service.PermissionRepo.FindUserRole(user.ID)
	if err != nil {
		log.Println("Error finding user role: ", err)
		return models.UserLoginResponse{}, errors.New("failed to find user role")
	}

	accessToken, err := s.service.Generator.GenerateJWT(user.ID, user.Email, role.RoleName)
	if err != nil {
		log.Println("Error generating JWT: ", err)
		return models.UserLoginResponse{}, errors.New("failed to generate access token")
	}

	refreshToken, err := s.service.Generator.GenerateRefreshToken(user.ID)
	if err != nil {
		log.Println("Error generating refresh token: ", err)
		return models.UserLoginResponse{}, errors.New("failed to generate refresh token")
	}

	permissions, err := s.service.PermissionRepo.FindPermissionsForUser(user.ID)
	if err != nil {
		log.Println("Error finding user permissions: ", err)
		return models.UserLoginResponse{}, errors.New("failed to find user permissions")
	}

	result := models.UserLoginResponse{
		UserID:       user.ID,
		RoleName:     role.RoleName,
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

	role, err := s.service.PermissionRepo.FindUserRole(user.ID)
	if err != nil {
		log.Println("Error finding user role: ", err)
		return models.UserLoginResponse{}, errors.New("failed to find user role")
	}

	accessToken, err = s.service.Generator.GenerateJWT(user.ID, user.Email, role.RoleName)
	if err != nil {
		log.Println("Error generating JWT: ", err)
		return models.UserLoginResponse{}, errors.New("failed to generate access token")
	}

	refreshToken, err := s.service.Generator.GenerateRefreshToken(user.ID)
	if err != nil {
		log.Println("Error generating refresh token: ", err)
		return models.UserLoginResponse{}, errors.New("failed to generate refresh token")
	}

	permissions, err := s.service.PermissionRepo.FindPermissionsForUser(user.ID)
	if err != nil {
		log.Println("Error finding user permissions: ", err)
		return models.UserLoginResponse{}, errors.New("failed to find user permissions")
	}

	result := models.UserLoginResponse{
		UserID:       user.ID,
		RoleName:     role.RoleName,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Permission:   permissions,
	}

	return result, nil
}

func (s UserService) GenerateOTP(req models.UserGenerateOTPRequest) (string, error) {
	otp, err := s.service.Generator.GenerateOTP(6)
	if err != nil {
		log.Println("Error generating OTP: ", err)
		return "", errors.New("failed to generate OTP")
	}

	// otpHash, err := s.service.Generator.GenerateHash(otp)
	// if err != nil {
	// 	log.Println("Error generating hash: ", err)
	// 	return "", errors.New("failed to generate hash")
	// }

	newData := models.OTPModels{
		UserID:    1,
		OtpHash:   otp,
		CreatedAt: helpers.TimeStampNow(),
		ExpiresAt: time.Now().Add(1000 * time.Second),
		IsUsed:    false,
	}

	err = s.service.UserRepo.SaveOtp(newData)
	if err != nil {
		log.Println("Error saving OTP: ", err)
		return "", errors.New("failed to save OTP")
	}

	return otp, nil
}

func (s UserService) ValidateOtp(req models.UserValidateOtpRequest) (bool, error) {


	otpStatus, err := s.service.UserRepo.CheckOtpStatus(1, req.OtpHash)
	if err != nil {
		log.Println("Error checking OTP status: ", err)
		return false, errors.New("otp not found")
	}

	if otpStatus.IsUsed {
		return false, errors.New("otp already used")
	}
	if time.Now().After(otpStatus.ExpiresAt) {
		return false, errors.New("otp expired")
	}

	return true, nil
}
