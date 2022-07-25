package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)
func HealthCheck(ctx echo.Context) error{
	return ctx.HTML(http.StatusOK,"OK HealthCheck")
}