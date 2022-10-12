package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Checks that the server is up and running
func pingHandler(c echo.Context) error {
	message := "AB-tests service. Version 1.0.0"
	return c.String(http.StatusOK, message)
}

func createHypothesisHandler(c echo.Context) error {
	var jsonMap map[string]interface{}
	if err := c.Bind(&jsonMap); err != nil {
		return err
	}
	hypothesis := createHypothesis(jsonMap)
	return c.JSON(http.StatusOK, hypothesis)
}

func getHypothesisHandler(c echo.Context) error {
	hypothesis := c.QueryParam("hypothesis")
	userId := c.QueryParam("user")
	option, err := getHypothesis(hypothesis, userId)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.String(http.StatusOK, option)
}

func getAllHypothesisHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, repository.GetAll())
}

func deleteOneHandler(c echo.Context) error {
	hypothesis := c.QueryParam("hypothesis")
	if hypothesis == "" {
		return c.String(http.StatusBadRequest, "No hypothesis for deletion")
	}
	err := deleter(hypothesis)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.NoContent(http.StatusOK)
}

func deleteAllHandler(c echo.Context) error {
	err := deleter("")
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.NoContent(http.StatusOK)
}

// The declaration of all routes comes from it
func routes(e *echo.Echo) {
	e.GET("/", pingHandler)
	e.GET("/ping", pingHandler)
	e.GET("/all", getAllHypothesisHandler)
	e.POST("/create", createHypothesisHandler)
	e.PATCH("/forUser", getHypothesisHandler)
	e.DELETE("/one", deleteOneHandler)
	e.DELETE("/", deleteAllHandler)
}

func server() {
	e := echo.New()
	routes(e)
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(1000)))
	log.Fatal(e.Start(":" + getEnvValue("PORT")))
}
