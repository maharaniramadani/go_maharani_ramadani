package main

import (
	"praktikum/config"
	"praktikum/routes"
)

func main(){
	config.InitDB()
	e := routes.New()
	e.Start(":8000")
}