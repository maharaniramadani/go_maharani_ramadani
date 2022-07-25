package controller

import (
	"net/http"
	"praktikum/config"
	"praktikum/model"

	"github.com/labstack/echo/v4"
)

func GetBooksController(c echo.Context) error {
	var books []model.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK,books)
}

func GetBookController(c echo.Context) error {
	id := c.Param("id")
	book := model.Book{}

	config.DB.Where("ID = ?", id).First(&book)

	if book.ID == 0 {
		return c.JSON(http.StatusNotFound,nil)
	}

	return c.JSON(http.StatusOK,book)
}

func CreateBookController(c echo.Context) error {
	book := model.Book{}
	c.Bind(&book)

	if err := config.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, book)
}

func DeleteBookController(c echo.Context) error {
	id := c.Param("id")

	config.DB.Delete(&model.Book{}, id)

	return c.JSON(http.StatusOK, nil)
}

func UpdateBookController(c echo.Context) error {
	id := c.Param("id")
	book := model.Book{}

	config.DB.Where("ID = ?", id).First(&book)

	if book.ID == 0 {
		return c.JSON(http.StatusNotFound, nil)
	}

	payload := model.Book{}
	c.Bind(&payload)

	book.Title = payload.Title
	book.Author = payload.Author
	book.Year = payload.Year
	config.DB.Save(&book)

	return c.JSON(http.StatusOK,book)
}