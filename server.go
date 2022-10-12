package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Checks that the server is up and running
func pingHandler(c echo.Context) error {
	message := "Bill building service. Version 0.0.1"
	return c.String(http.StatusOK, message)
}

// The declaration of all routes comes from it
func routes(e *echo.Echo) {
	e.GET("/", pingHandler)
	e.GET("/ping", pingHandler)
}

func server() {
	e := echo.New()
	routes(e)
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(1000)))
	log.Fatal(e.Start(":" + getEnvValue("PORT")))
}
