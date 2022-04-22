package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	Id int `json:"id" form:"id"`
	Name string `json: "name" form:"name"`
	Email string `json: "email" form:"email"`
	Password string `json: "password" form:"password"`
	
}

// ------------------controller

// get all users
func GetUsersController(c echo.Context)error  {
	var users []User

	if err := DB.Find(&users).Error;err !=nil{
		return echo.NewHTTPError(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,map[string]interface{}{
		"messages" : "success get all users",
		"users": users,
	})
}


// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	id := c.Param("id")
	user := User{}

	DB.Where("Id=?",id).First(&user)

	if user.Id ==0 {
		return c.JSON(http.StatusNotFound,nil)
	}
	return c.JSON(http.StatusOK, (user))
  }


  // delete user by id
  func DeleteUserController(c echo.Context) error {
	// your solution here
	id := c.Param("id")
	DB.Delete(&User{},id)
	return c.JSON(http.StatusOK,nil)
  }

  // update user by id
  func UpdateUserController(c echo.Context) error {
	// your solution here
	id := c.Param("Id")
	user := User{}

	DB.Where("Id = ?", id).First(&user)

	if user.Id == 0 {
		return c.JSON(http.StatusNotFound,nil)
	}

	payload := User{}
	c.Bind(&payload)

	user.Id = payload.Id
	user.Name = payload.Name
	user.Email = payload.Email
	user.Password = payload.Password
	DB.Save(&user)

	return c.JSON(http.StatusOK,(user))
  }
  
  // create new user
  func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)
  
	if err := DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	  }
	  return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	  })
  }
  // ---------------------------------------------------
var ( DB *gorm.DB)

func initDB(){
	var err error
	DB, err = gorm.Open(mysql.Open("root:@/test?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil{
		panic(err)
	}
  }

func  initMigrate(){
	DB.AutoMigrate(&User{})
}

// ---------------------------------------------------
  func main() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "ini adalah format yang beda : ${remote_ip} ${status}",
	  }))	  

	e.GET("/", func (ctx echo.Context) error  {
		data := "hello from/index"
		return ctx.String(http.StatusOK,data)
		
	})
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	e.GET("/users/:id", GetUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("/users/:id", DeleteUserController)
  
	// start the server, and log if it fails
	e.Start(":8000")
  }
  