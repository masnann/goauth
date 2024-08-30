package main

import (
	"context"
	"fmt"
	"go-auth/app"
	"go-auth/config"
	"go-auth/helpers"
	"go-auth/repository"
	"go-auth/routes"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := config.OpenConnection(); err != nil {
		panic(fmt.Sprintf("Open Connection Failed: %s", err.Error()))
	}
	defer config.CloseConnectionDB()

	DB := config.DBConnection()

	mongo := config.ConnectMongo(context.Background())

	//Initialize repository and service
	repo := repository.NewRepository(DB, mongo)
	handler := app.SetupApp(repo)

	e := echo.New()
	routes.ApiRoutes(e, handler)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		result := helpers.ResponseJSON(false, strconv.Itoa(report.Code), err.Error(), nil)
		c.Logger().Error(report)
		c.JSON(report.Code, result)
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, GO AUTH !")
	})
	// Start server
	port := fmt.Sprintf(":%s", config.GetEnv("APP_PORT", "8080"))

	e.Logger.Fatal(e.Start(port))
}
