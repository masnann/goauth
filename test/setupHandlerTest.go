package test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"go-auth/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
)

func SetupMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}

func NewRequestRecorder(req models.TestingHandlerRequest) (*http.Request, *httptest.ResponseRecorder, echo.Context) {
	reqBodyBytes, _ := json.Marshal(req.Body)
	request := httptest.NewRequest(req.Method, req.Path, bytes.NewReader(reqBodyBytes))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e := echo.New()
	ctx := e.NewContext(request, rec)
	return request, rec, ctx
}
