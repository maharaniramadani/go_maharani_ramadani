package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type User struct{
	Id int `json:"id" form:"id"`
	Name string `json: "name" form:"name"`
	Email string `json: "email" form:"email"`
	Password string `json: "password" form:"password"`
	
}

var users []User
// ------------------controller

// get all users
func GetUsersController(c echo.Context)error  {
	return c.JSON(http.StatusOK,map[string]interface{}{
		"messages" : "success get all users",
		"users": users,
	})
}


// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	id,_ := strconv.Atoi(c.Param("id"))

	for _, user := range users {
		if user.Id == id {
			return c.JSON(http.StatusOK, (user))
		}
	}
	return c.JSON(http.StatusNotFound, (nil))
  }

func SuccessResponse(user User) {
	panic("unimplemented")
}

  // delete user by id
  func DeleteUserController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	
	for index, user := range users{
		if user.Id == id {
			users = append(users[:index], users[index+1:]...)
			return c.JSON(http.StatusOK, (user))
		}
	}
	return c.JSON(http.StatusNotFound, (nil))
  }

  // update user by id
  func UpdateUserController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))

	for index, user := range users {
		if user.Id == id {
			payload := User{}
			c.Bind(&payload)

			users[index].Name = payload.Name
			users[index].Email = payload.Email
			users[index].Password = payload.Password

			return c.JSON(http.StatusOK,(users[index]))
		}
	}
	return c.JSON(http.StatusNotFound,(nil))
  }
  
  // create new user
  func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)
  
	if len(users) == 0 {
	  user.Id = 1
	} else {
	  newId := users[len(users)-1].Id + 1
	  user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "messages": "success create user",
	  "user": user,
	})
  }
  // ---------------------------------------------------
  func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	e.GET("/users/:id", GetUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("/users/:id", DeleteUserController)
  
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
  }
  

// func main() {
// 	e := echo.New()
	
// 	e.GET("/hello",helloController)
	
// 	e.Start(":8080")
// }

// func helloController(e echo.Context)error {
// 	return e.String(http.StatusOK,"welcome echo..")
// }