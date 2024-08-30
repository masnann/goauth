package userhandler

import (
	"context"
	"encoding/json"
	"errors"
	"go-auth/app"
	"go-auth/config"
	"go-auth/constants"
	"go-auth/helpers"
	"go-auth/models"
	"go-auth/repository"
	"go-auth/test"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Register(t *testing.T) {
	db, mock := test.SetupMockDB(t)
	defer db.Close()

	mongo := config.ConnectMongo(context.Background())

	repo := repository.NewRepository(db, mongo)
	handler := app.SetupApp(repo)
	userHandler := NewUserHandler(handler)

	req := models.UserRegisterRequest{
		Username:  "testuser",
		Email:     "testuser@example.com",
		Password:  "testpassword",
		Status:    "active",
		CreatedAt: helpers.TimeStampNow(),
		UpdatedAt: helpers.TimeStampNow(),
	}

	// Define query pattern for tests, matching the repository implementation
	query := `
    INSERT INTO users (username, email, password, status, created_at, updated_at) 
    VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING id`

	executeRequest := func(
		t *testing.T,
		reqBody models.UserRegisterRequest,
		expected models.TestingHandlerExpected,
	) {
		_, rec, ctx := test.NewRequestRecorder(
			models.TestingHandlerRequest{
				Method: http.MethodPost,
				Path:   "/api/v1/private/user/register",
				Body:   reqBody,
			},
		)

		// Run the handler
		err := userHandler.Register(ctx)
		assert.Nil(t, err)
		assert.Equal(t, expected.StatusCode, rec.Code)

		// Check the response
		var result models.Response
		err = json.Unmarshal(rec.Body.Bytes(), &result)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expected.Success, result.Success)
		assert.Equal(t, expected.Code, result.StatusCode)
		assert.Contains(t, result.Message, expected.Message)
	}

	t.Run("Error Validation", func(t *testing.T) {
		executeRequest(t,
			models.UserRegisterRequest{},
			models.TestingHandlerExpected{
				StatusCode: http.StatusBadRequest,
				Code:       constants.VALIDATION_ERROR_CODE,
				Success:    false,
				Message:    "Validation error: Field 'Username' is required",
			})
	})

	t.Run("Database Error Case", func(t *testing.T) {
		// Simulate a database error during registration
		mock.ExpectQuery(query).
			WithArgs(req.Username, req.Email, req.Password, req.Status, req.CreatedAt, req.UpdatedAt).
			WillReturnError(errors.New("error inserting user"))

		executeRequest(t,
			req,
			models.TestingHandlerExpected{
				StatusCode: http.StatusInternalServerError,
				Code:       constants.DATABASE_ERROR_CODE,
				Success:    false,
				Message:    "failed to register user",
			})
	})

}
