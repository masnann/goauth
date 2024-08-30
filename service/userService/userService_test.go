package userservice_test

import (
	"errors"
	logerror "go-auth/helpers/logError"
	"go-auth/models"
	"go-auth/test"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Register(t *testing.T) {
	ts := test.SetupTestCase(t)

	req := models.UserRegisterRequest{
		Username: "test_user",
		Email:    "test_user@example.com",
		Password: "test123",
	}

	t.Run("Error Generate Hash", func(t *testing.T) {
		expectedErr := logerror.NewBusinessError("error generating hash", errors.New("hash generation failed"))

		ts.Generator.On("GenerateHash", req.Password).Return("", expectedErr).Once()

		result, err := ts.UserService.Register(req)

		assert.Error(t, err)
		assert.Equal(t, int64(0), result)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("Error Register User", func(t *testing.T) {
		hash := "hashedpassword"
		expectedMsg := "failed to register user"
		expectedErr := logerror.NewDatabaseError(expectedMsg, errors.New("insert failed"))

		ts.Generator.On("GenerateHash", req.Password).Return(hash, nil).Once()
		ts.UserRepo.On("Register", mock.Anything).Return(int64(0), expectedErr).Once()

		result, err := ts.UserService.Register(req)

		assert.Error(t, err)
		assert.Equal(t, int64(0), result)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("Error Assign Role", func(t *testing.T) {
		hash := "hashedpassword"
		userID := int64(1)
		expectedMsg := "failed to assign role"
		expectedErr := logerror.NewDatabaseError(expectedMsg, errors.New("assign role failed"))

		ts.Generator.On("GenerateHash", req.Password).Return(hash, nil).Once()
		ts.UserRepo.On("Register", mock.Anything).Return(userID, nil).Once()
		ts.PermissionRepo.On("AssignRoleToUserRequest", mock.Anything).Return(expectedErr).Once()

		result, err := ts.UserService.Register(req)

		ts.Generator.AssertExpectations(t)
		ts.UserRepo.AssertExpectations(t)
		ts.PermissionRepo.AssertExpectations(t)

		assert.Error(t, err)
		assert.Equal(t, int64(0), result)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("Success Case", func(t *testing.T) {
		hash := "hashedpassword"
		userID := int64(1)

		ts.Generator.On("GenerateHash", req.Password).Return(hash, nil).Once()
		ts.UserRepo.On("Register", mock.Anything).Return(userID, nil).Once()
		ts.PermissionRepo.On("AssignRoleToUserRequest", mock.Anything).Return(nil).Once()

		result, err := ts.UserService.Register(req)

		ts.Generator.AssertExpectations(t)
		ts.UserRepo.AssertExpectations(t)
		ts.PermissionRepo.AssertExpectations(t)

		assert.NoError(t, err)
		assert.Equal(t, userID, result)

	})
}

func Test_FindByID(t *testing.T) {
	ts := test.SetupTestCase(t)
	req := models.RequestID{ID: 1}
	expectedUser := models.UserModels{
		ID:       1,
		Username: "Test User",
		Email:    "testuser@example.com",
	}

	t.Run("User Not Found", func(t *testing.T) {
		ts := test.SetupTestCase(t)
		expectedMsg := "user not found"
		expectedErr := logerror.NewDatabaseError(expectedMsg, errors.New("no rows in result set"))

		ts.UserRepo.On("FindUserByID", req.ID).Return(models.UserModels{}, expectedErr)

		result, err := ts.UserService.FindUserByID(req)

		ts.UserRepo.AssertExpectations(t)

		assert.Error(t, err)
		assert.Equal(t, models.UserModels{}, result)
		assert.Equal(t, expectedErr, err)

	})

	t.Run("User Found", func(t *testing.T) {

		ts.UserRepo.On("FindUserByID", req.ID).Return(expectedUser, nil)

		result, err := ts.UserService.FindUserByID(req)

		ts.UserRepo.AssertExpectations(t)

		assert.NoError(t, err)
		assert.Equal(t, expectedUser, result)
	})
}
