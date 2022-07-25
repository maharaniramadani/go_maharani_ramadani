package main

import (
	"docker/controller"

	"github.com/labstack/echo/v4"
)
func main() {
	e := echo.New()

	e.GET("/healthcheck", controller.HealthCheck)
	e.Start(":8080")
}