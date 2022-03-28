package main

import "net/http"


func main() {
	e := echo.New()
	
	e.GET("/hello",helloController)
	
	e.Start(":8080")
}

func helloController(e echo.Context)error {
	return e.String(http.StatusOK,"welcome echo..")
}