package controller

import (
	"net/http"
	"strconv"

	"praktikum/config"
	"praktikum/middleware"
	"praktikum/model"

	"github.com/labstack/echo/v4"
)

func LoginController(c echo.Context) error {

	user := model.User{}
	c.Bind(&user)

	err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, (map[string]interface{}{"error": err.Error()}))
	}

	token, err := middleware.CreateToken(strconv.Itoa(int(user.ID)), user.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, (map[string]interface{}{"error": err.Error()}))
	}

	responseData := model.UserLogin{}
	responseData.ID = user.ID
	responseData.Name = user.Name
	responseData.Email = user.Email
	responseData.Token = token

	return c.JSON(http.StatusOK,(responseData))
}