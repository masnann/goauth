package test

import (
	generator "go-auth/helpers/mocks"
	"go-auth/repository/mocks"
	"go-auth/service"
	userservice "go-auth/service/userService"
	"testing"
)

type TestSuite struct {
	Generator      *generator.GeneratorInterface
	UserRepo       *mocks.UserRepositoryInterface
	PermissionRepo *mocks.PermissionRepositoryInterface

	Service     service.Service
	UserService userservice.UserService
}

func SetupTestCase(t *testing.T) *TestSuite {

	generator := generator.NewGeneratorInterface(t)
	userRepo := mocks.NewUserRepositoryInterface(t)
	permissionRepo := mocks.NewPermissionRepositoryInterface(t)

	svc := service.NewService(generator, userRepo, permissionRepo)

	userService := userservice.NewUserService(svc)

	return &TestSuite{
		Generator:      generator,
		UserRepo:       userRepo,
		PermissionRepo: permissionRepo,

		Service:     svc,
		UserService: userService,
	}
}
