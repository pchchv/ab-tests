package main

import (
	"log"

	"github.com/labstack/echo/v4"
)

func server() {
	e := echo.New()
	log.Fatal(e.Start(":" + getEnvValue("PORT")))
}
