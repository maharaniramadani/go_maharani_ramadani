package routes

import (
	"praktikum/config"
	"praktikum/controller"
	"praktikum/middleware"

	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
)
func New() *echo.Echo {
	e := echo.New()
	// user routing
	// e.GET("/users", c.GetUsersController)
	// e.GET("/users/:id", c.GetUserController)
	// e.POST("/users", c.CreateUserController)
	// e.DELETE("/users/:id", c.DeleteUserController)
	// e.PUT("/users/:id", c.UpdateUserController)
   
	
	middleware.LogMiddleware(e)

	e.POST("/auth/login", controller.LoginController)
	e.POST("/users", controller.CreateUserController)
	e.GET("/books", controller.GetBooksController)
	e.GET("/books/:id", controller.GetBookController)

	jwtAuth := e.Group("/jwt")
	jwtAuth.Use(middleware.JWT([]byte(config.JwtSecret)))
	
	jwtAuth.GET("/users", controller.GetUsersController)
	jwtAuth.GET("/users/:id", controller.GetUserController)
	jwtAuth.DELETE("/users/:id", controller.DeleteUserController)
	jwtAuth.PUT("/users/:id", controller.UpdateUserController)

	jwtAuth.POST("/books", controller.CreateBookController)
	jwtAuth.DELETE("/books/:id", controller.DeleteBookController)
	jwtAuth.PUT("/books/:id", controller.UpdateBookController)

	return e

  }