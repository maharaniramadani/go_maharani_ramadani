package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var JwtSecret string

func InitDB(){
	connectionString:= "root:@/middleware?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString),&gorm.Config{})
	if err != nil{
		panic(err)
	}
	
  }

